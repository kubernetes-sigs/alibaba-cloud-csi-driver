package dara

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/proxy"
)

type WebSocketMessageType int

const (
	WebSocketMessageTypeText WebSocketMessageType = iota
	WebSocketMessageTypeBinary
	WebSocketMessageTypePing
	WebSocketMessageTypePong
	WebSocketMessageTypeClose
)

type WebSocketMessage struct {
	Type      WebSocketMessageType
	Payload   []byte
	Headers   map[string]string
	Timestamp time.Time
}

type WebSocketSessionInfo struct {
	SessionID   string // Session ID from server (x-acs-ws-session-id header)
	RequestID   string // Request ID from server (x-acs-request-id header)
	ConnectedAt time.Time
	RemoteAddr  string
	LocalAddr   string
	Attributes  map[string]interface{}
}

func GetWebSocketPingInterval(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketPingInterval
	}
	return nil
}

func GetWebSocketPongTimeout(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketPongTimeout
	}
	return nil
}

func GetWebSocketEnableReconnect(runtime interface{}) *bool {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketEnableReconnect
	}
	return nil
}

func GetWebSocketReconnectInterval(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketReconnectInterval
	}
	return nil
}

func GetWebSocketMaxReconnectTimes(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketMaxReconnectTimes
	}
	return nil
}

func GetWebSocketWriteTimeout(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketWriteTimeout
	}
	return nil
}

func GetWebSocketHandshakeTimeout(runtime interface{}) *int {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		return rt.WebSocketHandshakeTimeout
	}
	return nil
}

// GetWebSocketHandler safely extracts WebSocketHandler from runtime
// Returns the handler as WebSocketHandler interface
func GetWebSocketHandler(runtime interface{}) WebSocketHandler {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeOptions); ok {
		if handler, ok := rt.WebSocketHandler.(WebSocketHandler); ok {
			return handler
		}
	}
	return nil
}

func GetWebsocketSubProtocol(runtime interface{}) *string {
	if runtime == nil {
		return nil
	}
	if rt, ok := runtime.(*RuntimeObject); ok {
		return rt.WebsocketSubProtocol
	}
	return nil
}

// NewWebSocketResponse creates a Response from WebSocket handshake HTTP response
func NewWebSocketResponse(httpResponse *http.Response) *Response {
	res := &Response{}
	res.Headers = make(map[string]*string)
	res.StatusCode = Int(httpResponse.StatusCode)
	res.StatusMessage = String(httpResponse.Status)
	res.Body = httpResponse.Body

	for key, values := range httpResponse.Header {
		if len(values) > 0 {
			res.Headers[strings.ToLower(key)] = String(values[0])
		}
	}

	return res
}

type WebSocketHandler interface {
	AfterConnectionEstablished(session *WebSocketSessionInfo) error
	HandleRawMessage(session *WebSocketSessionInfo, message *WebSocketMessage) error
	HandleError(session *WebSocketSessionInfo, err error) error
	AfterConnectionClosed(session *WebSocketSessionInfo, code int, reason string) error
}

// AbstractWebSocketHandler provides default implementations for WebSocketHandler interface
// Users can embed this struct in their custom handlers to avoid implementing all methods
type AbstractWebSocketHandler struct{}

func (h *AbstractWebSocketHandler) AfterConnectionEstablished(session *WebSocketSessionInfo) error {
	return nil
}

func (h *AbstractWebSocketHandler) HandleRawMessage(session *WebSocketSessionInfo, message *WebSocketMessage) error {
	return nil
}

func (h *AbstractWebSocketHandler) HandleError(session *WebSocketSessionInfo, err error) error {
	return nil
}

func (h *AbstractWebSocketHandler) AfterConnectionClosed(session *WebSocketSessionInfo, code int, reason string) error {
	return nil
}

type WebSocketClient interface {
	Connect(ctx context.Context, request *Request, runtimeObject *RuntimeObject) (*Response, error)
	Disconnect() error
	Reconnect() (*Response, error)
	ReconnectGracefully() (*Response, error) // Graceful reconnection with session ID
	IsConnected() bool
	SendText(text string) error
	SendBinary(data []byte) error
	GetSessionInfo() *WebSocketSessionInfo
	Close() error
}

