package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	flag "github.com/spf13/pflag"
	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/alinas"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs2"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

var (
	socketPath     string
	drivers        []string
	handleTimeout  time.Duration
	enableNftables bool
)

func main() {
	flag.StringVar(&socketPath, "socket", "/var/run/csi/mounter.sock", "socket path")
	flag.StringSliceVar(&drivers, "driver", nil, "drivers to enable (e.g. 'ossfs,alinas')")
	flag.DurationVar(&handleTimeout, "timeout", time.Second*30, "timeout for connection")
	flag.BoolVar(&enableNftables, "enable-nftables", false, "enable nftables rules to restrict mount proxy port access (default: false)")
	utils.AddKlogFlags(flag.CommandLine)
	utils.AddGoFlags(flag.CommandLine)
	flag.Parse()

	_ = os.Remove(socketPath)
	server.Init(drivers)

	listener, err := listen(socketPath)
	if err != nil {
		klog.ErrorS(err, "Failed to listen", "socket", socketPath)
		os.Exit(1)
	}
	defer os.Remove(socketPath)

	klog.InfoS("Listening on socket", "socket", socketPath)

	if enableNftables {
		klog.InfoS("Init nftables", "enabled", true)
		if err := setupNftables(); err != nil {
			klog.ErrorS(err, "Failed to setup nftables")
			os.Exit(1)
		}
	} else {
		klog.InfoS("Skipping nftables setup", "enabled", false)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		klog.InfoS("exiting", "signal", sig)
		// close socket
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			klog.ErrorS(err, "Failed to accept connection")
			continue
		}
		go handle(conn)
	}
	server.Terminate(drivers)
}

func handle(conn net.Conn) {
	defer conn.Close()
	err := server.Handle(conn, handleTimeout)
	if err != nil {
		klog.ErrorS(err, "Failed to handle")
	}
}

func listen(socketPath string) (net.Listener, error) {
	if len(socketPath) < 100 {
		return net.Listen("unix", socketPath)
	}

	// Need to change the current working directory to the temp volume base path,
	// because the socket absolute path is longer than 108(104 on OSX) characters,
	// which will cause "bind: invalid argument" errors.
	exPwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get the current directory: %w", err)
	}

	socketDir := filepath.Dir(socketPath)

	if err = os.Chdir(socketDir); err != nil {
		return nil, fmt.Errorf("failed to change directory to %q: %w", socketDir, err)
	}

	defer func() {
		if err := os.Chdir(exPwd); err != nil {
			klog.ErrorS(err, "Failed to change directory back")
		}
	}()

	return net.Listen("unix", filepath.Base(socketPath))
}

const (
	nftablesTableName  = "csi_mount_proxy"
	nftablesChainName  = "early_filter"
	nftablesTargetPort = 12049
)

// setupNftables configures nftables to block local access to the mount proxy port
// from processes outside the target cgroup. This prevents unauthorized access.
func setupNftables() error {
	conn, err := nftables.New()
	if err != nil {
		return fmt.Errorf("failed to create nftables connection: %w", err)
	}

	// Check if table already exists — if so, log existing chains/rules and skip
	tables, err := conn.ListTablesOfFamily(nftables.TableFamilyINet)
	if err != nil {
		return fmt.Errorf("failed to list nftables tables: %w", err)
	}
	for _, t := range tables {
		if t.Name == nftablesTableName {
			chains, _ := conn.ListChainsOfTableFamily(nftables.TableFamilyINet)
			for _, c := range chains {
				if c.Table.Name == nftablesTableName {
					rules, _ := conn.GetRules(t, c)
					klog.InfoS("nftables table already exists, skipping setup",
						"table", t.Name, "chain", c.Name, "ruleCount", len(rules))
				}
			}
			return nil
		}
	}

	// Create table
	table := conn.AddTable(&nftables.Table{
		Family: nftables.TableFamilyINet,
		Name:   nftablesTableName,
	})

	// Create chain: type filter, hook output, priority -150, policy accept
	policy := nftables.ChainPolicyAccept
	chain := conn.AddChain(&nftables.Chain{
		Name:     nftablesChainName,
		Table:    table,
		Type:     nftables.ChainTypeFilter,
		Hooknum:  nftables.ChainHookOutput,
		Priority: nftables.ChainPriorityRef(-150),
		Policy:   &policy,
	})

	// Build port bytes (big-endian)
	portBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(portBytes, nftablesTargetPort)

	// Add rule: oif lo tcp dport 12049 meta cgroup != 0 drop
	conn.AddRule(&nftables.Rule{
		Table: table,
		Chain: chain,
		Exprs: []expr.Any{
			// Match output interface "lo"
			&expr.Meta{Key: expr.MetaKeyOIFNAME, Register: 1},
			&expr.Cmp{Op: expr.CmpOpEq, Register: 1, Data: []byte("lo\x00")},
			// Match TCP protocol
			&expr.Meta{Key: expr.MetaKeyL4PROTO, Register: 1},
			&expr.Cmp{Op: expr.CmpOpEq, Register: 1, Data: []byte{unix.IPPROTO_TCP}},
			// Match destination port 12049
			&expr.Payload{DestRegister: 1, Base: expr.PayloadBaseTransportHeader, Offset: 2, Len: 2},
			&expr.Cmp{Op: expr.CmpOpEq, Register: 1, Data: portBytes},
			// Match cgroup != 0
			&expr.Meta{Key: expr.MetaKeyCGROUP, Register: 1},
			&expr.Cmp{Op: expr.CmpOpNeq, Register: 1, Data: []byte{0, 0, 0, 0}},
			// Verdict: drop
			&expr.Verdict{Kind: expr.VerdictDrop},
		},
	})

	if err := conn.Flush(); err != nil {
		return fmt.Errorf("failed to apply nftables rules: %w", err)
	}
	return nil
}