type DefaultWebSocketClient struct {
	handler      WebSocketHandler // 必需：消息处理器
	stopChan     chan struct{}    // 必需：内部 goroutine 的快速停止
	pongReceived chan struct{}    // 必需：ping/pong 机制
	state        int32            // 必需：连接状态 (0=disconnected, 1=connecting, 2=connected, 3=disconnecting)

	conn          *websocket.Conn       // 连接时建立
	session       *WebSocketSessionInfo // 连接时创建
	request       *Request              // 连接时传入
	runtimeObject *RuntimeObject        // 连接时传入
	ctx           context.Context       // 连接时创建, 管理连接生命周期
	cancel        context.CancelFunc    // 连接时创建, 管理连接生命周期

	// 运行时赋值的字段
	reconnectCount int          // 重连时递增
	pingTicker     *time.Ticker // 启动 ping/pong 时创建

	// 超时配置（在 Connect 时一次性计算，后续直接使用）
	pingInterval      time.Duration // Ping 间隔
	reconnectInterval time.Duration // 重连间隔
	writeTimeout      time.Duration // 写入超时
	readTimeout       time.Duration // 读取超timeout
	pongTimeout       time.Duration // Pong 超时
	maxReconnectTimes int           // 最大重连次数

	// 零值可用的字段（sync 类型）
	reconnectMu          sync.Mutex
	closeMu              sync.Mutex
	connMu               sync.RWMutex // Protects conn field from data races
	wg                   sync.WaitGroup
	websocketSubProtocol *string // 在 Connect 时从 runtimeObject 获取并缓存
}

func NewDefaultWebSocketClient(handler WebSocketHandler) (*DefaultWebSocketClient, error) {
	if handler == nil {
		return nil, errors.New("handler cannot be nil")
	}

	client := &DefaultWebSocketClient{
		handler:      handler,
		stopChan:     make(chan struct{}),
		pongReceived: make(chan struct{}, 1),
		state:        0, // disconnected
	}

	return client, nil
}

func NewWebSocketClientAndConnect(request *Request, runtimeObject *RuntimeObject) (*DefaultWebSocketClient, *Response, error) {
	if runtimeObject == nil {
		return nil, nil, errors.New("runtimeObject cannot be nil")
	}

	handler := runtimeObject.WebSocketHandler
	if handler == nil {
		return nil, nil, errors.New("WebSocketHandler is required: please set it in runtimeObject.WebSocketHandler")
	}

	ctx := context.Background()

	client, err := NewDefaultWebSocketClient(handler)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.Connect(ctx, request, runtimeObject)
	if err != nil {
		return nil, nil, err
	}

	return client, response, nil
}

func NewWebSocketClientAndConnectWithContext(ctx context.Context, request *Request, runtimeObject *RuntimeObject) (*DefaultWebSocketClient, *Response, error) {
	if runtimeObject == nil {
		return nil, nil, errors.New("runtimeObject cannot be nil")
	}

	handler := runtimeObject.WebSocketHandler
	if handler == nil {
		return nil, nil, errors.New("WebSocketHandler is required: please set it in runtimeObject.WebSocketHandler")
	}

	client, err := NewDefaultWebSocketClient(handler)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.Connect(ctx, request, runtimeObject)
	if err != nil {
		return nil, nil, err
	}

	return client, response, nil
}

func buildWebSocketURL(request *Request) (string, error) {
	if request == nil {
		return "", errors.New("request cannot be nil")
	}

	if request.Protocol == nil {
		request.Protocol = String("ws")
	} else {
		protocol := strings.ToLower(StringValue(request.Protocol))
		if protocol == "http" {
			protocol = "ws"
		} else if protocol == "https" {
			protocol = "wss"
		}
		request.Protocol = String(protocol)
	}

	if request.Domain == nil {
		if host, ok := request.Headers["host"]; ok && host != nil {
			request.Domain = host
		} else {
			return "", errors.New("domain is required (set in request.Headers[\"host\"] or request.Domain)")
		}
	}

	requestURL := fmt.Sprintf("%s://%s", StringValue(request.Protocol), StringValue(request.Domain))

	if request.Pathname != nil {
		requestURL += StringValue(request.Pathname)
	} else {
		requestURL += "/"
	}

	if len(request.Query) > 0 {
		q := url.Values{}
		for key, value := range request.Query {
			if value != nil {
				q.Add(key, StringValue(value))
			}
		}
		querystring := q.Encode()
		if len(querystring) > 0 {
			if strings.Contains(requestURL, "?") {
				requestURL = fmt.Sprintf("%s&%s", requestURL, querystring)
			} else {
				requestURL = fmt.Sprintf("%s?%s", requestURL, querystring)
			}
		}
	}

	return requestURL, nil
}

func (c *DefaultWebSocketClient) updateTimeoutConfig(runtimeObject *RuntimeObject) {
	c.pingInterval = time.Duration(IntValue(runtimeObject.WebSocketPingInterval)) * time.Millisecond
	if c.pingInterval <= 0 {
		c.pingInterval = 30 * time.Second // Default value
	}

	c.reconnectInterval = time.Duration(IntValue(runtimeObject.WebSocketReconnectInterval)) * time.Millisecond
	if c.reconnectInterval <= 0 {
		c.reconnectInterval = 5 * time.Second // Default
	}

	c.writeTimeout = time.Duration(IntValue(runtimeObject.WebSocketWriteTimeout)) * time.Millisecond
	if c.writeTimeout <= 0 {
		c.writeTimeout = 30 * time.Second // Default
	}

	c.readTimeout = time.Duration(IntValue(runtimeObject.ReadTimeout)) * time.Millisecond
	if c.readTimeout <= 0 {
		c.readTimeout = 0 // No read timeout if not configured
	}

	c.pongTimeout = time.Duration(IntValue(runtimeObject.WebSocketPongTimeout)) * time.Millisecond
	if c.pongTimeout <= 0 {
		c.pongTimeout = 5 * time.Second // Default
	}

	c.maxReconnectTimes = IntValue(runtimeObject.WebSocketMaxReconnectTimes)
	if c.maxReconnectTimes <= 0 {
		c.maxReconnectTimes = 5 // Default
	}
}

func (c *DefaultWebSocketClient) Connect(ctx context.Context, request *Request, runtimeObject *RuntimeObject) (*Response, error) {
	if request == nil {
		return nil, errors.New("request cannot be nil")
	}
	if runtimeObject == nil {
		return nil, errors.New("runtimeObject cannot be nil")
	}

	c.request = request
	c.runtimeObject = runtimeObject

	c.updateTimeoutConfig(runtimeObject)

	c.websocketSubProtocol = GetWebsocketSubProtocol(runtimeObject)

	atomic.StoreInt32(&c.state, 1) // connecting

	// Create a cancellable context for this connection
	// This allows us to cancel reconnect operations when the client is closed
	c.closeMu.Lock()
	if c.cancel != nil {
		c.cancel() // Cancel previous context if exists
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.closeMu.Unlock()

	cleanupContext := func() {
		c.closeMu.Lock()
		if c.cancel != nil {
			c.cancel()
			c.cancel = nil
		}
		c.closeMu.Unlock()
	}

	requestURL, err := buildWebSocketURL(request)
	if err != nil {
		atomic.StoreInt32(&c.state, 0) // disconnected
		cleanupContext()
		return nil, err
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		atomic.StoreInt32(&c.state, 0) // disconnected
		cleanupContext()
		return nil, err
	}

	handshakeTimeout := time.Duration(IntValue(runtimeObject.WebSocketHandshakeTimeout)) * time.Millisecond
	if handshakeTimeout <= 0 {
		handshakeTimeout = 30 * time.Second // Default
	}
	connectTimeout := time.Duration(IntValue(runtimeObject.ConnectTimeout)) * time.Millisecond
	if connectTimeout <= 0 {
		connectTimeout = 10 * time.Second // Default
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: handshakeTimeout,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}

	if u.Scheme == "wss" || u.Scheme == "https" {
		tlsConfig, err := c.configureTLS(runtimeObject)
		if err != nil {
			atomic.StoreInt32(&c.state, 0) // disconnected
			cleanupContext()
			return nil, err
		}
		dialer.TLSClientConfig = tlsConfig
	}

	// Configure proxy and network dialer (Priority: SOCKS5 > HTTP/HTTPS proxy)
	if err := c.configureProxy(&dialer, u, runtimeObject, request, connectTimeout); err != nil {
		atomic.StoreInt32(&c.state, 0) // disconnected
		cleanupContext()
		return nil, err
	}

	header := http.Header{}
	if request.Headers != nil {
		for k, v := range request.Headers {
			if v != nil && k != "host" && k != "content-length" {
				header.Set(k, StringValue(v))
			}
		}
	}

	if c.websocketSubProtocol != nil && StringValue(c.websocketSubProtocol) != "" {
		header.Set("Sec-WebSocket-Protocol", StringValue(c.websocketSubProtocol))
	}

	debugLog("[WebSocket] Handshake headers:")
	for k, v := range header {
		debugLog("  %s: %v", k, v)
	}

	connectCtx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()

	conn, resp, err := dialer.DialContext(connectCtx, requestURL, header)
	if err != nil {
		atomic.StoreInt32(&c.state, 0) // disconnected
		if resp != nil {
			debugLog("[WebSocket] Handshake failed. Response status: %s", resp.Status)
			debugLog("[WebSocket] Response headers:")
			for k, v := range resp.Header {
				debugLog("  %s: %v", k, v)
			}
		}
		cleanupContext()
		return nil, err
	}

	c.connMu.Lock()
	c.conn = conn
	c.connMu.Unlock()
	atomic.StoreInt32(&c.state, 2) // connected

	c.setupPongHandler()

	// HTTP headers are case-insensitive, Go's Header.Get() is case-insensitive
	sessionID := ""
	requestID := ""
	if resp != nil && resp.Header != nil {
		sessionID = resp.Header.Get("x-acs-ws-session-id")
		requestID = resp.Header.Get("x-acs-request-id")
	}

	if sessionID == "" {
		sessionID = generateSessionID()
	}

	c.session = &WebSocketSessionInfo{
		SessionID:   sessionID,
		RequestID:   requestID,
		ConnectedAt: time.Now(),
		RemoteAddr:  conn.RemoteAddr().String(),
		LocalAddr:   conn.LocalAddr().String(),
		Attributes:  make(map[string]interface{}),
	}

	c.startMessageHandlers()

	// Start ping/pong if configured
	if c.pingInterval > 0 {
		c.startPingPong()
	}

	if err := c.handler.AfterConnectionEstablished(c.session); err != nil {
		// Cleanup connection and context on handler error
		c.connMu.Lock()
		if c.conn != nil {
			c.conn.Close()
			c.conn = nil
		}
		c.connMu.Unlock()
		atomic.StoreInt32(&c.state, 0) // disconnected
		cleanupContext()
		return nil, err
	}

	response := NewWebSocketResponse(resp)
	return response, nil
}

func (c *DefaultWebSocketClient) Disconnect() error {
	return c.disconnect(1000, "Normal closure")
}

func (c *DefaultWebSocketClient) disconnect(code int, reason string) error {
	// state == 0 means disconnected
	currentState := atomic.LoadInt32(&c.state)
	if currentState == 0 {
		c.closeMu.Lock()
		hasCancel := c.cancel != nil
		c.closeMu.Unlock()
		// no cancel function, it means never connected or already fully closed
		if !hasCancel {
			return nil
		}
		// If state is 0 but cancel exists, try to cancel and clean up, but don't error if already done
	}

	atomic.StoreInt32(&c.state, 3) // disconnecting

	c.stopPingPong()

	// Signal goroutines to stop first (before closing connection)
	// This allows goroutines to exit gracefully
	select {
	case <-c.stopChan:
		// Already closed, don't close againq
	default:
		close(c.stopChan)
	}

	// Close connection (this will cause ReadMessage to return error)
	c.connMu.Lock()
	if c.conn != nil {
		deadline := time.Now().Add(time.Second)
		c.conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(code, reason),
			deadline)
		c.conn.Close()
		c.conn = nil
	}
	c.connMu.Unlock()

	// Wait for all goroutines to finish (with timeout to avoid deadlock)
	// This is done OUTSIDE the lock to avoid blocking other operations
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All goroutines finished
	case <-time.After(5 * time.Second):
		// Timeout - log warning but continue
		debugLog("[WebSocket] Warning: timeout waiting for goroutines to finish")
	}

	if c.session != nil {
		c.handler.AfterConnectionClosed(c.session, code, reason)
	}

	atomic.StoreInt32(&c.state, 0) // disconnected

	// Cancel context (with lock)
	c.closeMu.Lock()
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	c.closeMu.Unlock()

	return nil
}

func (c *DefaultWebSocketClient) Reconnect() (*Response, error) {
	return c.reconnectInternal(false)
}

// ReconnectGracefully performs a graceful reconnection (server-initiated via RECONNECT control message, with session ID)
func (c *DefaultWebSocketClient) ReconnectGracefully() (*Response, error) {
	return c.reconnectInternal(true)
}

// graceful: true for graceful reconnection (uses session ID), false for normal reconnection (doesn't use session ID)
func (c *DefaultWebSocketClient) reconnectInternal(graceful bool) (*Response, error) {
	// Use minimal lock to check conditions and prevent concurrent reconnection
	c.reconnectMu.Lock()

	// For normal reconnection, skip if already connected
	// For graceful reconnection (server-initiated), allow reconnection even if connected
	if !graceful && c.IsConnected() {
		c.reconnectMu.Unlock()
		debugLog("[WebSocket] Already connected, skipping normal reconnect")
		return nil, errors.New("already connected")
	}

	// Check if context is already cancelled (client might be closing)
	c.reconnectMu.Unlock()

	// Read ctx with proper lock
	c.closeMu.Lock()
	ctx := c.ctx
	c.closeMu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Re-acquire lock for state changes
	c.reconnectMu.Lock()

	if !BoolValue(c.runtimeObject.WebSocketEnableReconnect) {
		c.reconnectMu.Unlock()
		return nil, errors.New("reconnect is disabled")
	}

	if c.reconnectCount >= c.maxReconnectTimes {
		c.reconnectMu.Unlock()
		return nil, fmt.Errorf("max reconnect times reached: %d", c.maxReconnectTimes)
	}

	// Save previous session ID before cleanup (only for graceful reconnection)
	previousSessionID := ""
	if graceful {
		if c.session != nil && c.session.SessionID != "" {
			previousSessionID = c.session.SessionID
		} else {
			c.reconnectMu.Unlock()
			return nil, errors.New("graceful reconnection requires existing session ID")
		}
	}

	c.cleanupResources()

	// Reset state for reconnection (state will be set to connecting in Connect())
	c.stopChan = make(chan struct{})
	c.reconnectCount++

	if c.request == nil {
		c.reconnectMu.Unlock()
		return nil, errors.New("request is nil, cannot reconnect")
	}

	// For graceful reconnection, add previous session ID to request headers
	if graceful && previousSessionID != "" {
		if c.request.Headers == nil {
			c.request.Headers = make(map[string]*string)
		}
		c.request.Headers["X-Acs-Ws-Session-Id"] = String(previousSessionID)
		debugLog("[WebSocket] Graceful reconnection with session ID: %s", previousSessionID)
	} else {
		if c.request.Headers != nil {
			delete(c.request.Headers, "X-Acs-Ws-Session-Id")
		}
		debugLog("[WebSocket] Normal reconnection (without session ID)")
	}

	reconnectRequest := c.request
	reconnectRuntime := c.runtimeObject

	// Release lock before network I/O to avoid blocking other operations
	c.reconnectMu.Unlock()

	c.closeMu.Lock()
	waitCtx := c.ctx
	c.closeMu.Unlock()

	// Use context-aware sleep to allow cancellation during reconnect interval
	// time.sleep is a blocking operation, and context cancellation is not supported
	select {
	case <-waitCtx.Done():
		// Context was cancelled during wait, return error
		return nil, waitCtx.Err()
	case <-time.After(c.reconnectInterval):
		// Wait completed, continue with reconnect
	}

	c.closeMu.Lock()
	reconnectCtx := c.ctx
	c.closeMu.Unlock()

	// Check context before attempting connection
	select {
	case <-reconnectCtx.Done():
		return nil, reconnectCtx.Err()
	default:
	}

	result, err := c.Connect(reconnectCtx, reconnectRequest, reconnectRuntime)

	if err == nil {
		c.reconnectMu.Lock()
		c.reconnectCount = 0 // Reset on success
		c.reconnectMu.Unlock()
	}

	return result, err
}

// cleanupResources cleans up all resources (ping/pong, goroutines, connection, session)
// This is used before reconnecting to ensure a clean state
func (c *DefaultWebSocketClient) cleanupResources() {
	atomic.StoreInt32(&c.state, 3) // disconnecting

	c.stopPingPong()

	// Signal goroutines to stop first (before closing connection)
	// This allows goroutines to exit gracefully
	select {
	case <-c.stopChan:
		// Already closed, don't close again
	default:
		close(c.stopChan)
	}

	c.connMu.Lock()
	if c.conn != nil {
		deadline := time.Now().Add(time.Second)
		c.conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseGoingAway, "Reconnecting"),
			deadline)
		c.conn.Close()
		c.conn = nil
	}
	c.connMu.Unlock()

	// Wait for all goroutines to finish (with timeout to avoid deadlock)
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All goroutines finished
	case <-time.After(5 * time.Second):
		// Timeout - log warning but continue
		debugLog("[WebSocket] Warning: timeout waiting for goroutines to finish during cleanup")
	}

	c.session = nil
	atomic.StoreInt32(&c.state, 0) // disconnected
}

func (c *DefaultWebSocketClient) IsConnected() bool {
	return atomic.LoadInt32(&c.state) == 2
}

func (c *DefaultWebSocketClient) SendText(text string) error {
	if !c.IsConnected() {
		return errors.New("not connected")
	}

	c.connMu.RLock()
	conn := c.conn
	if conn == nil {
		c.connMu.RUnlock()
		return errors.New("connection is nil")
	}
	if c.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(text))
	c.connMu.RUnlock()
	return err
}

func (c *DefaultWebSocketClient) SendBinary(data []byte) error {
	if !c.IsConnected() {
		return errors.New("not connected")
	}

	c.connMu.RLock()
	conn := c.conn
	if conn == nil {
		c.connMu.RUnlock()
		return errors.New("connection is nil")
	}
	if c.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
	}
	err := conn.WriteMessage(websocket.BinaryMessage, data)
	c.connMu.RUnlock()
	return err
}

func (c *DefaultWebSocketClient) GetSessionInfo() *WebSocketSessionInfo {
	return c.session
}

func (c *DefaultWebSocketClient) Close() error {
	return c.disconnect(1000, "Client closed")
}

func (c *DefaultWebSocketClient) configureTLS(runtimeObject *RuntimeObject) (*tls.Config, error) {
	if BoolValue(runtimeObject.IgnoreSSL) {
		return &tls.Config{
			InsecureSkipVerify: true,
		}, nil
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	// Client certificate and key
	if runtimeObject.Key != nil && runtimeObject.Cert != nil &&
		StringValue(runtimeObject.Key) != "" && StringValue(runtimeObject.Cert) != "" {
		cert, err := tls.X509KeyPair([]byte(StringValue(runtimeObject.Cert)), []byte(StringValue(runtimeObject.Key)))
		if err != nil {
			return nil, fmt.Errorf("failed to load client certificate: %w", err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// Root CA certificate
	if runtimeObject.Ca != nil && StringValue(runtimeObject.Ca) != "" {
		clientCertPool := x509.NewCertPool()
		if !clientCertPool.AppendCertsFromPEM([]byte(StringValue(runtimeObject.Ca))) {
			return nil, errors.New("failed to parse root certificate")
		}
		tlsConfig.RootCAs = clientCertPool
	}

	return tlsConfig, nil
}

// configureProxy configures proxy and network dialer for WebSocket connections
// Priority: SOCKS5 > HTTP/HTTPS proxy > Direct connection
func (c *DefaultWebSocketClient) configureProxy(
	dialer *websocket.Dialer,
	u *url.URL,
	runtimeObject *RuntimeObject,
	request *Request,
	connectTimeout time.Duration,
) error {
	// Priority 1: SOCKS5 proxy (works at TCP layer, uses NetDialContext)
	if runtimeObject.Socks5Proxy != nil && StringValue(runtimeObject.Socks5Proxy) != "" {
		return c.configureSOCKS5Proxy(dialer, runtimeObject, connectTimeout)
	}

	// Priority 2: HTTP/HTTPS proxy (works at HTTP layer, uses Proxy function)
	if err := c.configureHTTPProxy(dialer, u, runtimeObject, request); err != nil {
		return err
	}

	// Priority 3: Direct connection with optional local address
	c.configureNetDialer(dialer, runtimeObject, connectTimeout)
	return nil
}

// configureSOCKS5Proxy configures SOCKS5 proxy via NetDialContext
func (c *DefaultWebSocketClient) configureSOCKS5Proxy(
	dialer *websocket.Dialer,
	runtimeObject *RuntimeObject,
	connectTimeout time.Duration,
) error {
	socks5Proxy, err := url.Parse(StringValue(runtimeObject.Socks5Proxy))
	if err != nil {
		return fmt.Errorf("failed to parse SOCKS5 proxy URL: %w", err)
	}

	var auth *proxy.Auth
	if socks5Proxy.User != nil {
		password, _ := socks5Proxy.User.Password()
		auth = &proxy.Auth{
			User:     socks5Proxy.User.Username(),
			Password: password,
		}
	}

	socks5Network := strings.ToLower(StringValue(runtimeObject.Socks5NetWork))
	if socks5Network == "" {
		socks5Network = "tcp"
	}

	socks5Dialer, err := proxy.SOCKS5(
		socks5Network,
		socks5Proxy.Host,
		auth,
		&net.Dialer{
			Timeout:   connectTimeout,
			DualStack: true,
			LocalAddr: getLocalAddr(StringValue(runtimeObject.LocalAddr)),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to setup SOCKS5 proxy: %w", err)
	}

	dialer.NetDialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return socks5Dialer.Dial(network, addr)
	}

	return nil
}

// configureHTTPProxy configures HTTP/HTTPS proxy via Proxy function
func (c *DefaultWebSocketClient) configureHTTPProxy(
	dialer *websocket.Dialer,
	u *url.URL,
	runtimeObject *RuntimeObject,
	request *Request,
) error {
	protocol := u.Scheme
	host := u.Hostname()

	// Check if host is in no-proxy list
	noProxyList := getNoProxy(protocol, runtimeObject)
	for _, noProxyHost := range noProxyList {
		if noProxyHost == host {
			return nil // Skip proxy for this host
		}
	}

	// Determine proxy URL based on protocol
	var proxyURL *string
	if protocol == "wss" || protocol == "https" {
		proxyURL = runtimeObject.HttpsProxy
	} else {
		proxyURL = runtimeObject.HttpProxy
	}

	if proxyURL == nil || StringValue(proxyURL) == "" {
		return nil // No proxy configured
	}

	proxyProtocol := protocol
	if protocol == "wss" {
		proxyProtocol = "https"
	} else if protocol == "ws" {
		proxyProtocol = "http"
	}

	httpProxy, err := getHttpProxy(proxyProtocol, host, runtimeObject)
	if err != nil {
		return fmt.Errorf("failed to get HTTP proxy: %w", err)
	}

	if httpProxy == nil {
		return nil // No proxy needed
	}

	// Set proxy function (gorilla/websocket will handle the connection)
	dialer.Proxy = http.ProxyURL(httpProxy)

	// Add Proxy-Authorization header if proxy requires authentication
	if httpProxy.User != nil {
		password, _ := httpProxy.User.Password()
		auth := httpProxy.User.Username() + ":" + password
		basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		if request.Headers == nil {
			request.Headers = make(map[string]*string)
		}
		request.Headers["Proxy-Authorization"] = String(basic)
	}

	return nil
}

// configureNetDialer configures network dialer for direct connections
func (c *DefaultWebSocketClient) configureNetDialer(
	dialer *websocket.Dialer,
	runtimeObject *RuntimeObject,
	connectTimeout time.Duration,
) {
	// Only configure if NetDialContext is not already set (e.g., by SOCKS5)
	if dialer.NetDialContext != nil {
		return
	}

	localAddr := getLocalAddr(StringValue(runtimeObject.LocalAddr))
	dialer.NetDialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		d := &net.Dialer{
			Timeout:   connectTimeout,
			DualStack: true,
		}
		if localAddr != nil {
			d.LocalAddr = localAddr
		}
		return d.DialContext(ctx, network, addr)
	}
}

func (c *DefaultWebSocketClient) startMessageHandlers() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.readMessages()
	}()
}

func (c *DefaultWebSocketClient) readMessages() {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic in readMessages: %v", r)
			if c.session != nil {
				c.handler.HandleError(c.session, err)
			}
		}
	}()

	for {
		// Check stopChan first
		select {
		case <-c.stopChan:
			return
		default:
		}

		// Safely get conn reference with read lock
		c.connMu.RLock()
		conn := c.conn
		if conn == nil {
			c.connMu.RUnlock()
			return
		}
		// Set read deadline while holding lock
		if c.readTimeout > 0 {
			conn.SetReadDeadline(time.Now().Add(c.readTimeout))
		}
		c.connMu.RUnlock()

		// Read message without holding lock (blocking operation)
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			// Check if we should stop (connection might be closed)
			select {
			case <-c.stopChan:
				return
			default:
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				if c.session != nil {
					c.handler.HandleError(c.session, err) // might be already disconnected
				}
			}

			// Try to reconnect only if not stopping
			select {
			case <-c.stopChan:
				return
			default:
				if BoolValue(c.runtimeObject.WebSocketEnableReconnect) {
					if c.request != nil {
						go c.Reconnect()
					}
				}
			}
			return
		}

		msg := &WebSocketMessage{
			Type:      convertToWebSocketMessageType(messageType),
			Payload:   data,
			Headers:   make(map[string]string),
			Timestamp: time.Now(),
		}

		debugLog("[WebSocket] Received message: type=%d, size=%d bytes", messageType, len(data))

		if c.session == nil {
			debugLog("[WebSocket] Warning: session is nil, cannot handle message")
			continue
		}

		if err := c.handler.HandleRawMessage(c.session, msg); err != nil {
			debugLog("[WebSocket] HandleRawMessage error: %v", err)
			c.handler.HandleError(c.session, err)
		}
	}
}

func (c *DefaultWebSocketClient) startPingPong() {
	if c.pingInterval <= 0 {
		return
	}
	c.pingTicker = time.NewTicker(c.pingInterval)

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		for {
			select {
			case <-c.stopChan:
				return
			case <-c.pingTicker.C:
				c.connMu.RLock()
				conn := c.conn
				if conn == nil {
					c.connMu.RUnlock()
					return
				}
				deadline := time.Now().Add(c.writeTimeout)
				err := conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
				c.connMu.RUnlock()
				if err != nil {
					if c.session != nil {
						c.handler.HandleError(c.session, err)
					}
					return
				}

				select {
				case <-c.pongReceived:
					// Pong received
				case <-time.After(c.pongTimeout):
					// Pong timeout, try to reconnect
					if BoolValue(c.runtimeObject.WebSocketEnableReconnect) {
						go c.Reconnect()
					}
					return
				}
			}
		}
	}()
}

// This should be called immediately after connection is established to avoid race conditions
func (c *DefaultWebSocketClient) setupPongHandler() {
	c.connMu.RLock()
	conn := c.conn
	if conn == nil {
		c.connMu.RUnlock()
		return
	}
	conn.SetPongHandler(func(appData string) error {
		select {
		case c.pongReceived <- struct{}{}:
		default:
			// Channel is full, drop the pong signal (connection is already known to be alive)
		}
		return nil
	})
	c.connMu.RUnlock()
}

func (c *DefaultWebSocketClient) stopPingPong() {
	if c.pingTicker != nil {
		c.pingTicker.Stop()
	}
}

func convertToWebSocketMessageType(mt int) WebSocketMessageType {
	switch mt {
	case websocket.TextMessage:
		return WebSocketMessageTypeText
	case websocket.BinaryMessage:
		return WebSocketMessageTypeBinary
	case websocket.PingMessage:
		return WebSocketMessageTypePing
	case websocket.PongMessage:
		return WebSocketMessageTypePong
	case websocket.CloseMessage:
		return WebSocketMessageTypeClose
	default:
		return WebSocketMessageTypeBinary
	}
}

func generateSessionID() string {
	return fmt.Sprintf("ws-session-%d", time.Now().UnixNano())
}
