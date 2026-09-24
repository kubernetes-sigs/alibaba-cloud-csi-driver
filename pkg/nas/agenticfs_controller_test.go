//go:build !windows

package nas

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
)

const (
	testAgenticFsFilesystemID   = "0025g9v3dzw1ozfl3w0"
	testAgenticFsAgenticSpaceID = "as-test-space-id"
	testAgenticFsAccessPointID  = "ap-test-ap-id"
	testAgenticFsAPDomain       = "ap-test-ap-id.0025g9v3dzw1ozfl3w0.cn-hangzhou.nas.aliyuncs.com"
	testAgenticFsPVName         = "pvc-00000000-0000-0000-0000-000000000000"
	testAgenticFsCNFSName       = "agenticfs-cnfs"

	// Differs from the CNFS fixture's filesystem, so a test can tell which of the two sources was used.
	testAgenticFsDirectFilesystemID = "0025directfsid0000001"

	// The zone of the AgenticSpace and the VPC/vSwitch of its accesspoint.
	testAgenticFsZoneID    = "cn-hangzhou-k"
	testAgenticFsVpcID     = "vpc-test-001"
	testAgenticFsVSwitchID = "vsw-test-001"

	// fakeStatusNotFound makes DescribeAccesspoint answer "the accesspoint does not exist", as the real API
	fakeStatusNotFound = "NotFound"
)

// The generated MockNasClientV2Interface exposes no EXPECT(), so a hand-written fake is used.
type fakeNasClientV2 struct {
	createAgenticSpaceReqs []*sdk.CreateAgenticSpaceRequest
	createAccessPointReqs  []*sdk.CreateAccessPointRequest
	deleteAgenticSpaceReqs []*sdk.DeleteAgenticSpaceRequest
	getAgenticSpaceReqs    []*sdk.GetAgenticSpaceRequest
	setQuotaReqs           []*sdk.SetAgenticSpaceQuotaRequest
	listAccessPointsReqs   []*sdk.ListAccessPointsRequest
	deleteAccessPointIDs   []string
	callOrder              []string
	describeCalls          int

	createAgenticSpaceResp *sdk.CreateAgenticSpaceResponse
	createAgenticSpaceErr  error
	createAccessPointResp  *sdk.CreateAccessPointResponse
	createAccessPointErr   error
	deleteAccessPointErr   error
	// Called synchronously; tests running CreateVolume in a goroutine must join it before reading calls.
	deleteAccessPointHook func(ctx context.Context, filesystemId, accessPointId string) error
	deleteAgenticSpaceErr error
	getAgenticSpaceResp   *sdk.GetAgenticSpaceResponse
	getAgenticSpaceErr    error
	setQuotaErr           error
	listAccessPointsErr   error
	// listNilBody makes EVERY ListAccesspoints answer with Body == nil; use a nil entry in listPages for one page.
	listNilBody bool
	// The resp == nil half of the malformed-response guard that listNilBody (nil Body) cannot reach.
	listNilResp bool

	// The last entry repeats once exhausted, so a page that keeps returning a NextToken drives the loop into its cap.
	listPages []*sdk.ListAccessPointsResponseBody
	listIdx   int

	// Read-after-write: an accesspoint returned by CreateAccesspoint shows up in every later list for that space.
	createdAccessPoints []*sdk.ListAccessPointsResponseBodyAccessPoints

	// The last entry is repeated once exhausted.
	apStatuses  []string
	describeIdx int
	// describeErrCalls limits describeErr to the first N calls (0 = every call).
	describeErr      error
	describeErrCalls int
	// apDomain overrides the reported DomainName; an empty domain has to be requested through apDomainSet.
	apDomain    string
	apDomainSet bool
}

func newFakeNasClientV2() *fakeNasClientV2 {
	return &fakeNasClientV2{
		createAgenticSpaceResp: &sdk.CreateAgenticSpaceResponse{
			Body: &sdk.CreateAgenticSpaceResponseBody{AgenticSpaceId: tea.String(testAgenticFsAgenticSpaceID)},
		},
		createAccessPointResp: &sdk.CreateAccessPointResponse{
			Body: &sdk.CreateAccessPointResponseBody{
				AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{
					AccessPointId:     tea.String(testAgenticFsAccessPointID),
					AccessPointDomain: tea.String(testAgenticFsAPDomain),
				},
			},
		},
		apStatuses: []string{accessPointStatusActive},
	}
}

// Accesspoints disappear on the first DescribeAccesspoint after DeleteAccesspoint.
func newDeleteFakeNasClientV2() *fakeNasClientV2 {
	fake := newFakeNasClientV2()
	fake.apStatuses = []string{fakeStatusNotFound}
	return fake
}

func (f *fakeNasClientV2) CreateAgenticSpace(_ context.Context, req *sdk.CreateAgenticSpaceRequest) (*sdk.CreateAgenticSpaceResponse, error) {
	f.createAgenticSpaceReqs = append(f.createAgenticSpaceReqs, req)
	f.callOrder = append(f.callOrder, "CreateAgenticSpace")
	return f.createAgenticSpaceResp, f.createAgenticSpaceErr
}

func (f *fakeNasClientV2) GetAgenticSpace(_ context.Context, req *sdk.GetAgenticSpaceRequest) (*sdk.GetAgenticSpaceResponse, error) {
	f.getAgenticSpaceReqs = append(f.getAgenticSpaceReqs, req)
	f.callOrder = append(f.callOrder, "GetAgenticSpace")
	return f.getAgenticSpaceResp, f.getAgenticSpaceErr
}

func (f *fakeNasClientV2) DeleteAgenticSpace(_ context.Context, req *sdk.DeleteAgenticSpaceRequest) (*sdk.DeleteAgenticSpaceResponse, error) {
	f.deleteAgenticSpaceReqs = append(f.deleteAgenticSpaceReqs, req)
	f.callOrder = append(f.callOrder, "DeleteAgenticSpace")
	return &sdk.DeleteAgenticSpaceResponse{Body: &sdk.DeleteAgenticSpaceResponseBody{}}, f.deleteAgenticSpaceErr
}

func (f *fakeNasClientV2) SetAgenticSpaceQuota(_ context.Context, req *sdk.SetAgenticSpaceQuotaRequest) (*sdk.SetAgenticSpaceQuotaResponse, error) {
	f.setQuotaReqs = append(f.setQuotaReqs, req)
	f.callOrder = append(f.callOrder, "SetAgenticSpaceQuota")
	return &sdk.SetAgenticSpaceQuotaResponse{Body: &sdk.SetAgenticSpaceQuotaResponseBody{}}, f.setQuotaErr
}

func (f *fakeNasClientV2) CreateAccesspoint(_ context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	f.createAccessPointReqs = append(f.createAccessPointReqs, req)
	f.callOrder = append(f.callOrder, "CreateAccesspoint")
	if f.createAccessPointErr == nil && f.createAccessPointResp != nil && f.createAccessPointResp.Body != nil &&
		f.createAccessPointResp.Body.AccessPoint != nil {
		ap := f.createAccessPointResp.Body.AccessPoint
		if id := tea.StringValue(ap.AccessPointId); id != "" {
			f.createdAccessPoints = append(f.createdAccessPoints,
				apItem(id, accessPointStatusActive, tea.StringValue(ap.AccessPointDomain)))
		}
	}
	return f.createAccessPointResp, f.createAccessPointErr
}

func (f *fakeNasClientV2) DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error {
	f.deleteAccessPointIDs = append(f.deleteAccessPointIDs, accessPointId)
	f.callOrder = append(f.callOrder, "DeleteAccesspoint")
	if f.deleteAccessPointHook != nil {
		return f.deleteAccessPointHook(ctx, filesystemId, accessPointId)
	}
	return f.deleteAccessPointErr
}

func (f *fakeNasClientV2) ListAccesspoints(_ context.Context, req *sdk.ListAccessPointsRequest) (*sdk.ListAccessPointsResponse, error) {
	f.listAccessPointsReqs = append(f.listAccessPointsReqs, req)
	f.callOrder = append(f.callOrder, "ListAccesspoints")
	if f.listAccessPointsErr != nil {
		return nil, f.listAccessPointsErr
	}
	if f.listNilResp {
		return nil, nil
	}
	if f.listNilBody {
		return &sdk.ListAccessPointsResponse{}, nil
	}
	page := &sdk.ListAccessPointsResponseBody{}
	if len(f.listPages) > 0 {
		idx := f.listIdx
		if idx >= len(f.listPages) {
			idx = len(f.listPages) - 1
		}
		page = f.listPages[idx]
		f.listIdx++
		// A nil entry models a page whose Body is absent; skip the merge below, which would dereference it.
		if page == nil {
			return &sdk.ListAccessPointsResponse{}, nil
		}
	}
	if len(f.createdAccessPoints) > 0 {
		// A new body is built so the scripted pages, which tests reuse by pointer, are never mutated.
		listed := make(map[string]bool, len(page.AccessPoints))
		for _, ap := range page.AccessPoints {
			listed[tea.StringValue(ap.AccessPointId)] = true
		}
		merged := append([]*sdk.ListAccessPointsResponseBodyAccessPoints{}, page.AccessPoints...)
		for _, ap := range f.createdAccessPoints {
			if !listed[tea.StringValue(ap.AccessPointId)] {
				merged = append(merged, ap)
			}
		}
		page = &sdk.ListAccessPointsResponseBody{AccessPoints: merged, NextToken: page.NextToken}
	}
	return &sdk.ListAccessPointsResponse{Body: page}, nil
}

func (f *fakeNasClientV2) DescribeAccesspoint(_ context.Context, _, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	f.describeCalls++
	f.callOrder = append(f.callOrder, "DescribeAccesspoint")
	if f.describeErr != nil && (f.describeErrCalls == 0 || f.describeCalls <= f.describeErrCalls) {
		return nil, f.describeErr
	}
	st := "Pending"
	if len(f.apStatuses) > 0 {
		idx := f.describeIdx
		if idx >= len(f.apStatuses) {
			idx = len(f.apStatuses) - 1
		}
		st = f.apStatuses[idx]
		f.describeIdx++
	}
	if st == fakeStatusNotFound {
		return nil, wrap.ErrorCode("NotFound")
	}
	domain := testAgenticFsAPDomain
	if f.apDomainSet {
		domain = f.apDomain
	}
	return &sdk.DescribeAccessPointResponse{
		Body: &sdk.DescribeAccessPointResponseBody{
			AccessPoint: &sdk.DescribeAccessPointResponseBodyAccessPoint{
				AccessPointId: tea.String(accessPointId),
				DomainName:    tea.String(domain),
				Status:        tea.String(st),
			},
		},
	}, nil
}

// Unused interface methods.
func (f *fakeNasClientV2) CreateDir(_ context.Context, _ *sdk.CreateDirRequest) error { return nil }
func (f *fakeNasClientV2) SetDirQuota(_ context.Context, _ *sdk.SetDirQuotaRequest) error {
	return nil
}
func (f *fakeNasClientV2) CancelDirQuota(_ context.Context, _ *sdk.CancelDirQuotaRequest) error {
	return nil
}
func (f *fakeNasClientV2) GetRecycleBinAttribute(_ context.Context, _ string) (*sdk.GetRecycleBinAttributeResponse, error) {
	return nil, nil
}

// Recorded in callOrder so a test can prove the controller never pre-validates the filesystem.
func (f *fakeNasClientV2) DescribeFileSystems(_ context.Context, _ string) (*sdk.DescribeFileSystemsResponse, error) {
	f.callOrder = append(f.callOrder, "DescribeFileSystems")
	return nil, nil
}

// Compile-time proof the fake covers the interface, so a new method breaks the build.
var _ interfaces.NasClientV2Interface = (*fakeNasClientV2)(nil)

type fakeNasClientFactory struct {
	v2  interfaces.NasClientV2Interface
	err error
}

func (f *fakeNasClientFactory) V1(string) (interfaces.NasV1Interface, error) { return nil, nil }
func (f *fakeNasClientFactory) V2(string) (interfaces.NasClientV2Interface, error) {
	return f.v2, f.err
}

// Returns a real apierrors NotFound; the shared fakeCNFSGetter returns a plain error, which maps to Internal.
type notFoundCNFSGetter struct{}

func (notFoundCNFSGetter) GetCNFS(_ context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
	return nil, apierrors.NewNotFound(cnfsv1beta1.GVR.GroupResource(), name)
}

// Fails the test if the controller consults a CNFS at all.
type forbiddenCNFSGetter struct {
	t *testing.T
}

func (g forbiddenCNFSGetter) GetCNFS(_ context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
	g.t.Errorf("GetCNFS(%q) must not be called: storageclass parameters.%s already names the filesystem",
		name, filesystemIDKey)
	return nil, apierrors.NewNotFound(cnfsv1beta1.GVR.GroupResource(), name)
}

// Exposes ErrorCode() so the classification helpers can be tested without a real tea.SDKError.
type fakeAliError struct {
	code string
}

func (e *fakeAliError) Error() string     { return "OpenAPI returned error: fake (" + e.code + ")" }
func (e *fakeAliError) ErrorCode() string { return e.code }
func (e *fakeAliError) Message() string   { return "fake" }

func aliErr(code string) error { return &fakeAliError{code: code} }

// Carries a message that differs from its code, pinning "classify on the code, never on the message".
type fakeAliErrorWithMessage struct {
	code    string
	message string
}

func (e *fakeAliErrorWithMessage) Error() string {
	return "OpenAPI returned error: " + e.code + " (" + e.message + ")"
}
func (e *fakeAliErrorWithMessage) ErrorCode() string { return e.code }
func (e *fakeAliErrorWithMessage) Message() string   { return e.message }

func agenticfsCNFS(name string) *cnfsv1beta1.ContainerNetworkFileSystem {
	return &cnfsv1beta1.ContainerNetworkFileSystem{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: cnfsv1beta1.ContainerNetworkFileSystemSpec{
			StorageType: cnfsSpecTypeAgenticFS,
		},
		Status: cnfsv1beta1.ContainerNetworkFileSystemStatus{
			Status: cnfsv1beta1.StatusAvailable,
			FsAttributes: cnfsv1beta1.FsAttributes{
				FilesystemID:   testAgenticFsFilesystemID,
				FilesystemType: cloud.FilesystemTypeStandard, // "standard": must NOT be used to detect AgenticFS
				StorageType:    cloud.StorageTypeAgentic,     // "Agentic"
				ProtocolType:   "NFS",
				RegionID:       "cn-hangzhou",
				// StorageClass parameters: a controller that still read them from the CNFS would fail every case below.
				Server: "", // AgenticFS has no server; controller must not read it
			},
		},
	}
}

func newTestAgenticfsController(t *testing.T, fake *fakeNasClientV2, getter cnfsv1beta1.CNFSGetter) *agenticfsController {
	t.Helper()
	c, err := newAgenticfsController(&internal.ControllerConfig{
		Metadata:         testMetadata,
		NasClientFactory: &fakeNasClientFactory{v2: fake},
		CNFSGetter:       getter,
	})
	require.NoError(t, err)
	ctrl, ok := c.(*agenticfsController)
	require.True(t, ok)
	ctrl.apPollInterval = time.Millisecond
	ctrl.apPollTimeout = 100 * time.Millisecond
	return ctrl
}

func newAgenticfsCtrl(t *testing.T, fake *fakeNasClientV2) *agenticfsController {
	t.Helper()
	return newTestAgenticfsController(t, fake, newFakeCNFSGetter(agenticfsCNFS(testAgenticFsCNFSName)))
}

func newAgenticfsDirectCtrl(t *testing.T, fake *fakeNasClientV2) *agenticfsController {
	t.Helper()
	return newTestAgenticfsController(t, fake, forbiddenCNFSGetter{t})
}

func agenticfsCreateReq(pvName string, requiredBytes int64, params map[string]string) *csi.CreateVolumeRequest {
	if params == nil {
		params = map[string]string{}
	}
	params[paramContainerNetworkFileSystem] = testAgenticFsCNFSName
	params["volumeAs"] = agenticFsVolumeAs
	// Set only when absent, so a case can pass an explicit "" to exercise missing-parameter validation.
	for k, v := range map[string]string{
		ZoneID:    testAgenticFsZoneID,
		VpcID:     testAgenticFsVpcID,
		VSwitchID: testAgenticFsVSwitchID,
	} {
		if _, ok := params[k]; !ok {
			params[k] = v
		}
	}
	return &csi.CreateVolumeRequest{
		Name:          pvName,
		CapacityRange: &csi.CapacityRange{RequiredBytes: requiredBytes},
		Parameters:    params,
	}
}

func agenticfsDirectCreateReq(pvName string, requiredBytes int64, params map[string]string) *csi.CreateVolumeRequest {
	req := agenticfsCreateReq(pvName, requiredBytes, params)
	delete(req.Parameters, paramContainerNetworkFileSystem)
	req.Parameters[filesystemIDKey] = testAgenticFsDirectFilesystemID
	return req
}

func agenticfsCreateReqRange(pvName string, cr *csi.CapacityRange, params map[string]string) *csi.CreateVolumeRequest {
	req := agenticfsCreateReq(pvName, cr.GetRequiredBytes(), params)
	req.CapacityRange = cr
	return req
}

func apItem(id, st, domain string) *sdk.ListAccessPointsResponseBodyAccessPoints {
	return &sdk.ListAccessPointsResponseBodyAccessPoints{
		AccessPointId:  tea.String(id),
		AgenticSpaceId: tea.String(testAgenticFsAgenticSpaceID),
		DomainName:     tea.String(domain),
		Status:         tea.String(st),
	}
}

func apPage(aps ...*sdk.ListAccessPointsResponseBodyAccessPoints) *sdk.ListAccessPointsResponseBody {
	return &sdk.ListAccessPointsResponseBody{AccessPoints: aps}
}

func activeApPage() *sdk.ListAccessPointsResponseBody {
	return apPage(apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain))
}

func countCalls(fake *fakeNasClientV2, name string) int {
	n := 0
	for _, c := range fake.callOrder {
		if c == name {
			n++
		}
	}
	return n
}

// fileSystemPath is the only key a reaper can use for a space whose ID was never read back; "cause" is its own key.
var orphanLogFields = []string{"fileSystemId", "agenticSpaceId", "fileSystemPath", "accesspointId", "region", "volumeHandle", "reason", "cause"}

// BufferLogs must be turned on: ktesting leaves it off, which would make every assertion vacuous.
func newLogCapture(t *testing.T) (klog.Logger, context.Context) {
	t.Helper()
	logger := ktesting.NewLogger(t, ktesting.NewConfig(ktesting.BufferLogs(true)))
	return logger, klog.NewContext(context.Background(), logger)
}

// ktesting's sink implements Underlier; the assertion fails loudly rather than on an empty string if that changes.
func logText(logger klog.Logger) string {
	return logger.GetSink().(ktesting.Underlier).GetBuffer().String()
}

// LINE-LEVEL: scanning the whole buffer would pass if the fields leaked onto an Info line.
func assertOrphanLog(t *testing.T, logs string) {
	t.Helper()
	var orphanLine string
	for _, line := range strings.Split(logs, "\n") {
		if strings.Contains(line, orphanLogPrefix) {
			orphanLine = line
			break
		}
	}
	require.NotEmpty(t, orphanLine, "a terminal provisioning failure must be reported for reconciliation")
	assert.Contains(t, orphanLine, "ERROR", "the orphan line must be Error level, not Info")
	for _, field := range orphanLogFields {
		assert.Contains(t, orphanLine, field, "every orphan line carries all contract fields on the SAME line")
	}
	assert.Contains(t, orphanLine, "/"+testAgenticFsPVName, "fileSystemPath value must never be dropped")
}

// A reaper counts LINES; counting occurrences in the buffer would not distinguish one line from two.
func countLogLines(logs, prefix string) int {
	n := 0
	for _, line := range strings.Split(logs, "\n") {
		if strings.Contains(line, prefix) {
			n++
		}
	}
	return n
}

func assertExactlyOneFailureReport(t *testing.T, logs string) {
	t.Helper()
	assert.Equal(t, 1, countLogLines(logs, orphanLogPrefix)+countLogLines(logs, retainedForRetryLogPrefix))
	assert.NotContains(t, logs, "agenticfs-orphan-resolved", "CreateVolume never performs cleanup")
}

func TestAgenticfsVolumeAs(t *testing.T) {
	ctrl := newAgenticfsCtrl(t, newFakeNasClientV2())
	assert.Equal(t, "Agentic", ctrl.VolumeAs())
}

func TestAgenticfsCreateVolumeSuccess(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	require.NotNil(t, resp)

	vc := resp.Volume.VolumeContext
	// Must use the package-wide filesystemIDKey: getNASIDFromMapOrServer would otherwise split the AP domain.
	assert.Equal(t, testAgenticFsAPDomain, vc[vcKeyServer])
	assert.Equal(t, "/", vc[vcKeyPath])
	assert.Equal(t, mountProtocolAlinas, vc[vcKeyMountProtocol])
	assert.Equal(t, defaultAgenticFsMountOptions, vc[vcKeyOptions])
	assert.Equal(t, testAgenticFsFilesystemID, vc[filesystemIDKey])
	assert.Equal(t, testAgenticFsAccessPointID, vc[vcKeyAccesspointId])
	assert.Equal(t, testAgenticFsAgenticSpaceID, vc[vcKeyAgenticSpaceId])
	assert.Equal(t, int64(20*GiB), resp.Volume.CapacityBytes)
	assert.Equal(t, testAgenticFsPVName, resp.Volume.VolumeId)

	assert.NotContains(t, vc, "accesspoint")
	assert.NotContains(t, vc, "containerNetworkFileSystem")
	assert.NotContains(t, vc, vcKeyFilesystemIdLegacy)

	// Phase-1 guardrail: the jwtauth keys are only unlocked after the phase-2 node-side changes ship.
	assert.NotContains(t, vc, "authType")
	assert.NotContains(t, vc, "credentialProviderName")
	assert.NotContains(t, vc, "sandboxId")

	require.Len(t, fake.createAgenticSpaceReqs, 1)
	spaceReq := fake.createAgenticSpaceReqs[0]
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(spaceReq.FileSystemId))
	assert.Equal(t, "/"+testAgenticFsPVName+"/", tea.StringValue(spaceReq.FileSystemPath))
	assert.Equal(t, testAgenticFsZoneID, tea.StringValue(spaceReq.Azone))
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(spaceReq.ClientToken))
	require.NotNil(t, spaceReq.Quota)
	assert.Equal(t, int64(20*GiB), tea.Int64Value(spaceReq.Quota.SizeLimit))
	assert.Equal(t, defaultAgenticSpaceFileCountLimit, tea.Int64Value(spaceReq.Quota.FileCountLimit))
	// CreateAgenticSpaceRequest has no RegionId field; region rides in the client's GlobalParameters.

	// The idempotence key: it must ask for THIS space's accesspoints, server-side, and page through them.
	require.Len(t, fake.listAccessPointsReqs, 1)
	listReq := fake.listAccessPointsReqs[0]
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(listReq.FileSystemId))
	assert.Equal(t, apListMaxResults, tea.Int32Value(listReq.MaxResults))
	assert.Empty(t, tea.StringValue(listReq.NextToken))
	require.Len(t, listReq.Filters, 1)
	assert.Equal(t, apListFilterAgenticSpaceId, tea.StringValue(listReq.Filters[0].Name))
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(listReq.Filters[0].Value))

	require.Len(t, fake.createAccessPointReqs, 1)
	apReq := fake.createAccessPointReqs[0]
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(apReq.FileSystemId))
	assert.NotEmpty(t, tea.StringValue(apReq.AgenticSpaceId))
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(apReq.AgenticSpaceId))
	assert.Equal(t, testAgenticFsVSwitchID, tea.StringValue(apReq.VswId))
	assert.Equal(t, testAgenticFsVpcID, tea.StringValue(apReq.VpcId))
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(apReq.AccessPointName))
	assert.True(t, tea.BoolValue(apReq.EnabledRam))
	// The volume tag is observability only; it must still be written.
	require.Len(t, apReq.Tag, 1)
	assert.Equal(t, apTagKeyVolumeID, tea.StringValue(apReq.Tag[0].Key))
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(apReq.Tag[0].Value))
	// AgenticFS-unsupported fields must all stay nil.
	assert.Nil(t, apReq.AccessGroup)
	assert.Nil(t, apReq.RootDirectory)
	assert.Nil(t, apReq.OwnerUserId)
	assert.Nil(t, apReq.OwnerGroupId)
	assert.Nil(t, apReq.Permission)
	assert.Nil(t, apReq.PosixUserId)
	assert.Nil(t, apReq.PosixGroupId)
	assert.Nil(t, apReq.PosixSecondaryGroupIds)

	assert.Equal(t,
		[]string{"CreateAgenticSpace", "ListAccesspoints", "CreateAccesspoint", "DescribeAccesspoint"},
		fake.callOrder)
}

func TestAgenticfsCreateVolumeDirectFilesystemID(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsDirectCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsDirectCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, testAgenticFsDirectFilesystemID, resp.Volume.VolumeContext[filesystemIDKey])
	require.Len(t, fake.createAgenticSpaceReqs, 1)
	assert.Equal(t, testAgenticFsDirectFilesystemID, tea.StringValue(fake.createAgenticSpaceReqs[0].FileSystemId))
	assert.Equal(t, "/"+testAgenticFsPVName+"/", tea.StringValue(fake.createAgenticSpaceReqs[0].FileSystemPath))
	require.Len(t, fake.listAccessPointsReqs, 1)
	assert.Equal(t, testAgenticFsDirectFilesystemID, tea.StringValue(fake.listAccessPointsReqs[0].FileSystemId))
	require.Len(t, fake.createAccessPointReqs, 1)
	assert.Equal(t, testAgenticFsDirectFilesystemID, tea.StringValue(fake.createAccessPointReqs[0].FileSystemId))

	// The filesystem type is deliberately not pre-checked, so DescribeFileSystems never appears.
	assert.Equal(t,
		[]string{"CreateAgenticSpace", "ListAccesspoints", "CreateAccesspoint", "DescribeAccesspoint"},
		fake.callOrder)
}

func TestAgenticfsCreateVolumeFilesystemIDParameterWins(t *testing.T) {
	// fileSystemId must win, so a StorageClass can be migrated by adding it before removing the CNFS parameter.
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsDirectCtrl(t, fake)

	req := agenticfsCreateReq(testAgenticFsPVName, 20*GiB, map[string]string{
		filesystemIDKey: testAgenticFsDirectFilesystemID,
	})
	resp, err := ctrl.CreateVolume(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, testAgenticFsDirectFilesystemID, resp.Volume.VolumeContext[filesystemIDKey])
	require.Len(t, fake.createAgenticSpaceReqs, 1)
	assert.Equal(t, testAgenticFsDirectFilesystemID,
		tea.StringValue(fake.createAgenticSpaceReqs[0].FileSystemId),
		"fileSystemId must take precedence over the CNFS-owned filesystem")
}

func TestAgenticfsCreateVolumeNoFilesystemSource(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	req := agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil)
	delete(req.Parameters, paramContainerNetworkFileSystem)

	_, err := ctrl.CreateVolume(context.Background(), req)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	// The message must name both accepted keys, or the operator cannot tell which one to add.
	msg := status.Convert(err).Message()
	assert.Contains(t, msg, filesystemIDKey)
	assert.Contains(t, msg, paramContainerNetworkFileSystem)
	assert.Empty(t, fake.callOrder, "no cloud call may happen before the parameters are validated")
}

func TestAgenticfsCreateVolumeCNFSNotFound(t *testing.T) {
	ctrl := newTestAgenticfsController(t, newFakeNasClientV2(), notFoundCNFSGetter{})
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestAgenticfsCreateVolumeCNFSGetError(t *testing.T) {
	// newFakeCNFSGetter with no object returns a plain (non-NotFound) error.
	ctrl := newTestAgenticfsController(t, newFakeNasClientV2(), newFakeCNFSGetter())
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
}

func TestAgenticfsCreateVolumeTypeValidation(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(cnfs *cnfsv1beta1.ContainerNetworkFileSystem)
	}{
		{
			name:   "specTypeIsNotAgenticfs",
			mutate: func(c *cnfsv1beta1.ContainerNetworkFileSystem) { c.Spec.StorageType = "nas" },
		},
		{
			name:   "fsAttributesStorageTypeIsNotAgentic",
			mutate: func(c *cnfsv1beta1.ContainerNetworkFileSystem) { c.Status.FsAttributes.StorageType = "" },
		},
		{
			// FilesystemType is "standard" for plain NAS too, so it is not a valid AgenticFS discriminator.
			name: "filesystemTypeStandardButStorageTypeWrongMustStillFail",
			mutate: func(c *cnfsv1beta1.ContainerNetworkFileSystem) {
				c.Status.FsAttributes.FilesystemType = cloud.FilesystemTypeStandard
				c.Spec.StorageType = "cpfs"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnfs := agenticfsCNFS(testAgenticFsCNFSName)
			tt.mutate(cnfs)
			ctrl := newTestAgenticfsController(t, newFakeNasClientV2(), newFakeCNFSGetter(cnfs))
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
		})
	}
}

func TestAgenticfsCreateVolumeCNFSFilesystemIDEmpty(t *testing.T) {
	// An empty status.fsAttributes ID means the CNFS controller has not finished creating the filesystem.
	cnfs := agenticfsCNFS(testAgenticFsCNFSName)
	cnfs.Status.FsAttributes.FilesystemID = ""
	fake := newFakeNasClientV2()
	ctrl := newTestAgenticfsController(t, fake, newFakeCNFSGetter(cnfs))

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Empty(t, fake.callOrder, "no billable resource may be created from an incomplete CNFS")
}

func TestAgenticfsCreateVolumeMissingPlacementParameters(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{"missingZoneId", ZoneID},
		{"missingVpcId", VpcID},
		{"missingVSwitchId", VSwitchID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)

			// An explicit "" overrides the placement default agenticfsCreateReq injects.
			_, err := ctrl.CreateVolume(context.Background(),
				agenticfsCreateReq(testAgenticFsPVName, 20*GiB, map[string]string{tt.key: ""}))
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Empty(t, fake.callOrder,
				"no billable resource may be created from an incomplete StorageClass")
		})
	}
}

// agenticSpaceSizeLimit is an upper bound that is actually enforced, not a fallback dead in practice.
func TestAgenticfsCreateVolumeQuotaValidation(t *testing.T) {
	tests := []struct {
		name          string
		requiredBytes int64
		params        map[string]string
		wantCode      codes.Code
	}{
		{"sizeBelow10GiBRejected", 5 * GiB, nil, codes.InvalidArgument},
		{"fileCountBelow10000Rejected", 20 * GiB, map[string]string{paramAgenticSpaceFileCountLimit: "5000"}, codes.InvalidArgument},
		// The OpenAPI maximum is 1e8, not the 1e9 an earlier draft documented.
		{"fileCountAbove1e8Rejected", 20 * GiB, map[string]string{paramAgenticSpaceFileCountLimit: "100000001"}, codes.InvalidArgument},
		{"fileCountOf1e9Rejected", 20 * GiB, map[string]string{paramAgenticSpaceFileCountLimit: "1000000001"}, codes.InvalidArgument},
		{"invalidSizeLimitParamRejected", 0, map[string]string{paramAgenticSpaceSizeLimit: "not-a-quantity"}, codes.InvalidArgument},
		{"invalidFileCountParamRejected", 20 * GiB, map[string]string{paramAgenticSpaceFileCountLimit: "many"}, codes.InvalidArgument},
		{"noCapacityAtAllRejected", 0, nil, codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, tt.requiredBytes, tt.params))
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.callOrder, "an invalid StorageClass must not create a billable resource")
		})
	}
}

func TestAgenticfsCreateVolumeSizeLimitIsACap(t *testing.T) {
	tests := []struct {
		name          string
		requiredBytes int64
		capParam      string
		wantErr       bool
		wantSize      int64
	}{
		{"requestBelowTheCapIsHonored", 20 * GiB, "100Gi", false, 20 * GiB},
		{"requestEqualToTheCapIsHonored", 100 * GiB, "100Gi", false, 100 * GiB},
		// M8: this is the case the fallback-only reading silently allowed.
		{"requestAboveTheCapIsRejected", 100 * GiB, "50Gi", true, 0},
		{"capUsedWhenNoCapacityIsRequested", 0, "10Gi", false, 10 * GiB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)
			resp, err := ctrl.CreateVolume(context.Background(),
				agenticfsCreateReq(testAgenticFsPVName, tt.requiredBytes,
					map[string]string{paramAgenticSpaceSizeLimit: tt.capParam}))
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, codes.InvalidArgument, status.Code(err))
				assert.Empty(t, fake.createAgenticSpaceReqs)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantSize, resp.Volume.CapacityBytes)
			require.Len(t, fake.createAgenticSpaceReqs, 1)
			assert.Equal(t, tt.wantSize, tea.Int64Value(fake.createAgenticSpaceReqs[0].Quota.SizeLimit))
		})
	}
}

func TestAgenticfsCreateVolumeSizeLimitGiBAlignment(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 10*GiB+1, nil))
	require.NoError(t, err)
	require.Len(t, fake.createAgenticSpaceReqs, 1)
	assert.Equal(t, int64(11*GiB), tea.Int64Value(fake.createAgenticSpaceReqs[0].Quota.SizeLimit))
	assert.Equal(t, int64(11*GiB), resp.Volume.CapacityBytes)
}

func TestAgenticfsCreateVolumeLimitBytes(t *testing.T) {
	tests := []struct {
		name     string
		cr       *csi.CapacityRange
		wantCode codes.Code
	}{
		{
			// Rounding up to the 1 GiB boundary pushes the capacity over limit_bytes, which CSI forbids returning.
			name:     "roundingOverLimitBytesIsRejected",
			cr:       &csi.CapacityRange{RequiredBytes: 20*GiB + 1, LimitBytes: 20 * GiB},
			wantCode: codes.OutOfRange,
		},
		{
			name:     "limitBytesAboveRequestIsFine",
			cr:       &csi.CapacityRange{RequiredBytes: 20 * GiB, LimitBytes: 30 * GiB},
			wantCode: codes.OK,
		},
		{
			name:     "capacityThatOverflowsRoundingIsRejected",
			cr:       &csi.CapacityRange{RequiredBytes: math.MaxInt64},
			wantCode: codes.OutOfRange,
		},
		{
			// The cloud would answer with an opaque retryable error; the driver rejects it as permanent instead.
			name:     "capacityAboveAgenticSpaceMaximumIsRejected",
			cr:       &csi.CapacityRange{RequiredBytes: maxAgenticSpaceSizeLimit + GiB},
			wantCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)
			resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReqRange(testAgenticFsPVName, tt.cr, nil))
			if tt.wantCode == codes.OK {
				require.NoError(t, err)
				assert.Equal(t, tt.cr.RequiredBytes, resp.Volume.CapacityBytes)
				return
			}
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.createAgenticSpaceReqs)
		})
	}
}

// The ordinary lifecycle is one Space with at most one AccessPoint. Retrying
// an unavailable or pending accesspoint must not create a replacement.
func TestAgenticfsCreateVolumeSingleAccessPointLifecycle(t *testing.T) {
	for _, tt := range []struct {
		name          string
		state         string
		wantCode      codes.Code
		wantCreates   int
		wantDescribes int
	}{
		{"absent", "", codes.OK, 1, 1},
		{"active", accessPointStatusActive, codes.OK, 0, 1},
		{"pending", "Pending", codes.OK, 0, 2},
		{"inactive", accessPointStatusInactive, codes.Aborted, 0, 0},
		{"deleting", accessPointStatusDeleting, codes.Aborted, 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			if tt.state != "" {
				fake.listPages = []*sdk.ListAccessPointsResponseBody{
					apPage(apItem(testAgenticFsAccessPointID, tt.state, testAgenticFsAPDomain)),
				}
			}
			if tt.state == "Pending" {
				fake.apStatuses = []string{"Pending", accessPointStatusActive}
			}
			ctrl := newAgenticfsCtrl(t, fake)
			resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Equal(t, tt.wantCode, status.Code(err), "%v", err)
			if tt.wantCode == codes.OK {
				require.NotNil(t, resp)
				assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
			} else {
				assert.Nil(t, resp)
			}
			assert.Len(t, fake.createAccessPointReqs, tt.wantCreates)
			assert.Equal(t, tt.wantDescribes, fake.describeCalls)
			assert.Empty(t, fake.deleteAccessPointIDs)
			assert.Empty(t, fake.deleteAgenticSpaceReqs)
		})
	}
}

func TestAgenticfsCreateVolumeReusesExistingAccessPoint(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{activeApPage()}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Empty(t, fake.createAccessPointReqs, "an existing accesspoint must be reused, not duplicated")
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Equal(t, []string{"CreateAgenticSpace", "ListAccesspoints", "DescribeAccesspoint"}, fake.callOrder)
}

func TestAgenticfsCreateVolumeRetryCreatesOneAccessPoint(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		{}, // first attempt: the space has no accesspoint yet
		activeApPage(),
	}
	ctrl := newAgenticfsCtrl(t, fake)

	first, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	second, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)

	assert.Len(t, fake.createAccessPointReqs, 1, "the retry must not stack a second accesspoint on the space")
	assert.Len(t, fake.createAgenticSpaceReqs, 2, "CreateAgenticSpace is retried but is idempotent through ClientToken")
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(fake.createAgenticSpaceReqs[0].ClientToken))
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(fake.createAgenticSpaceReqs[1].ClientToken))
	assert.Equal(t, first.Volume.VolumeContext, second.Volume.VolumeContext)
}

func TestAgenticfsCreateVolumePrefersAnActiveAccessPoint(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(
		apItem("ap-pending", "Pending", "pending.example.com"),
		apItem("ap-empty-id", accessPointStatusActive, ""),
		apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain),
	)}
	// An entry with an empty ID must be skipped instead of being picked first.
	fake.listPages[0].AccessPoints[1].AccessPointId = nil
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
}

// Creating a new accesspoint while a previous one is still tearing down would leave two behind.
func TestAgenticfsCreateVolumeDeletingAccessPointAborts(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		apPage(apItem(testAgenticFsAccessPointID, accessPointStatusDeleting, testAgenticFsAPDomain)),
	}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Empty(t, fake.createAccessPointReqs)
	// Retryable failures retain resources for the next request.
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

// Mounting another space's accesspoint is a cross-volume data-plane exposure, not merely a leak.
func TestAgenticfsCreateVolumeListFilterMismatchAborts(t *testing.T) {
	fake := newFakeNasClientV2()
	other := apItem("ap-other", accessPointStatusActive, "other.example.com")
	other.AgenticSpaceId = tea.String("as-someone-elses-space")
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(other)}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Contains(t, err.Error(), "nas:ListAccesspoints:")
	assert.Contains(t, err.Error(), "did not take effect")
	// The only diagnostic an operator gets here, so it must name the accesspoint AND both space ids.
	assert.Contains(t, err.Error(), foreignAccessPointID)
	assert.Contains(t, err.Error(), "as-someone-elses-space", "the space the accesspoint really belongs to")
	assert.Contains(t, err.Error(), testAgenticFsAgenticSpaceID, "the space that was requested")
	// No new accesspoint is created and the existing resources are retained.
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

const foreignAccessPointID = "ap-other"

// Even when the first entry is Active, validate the whole listing before reusing it.
func TestAgenticfsCreateVolumeValidatesOwnershipBeforeSelectingActiveAccessPoint(t *testing.T) {
	fake := newFakeNasClientV2()
	foreign := apItem("ap-foreign", accessPointStatusActive, "foreign.example.com")
	foreign.AgenticSpaceId = tea.String("as-other-space")
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(
		apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain), foreign,
	)}
	ctrl := newAgenticfsCtrl(t, fake)
	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Equal(t, codes.Aborted, status.Code(err))
	assert.Nil(t, resp)
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Zero(t, fake.describeCalls)
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

// AgenticSpaceId is optional on the wire; warn-and-skip was rejected, an unproven accesspoint would be reused.
func TestAgenticfsCreateVolumeListMissingAgenticSpaceIdAborts(t *testing.T) {
	newFake := func() *fakeNasClientV2 {
		fake := newFakeNasClientV2()
		noOwner := apItem("ap-no-owner", accessPointStatusActive, "no-owner.example.com")
		noOwner.AgenticSpaceId = nil // the field is absent, not mismatched
		fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(noOwner)}
		return fake
	}

	t.Run("createVolume", func(t *testing.T) {
		fake := newFake()
		ctrl := newAgenticfsCtrl(t, fake)
		_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		require.Error(t, err)
		assert.Equal(t, codes.Aborted, status.Code(err), "same code as the cross-volume case, only the message differs")
		assert.Contains(t, err.Error(), "without an AgenticSpaceId")
		assert.Contains(t, err.Error(), "does not return the field", "the operator must be pointed at an API-version gap")
		assert.Contains(t, err.Error(), "ap-no-owner")
		assert.Contains(t, err.Error(), testAgenticFsAgenticSpaceID)
		assert.NotContains(t, err.Error(), "did not take effect",
			"the cross-volume wording must not be used when the field is merely absent")
		assert.Empty(t, fake.createAccessPointReqs)
		assert.Empty(t, fake.deleteAgenticSpaceReqs)
		assert.Empty(t, fake.deleteAccessPointIDs)
	})

	t.Run("deleteVolumeIsWedgedTheSameWay", func(t *testing.T) {
		fake := newFake()
		fake.apStatuses = []string{fakeStatusNotFound}
		ctrl := newAgenticfsCtrl(t, fake)
		_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
		require.Error(t, err)
		assert.Equal(t, codes.Aborted, status.Code(err))
		assert.Contains(t, err.Error(), "without an AgenticSpaceId")
		assert.Empty(t, fake.deleteAccessPointIDs, "an unproven accesspoint must not be deleted")
		assert.Empty(t, fake.deleteAgenticSpaceReqs)
	})
}

func TestAgenticfsCreateVolumeInactiveAccessPointIsUnusable(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		apPage(apItem(testAgenticFsAccessPointID, accessPointStatusInactive, testAgenticFsAPDomain)),
	}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Contains(t, err.Error(), "nas:ListAccesspoints:")
	assert.Contains(t, err.Error(), "not usable yet")
	assert.Contains(t, err.Error(), testAgenticFsAccessPointID+"("+accessPointStatusInactive+")")
	// A Pending PVC may never receive DeleteVolume: direct operators to verified cleanup.
	assert.Contains(t, err.Error(), "NAS console")
	assert.Contains(t, err.Error(), "examples/nas/agenticfs/README.md §4.1")
	assert.Contains(t, err.Error(), "deleting a Pending PVC does not guarantee cloud resource cleanup")
	assert.NotContains(t, err.Error(), "delete the PVC to release the agenticspace")
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

// Proves Inactive is filtered into the unusable set rather than aborting the whole discovery.
func TestAgenticfsCreateVolumeActiveAccessPointIsPreferredOverInactive(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(
		apItem("ap-inactive", accessPointStatusInactive, "inactive.example.com"),
		apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain),
	)}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
}

func TestAgenticfsCreateVolumeListPagination(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		{
			AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{apItem("ap-page1", "Pending", "page1.example.com")},
			NextToken:    tea.String("token-1"),
		},
		activeApPage(),
	}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	require.Len(t, fake.listAccessPointsReqs, 2, "NextToken must be followed to the end; reading one page would miss accesspoints")
	assert.Empty(t, tea.StringValue(fake.listAccessPointsReqs[0].NextToken))
	assert.Equal(t, "token-1", tea.StringValue(fake.listAccessPointsReqs[1].NextToken))
	// A one-page read would never have seen it and would have answered with a duplicate CreateAccesspoint.
	assert.Empty(t, fake.createAccessPointReqs)
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
}

func TestAgenticfsCreateVolumeListPaginationCapIsRetryable(t *testing.T) {
	fake := newFakeNasClientV2()
	// The loop must give up with a retryable error rather than spin forever or silently use page 1.
	fake.listPages = []*sdk.ListAccessPointsResponseBody{{NextToken: tea.String("forever")}}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Aborted, status.Code(err))
	// These calls come only from discovery reaching its page cap.
	assert.Len(t, fake.listAccessPointsReqs, apListMaxPages)
	msg := err.Error()
	assert.Contains(t, msg, fmt.Sprint(apListMaxPages), "the page count helps an operator tell a runaway listing apart")
	assert.Contains(t, msg, testAgenticFsFilesystemID)
	assert.Empty(t, fake.createAccessPointReqs)
}

// On page >= 2 an empty body means the API truncated a list it claimed exhaustive, which stacks a duplicate.
func TestAgenticfsCreateVolumeListSecondPageEmptyBodyAborts(t *testing.T) {
	newFake := func() *fakeNasClientV2 {
		fake := newFakeNasClientV2()
		fake.listPages = []*sdk.ListAccessPointsResponseBody{
			{
				AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{apItem("ap-page1", "Pending", "page1.example.com")},
				NextToken:    tea.String("token-1"),
			},
			nil, // page 2 comes back with no body at all
		}
		return fake
	}

	t.Run("createVolume", func(t *testing.T) {
		fake := newFake()
		ctrl := newAgenticfsCtrl(t, fake)
		_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		require.Error(t, err)
		assert.Equal(t, codes.Aborted, status.Code(err))
		assert.Contains(t, err.Error(), "page 2", "the page number tells an operator where the enumeration broke")
		assert.Contains(t, err.Error(), testAgenticFsAgenticSpaceID)
		assert.Contains(t, err.Error(), "refusing to act on a partial list")
		assert.Empty(t, fake.createAccessPointReqs, "no accesspoint may be created from a truncated list")
		assert.Empty(t, fake.deleteAgenticSpaceReqs)
		assert.Empty(t, fake.deleteAccessPointIDs)
	})

	t.Run("deleteVolume", func(t *testing.T) {
		fake := newFake()
		fake.apStatuses = []string{fakeStatusNotFound}
		ctrl := newAgenticfsCtrl(t, fake)
		_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
		require.Error(t, err)
		assert.Equal(t, codes.Aborted, status.Code(err))
		assert.Contains(t, err.Error(), "page 2")
		assert.Empty(t, fake.deleteAccessPointIDs, "nothing may be deleted from an incomplete listing")
		assert.Empty(t, fake.deleteAgenticSpaceReqs,
			"DeleteAgenticSpace would be refused anyway while a missed accesspoint survives")
	})
}

func TestAgenticfsCreateVolumeListError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"transientErrorIsRetryable", errors.New("connection reset"), codes.Internal},
		{"throttlingIsRetryable", aliErr("Throttling.User"), codes.Internal},
		{"permanentErrorIsTerminal", aliErr("InvalidParameter.Filter"), codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.listAccessPointsErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.createAccessPointReqs)
			// On the terminal case the failure happened during discovery, before any accesspoint existed.
			assert.Empty(t, fake.deleteAgenticSpaceReqs)
			assert.Empty(t, fake.deleteAccessPointIDs)
		})
	}
}

func TestAgenticfsCreateVolumeCreateAgenticSpaceError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"transientErrorIsRetryable", errors.New("dial tcp: i/o timeout"), codes.Internal},
		{"invalidPathIsTerminal", aliErr("InvalidFileSystemPath.InvalidCharacters"), codes.InvalidArgument},
		// TERMINAL: retrying makes external-provisioner back off to ~16 minutes and loop forever, burning quota.
		{"regionNotSupportedIsTerminal", aliErr("OperationDenied.RegionNotSupported"), codes.InvalidArgument},
		{"unsupportedRegionIsTerminal", aliErr("UnsupportedRegion"), codes.InvalidArgument},
		// The state-class half stays retryable - it flips without the request changing.
		{"protocolNotSupportedIsRetryable", aliErr("InvalidProtocolType.NotSupported"), codes.Internal},
		{"accesspointQuotaExceededIsRetryable", aliErr("OperationDenied.AccessPointCountsExceeded"), codes.Internal},
		// The driver does NOT hardcode a region whitelist, it forwards the cloud's own verdict.
		{"invalidRegionIsTerminal", aliErr("InvalidRegionId.NotFound"), codes.InvalidArgument},
		{"invalidZoneIsTerminal", aliErr("InvalidZoneId.NotFound"), codes.InvalidArgument},
		{"forbiddenRegionIsTerminal", aliErr("Forbidden.RegionDisabled"), codes.InvalidArgument},
		// The padding NAS adds to codes must not flip the verdict in either direction.
		{"paddedRegionNotSupportedIsTerminal", aliErr("OperationDenied.RegionNotSupported\n"), codes.InvalidArgument},
		{"paddedThrottlingIsRetryable", aliErr("\tThrottling.User "), codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.createAgenticSpaceErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.deleteAgenticSpaceReqs, "creation never performs cleanup")
		})
	}
}

func TestAgenticfsCreateVolumeEmptyAgenticSpaceId(t *testing.T) {
	tests := []struct {
		name    string
		resp    *sdk.CreateAgenticSpaceResponse
		wantMsg string
	}{
		// They share ONE orphan log, but their error messages must be distinguishable.
		{"nilResponse", nil, "malformed response"},
		{"nilBody", &sdk.CreateAgenticSpaceResponse{}, "malformed response"},
		{"nilField", &sdk.CreateAgenticSpaceResponse{Body: &sdk.CreateAgenticSpaceResponseBody{}}, "empty AgenticSpaceId"},
		{"emptyString", &sdk.CreateAgenticSpaceResponse{Body: &sdk.CreateAgenticSpaceResponseBody{AgenticSpaceId: tea.String("")}}, "empty AgenticSpaceId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.createAgenticSpaceResp = tt.resp
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			// An empty ID is ambiguous; keep the token so a later request can recover it.
			assert.Equal(t, codes.Internal, status.Code(err))
			assert.Contains(t, err.Error(), tt.wantMsg, "The two failure modes must be tellable apart")
			assert.Empty(t, fake.createAccessPointReqs)
			assert.Empty(t, fake.deleteAgenticSpaceReqs)
		})
	}
}

func TestAgenticfsCreateVolumeCreateAccessPointError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"transientErrorIsRetryable", errors.New("connection reset"), codes.Internal},
		// An externally deleted space can still be referenced by a replayed ClientToken.
		{"missingSpaceIsTerminal", aliErr("InvalidAgenticSpaceId.NotFound"), codes.InvalidArgument},
		{"permanentParameterErrorIsTerminal", aliErr("InvalidParameter.VSwitchId"), codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.createAccessPointErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

func TestAgenticfsCreateVolumeEmptyAccessPointResponse(t *testing.T) {
	tests := []struct {
		name string
		resp *sdk.CreateAccessPointResponse
	}{
		{"nilResponse", nil},
		{"nilBody", &sdk.CreateAccessPointResponse{}},
		{"nilAccesspoint", &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{}}},
		{"emptyAccesspointId", &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{
			AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{AccessPointId: tea.String("")},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.createAccessPointResp = tt.resp
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			// The accesspoint could then never be found or deleted again.
			assert.Equal(t, codes.Internal, status.Code(err))
		})
	}
}

func TestAgenticfsCreateVolumeEmptyDomain(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAccessPointResp = &sdk.CreateAccessPointResponse{
		Body: &sdk.CreateAccessPointResponseBody{
			AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{
				AccessPointId:     tea.String(testAgenticFsAccessPointID),
				AccessPointDomain: tea.String(""),
			},
		},
	}
	fake.apDomainSet = true
	fake.apDomain = ""
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "no domain")
}

func TestAgenticfsCreateVolumeDomainFilledFromPolling(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAccessPointResp = &sdk.CreateAccessPointResponse{
		Body: &sdk.CreateAccessPointResponseBody{
			AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{
				AccessPointId: tea.String(testAgenticFsAccessPointID),
			},
		},
	}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
}

func TestAgenticfsCreateVolumeAPBecomesActiveAfterPending(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.apStatuses = []string{"Pending", "Pending", accessPointStatusActive}
	ctrl := newAgenticfsCtrl(t, fake)
	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, 3, fake.describeCalls)
}

func TestAgenticfsCreateVolumeAPNeverActiveTimesOut(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.apStatuses = []string{"Pending"}
	ctrl := newAgenticfsCtrl(t, fake)

	done := make(chan error, 1)
	go func() {
		_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		done <- err
	}()
	select {
	case err := <-done:
		require.Error(t, err)
		// The accesspoint is still converging, so the provisioner should retry.
		assert.Equal(t, codes.DeadlineExceeded, status.Code(err))
		assert.Contains(t, err.Error(), "to become Active")
	case <-time.After(5 * time.Second):
		t.Fatal("CreateVolume did not return; AP polling did not respect the timeout")
	}
	// Retryable, so the space and accesspoint are left for the next attempt, which reuses them via discovery.
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

func TestAgenticfsCreateVolumeTransientDescribeErrorIsRetried(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeErr = errors.New("ServiceUnavailable: request throttled")
	fake.describeErrCalls = 2
	fake.apStatuses = []string{accessPointStatusActive}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, 3, fake.describeCalls, "the two failures must be retried inside the poll budget")
}

func TestAgenticfsCreateVolumeDescribeNotFoundIsReadAfterWrite(t *testing.T) {
	fake := newFakeNasClientV2()
	// Right after CreateAccesspoint the read API may not see the accesspoint yet.
	fake.apStatuses = []string{fakeStatusNotFound, fakeStatusNotFound, accessPointStatusActive}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, 3, fake.describeCalls)
}

func TestAgenticfsCreateVolumePermanentDescribeErrorIsTerminal(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeErr = aliErr("InvalidParameter.AccessPointId")
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Equal(t, 1, fake.describeCalls, "a request the cloud will never accept must not be retried")
}

// Proves the poll loop inspects ctx before spending an API call.
func TestAgenticfsCreateVolumeCancelledContextChecksBeforeFirstCall(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ctrl.CreateVolume(ctx, agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.DeadlineExceeded, status.Code(err))
	assert.Contains(t, err.Error(), "to become Active")
	assert.Zero(t, fake.describeCalls, "ctx.Done must be checked before the first DescribeAccesspoint")
	// DeadlineExceeded is retryable, so nothing is destroyed even though ctx is already dead.
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

func TestAgenticfsCreateVolumePostSpaceFailuresLeaveResources(t *testing.T) {
	tests := []struct {
		name     string
		mutate   func(fake *fakeNasClientV2)
		wantCode codes.Code
		// wantCreateAPCalls is how often CreateAccesspoint was attempted, wantExistingAPs how many came into existence.
		wantCreateAPCalls int
		wantExistingAPs   int
	}{
		{"terminalListFailure", func(f *fakeNasClientV2) { f.listAccessPointsErr = aliErr("Forbidden.RAM") }, codes.InvalidArgument, 0, 0},
		{"terminalCreateAPFailure", func(f *fakeNasClientV2) { f.createAccessPointErr = aliErr("InvalidParameter.VpcId") }, codes.InvalidArgument, 1, 0},
		{"terminalDescribeFailure", func(f *fakeNasClientV2) { f.describeErr = aliErr("Forbidden.RAM") }, codes.InvalidArgument, 1, 1},
		{"listFails", func(f *fakeNasClientV2) { f.listAccessPointsErr = errors.New("boom") }, codes.Internal, 0, 0},
		{"createAccesspointFails", func(f *fakeNasClientV2) { f.createAccessPointErr = errors.New("boom") }, codes.Internal, 1, 0},
		{"emptyAccesspointId", func(f *fakeNasClientV2) {
			f.createAccessPointResp = &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{
				AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{AccessPointId: tea.String("")},
			}}
		}, codes.Internal, 1, 0},
		{"accesspointNeverBecomesActive", func(f *fakeNasClientV2) { f.apStatuses = []string{"Pending"} }, codes.DeadlineExceeded, 1, 1},
		{"activeWithoutADomain", func(f *fakeNasClientV2) {
			f.createAccessPointResp = &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{
				AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{AccessPointId: tea.String(testAgenticFsAccessPointID)},
			}}
			f.apDomainSet = true
			f.apDomain = ""
		}, codes.Internal, 1, 1},
		// The poll budget is consumed by transient describe failures first; same outcome, different route.
		{"accesspointNeverBecomesActiveAfterTransientDescribeErrors", func(f *fakeNasClientV2) {
			f.describeErr = errors.New("ServiceUnavailable: request throttled")
			f.describeErrCalls = 2
			f.apStatuses = []string{"Pending"}
		}, codes.DeadlineExceeded, 1, 1},
		// Discovery is pinned to an explicitly empty page, so the accesspoint came from THIS call's CreateAccesspoint.
		{"activeWithoutADomainAfterAnExplicitlyEmptyDiscoveryPage", func(f *fakeNasClientV2) {
			f.listPages = []*sdk.ListAccessPointsResponseBody{{}}
			f.createAccessPointResp = &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{
				AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{AccessPointId: tea.String(testAgenticFsAccessPointID)},
			}}
			f.apDomainSet = true
			f.apDomain = ""
		}, codes.Internal, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			tt.mutate(fake)
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.deleteAgenticSpaceReqs, "CreateVolume never deletes the space")
			assert.Empty(t, fake.deleteAccessPointIDs, "every failure leaves the accesspoint for recovery")
			assert.Len(t, fake.createAccessPointReqs, tt.wantCreateAPCalls,
				"How many accesspoints this attempt really tried to create")
			if tt.wantExistingAPs == 1 {
				// Verify a real fake resource remains discoverable, not just an empty delete call list.
				require.Len(t, fake.createdAccessPoints, 1)
				assert.Equal(t, testAgenticFsAccessPointID, tea.StringValue(fake.createdAccessPoints[0].AccessPointId))
				assert.Equal(t, accessPointStatusActive, tea.StringValue(fake.createdAccessPoints[0].Status),
					"the retained accesspoint is discoverable, so the next attempt reuses it instead of stacking a duplicate")
			} else {
				assert.Empty(t, fake.createdAccessPoints, "no accesspoint came into existence on this row")
			}
		})
	}
}

// Retaining the accesspoint only pays off if the next attempt reuses it.
func TestAgenticfsCreateVolumeRetryReusesTheAccessPointTheFailedAttemptCreated(t *testing.T) {
	tests := []struct {
		name       string
		breakFirst func(f *fakeNasClientV2)
		fixRetry   func(f *fakeNasClientV2)
	}{
		{
			name:       "terminalDescribeFailureThenPermissionRestored",
			breakFirst: func(f *fakeNasClientV2) { f.describeErr = aliErr("Forbidden.RAM") },
			fixRetry:   func(f *fakeNasClientV2) { f.describeErr = nil },
		},
		{
			name:       "neverBecameActiveThenDoes",
			breakFirst: func(f *fakeNasClientV2) { f.apStatuses = []string{"Pending"} },
			fixRetry:   func(f *fakeNasClientV2) { f.apStatuses = []string{accessPointStatusActive} },
		},
		{
			name: "hadNoDomainThenDoes",
			breakFirst: func(f *fakeNasClientV2) {
				// The fake's default CreateAccesspoint response carries a domain, which CreateVolume would use without polling.
				f.createAccessPointResp = &sdk.CreateAccessPointResponse{Body: &sdk.CreateAccessPointResponseBody{
					AccessPoint: &sdk.CreateAccessPointResponseBodyAccessPoint{AccessPointId: tea.String(testAgenticFsAccessPointID)},
				}}
				f.apDomainSet = true
				f.apDomain = ""
			},
			fixRetry: func(f *fakeNasClientV2) {
				f.apDomainSet = true
				f.apDomain = testAgenticFsAPDomain
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			tt.breakFirst(fake)
			ctrl := newAgenticfsCtrl(t, fake)

			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err)
			require.Len(t, fake.createAccessPointReqs, 1, "the first attempt created the accesspoint")
			assert.Empty(t, fake.deleteAccessPointIDs, "It was retained, not destroyed")

			tt.fixRetry(fake)
			fake.describeIdx = 0
			resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.NoError(t, err)
			assert.Len(t, fake.createAccessPointReqs, 1,
				"the retry must reuse the retained accesspoint instead of stacking a second one on the space")
			assert.Empty(t, fake.deleteAccessPointIDs)
			assert.Empty(t, fake.deleteAgenticSpaceReqs, "the space is never deleted")
			assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
			require.Len(t, fake.createAgenticSpaceReqs, 2)
			for _, req := range fake.createAgenticSpaceReqs {
				assert.Equal(t, testAgenticFsPVName, tea.StringValue(req.ClientToken), "retry must replay the same token")
			}
		})
	}
}

// A lost AccessPoint response cannot be replayed with a token. Recovery must discover
// the resource accepted by NAS instead of creating a duplicate.
func TestAgenticfsCreateVolumeLostAccessPointResponseIsRecoveredByDiscovery(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAccessPointErr = errors.New("response lost after NAS accepted CreateAccessPoint")
	ctrl := newAgenticfsCtrl(t, fake)
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Equal(t, codes.Internal, status.Code(err))
	require.Len(t, fake.createAccessPointReqs, 1)

	// Model the cloud-side success that the client could not observe on the first call.
	fake.createdAccessPoints = []*sdk.ListAccessPointsResponseBodyAccessPoints{
		apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain),
	}
	fake.createAccessPointErr = nil
	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Len(t, fake.createAccessPointReqs, 1, "recovery must not issue another CreateAccessPoint")
	require.Len(t, fake.createAgenticSpaceReqs, 2)
	assert.Equal(t, tea.StringValue(fake.createAgenticSpaceReqs[0].ClientToken), tea.StringValue(fake.createAgenticSpaceReqs[1].ClientToken))
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

// A failed create must preserve its error without ever calling the delete API.
func TestAgenticfsCreateVolumeFailureDoesNotCallDelete(t *testing.T) {
	fake := newFakeNasClientV2()
	// Even a terminal failure after creation must retain the accesspoint for recovery.
	fake.describeErr = aliErr("InvalidParameter.AccessPointId")
	fake.deleteAccessPointErr = errors.New("delete accesspoint failed")
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Contains(t, err.Error(), "InvalidParameter.AccessPointId")
	assert.NotContains(t, err.Error(), "delete accesspoint failed")
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

// Both newly created and reused accesspoints survive terminal failures.
func TestAgenticfsCreateVolumeRetainsNewAndReusedAccessPoints(t *testing.T) {
	t.Run("reusedAccesspointIsNeverDeleted", func(t *testing.T) {
		fake := newFakeNasClientV2()
		// A terminal failure after discovery must not destroy the reused accesspoint.
		fake.listPages = []*sdk.ListAccessPointsResponseBody{activeApPage()}
		fake.describeErr = aliErr("InvalidParameter.AccessPointId") // terminal InvalidArgument
		ctrl := newAgenticfsCtrl(t, fake)

		_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		assert.Empty(t, fake.createAccessPointReqs, "the existing accesspoint must be reused, not duplicated")
		assert.Empty(t, fake.deleteAccessPointIDs, "A reused accesspoint is not owned by this call, so it is never deleted")
		assert.Empty(t, fake.deleteAgenticSpaceReqs, "The space is never deleted")
	})

	t.Run("newAccesspointIsNeverDeleted", func(t *testing.T) {
		fake := newFakeNasClientV2()
		// A terminal describe failure must retain this call's newly created accesspoint.
		fake.describeErr = aliErr("InvalidParameter.AccessPointId") // terminal InvalidArgument
		ctrl := newAgenticfsCtrl(t, fake)

		_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		require.Len(t, fake.createAccessPointReqs, 1)
		assert.Empty(t, fake.deleteAccessPointIDs, "new accesspoints must also survive the failed request")
		assert.Empty(t, fake.deleteAgenticSpaceReqs, "The space is never deleted")
		// The only ListAccesspoints call is discovery.
		assert.Equal(t, 1, countCalls(fake, "ListAccesspoints"))
	})
}

func TestAgenticfsCreateVolumeSuccessNeverDeletesResources(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

func TestAgenticfsCreateVolumeInvalidVolumeName(t *testing.T) {
	tests := []struct {
		name     string
		volumeID string
	}{
		{"empty", ""},
		{"tooLong", strings.Repeat("a", maxVolumeNameLen+1)},
		{"forwardSlash", "pvc/escape"},
		{"backSlash", `pvc\escape`},
		{"singleDot", "."},
		{"doubleDot", ".."},
		{"nonAscii", "pvc-中文"},
		{"nonPrintable", "pvc-\x01"},
		// A real PV name is pvc-<uuid> and never contains a space, so this is zero-risk.
		{"leadingSpace", " pvc-x"},
		{"trailingSpace", "pvc-x "},
		{"onlyWhitespace", "   "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(tt.volumeID, 20*GiB, nil))
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Empty(t, fake.callOrder)
		})
	}
}

func TestAgenticfsCreateVolumeMaxLengthVolumeNameIsAccepted(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	name := strings.Repeat("a", maxVolumeNameLen)
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(name, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, name, tea.StringValue(fake.createAgenticSpaceReqs[0].ClientToken))
}

func agenticfsPV(attrs map[string]string) *corev1.PersistentVolume {
	return &corev1.PersistentVolume{
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				CSI: &corev1.CSIPersistentVolumeSource{
					Driver:           "nasplugin.csi.alibabacloud.com",
					VolumeHandle:     testAgenticFsPVName,
					VolumeAttributes: attrs,
				},
			},
		},
	}
}

func agenticfsDeletePV() *corev1.PersistentVolume {
	return agenticfsPV(map[string]string{
		"volumeAs":          agenticFsVolumeAs,
		filesystemIDKey:     testAgenticFsFilesystemID,
		vcKeyAccesspointId:  testAgenticFsAccessPointID,
		vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
	})
}

func agenticfsDeleteReq() *csi.DeleteVolumeRequest {
	return &csi.DeleteVolumeRequest{VolumeId: testAgenticFsPVName}
}

func TestAgenticfsDeleteVolumeSuccess(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.NoError(t, err)
	// The OpenAPI refuses DeleteAgenticSpace while any accesspoint still exists.
	assert.Equal(t,
		[]string{"ListAccesspoints", "DeleteAccesspoint", "DescribeAccesspoint", "DeleteAgenticSpace"},
		fake.callOrder)
	require.Len(t, fake.deleteAgenticSpaceReqs, 1)
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(fake.deleteAgenticSpaceReqs[0].AgenticSpaceId))
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(fake.deleteAgenticSpaceReqs[0].FileSystemId))
	assert.Equal(t, testAgenticFsPVName, tea.StringValue(fake.deleteAgenticSpaceReqs[0].ClientToken))
	require.Len(t, fake.listAccessPointsReqs, 1)
	require.Len(t, fake.listAccessPointsReqs[0].Filters, 1)
	assert.Equal(t, apListFilterAgenticSpaceId, tea.StringValue(fake.listAccessPointsReqs[0].Filters[0].Name))
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(fake.listAccessPointsReqs[0].Filters[0].Value))
}

// DeleteAgenticSpace requires ALL accesspoints detached, so deleting only the recorded one wedges DeleteVolume.
func TestAgenticfsDeleteVolumeDeletesEveryAccessPointOfTheSpace(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(
		apItem("ap-orphan-1", accessPointStatusActive, "orphan1.example.com"),
		apItem(testAgenticFsAccessPointID, accessPointStatusActive, testAgenticFsAPDomain),
		apItem("ap-orphan-2", accessPointStatusActive, "orphan2.example.com"),
	)}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.NoError(t, err)
	// The list API can lag a fresh accesspoint, so the recorded ID goes first, with no duplicates.
	assert.Equal(t,
		[]string{testAgenticFsAccessPointID, "ap-orphan-1", "ap-orphan-2"},
		fake.deleteAccessPointIDs)
	assert.Len(t, fake.deleteAgenticSpaceReqs, 1)
	assert.Equal(t, 3, countCalls(fake, "DescribeAccesspoint"), "every accesspoint must be waited for")
}

func TestAgenticfsDeleteVolumeIdempotent(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	fake.deleteAccessPointErr = wrap.ErrorCode("NotFound")
	fake.deleteAgenticSpaceErr = wrap.ErrorCode("NotFound")
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.NoError(t, err)
	require.NotNil(t, resp)
	// The wait still runs on NotFound: an accesspoint mid-Deleting would otherwise fail DeleteAgenticSpace for rounds.
	assert.Equal(t,
		[]string{"ListAccesspoints", "DeleteAccesspoint", "DescribeAccesspoint", "DeleteAgenticSpace"},
		fake.callOrder, "An already-gone accesspoint is still waited for")
}

func TestAgenticfsDeleteVolumeSpaceAlreadyGone(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	fake.listAccessPointsErr = aliErr("InvalidAgenticSpaceId.NotFound")
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

func TestAgenticfsDeleteVolumeMissingFilesystemId(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	pv := agenticfsPV(map[string]string{
		vcKeyAccesspointId:  testAgenticFsAccessPointID,
		vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
	})
	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), pv)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Contains(t, err.Error(), filesystemIDKey)
	assert.Empty(t, fake.callOrder)
}

func TestAgenticfsDeleteVolumeLegacyFilesystemIdKey(t *testing.T) {
	// An earlier (never released) build used the "filesystemId" spelling; it must survive a rolling upgrade.
	fake := newDeleteFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	pv := agenticfsPV(map[string]string{
		vcKeyFilesystemIdLegacy: testAgenticFsFilesystemID,
		vcKeyAccesspointId:      testAgenticFsAccessPointID,
		vcKeyAgenticSpaceId:     testAgenticFsAgenticSpaceID,
	})
	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), pv)
	require.NoError(t, err)
	require.Len(t, fake.deleteAgenticSpaceReqs, 1)
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(fake.deleteAgenticSpaceReqs[0].FileSystemId))
}

// Without the space ID nothing can be deleted, so reporting success would leak a billable AgenticSpace.
func TestAgenticfsDeleteVolumeMissingAgenticSpaceId(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	pv := agenticfsPV(map[string]string{
		filesystemIDKey:    testAgenticFsFilesystemID,
		vcKeyAccesspointId: testAgenticFsAccessPointID,
	})
	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), pv)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	// The message echoes all read IDs and points the operator at the documented cleanup procedure.
	assert.Contains(t, err.Error(), testAgenticFsFilesystemID)
	assert.Contains(t, err.Error(), testAgenticFsAccessPointID)
	assert.Contains(t, err.Error(), testAgenticFsPVName)
	assert.Contains(t, err.Error(), "README.md §4.1")
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

func TestAgenticfsDeleteVolumeMissingAccesspointId(t *testing.T) {
	t.Run("noAccesspointListedEitherIsRetryable", func(t *testing.T) {
		fake := newDeleteFakeNasClientV2()
		ctrl := newAgenticfsCtrl(t, fake)
		pv := agenticfsPV(map[string]string{
			filesystemIDKey:     testAgenticFsFilesystemID,
			vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
		})
		_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), pv)
		require.Error(t, err)
		// ListAccesspoints can lag a freshly created accesspoint, so this may self-heal: Aborted, not InvalidArgument.
		assert.Equal(t, codes.Aborted, status.Code(err))
		assert.Empty(t, fake.deleteAgenticSpaceReqs,
			"deleting the space while an unlisted accesspoint may exist would leak it")
	})

	t.Run("fallsBackToListedAccesspoints", func(t *testing.T) {
		fake := newDeleteFakeNasClientV2()
		fake.listPages = []*sdk.ListAccessPointsResponseBody{activeApPage()}
		ctrl := newAgenticfsCtrl(t, fake)
		pv := agenticfsPV(map[string]string{
			filesystemIDKey:     testAgenticFsFilesystemID,
			vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
		})
		_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), pv)
		require.NoError(t, err)
		assert.Equal(t, []string{testAgenticFsAccessPointID}, fake.deleteAccessPointIDs)
		assert.Len(t, fake.deleteAgenticSpaceReqs, 1)
	})
}

// DeleteAgenticSpace is refused while the accesspoint is still detaching, and failing to converge is retryable.
func TestAgenticfsDeleteVolumeWaitsForAccessPointGone(t *testing.T) {
	t.Run("deletedAfterAccesspointDisappears", func(t *testing.T) {
		fake := newFakeNasClientV2()
		fake.apStatuses = []string{accessPointStatusDeleting, accessPointStatusDeleting, fakeStatusNotFound}
		ctrl := newAgenticfsCtrl(t, fake)

		_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
		require.NoError(t, err)
		assert.Equal(t, 3, fake.describeCalls)
		require.Len(t, fake.deleteAgenticSpaceReqs, 1)
		assert.Equal(t,
			[]string{"ListAccesspoints", "DeleteAccesspoint", "DescribeAccesspoint", "DescribeAccesspoint",
				"DescribeAccesspoint", "DeleteAgenticSpace"},
			fake.callOrder)
	})

	t.Run("stillDeletingIsRetryableAndDoesNotTouchSpace", func(t *testing.T) {
		fake := newFakeNasClientV2()
		fake.apStatuses = []string{accessPointStatusDeleting}
		ctrl := newAgenticfsCtrl(t, fake)

		done := make(chan error, 1)
		go func() {
			_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			done <- err
		}()
		var err error
		select {
		case err = <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("DeleteVolume did not return; the accesspoint wait did not respect the timeout")
		}
		require.Error(t, err)
		// Aborted, not Internal: the accesspoint is converging towards gone.
		assert.Equal(t, codes.Aborted, status.Code(err))
		assert.Empty(t, fake.deleteAgenticSpaceReqs,
			"DeleteAgenticSpace would fail anyway while an accesspoint is attached")
	})
}

func TestAgenticfsDeleteVolumeAPIErrors(t *testing.T) {
	tests := []struct {
		name     string
		mutate   func(fake *fakeNasClientV2)
		wantCode codes.Code
	}{
		{"listFailsTransiently", func(f *fakeNasClientV2) { f.listAccessPointsErr = errors.New("boom") }, codes.Internal},
		{"listFailsPermanently", func(f *fakeNasClientV2) { f.listAccessPointsErr = aliErr("InvalidParameter.Filter") }, codes.InvalidArgument},
		{"deleteAccesspointFails", func(f *fakeNasClientV2) { f.deleteAccessPointErr = errors.New("boom") }, codes.Internal},
		// "*.NotSupported" is retryable now, so a permanent delete failure uses a malformed-parameter code.
		{"deleteAccesspointFailsPermanently", func(f *fakeNasClientV2) {
			f.deleteAccessPointErr = aliErr("InvalidParameter.AccessPointId")
		}, codes.InvalidArgument},
		{"deleteAccesspointNotSupportedIsRetryable", func(f *fakeNasClientV2) {
			f.deleteAccessPointErr = aliErr("OperationDenied.NotSupported")
		}, codes.Internal},
		{"deleteAgenticspaceFails", func(f *fakeNasClientV2) { f.deleteAgenticSpaceErr = errors.New("boom") }, codes.Internal},
		{"deleteAgenticspaceFailsPermanently", func(f *fakeNasClientV2) {
			f.deleteAgenticSpaceErr = aliErr("InvalidParameter.AgenticSpaceId")
		}, codes.InvalidArgument},
		{"deleteAgenticspaceNotSupportedIsRetryable", func(f *fakeNasClientV2) {
			f.deleteAgenticSpaceErr = aliErr("InvalidAgenticSpaceId.NotSupported")
		}, codes.Internal},
		// Terminal on the delete path too, so the operator gets an event instead of a PV retrying its own deletion.
		{"deleteAccesspointRegionNotSupportedIsTerminal", func(f *fakeNasClientV2) {
			f.deleteAccessPointErr = aliErr("OperationDenied.RegionNotSupported")
		}, codes.InvalidArgument},
		{"deleteAgenticspaceRegionNotSupportedIsTerminal", func(f *fakeNasClientV2) {
			f.deleteAgenticSpaceErr = aliErr("OperationDenied.RegionNotSupported")
		}, codes.InvalidArgument},
		// Padding must not flip a state-class code into a terminal one, nor hide a terminal one.
		{"deleteAccesspointPaddedRegionNotSupportedIsTerminal", func(f *fakeNasClientV2) {
			f.deleteAccessPointErr = aliErr("OperationDenied.RegionNotSupported\n")
		}, codes.InvalidArgument},
		{"deleteAccesspointPaddedNotSupportedIsRetryable", func(f *fakeNasClientV2) {
			f.deleteAccessPointErr = aliErr("\tOperationDenied.NotSupported ")
		}, codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newDeleteFakeNasClientV2()
			tt.mutate(fake)
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

func TestAgenticfsDeleteVolumePaginationCapIsRetryable(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	fake.listPages = []*sdk.ListAccessPointsResponseBody{{NextToken: tea.String("forever")}}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.Error(t, err)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Empty(t, fake.deleteAccessPointIDs, "nothing may be deleted from an incomplete listing")
}

func agenticfsExpandPV() *corev1.PersistentVolume {
	return agenticfsPV(map[string]string{
		"volumeAs":          agenticFsVolumeAs,
		filesystemIDKey:     testAgenticFsFilesystemID,
		vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
	})
}

func agenticfsExpandReq(requiredBytes int64) *csi.ControllerExpandVolumeRequest {
	return &csi.ControllerExpandVolumeRequest{
		VolumeId:      testAgenticFsPVName,
		CapacityRange: &csi.CapacityRange{RequiredBytes: requiredBytes},
	}
}

func getSpaceResp(sizeLimit, fileCountLimit, spaceUsage, fileCountUsage int64) *sdk.GetAgenticSpaceResponse {
	return &sdk.GetAgenticSpaceResponse{
		Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota: &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{
					SizeLimit:      tea.Int64(sizeLimit),
					FileCountLimit: tea.Int64(fileCountLimit),
				},
				SpaceUsage:     tea.Int64(spaceUsage),
				FileCountUsage: tea.Int64(fileCountUsage),
			},
		},
	}
}

func TestAgenticfsControllerExpandVolume(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.getAgenticSpaceResp = getSpaceResp(20*GiB, 500000, 1*GiB, 100)
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
	require.NoError(t, err)
	assert.Equal(t, int64(30*GiB), resp.CapacityBytes)

	require.Len(t, fake.setQuotaReqs, 1)
	q := fake.setQuotaReqs[0]
	// Flat fields (not a nested Quota), not SetDirQuota.
	assert.Equal(t, int64(30*GiB), tea.Int64Value(q.SizeLimit))
	// Passed through unchanged instead of being silently rewritten to the default.
	assert.Equal(t, int64(500000), tea.Int64Value(q.FileCountLimit))
	assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(q.FileSystemId))
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(q.AgenticSpaceId))
	assert.Equal(t, []string{"GetAgenticSpace", "SetAgenticSpaceQuota"}, fake.callOrder)
}

func TestAgenticfsControllerExpandVolumeMissingAttributes(t *testing.T) {
	tests := []struct {
		name string
		pv   *corev1.PersistentVolume
	}{
		{"missingFilesystemId", agenticfsPV(map[string]string{vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID})},
		{"missingAgenticSpaceId", agenticfsPV(map[string]string{filesystemIDKey: testAgenticFsFilesystemID})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), tt.pv)
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Empty(t, fake.callOrder)
		})
	}
}

// SetAgenticSpaceQuota always sends BOTH limits, so an unreadable file count limit must not be guessed.
func TestAgenticfsControllerExpandVolumeIncompleteQuota(t *testing.T) {
	tests := []struct {
		name string
		resp *sdk.GetAgenticSpaceResponse
	}{
		{"nilResponse", nil},
		{"nilBody", &sdk.GetAgenticSpaceResponse{}},
		{"nilAgenticspace", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{}}},
		{"nilQuota", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{},
		}}},
		// Trips the nil-usage guard before reaching the omit-the-unreadable-file-count-limit branch.
		{"nilUsages", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota: &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{SizeLimit: tea.Int64(20 * GiB)},
			},
		}}},
		{"nilSizeLimit", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota: &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{FileCountLimit: tea.Int64(500000)},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceResp = tt.resp
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
			require.Error(t, err)
			assert.Equal(t, codes.Internal, status.Code(err))
			assert.Contains(t, err.Error(), "refusing to guess")
			assert.Empty(t, fake.setQuotaReqs)
		})
	}
}

func TestAgenticfsControllerExpandVolumeGetError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"transient", errors.New("boom"), codes.Internal},
		{"permanent", aliErr("InvalidAgenticSpaceId.NotFound"), codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.setQuotaReqs)
		})
	}
}

func TestAgenticfsControllerExpandVolumeProtections(t *testing.T) {
	tests := []struct {
		name          string
		space         *sdk.GetAgenticSpaceResponse
		requiredBytes int64
		limitBytes    int64
		wantCode      codes.Code
	}{
		{
			name:          "below10GibMinimum",
			space:         getSpaceResp(20*GiB, 500000, 0, 0),
			requiredBytes: 5 * GiB,
			wantCode:      codes.InvalidArgument,
		},
		{
			name:          "shrinkIsRefused",
			space:         getSpaceResp(30*GiB, 500000, 0, 0),
			requiredBytes: 20 * GiB,
			wantCode:      codes.OutOfRange,
		},
		{
			name:          "belowSpaceUsageIsRefused",
			space:         getSpaceResp(20*GiB, 500000, 25*GiB, 0),
			requiredBytes: 20 * GiB,
			wantCode:      codes.OutOfRange,
		},
		{
			name:          "roundingOverLimitBytesIsRefused",
			space:         getSpaceResp(20*GiB, 500000, 0, 0),
			requiredBytes: 30*GiB + 1,
			limitBytes:    30 * GiB,
			wantCode:      codes.OutOfRange,
		},
		{
			name:          "noRequestedCapacity",
			space:         getSpaceResp(20*GiB, 500000, 0, 0),
			requiredBytes: 0,
			wantCode:      codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceResp = tt.space
			ctrl := newAgenticfsCtrl(t, fake)
			req := agenticfsExpandReq(tt.requiredBytes)
			if tt.limitBytes > 0 {
				req.CapacityRange.LimitBytes = tt.limitBytes
			}
			_, err := ctrl.ControllerExpandVolume(context.Background(), req, agenticfsExpandPV())
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Empty(t, fake.setQuotaReqs)
		})
	}
}

func TestAgenticfsControllerExpandVolumeSameSizeIsAccepted(t *testing.T) {
	// external-resizer re-issues an expansion that already succeeded; a no-op grow must not read as a shrink.
	fake := newFakeNasClientV2()
	fake.getAgenticSpaceResp = getSpaceResp(30*GiB, 500000, 0, 0)
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
	require.NoError(t, err)
	assert.Equal(t, int64(30*GiB), resp.CapacityBytes)
}

// Expand only grows SizeLimit and passes FileCountLimit through, so an out-of-range read-back is not this
// expansion's fault. CreateVolume keeps the strict validator, where the value IS the caller's.
func TestAgenticfsControllerExpandVolumeFileCountLimitOutOfRangeIsOmitted(t *testing.T) {
	fake := newFakeNasClientV2()
	// Below the OpenAPI minimum yet non-nil and non-zero, so the unreadable-value branch does not omit it.
	fake.getAgenticSpaceResp = getSpaceResp(20*GiB, 5000, 0, 0)
	ctrl := newAgenticfsCtrl(t, fake)
	logger, ctx := newLogCapture(t)

	resp, err := ctrl.ControllerExpandVolume(ctx, agenticfsExpandReq(30*GiB), agenticfsExpandPV())
	require.NoError(t, err, "An out-of-range read-back FileCountLimit must NOT block the size expansion")
	assert.Equal(t, int64(30*GiB), resp.CapacityBytes)
	require.Len(t, fake.setQuotaReqs, 1, "the expansion must still reach SetAgenticSpaceQuota")
	assert.Equal(t, int64(30*GiB), tea.Int64Value(fake.setQuotaReqs[0].SizeLimit))
	assert.Nil(t, fake.setQuotaReqs[0].FileCountLimit,
		"The illegal read-back FileCountLimit is OMITTED, not forwarded to the OpenAPI")

	logs := logText(logger)
	require.NotEmpty(t, logs, "the capture is vacuous - ktesting buffered nothing")
	assert.Contains(t, logs, "WARNING", "The downgrade must be visible to an operator")
	assert.Contains(t, logs, "OMITTED", "The warning must say the field was omitted")
}

func TestAgenticfsControllerExpandVolumeFileCountLimitBelowUsageIsWarning(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.getAgenticSpaceResp = getSpaceResp(20*GiB, 10000, 0, 20000)
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
	require.NoError(t, err)
	assert.Equal(t, int64(30*GiB), resp.CapacityBytes)
	require.Len(t, fake.setQuotaReqs, 1)
	assert.Equal(t, int64(30*GiB), tea.Int64Value(fake.setQuotaReqs[0].SizeLimit))
	assert.Equal(t, int64(10000), tea.Int64Value(fake.setQuotaReqs[0].FileCountLimit))
}

// FileCountLimit is not stored in the VolumeContext and the OpenAPI minimum is 10000, so a nil or 0 read-back
// must be OMITTED rather than sent as 0.
func TestAgenticfsControllerExpandVolumeOmitsUnreadableFileCountLimit(t *testing.T) {
	tests := []struct {
		name string
		resp *sdk.GetAgenticSpaceResponse
	}{
		{"nilFileCountLimit", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota:          &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{SizeLimit: tea.Int64(20 * GiB)},
				SpaceUsage:     tea.Int64(0),
				FileCountUsage: tea.Int64(0),
			},
		}}},
		{"zeroFileCountLimit", getSpaceResp(20*GiB, 0, 0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceResp = tt.resp
			ctrl := newAgenticfsCtrl(t, fake)
			resp, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
			require.NoError(t, err)
			assert.Equal(t, int64(30*GiB), resp.CapacityBytes)
			require.Len(t, fake.setQuotaReqs, 1)
			assert.Nil(t, fake.setQuotaReqs[0].FileCountLimit, "An unreadable file count limit is omitted, never sent as 0")
		})
	}
}

// A nil usage would make the shrink/usage guards evaluate against 0 and let a bad write through.
func TestAgenticfsControllerExpandVolumeMissingUsageIsRefused(t *testing.T) {
	tests := []struct {
		name string
		resp *sdk.GetAgenticSpaceResponse
	}{
		{"nilSpaceUsage", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota: &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{
					SizeLimit:      tea.Int64(20 * GiB),
					FileCountLimit: tea.Int64(500000),
				},
				FileCountUsage: tea.Int64(0),
			},
		}}},
		{"nilFileCountUsage", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{
			AgenticSpace: &sdk.GetAgenticSpaceResponseBodyAgenticSpace{
				Quota: &sdk.GetAgenticSpaceResponseBodyAgenticSpaceQuota{
					SizeLimit:      tea.Int64(20 * GiB),
					FileCountLimit: tea.Int64(500000),
				},
				SpaceUsage: tea.Int64(0),
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceResp = tt.resp
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
			require.Error(t, err)
			assert.Equal(t, codes.Internal, status.Code(err))
			assert.Contains(t, err.Error(), "refusing to guess")
			assert.Empty(t, fake.setQuotaReqs)
		})
	}
}

func TestAgenticfsControllerExpandVolumeSetQuotaError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"transient", errors.New("boom"), codes.Internal},
		{"permanent", aliErr("InvalidParameter.SizeLimit"), codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.getAgenticSpaceResp = getSpaceResp(20*GiB, 500000, 0, 0)
			fake.setQuotaErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			_, err := ctrl.ControllerExpandVolume(context.Background(), agenticfsExpandReq(30*GiB), agenticfsExpandPV())
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
			assert.Len(t, fake.setQuotaReqs, 1)
		})
	}
}

func TestAgenticfsFilesystemID(t *testing.T) {
	tests := []struct {
		name  string
		attrs map[string]string
		want  string
	}{
		{"canonicalKey", map[string]string{filesystemIDKey: "fs-1"}, "fs-1"},
		{"legacyKey", map[string]string{vcKeyFilesystemIdLegacy: "fs-legacy"}, "fs-legacy"},
		{"canonicalWins", map[string]string{filesystemIDKey: "fs-1", vcKeyFilesystemIdLegacy: "fs-legacy"}, "fs-1"},
		{"emptyCanonicalFallsBack", map[string]string{filesystemIDKey: "", vcKeyFilesystemIdLegacy: "fs-legacy"}, "fs-legacy"},
		{"missing", map[string]string{}, ""},
		{"nil", nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, agenticfsFilesystemID(tt.attrs))
		})
	}
}

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"bareWrapErrorCode", wrap.ErrorCode("NotFound"), true},
		{"wrappedWrapErrorCode", fmt.Errorf("outer: %w", wrap.ErrorCode("NotFound")), true},
		{"exactCodeFromTheSdk", aliErr("NotFound"), true},
		{"documentedAgenticspace", aliErr("InvalidAgenticSpace.NotFound"), true},
		{"paddedDocumentedAgenticspace", aliErr("\tInvalidAgenticSpace.NotFound\n"), true},
		{"wrappedDocumentedAgenticspace", fmt.Errorf("outer: %w", aliErr("INVALIDAGENTICSPACE.NOTFOUND")), true},
		{"namespacedAgenticspace", aliErr("InvalidAgenticSpaceId.NotFound"), true},
		{"namespacedAccesspoint", aliErr("InvalidAccessPoint.NotFound"), true},
		{"namespacedAccesspointId", aliErr("InvalidAccessPointId.NotFound"), true},
		{"namespacedFilesystem", aliErr("InvalidFileSystem.NotFound"), true},
		{"namespacedFilesystemId", aliErr("InvalidFileSystemId.NotFound"), true},
		// Not one of the three managed resources, so it stays retryable instead of short-circuiting a delete.
		{"unrelatedNamespacedNotfound", aliErr("Quota.NotFound"), false},
		{"unrelatedResourceNotfound", aliErr("SomeResource.NotFound"), false},
		{"filesetIsOnlySpaceScoped", aliErr("InvalidFilesetId.NotFound"), false},
		{"fsetIsOnlySpaceScoped", aliErr("InvalidFsetId.NotFound"), false},
		// Contains "FileSystem", so the old substring logic swallowed a capacity/quota verdict as "already gone".
		{"capacityNotfoundIsNotGone", aliErr("InvalidFileSystemCapacity.NotFound"), false},
		{"unrelatedCode", aliErr("Throttling.User"), false},
		{"plainErrorMentioningNotFound", errors.New("NotFound"), false},
		{"invalidPathIsNotMissingResource", aliErr("InvalidFileSystemPath.InvalidCharacters"), false},
		{"wrappedNamespacedCode", fmt.Errorf("outer: %w", aliErr("InvalidAgenticSpaceId.NotFound")), true},
		// NAS pads error codes: one trailing newline defeated the exact comparisons while the substring test matched.
		{"trailingNewlineOnANamespacedCode", aliErr("InvalidAccessPointId.NotFound\n"), true},
		{"tabAndTrailingSpaceOnABareCode", aliErr("\tNotFound "), true},
		{"newlineOnANamespacedFilesystemCode", aliErr("InvalidFileSystem.NotFound\n"), true},
		{"whitespaceDoesNotSmuggleAnUnrelatedCodeIn", aliErr("Quota.NotFound\n"), false},
		{"whitespaceOnly", aliErr("  \n\t "), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isNotFoundError(tt.err))
		})
	}
}

// A region/zone verdict is TERMINAL while a state-class *.NotSupported stays retryable.
func TestIsPermanentAPIError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		// Rule 1: the exact permanent code set (lower-cased, TrimSpace'd).
		{"invalidRegion", aliErr("InvalidRegionId.NotFound"), true},
		{"invalidZone", aliErr("InvalidZoneId.NotFound"), true},
		{"forbiddenRegionDisabled", aliErr("Forbidden.RegionDisabled"), true},
		{"forbiddenRegion", aliErr("Forbidden.Region"), true},
		{"invalidPath", aliErr("InvalidFileSystemPath.InvalidCharacters"), true},
		{"paddedInvalidPath", aliErr("\tInvalidFileSystemPath.InvalidCharacters\n"), true},
		{"unknownPathErrorIsRetryable", aliErr("InvalidFileSystemPath.NotSupported"), false},
		// Permanent prefixes, so any suffix qualifies.
		{"invalidParameter", aliErr("InvalidParameter.SizeLimit"), true},
		{"invalidParamShortPrefix", aliErr("InvalidParam.Foo"), true},
		{"invalidRegionPrefix", aliErr("InvalidRegion.Something"), true},
		{"invalidZonePrefix", aliErr("InvalidZone.Something"), true},
		{"forbiddenPrefix", aliErr("Forbidden.SubUser"), true},
		{"lowercaseCode", aliErr("invalidparameter.foo"), true},
		// The PURE-zone rows are what make this real: the ZoneRegion row matches on its "region" substring, so alone it
		// gives false confidence about the zone dimension.
		{"regionNotSupportedIsTerminal", aliErr("OperationDenied.RegionNotSupported"), true},
		{"zoneAndRegionNotSupportedIsTerminalMatchesOnRegion", aliErr("OperationDenied.ZoneRegionNotSupported"), true},
		{"pureZoneNotSupportedIsTerminal", aliErr("OperationDenied.ZoneNotSupported"), true},
		{"unsupportedRegionIsTerminal", aliErr("UnsupportedRegion"), true},
		{"unsupportedDotRegionIsTerminal", aliErr("Unsupported.Region"), true},
		{"unsupportedDotZoneIsTerminal", aliErr("Unsupported.Zone"), true},
		// The rule needs BOTH halves, so a state-class verdict stays retryable.
		{"unsupportedWithoutRegionIsRetryable", aliErr("Unsupported.Operation"), false},
		// Rule 2c: a whitelisted NotFound is permanent (nothing to retry against).
		{"whitelistedNotfound", aliErr("InvalidAgenticSpaceId.NotFound"), true},
		{"bareNotfound", aliErr("NotFound"), true},
		{"paddedWhitelistedNotfound", aliErr("InvalidAccessPointId.NotFound\n"), true},
		// These stop the classification from creeping back to "anything containing notsupport/notfound is terminal".
		{"protocolNotSupportedIsRetryable", aliErr("InvalidProtocolType.NotSupported"), false},
		{"operationDeniedNotSupportedIsRetryable", aliErr("OperationDenied.NotSupported"), false},
		{"agenticspaceNotSupportedIsRetryable", aliErr("InvalidAgenticSpaceId.NotSupported"), false},
		{"accesspointQuotaIsRetryable", aliErr("OperationDenied.AccessPointCountsExceeded"), false},
		{"throttlingIsRetryable", aliErr("Throttling.User"), false},
		{"serviceUnavailableIsRetryable", aliErr("ServiceUnavailable"), false},
		{"internalErrorIsRetryable", aliErr("InternalError"), false},
		// Stays retryable rather than short-circuiting a delete into a false success or wedging a PVC on a quota hiccup.
		{"unrelatedNotfoundIsRetryable", aliErr("Quota.NotFound"), false},
		{"resourceNotfoundIsRetryable", aliErr("ResourceNotFound"), false},
		// The classification reads the CODE only, never the Message.
		{"messageIsNeverMatched", &fakeAliErrorWithMessage{
			code:    "Throttling.User",
			message: "region InvalidParameter Forbidden NotFound UnsupportedRegion notsupported",
		}, false},
		// Padding must not change the verdict in EITHER direction.
		{"paddedRegionNotSupportedIsTerminal", aliErr("OperationDenied.RegionNotSupported\n"), true},
		{"paddedInvalidParameterIsTerminal", aliErr("\tInvalidParameter.Foo "), true},
		{"paddedThrottlingIsRetryable", aliErr("Throttling.User\n"), false},
		{"noErrorCodeAtAll", errors.New("dial tcp: i/o timeout"), false},
		{"nil", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPermanentAPIError(tt.err))
		})
	}
}

func TestApiStatusError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{"nil", nil, codes.OK},
		{"cancelledContext", context.Canceled, codes.DeadlineExceeded},
		{"expiredContext", context.DeadlineExceeded, codes.DeadlineExceeded},
		{"wrappedCancelledContext", fmt.Errorf("nas call: %w", context.Canceled), codes.DeadlineExceeded},
		{"permanentOpenapiError", aliErr("InvalidParameter.Foo"), codes.InvalidArgument},
		{"invalidPathOpenapiError", aliErr("InvalidFileSystemPath.InvalidCharacters"), codes.InvalidArgument},
		// Terminal, so the provisioner stops instead of backing off to ~16 minutes and retrying forever.
		{"regionNotSupportedOpenapiErrorIsTerminal", aliErr("OperationDenied.RegionNotSupported"), codes.InvalidArgument},
		{"unsupportedRegionOpenapiErrorIsTerminal", aliErr("UnsupportedRegion"), codes.InvalidArgument},
		{"protocolNotSupportedOpenapiErrorIsRetryable", aliErr("InvalidProtocolType.NotSupported"), codes.Internal},
		{"accesspointQuotaOpenapiErrorIsRetryable", aliErr("OperationDenied.AccessPointCountsExceeded"), codes.Internal},
		{"transientOpenapiError", aliErr("Throttling.User"), codes.Internal},
		{"plainError", errors.New("boom"), codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := apiStatusError("nas:TestOp", tt.err)
			if tt.err == nil {
				assert.NoError(t, got)
				return
			}
			require.Error(t, got)
			assert.Equal(t, tt.wantCode, status.Code(got))
			assert.Contains(t, got.Error(), "nas:TestOp", "the operation name must survive the mapping")
		})
	}

	// Re-wrapping would turn the pagination cap's Aborted into an Internal and lose the operator's detail message.
	t.Run("alreadyClassifiedStatusIsPassedThrough", func(t *testing.T) {
		original := status.Error(codes.Aborted, "retry me")
		assert.Equal(t, original, apiStatusError("nas:TestOp", original))
	})
}

func TestRoundUpToGiBChecked(t *testing.T) {
	tests := []struct {
		name       string
		bytes      int64
		limitBytes int64
		want       int64
		wantCode   codes.Code
	}{
		{"zeroStaysZero", 0, 0, 0, codes.OK},
		{"oneByteRoundsTo1GiB", 1, 0, GiB, codes.OK},
		{"exactMultipleIsKept", 20 * GiB, 0, 20 * GiB, codes.OK},
		{"oneByteOverRoundsUp", 20*GiB + 1, 0, 21 * GiB, codes.OK},
		{"openapiMaximumIsKept", maxAgenticSpaceSizeLimit, 0, maxAgenticSpaceSizeLimit, codes.OK},
		{"aboveTheOpenapiMaximum", maxAgenticSpaceSizeLimit + 1, 0, 0, codes.InvalidArgument},
		{"roundingOverflow", math.MaxInt64, 0, 0, codes.OutOfRange},
		{"onePastTheRoundingOverflowBoundary", math.MaxInt64 - GiB + 2, 0, 0, codes.OutOfRange},
		// Rejected for exceeding the OpenAPI maximum instead - a permanent, not retryable, verdict.
		{"atTheRoundingOverflowBoundary", math.MaxInt64 - GiB + 1, 0, 0, codes.InvalidArgument},
		{"overLimitBytes", 20 * GiB, 10 * GiB, 0, codes.OutOfRange},
		{"roundingPushesOverLimitBytes", 10*GiB + 1, 10 * GiB, 0, codes.OutOfRange},
		{"equalToLimitBytes", 20 * GiB, 20 * GiB, 20 * GiB, codes.OK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := roundUpToGiBChecked(tt.bytes, tt.limitBytes)
			if tt.wantCode == codes.OK {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
				return
			}
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

func TestRoundUpToGiB(t *testing.T) {
	assert.Equal(t, int64(GiB), roundUpToGiB(1))
	assert.Equal(t, int64(GiB), roundUpToGiB(GiB))
	assert.Equal(t, int64(2*GiB), roundUpToGiB(GiB+1))
	assert.Equal(t, int64(0), roundUpToGiB(0))
}

func TestComputeAgenticSpaceSizeLimit(t *testing.T) {
	tests := []struct {
		name     string
		cr       *csi.CapacityRange
		param    string
		want     int64
		wantCode codes.Code
	}{
		{"noCapacityAndNoParam", nil, "", 0, codes.InvalidArgument},
		{"negativeCapacityAndNoParam", &csi.CapacityRange{RequiredBytes: -1}, "", 0, codes.InvalidArgument},
		{"paramUsedWhenCapacityAbsent", &csi.CapacityRange{RequiredBytes: 0}, "10Gi", 10 * GiB, codes.OK},
		{"paramUsedWhenCrIsNil", nil, "20Gi", 20 * GiB, codes.OK},
		{"capacityHonoredBelowTheCap", &csi.CapacityRange{RequiredBytes: 20 * GiB}, "100Gi", 20 * GiB, codes.OK},
		{"capacityEqualToTheCap", &csi.CapacityRange{RequiredBytes: 100 * GiB}, "100Gi", 100 * GiB, codes.OK},
		{"capacityAboveTheCap", &csi.CapacityRange{RequiredBytes: 101 * GiB}, "100Gi", 0, codes.InvalidArgument},
		{"roundedCapacityAboveFractionalCap", &csi.CapacityRange{RequiredBytes: 10*GiB + 1}, "10.5Gi", 0, codes.InvalidArgument},
		{"fractionalCapWithoutCapacity", nil, "10.5Gi", 0, codes.InvalidArgument},
		{"decimalCapExceededByRounding", &csi.CapacityRange{RequiredBytes: 12000000000}, "12G", 0, codes.InvalidArgument},
		{"roundedCapacityAtCap", &csi.CapacityRange{RequiredBytes: 10*GiB + 1}, "11Gi", 11 * GiB, codes.OK},
		{"wholeGiBBelowFractionalCap", &csi.CapacityRange{RequiredBytes: 10 * GiB}, "10.5Gi", 10 * GiB, codes.OK},
		{"unparsableCap", &csi.CapacityRange{RequiredBytes: 20 * GiB}, "10Gib", 0, codes.InvalidArgument},
		{"emptyCapStringIsIgnored", &csi.CapacityRange{RequiredBytes: 20 * GiB}, "", 20 * GiB, codes.OK},
		{"limitBytesIsHonored", &csi.CapacityRange{RequiredBytes: 20 * GiB, LimitBytes: 10 * GiB}, "", 0, codes.OutOfRange},
		{"negativeCapIsRejected", &csi.CapacityRange{RequiredBytes: 20 * GiB}, "-1Gi", 0, codes.InvalidArgument},
		{"zeroCapIsRejected", &csi.CapacityRange{RequiredBytes: 20 * GiB}, "0", 0, codes.InvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeAgenticSpaceSizeLimit(tt.cr, tt.param)
			if tt.wantCode == codes.OK {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
				return
			}
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

func TestComputeAgenticSpaceFileCountLimit(t *testing.T) {
	got, err := computeAgenticSpaceFileCountLimit("")
	require.NoError(t, err)
	assert.Equal(t, defaultAgenticSpaceFileCountLimit, got)

	got, err = computeAgenticSpaceFileCountLimit("20000")
	require.NoError(t, err)
	assert.Equal(t, int64(20000), got)

	_, err = computeAgenticSpaceFileCountLimit("not-a-number")
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestValidateAgenticSpaceQuota(t *testing.T) {
	tests := []struct {
		name           string
		sizeLimit      int64
		fileCountLimit int64
		wantErr        bool
	}{
		{"valid", 20 * GiB, defaultAgenticSpaceFileCountLimit, false},
		{"minimumSize", minAgenticSpaceSizeLimit, minAgenticSpaceFileCountLimit, false},
		{"maximumSizeAndFileCount", maxAgenticSpaceSizeLimit, maxAgenticSpaceFileCountLimit, false},
		{"sizeBelowMinimum", minAgenticSpaceSizeLimit - 1, defaultAgenticSpaceFileCountLimit, true},
		{"fileCountBelowMinimum", 20 * GiB, minAgenticSpaceFileCountLimit - 1, true},
		// The OpenAPI maximum is 1e8; 1e9 (documented by an earlier draft) is refused.
		{"fileCountAboveMaximum", 20 * GiB, maxAgenticSpaceFileCountLimit + 1, true},
		{"fileCountOf1e9", 20 * GiB, 1000000000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAgenticSpaceQuota(tt.sizeLimit, tt.fileCountLimit)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, codes.InvalidArgument, status.Code(err))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestComputeExpandSizeLimit(t *testing.T) {
	tests := []struct {
		name     string
		cr       *csi.CapacityRange
		want     int64
		wantCode codes.Code
	}{
		{"nilRange", nil, 0, codes.InvalidArgument},
		{"zeroRequiredBytes", &csi.CapacityRange{}, 0, codes.InvalidArgument},
		{"belowTheMinimum", &csi.CapacityRange{RequiredBytes: 5 * GiB}, 0, codes.InvalidArgument},
		{"valid", &csi.CapacityRange{RequiredBytes: 30 * GiB}, 30 * GiB, codes.OK},
		{"roundedUp", &csi.CapacityRange{RequiredBytes: 30*GiB + 1}, 31 * GiB, codes.OK},
		{"overLimitBytes", &csi.CapacityRange{RequiredBytes: 31 * GiB, LimitBytes: 30 * GiB}, 0, codes.OutOfRange},
		{"overTheOpenapiMaximum", &csi.CapacityRange{RequiredBytes: maxAgenticSpaceSizeLimit + GiB}, 0, codes.InvalidArgument},
		{"roundingOverflow", &csi.CapacityRange{RequiredBytes: math.MaxInt64}, 0, codes.OutOfRange},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeExpandSizeLimit(tt.cr)
			if tt.wantCode == codes.OK {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
				return
			}
			require.Error(t, err)
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

func TestValidateVolumeName(t *testing.T) {
	tests := []struct {
		name    string
		volume  string
		wantErr bool
	}{
		{"pvName", testAgenticFsPVName, false},
		{"simple", "vol-1", false},
		{"maxLength", strings.Repeat("a", maxVolumeNameLen), false},
		{"empty", "", true},
		{"tooLong", strings.Repeat("a", maxVolumeNameLen+1), true},
		{"forwardSlash", "a/b", true},
		{"backSlash", `a\b`, true},
		// '/' is already rejected and req.Name is a single first-level segment; only "." and ".." exactly are refused.
		{"embeddedDotDotIsAllowed", "a..b", false},
		{"trailingDotDotIsAllowed", "ab..", false},
		{"singleDot", ".", true},
		{"doubleDot", "..", true},
		{"nonAscii", "café", true},
		{"spaceIsPrintableSoItPasses", "a b", false},
		{"tabIsNotPrintable", "a\tb", true},
		{"controlCharacter", "a\x00b", true},
		// It would produce a ClientToken that a later attempt for the same PVC no longer matches.
		{"leadingSpace", " pvc-x", true},
		{"trailingSpace", "pvc-x ", true},
		{"leadingTab", "\tpvc-x", true},
		{"trailingNewline", "pvc-x\n", true},
		{"onlyWhitespace", "   ", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateVolumeName(tt.volume)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, codes.InvalidArgument, status.Code(err))
				return
			}
			assert.NoError(t, err)
		})
	}
}

// A silent edit here would break mounts or the OpenAPI quota bounds.
func TestAgenticfsConstants(t *testing.T) {
	assert.Equal(t, "fileSystemId", filesystemIDKey)
	assert.Equal(t, "alinas", mountProtocolAlinas)
	assert.Equal(t, "tls,vers=3,ram", defaultAgenticFsMountOptions)
	assert.Equal(t, "Agentic", cloud.StorageTypeAgentic)
	assert.Equal(t, cloud.StorageTypeAgentic, agenticFsVolumeAs)
	assert.Equal(t, cloud.StorageTypeAgentic, cnfsSpecTypeAgenticFS)
	assert.Equal(t, int64(10*GiB), minAgenticSpaceSizeLimit)
	assert.Equal(t, int64(1099511627776000), maxAgenticSpaceSizeLimit)
	assert.Equal(t, int64(10000), minAgenticSpaceFileCountLimit)
	assert.Equal(t, int64(100000000), maxAgenticSpaceFileCountLimit)
	assert.Equal(t, int32(100), apListMaxResults)
	// Cross-component values: a silent edit would go unnoticed if only asserted indirectly.
	assert.Equal(t, 3*time.Second, defaultApPollInterval)
	assert.Equal(t, 45*time.Second, defaultApPollTimeout)
	assert.Less(t, defaultApPollTimeout, 60*time.Second)

	// This bounds the HTTP exchange, not credential resolution or the whole RPC.
	const nasAPICallBound = 10 * time.Second
	sdkClient, err := cloud.NewNasClientV2("cn-hangzhou")
	require.NoError(t, err)
	assert.Equal(t, nasAPICallBound, time.Duration(tea.IntValue(sdkClient.ConnectTimeout))*time.Millisecond)
	assert.Zero(t, tea.IntValue(sdkClient.ReadTimeout))
	assert.Nil(t, sdkClient.RetryOptions)

	// KNOWN DEFICIT: external-nas-resizer has no --timeout, so it inherits upstream's 10s, which EQUALS
	// nasAPICallBound. Equality on purpose: adding --timeout must turn this red, do not weaken it to an inequality.
	const resizerRPCTimeout = 10 * time.Second // external-resizer upstream default; the chart does not override it
	// agenticfs expand = nas:GetAgenticSpace + nas:SetAgenticSpaceQuota, TWO serial NAS calls.
	assert.Equal(t, -10*time.Second, resizerRPCTimeout-2*nasAPICallBound,
		"KNOWN DEFICIT (agenticfs expand): the resizer's 10s budget minus two serial 10s-bounded calls "+
			"is -10s. The chart is deliberately unchanged for now; a follow-up adds --timeout to "+
			"external-nas-resizer, at which point this equality must become a positive-margin assertion")
	// subpath expand (volumeCapacity=true, already GA and running in production) = ONE NAS call.
	assert.Equal(t, time.Duration(0), resizerRPCTimeout-1*nasAPICallBound,
		"KNOWN DEFICIT (subpath expand): the resizer's 10s budget minus one 10s-bounded call is exactly "+
			"0 - no headroom for the endpoint to be even slightly slow. The chart is deliberately unchanged")

	assert.Equal(t, "agenticfs-resource-created", resourceCreatedLogPrefix)
	assert.Equal(t, "agenticfs-resource-retained-for-retry", retainedForRetryLogPrefix)
	assert.Equal(t, "agenticfs-orphan-resource", orphanLogPrefix)
	// Prefixes must remain distinguishable by log consumers.
	prefixContract := []struct{ name, prefix string }{
		{"resourceCreatedLogPrefix", resourceCreatedLogPrefix},
		{"retainedForRetryLogPrefix", retainedForRetryLogPrefix},
		{"orphanLogPrefix", orphanLogPrefix},
	}
	for _, a := range prefixContract {
		for _, b := range prefixContract {
			if a.name == b.name {
				continue
			}
			assert.NotContains(t, a.prefix, b.prefix,
				"%s (%q) must not contain %s (%q) as a substring, or grepping %s would also swallow every %s line",
				a.name, a.prefix, b.name, b.prefix, a.prefix, b.name)
		}
	}
}

func TestNewAgenticfsControllerErrors(t *testing.T) {
	t.Run("regionMetadataUnavailable", func(t *testing.T) {
		// A FakeProvider without values answers ErrUnknownMetadataKey for RegionID.
		_, err := newAgenticfsController(&internal.ControllerConfig{
			Metadata:         &metadata.FakeProvider{},
			NasClientFactory: &fakeNasClientFactory{v2: newFakeNasClientV2()},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "region")
	})
	t.Run("nasClientFactoryFails", func(t *testing.T) {
		_, err := newAgenticfsController(&internal.ControllerConfig{
			Metadata:         testMetadata,
			NasClientFactory: &fakeNasClientFactory{err: errors.New("no credential configured")},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "no credential configured")
	})
}

// A non-nil body with an empty list means "none yet"; a nil body/response is malformed and Aborted on ANY page.
func TestAgenticfsCreateVolumeListEmptyVsMalformedBody(t *testing.T) {
	t.Run("genuineEmptyBodyProceedsAndCreatesAccesspoint", func(t *testing.T) {
		fake := newFakeNasClientV2()
		fake.listPages = []*sdk.ListAccessPointsResponseBody{{}}
		ctrl := newAgenticfsCtrl(t, fake)

		resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		require.NoError(t, err, "an empty (but non-nil) body means the space has no accesspoint yet")
		assert.Len(t, fake.createAccessPointReqs, 1, "an empty list means no accesspoint exists yet, so one is created")
		assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
		assert.Empty(t, fake.deleteAgenticSpaceReqs)
	})

	// Malformed, never "genuinely empty", on ANY page: abort rather than create from a list it could not read.
	malformed := []struct {
		name   string
		mutate func(*fakeNasClientV2)
	}{
		{"nilBodyOnEveryPage", func(fake *fakeNasClientV2) { fake.listNilBody = true }},
		{"nilResponseOnEveryPage", func(fake *fakeNasClientV2) { fake.listNilResp = true }},
		{"nilBodyOnPage1", func(fake *fakeNasClientV2) {
			fake.listPages = []*sdk.ListAccessPointsResponseBody{nil}
		}},
	}
	for _, tt := range malformed {
		t.Run(tt.name+"Aborts", func(t *testing.T) {
			fake := newFakeNasClientV2()
			tt.mutate(fake)
			ctrl := newAgenticfsCtrl(t, fake)

			_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
			require.Error(t, err, "a nil body/response is a malformed answer, not a proven-empty list")
			assert.Equal(t, codes.Aborted, status.Code(err))
			assert.Contains(t, err.Error(), "malformed response (nil body)")
			assert.Empty(t, fake.createAccessPointReqs, "no accesspoint may be created from a list that could not be read")
			assert.Empty(t, fake.deleteAgenticSpaceReqs)
			assert.Empty(t, fake.deleteAccessPointIDs)
		})
	}
}

// The caller's deadline fires while the loop is sleeping, as opposed to the poll timeout.
func TestAgenticfsCreateVolumeContextCancelledDuringPoll(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.apStatuses = []string{"Pending"}
	ctrl := newAgenticfsCtrl(t, fake)
	ctrl.apPollInterval = time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	_, err := ctrl.CreateVolume(ctx, agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Equal(t, codes.DeadlineExceeded, status.Code(err))
	assert.Contains(t, err.Error(), "to become Active")
	assert.Contains(t, err.Error(), context.DeadlineExceeded.Error())
	assert.Positive(t, fake.describeCalls)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs, "A retryable code leaves the created accesspoint for the next attempt")
}

// Failed provisioning is observable without issuing any compensating cloud call.
func TestAgenticfsCreateVolumeThreeTierResourceLogging(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*fakeNasClientV2)
		// Asserted in both directions, so a prefix added to the wrong tier fails.
		wantCreated  bool
		wantRetained bool
		wantOrphan   bool
		// Two failure modes on one log line must still be tellable apart.
		wantOrphanMsg   string
		wantRetainedMsg string
	}{
		{
			name:        "successEmitsCreatedLineAndNothingElse",
			mutate:      nil,
			wantCreated: true,
		},
		{
			name:         "retryablePollTimeoutRetainsResources",
			mutate:       func(fake *fakeNasClientV2) { fake.apStatuses = []string{"Pending"} },
			wantCreated:  true,
			wantRetained: true,
		},
		{
			name:          "terminalDescribeErrorRequiresReconciliation",
			mutate:        func(fake *fakeNasClientV2) { fake.describeErr = aliErr("InvalidParameter.AccessPointId") },
			wantCreated:   true,
			wantOrphan:    true,
			wantOrphanMsg: "no resources deleted",
		},
		{
			// Classified as RETRYABLE codes.Internal (the ClientToken replay recovers the space), so tier-2 retained.
			name:            "ambiguousCreateAgenticSpaceFailureIsRetainedForRetry",
			mutate:          func(fake *fakeNasClientV2) { fake.createAgenticSpaceErr = errors.New("connection reset by peer") },
			wantRetained:    true,
			wantRetainedMsg: "connection reset by peer",
		},
		{
			// RETRYABLE codes.Internal: the space may exist but its ID was never read back, so tier 2.
			name:            "malformedCreateAgenticSpaceResponseIsRetainedForRetry",
			mutate:          func(fake *fakeNasClientV2) { fake.createAgenticSpaceResp = &sdk.CreateAgenticSpaceResponse{} },
			wantRetained:    true,
			wantRetainedMsg: "malformed response (nil body)",
		},
		{
			name: "emptyAgenticSpaceIdIsRetainedForRetry",
			mutate: func(fake *fakeNasClientV2) {
				fake.createAgenticSpaceResp = &sdk.CreateAgenticSpaceResponse{
					Body: &sdk.CreateAgenticSpaceResponseBody{AgenticSpaceId: tea.String("")},
				}
			},
			wantRetained:    true,
			wantRetainedMsg: "empty AgenticSpaceId",
		},
		{
			name:         "throttledCreateAgenticSpaceFailureIsRetainedForRetry",
			mutate:       func(fake *fakeNasClientV2) { fake.createAgenticSpaceErr = aliErr("Throttling.User") },
			wantRetained: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			if tt.mutate != nil {
				tt.mutate(fake)
			}
			ctrl := newAgenticfsCtrl(t, fake)
			logger, ctx := newLogCapture(t)

			_, _ = ctrl.CreateVolume(ctx, agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))

			logs := logText(logger)
			require.NotEmpty(t, logs, "the capture is vacuous - ktesting buffered nothing, so every assertion below would pass on a deleted log line")

			if tt.wantCreated {
				assert.Contains(t, logs, resourceCreatedLogPrefix,
					"tier 1: emitted on EVERY successful CreateAgenticSpace so a reaper can reconcile created-against-delivered")
			} else {
				assert.NotContains(t, logs, resourceCreatedLogPrefix,
					"tier 1 must not claim a space with a known ID exists when the ID was never read back")
			}
			if tt.wantRetained {
				assert.Contains(t, logs, retainedForRetryLogPrefix)
				assert.Contains(t, logs, "NOT an orphan", "the line must say out loud that nothing may be cleaned up by hand")
				if tt.wantRetainedMsg != "" {
					assert.Contains(t, logs, tt.wantRetainedMsg, "this retained path must be distinguishable from the others")
				}
				var retainedLine string
				for _, line := range strings.Split(logs, "\n") {
					if strings.Contains(line, retainedForRetryLogPrefix) {
						retainedLine = line
						break
					}
				}
				require.NotEmpty(t, retainedLine, "the retained-for-retry line must exist")
				for _, field := range orphanLogFields {
					assert.Contains(t, retainedLine, field, "D2: retained-for-retry carries the full contract schema")
				}
			} else {
				assert.NotContains(t, logs, retainedForRetryLogPrefix)
			}
			if tt.wantOrphan {
				assertOrphanLog(t, logs)
				assert.Contains(t, logs, tt.wantOrphanMsg, "this orphan path must be distinguishable from the others")
			} else {
				assert.NotContains(t, logs, orphanLogPrefix,
					"only terminal failures use the reconciliation prefix")
			}
			assert.NotContains(t, logs, "agenticfs-orphan-resolved")
			assert.Empty(t, fake.deleteAccessPointIDs)
			if tt.wantRetained || tt.wantOrphan {
				assertExactlyOneFailureReport(t, logs)
			}
			// The one value that survives every failure mode, including an unknown space ID.
			assert.Contains(t, logs, "/"+testAgenticFsPVName, "fileSystemPath must never be dropped")
			assert.Empty(t, fake.deleteAgenticSpaceReqs, "No path in CreateVolume ever deletes the space")
		})
	}
}

func TestAgenticfsCreateVolumeFailureReleasesLockWithoutCleanup(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeErr = aliErr("InvalidParameter.AccessPointId")
	fake.deleteAccessPointHook = func(context.Context, string, string) error {
		panic("CreateVolume must never call DeleteAccesspoint")
	}
	ctrl := newAgenticfsCtrl(t, fake)
	cs := &controllerServer{
		ControllerFactory: &internal.ControllerFactory{Modes: map[string]internal.Controller{agenticFsVolumeAs: ctrl}},
		locks:             utils.NewVolumeLocks(),
	}
	logger, ctx := newLogCapture(t)
	_, err := cs.CreateVolume(ctx, agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Equal(t, codes.InvalidArgument, status.Code(err))
	require.True(t, cs.locks.TryAcquire(testAgenticFsPVName))
	cs.locks.Release(testAgenticFsPVName)
	assertExactlyOneFailureReport(t, logText(logger))
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

// One trailing newline made external-provisioner delete the PV and leak the billable AgenticSpace.
func TestAgenticfsDeleteVolumePaddedNotFoundCodeIsIdempotent(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{"trailingNewline", aliErr("InvalidAccessPointId.NotFound\n")},
		{"leadingTabAndTrailingSpace", aliErr("\tNotFound ")},
		{"paddedAgenticspaceCode", aliErr("InvalidAgenticSpaceId.NotFound\n")},
		{"paddedFilesystemCode", aliErr("InvalidFileSystem.NotFound\n")},
		{"paddedOnBothEnds", aliErr("  InvalidAccessPointId.NotFound\t")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newDeleteFakeNasClientV2()
			fake.deleteAccessPointErr = tt.err
			// Accesspoint absence is idempotent only for the AP operation.
			fake.deleteAgenticSpaceErr = aliErr("InvalidAgenticSpaceId.NotFound\n")
			ctrl := newAgenticfsCtrl(t, fake)

			resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			require.NoError(t, err, "a padded NotFound code must still be recognised as already-gone")
			require.NotNil(t, resp)
			// The old assertion ran after require.NoError, so status.Code(nil) == OK made it a tautology.
			assert.Equal(t, codes.OK, status.Code(err),
				"a padded NotFound must never be classified permanent: terminal InvalidArgument here makes the provisioner delete the PV and leak the resources")
			assert.Positive(t, fake.describeCalls, "An already-gone accesspoint is still waited for")
			assert.Len(t, fake.deleteAgenticSpaceReqs, 1)
		})
	}
}

// isNotFoundError accepts three resource names, so the message must name what the cloud named.
func TestAgenticfsDeleteVolumeGoneMessageNamesTheActualResource(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantGone string
		notGone  []string
	}{
		{
			name:     "agenticSpaceNotFound",
			err:      aliErr("InvalidAgenticSpaceId.NotFound"),
			wantGone: "agenticspace is gone",
			notGone:  []string{"filesystem is gone", "accesspoint is gone"},
		},
		{
			name:     "filesystemNotFound",
			err:      aliErr("InvalidFileSystem.NotFound"),
			wantGone: "filesystem is gone",
			notGone:  []string{"agenticspace is gone", "accesspoint is gone"},
		},
		{
			// The switch lowercases first, so a lowercase-'s' wire spelling must still be labelled "filesystem".
			name:     "filesystemNotFoundSpelledWithLowercaseS",
			err:      aliErr("InvalidFilesystem.NotFound"),
			wantGone: "filesystem is gone",
			notGone:  []string{"agenticspace is gone", "accesspoint is gone"},
		},
		{
			name:     "paddedFilesystemNotFound",
			err:      aliErr("InvalidFileSystem.NotFound\n"),
			wantGone: "filesystem is gone",
			notGone:  []string{"agenticspace is gone"},
		},
		{
			name:     "paddedAgenticSpaceNotFound",
			err:      aliErr("\tInvalidAgenticSpaceId.NotFound\n"),
			wantGone: "agenticspace is gone",
			notGone:  []string{"filesystem is gone", "accesspoint is gone"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newDeleteFakeNasClientV2()
			fake.listAccessPointsErr = tt.err
			ctrl := newAgenticfsCtrl(t, fake)
			logger, ctx := newLogCapture(t)

			resp, err := ctrl.DeleteVolume(ctx, agenticfsDeleteReq(), agenticfsDeletePV())
			require.NoError(t, err)
			require.NotNil(t, resp)

			logs := logText(logger)
			require.NotEmpty(t, logs)
			assert.Contains(t, logs, tt.wantGone)
			for _, notWant := range tt.notGone {
				assert.NotContains(t, logs, notWant, "the message must name the resource the cloud actually named")
			}
			assert.Contains(t, logs, "errorCode", "the raw code is always logged next to the interpretation")
			assert.Empty(t, fake.getAgenticSpaceReqs, "explicit parent absence needs no verification call")
			assert.Empty(t, fake.deleteAccessPointIDs)
			assert.Empty(t, fake.deleteAgenticSpaceReqs)
		})
	}
}

func TestAgenticfsCreateVolumeUnclassifiedFailureIsVisible(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	logger, _ := newLogCapture(t)
	cause := errors.New("a bare error that never went through status.Errorf")
	require.Equal(t, codes.Unknown, status.Code(cause))
	ctrl.reportCreateVolumeFailure(logger, testAgenticFsFilesystemID, testAgenticFsAgenticSpaceID,
		"/"+testAgenticFsPVName, testAgenticFsAccessPointID, cause)

	logs := logText(logger)
	require.NotEmpty(t, logs, "the capture is vacuous - ktesting buffered nothing")
	assert.Equal(t, 1, countLogLines(logs, retainedForRetryLogPrefix), "one retained outcome, no separate diagnostic needed")
	var retainedLine string
	for _, line := range strings.Split(logs, "\n") {
		if strings.Contains(line, retainedForRetryLogPrefix) {
			retainedLine = line
			break
		}
	}
	require.NotEmpty(t, retainedLine)
	for _, field := range orphanLogFields {
		assert.Contains(t, retainedLine, field)
	}
	assert.Contains(t, retainedLine, "code")
	assert.Contains(t, retainedLine, "Unknown", "unclassified codes must remain visible")
	assert.Contains(t, retainedLine, cause.Error(), "preserve the original failure")
	// "Treating as retryable" means tier 2, never tier 3, and nothing is destroyed.
	assert.Contains(t, logs, retainedForRetryLogPrefix)
	assert.NotContains(t, logs, orphanLogPrefix, "an unclassified code is not a confirmed leak")
	assert.Empty(t, fake.deleteAccessPointIDs, "a retryable classification must not delete the accesspoint")
	assert.Empty(t, fake.deleteAgenticSpaceReqs, "failure reporting never deletes the agenticspace")
}

// A terminal retry does not disprove creation by a previous, unacknowledged call.
func TestAgenticfsCreateVolumeLostSpaceResponseThenTerminalRetryKeepsStateUnknown(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)
	fake.createAgenticSpaceErr = errors.New("response lost after server accepted creation")
	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Equal(t, codes.Internal, status.Code(err))

	fake.createAgenticSpaceErr = aliErr("Forbidden.RAM")
	logger, ctx := newLogCapture(t)
	_, err = ctrl.CreateVolume(ctx, agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Equal(t, codes.InvalidArgument, status.Code(err))
	logs := logText(logger)
	assertExactlyOneFailureReport(t, logs)
	assertOrphanLog(t, logs)
	assert.Contains(t, logs, "resourceState")
	assert.Contains(t, logs, "unknown")
	assert.NotContains(t, logs, "nothing to reap")
	assert.NotContains(t, logs, "no AgenticSpace in existence")
	require.Len(t, fake.createAgenticSpaceReqs, 2)
	for _, req := range fake.createAgenticSpaceReqs {
		assert.Equal(t, testAgenticFsPVName, tea.StringValue(req.ClientToken))
	}
	assert.Empty(t, fake.deleteAccessPointIDs)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
}

func TestAgenticfsOrphanLogFieldsPreserveVolumeHandle(t *testing.T) {
	ctrl := newAgenticfsCtrl(t, newFakeNasClientV2())
	assert.Equal(t, []any{
		"fileSystemId", testAgenticFsFilesystemID,
		"agenticSpaceId", testAgenticFsAgenticSpaceID,
		"fileSystemPath", "/" + testAgenticFsPVName + "/",
		"accesspointId", testAgenticFsAccessPointID,
		"region", ctrl.region,
		"volumeHandle", testAgenticFsPVName,
	}, ctrl.orphanLogFields(testAgenticFsFilesystemID, testAgenticFsAgenticSpaceID,
		"/"+testAgenticFsPVName+"/", testAgenticFsAccessPointID))
}

// The production limiter exposes a typed contract, independent of message wording.
func TestNasRateLimiterWaitErrorMatchesTheProductionClient(t *testing.T) {
	client, err := cloud.NewNasClientFactory().V2("cn-hangzhou")
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), -time.Second) // already expired
	defer cancel()
	require.Error(t, ctx.Err(), "the premise: the ctx is already done so limiter.Wait fails at once")

	err = client.DeleteAccesspoint(ctx, testAgenticFsFilesystemID, testAgenticFsAccessPointID)
	require.Error(t, err, "limiter.Wait on an expired ctx must fail before any wire request is built")
	assert.ErrorIs(t, err, cloud.ErrRateLimiterWait)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.ErrorIs(t, fmt.Errorf("outer wrapper: %w", err), cloud.ErrRateLimiterWait)
	assert.NotErrorIs(t, errors.New(cloud.ErrRateLimiterWait.Error()), cloud.ErrRateLimiterWait,
		"message text alone must not classify an SDK error as never sent")
}

// The controller completes missing vers/tls/ram after parameters.options
// overwrites VolumeContext, without replacing values the caller supplied.
func TestEnforceAgenticFsMountOptionsRestoresForcedOptions(t *testing.T) {
	tests := []struct {
		name    string
		options string
		want    []string // substrings the merged options must contain
		wantRaw string   // if set, the result must equal this EXACTLY (byte-identical no-op)
	}{
		{
			name:    "storageClassOptionsDropTlsAndRamBothAreRestored",
			options: "vers=3,noresvport",
			want:    []string{"tls", "ram", "vers=3", "noresvport"},
		},
		{
			name:    "defaultPathIsByteIdentical",
			options: defaultAgenticFsMountOptions,
			wantRaw: defaultAgenticFsMountOptions,
		},
		{
			name:    "emptyOptionsCompleted",
			options: "",
			wantRaw: defaultAgenticFsMountOptions,
		},
		{
			name:    "explicitVersionIsNotDefaulted",
			options: "vers=4.1,noresvport",
			wantRaw: "vers=4.1,noresvport,tls,ram",
		},
		{
			name:    "onlyMissingVersionIsAdded",
			options: "tls,ram",
			wantRaw: "tls,ram,vers=3",
		},
		{
			name:    "allExplicitValuesPreserved",
			options: "vers=4,tls=custom,ram=custom",
			wantRaw: "vers=4,tls=custom,ram=custom",
		},
		{
			name:    "noneIsCompletedBeforePVCreation",
			options: "none",
			wantRaw: defaultAgenticFsMountOptions,
		},
		{
			name:    "explicitTlsValueSetByCallerIsPreservedRamStillAdded",
			options: "tls=custom,vers=3",
			want:    []string{"tls=custom", "ram", "vers=3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := map[string]string{vcKeyOptions: tt.options}
			enforceAgenticFsMountOptions(agenticFsVolumeAs, vc)
			got := vc[vcKeyOptions]
			if tt.wantRaw != "" {
				assert.Equal(t, tt.wantRaw, got, "the default path must not be reordered or normalised")
				return
			}
			require.NotEmpty(t, got, "the merged options must not be empty")
			for _, w := range tt.want {
				assert.Contains(t, got, w, "A mandatory mount option must survive the overwrite")
			}
		})
	}
}

// For every mode but agenticfs the function returns on its first statement, so the VolumeContext is unchanged.
func TestEnforceAgenticFsMountOptionsIsANoOpForOtherVolumeModes(t *testing.T) {
	// The guard is an exact match, so a differently-cased value must be left untouched.
	for _, tc := range []struct{ name, volumeAs string }{
		{"volumeAsSubpath", "subpath"},
		{"volumeAsSharepath", "sharepath"},
		{"volumeAsAccesspoint", "accesspoint"},
		{"volumeAsFilesystem", "filesystem"},
		{"volumeAsEmpty", ""},
		{"volumeAsWrongCaseAgentic", "agentic"},
		{"volumeAsOldSpellingAgenticFS", "agenticfs"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			vc := map[string]string{vcKeyOptions: "vers=3,noresvport", vcKeyServer: "s", vcKeyPath: "/"}
			before := map[string]string{}
			for k, v := range vc {
				before[k] = v
			}
			enforceAgenticFsMountOptions(tc.volumeAs, vc)
			assert.Equal(t, before, vc,
				"a non-agenticfs mode must see enforceAgenticFsMountOptions as a byte-identical no-op")
		})
	}
	assert.NotPanics(t, func() { enforceAgenticFsMountOptions(agenticFsVolumeAs, nil) },
		"a nil VolumeContext must not panic")
}

// --- Tests for discoverAgenticSpaceByPath (Path already used fallback) ---

func TestAgenticfsCreateVolumePathAlreadyUsedDiscoversExistingSpace(t *testing.T) {
	fake := newFakeNasClientV2()
	// CreateAgenticSpace returns "Path already used", simulating the NAS API's non-idempotent behavior.
	fake.createAgenticSpaceErr = errors.New("InvalidArgument: Path already used")
	// ListAccesspoints returns an accesspoint whose RootPath matches the volume path.
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		{AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{
			{
				AccessPointId:  tea.String(testAgenticFsAccessPointID),
				AgenticSpaceId: tea.String(testAgenticFsAgenticSpaceID),
				DomainName:     tea.String(testAgenticFsAPDomain),
				Status:         tea.String(accessPointStatusActive),
				RootPath:       tea.String("/" + testAgenticFsPVName + "/"),
			},
		}},
	}
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, testAgenticFsAPDomain, resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, testAgenticFsAccessPointID, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Equal(t, testAgenticFsAgenticSpaceID, resp.Volume.VolumeContext[vcKeyAgenticSpaceId])
	// CreateAgenticSpace was called once (returned error), then discovery via ListAccesspoints happened twice:
	// once inside discoverAgenticSpaceByPath (unfiltered), once inside findReusableAccessPoint (filtered by space).
	assert.Len(t, fake.createAgenticSpaceReqs, 1)
	assert.GreaterOrEqual(t, len(fake.listAccessPointsReqs), 2,
		"ListAccesspoints must be called by discoverAgenticSpaceByPath and then again by findReusableAccessPoint")
	assert.Empty(t, fake.createAccessPointReqs, "existing accesspoint must be reused")
}

func TestAgenticfsCreateVolumePathAlreadyUsedNoMatchingSpace(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAgenticSpaceErr = errors.New("InvalidArgument: Path already used")
	// ListAccesspoints returns an accesspoint with a different path.
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		{AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{
			{
				AccessPointId:  tea.String("ap-other"),
				AgenticSpaceId: tea.String("as-other"),
				DomainName:     tea.String("other.example.com"),
				Status:         tea.String(accessPointStatusActive),
				RootPath:       tea.String("/some-other-path/"),
			},
		}},
	}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Path already used")
	assert.Contains(t, err.Error(), "no matching AgenticSpace found")
}

func TestAgenticfsCreateVolumePathAlreadyUsedListError(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAgenticSpaceErr = errors.New("InvalidArgument: Path already used")
	fake.listAccessPointsErr = errors.New("connection reset")
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "discovering existing AgenticSpace")
}

func TestAgenticfsCreateVolumePathAlreadyUsedTrailingSlashVariants(t *testing.T) {
	// The NAS API may return a path with or without a trailing slash; the match must be flexible.
	for _, tt := range []struct {
		name       string
		apiPath    string
		searchPath string
	}{
		{"exactMatch", "/pv-name/", "/pv-name/"},
		{"apiHasNoTrailingSlash", "/pv-name", "/pv-name/"},
		{"searchHasNoTrailingSlash", "/pv-name/", "/pv-name"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.createAgenticSpaceErr = errors.New("InvalidArgument: Path already used")
			fake.listPages = []*sdk.ListAccessPointsResponseBody{
				{AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{
					{
						AccessPointId:  tea.String(testAgenticFsAccessPointID),
						AgenticSpaceId: tea.String(testAgenticFsAgenticSpaceID),
						RootPath:       tea.String(tt.apiPath),
					},
				}},
			}
			ctrl := newAgenticfsCtrl(t, fake)
			spaceID, err := ctrl.discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, tt.searchPath)
			require.NoError(t, err)
			assert.Equal(t, testAgenticFsAgenticSpaceID, spaceID)
		})
	}
}

func TestAgenticfsCreateVolumePathAlreadyUsedEmptyAgenticSpaceIdSkipped(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAgenticSpaceErr = errors.New("InvalidArgument: Path already used")
	fake.listPages = []*sdk.ListAccessPointsResponseBody{
		{AccessPoints: []*sdk.ListAccessPointsResponseBodyAccessPoints{
			{
				AccessPointId:  tea.String("ap-empty-space"),
				AgenticSpaceId: tea.String(""),
				RootPath:       tea.String("/" + testAgenticFsPVName + "/"),
			},
		}},
	}
	ctrl := newAgenticfsCtrl(t, fake)

	_, err := ctrl.CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no matching AgenticSpace found")
}

func TestAgenticfsCreateVolumeServerPresentPassthrough(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	req := &csi.CreateVolumeRequest{
		Name:          testAgenticFsPVName,
		CapacityRange: &csi.CapacityRange{RequiredBytes: 10 * GiB},
		Parameters: map[string]string{
			"volumeAs": agenticFsVolumeAs,
			"server":   "ap-existing.nas.aliyuncs.com",
			"path":     "/user-data",
			"options":  "tls,vers=3,ram",
		},
	}

	resp, err := ctrl.CreateVolume(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "ap-existing.nas.aliyuncs.com", resp.Volume.VolumeContext[vcKeyServer])
	assert.Equal(t, "/user-data", resp.Volume.VolumeContext[vcKeyPath])
	assert.Equal(t, mountProtocolAlinas, resp.Volume.VolumeContext[vcKeyMountProtocol])
	assert.Empty(t, resp.Volume.VolumeContext[vcKeyAgenticSpaceId])
	assert.Empty(t, resp.Volume.VolumeContext[vcKeyAccesspointId])
	assert.Empty(t, fake.createAgenticSpaceReqs, "no NAS API calls when server is provided")
	assert.Empty(t, fake.createAccessPointReqs, "no NAS API calls when server is provided")
}

func TestAgenticfsCreateVolumeServerPresentDefaultPath(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	req := &csi.CreateVolumeRequest{
		Name:          testAgenticFsPVName,
		CapacityRange: &csi.CapacityRange{RequiredBytes: 10 * GiB},
		Parameters: map[string]string{
			"volumeAs": agenticFsVolumeAs,
			"server":   "ap-existing.nas.aliyuncs.com",
		},
	}

	resp, err := ctrl.CreateVolume(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, "/", resp.Volume.VolumeContext[vcKeyPath])
}

func TestAgenticfsCreateVolumeServerPresentWithFilesystemID(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	req := &csi.CreateVolumeRequest{
		Name:          testAgenticFsPVName,
		CapacityRange: &csi.CapacityRange{RequiredBytes: 10 * GiB},
		Parameters: map[string]string{
			"volumeAs":      agenticFsVolumeAs,
			"server":        "ap-existing.nas.aliyuncs.com",
			filesystemIDKey: "011wozfr3hjd4jxhdbz",
		},
	}

	resp, err := ctrl.CreateVolume(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, "011wozfr3hjd4jxhdbz", resp.Volume.VolumeContext[filesystemIDKey])
}

func TestAgenticfsDeleteVolumeServerPresentSkipsDeletion(t *testing.T) {
	fake := newFakeNasClientV2()
	ctrl := newAgenticfsCtrl(t, fake)

	pv := &corev1.PersistentVolume{
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				CSI: &corev1.CSIPersistentVolumeSource{
					VolumeAttributes: map[string]string{
						vcKeyServer:        "ap-existing.nas.aliyuncs.com",
						vcKeyPath:          "/user-data",
						vcKeyMountProtocol: mountProtocolAlinas,
					},
				},
			},
		},
	}

	resp, err := ctrl.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: testAgenticFsPVName}, pv)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Empty(t, fake.deleteAccessPointIDs, "no NAS API calls for passthrough volume deletion")
	assert.Empty(t, fake.deleteAgenticSpaceReqs, "no NAS API calls for passthrough volume deletion")
}

func TestAgenticfsDeleteVolumeCreatedModeStillDeletes(t *testing.T) {
	fake := newFakeNasClientV2()
	// After DeleteAccessPoint is called, DescribeAccessPoint should return NotFound
	// so waitAccessPointDeleted succeeds.
	fake.deleteAccessPointHook = func(_ context.Context, _, _ string) error {
		fake.apStatuses = []string{fakeStatusNotFound}
		return nil
	}
	ctrl := newAgenticfsCtrl(t, fake)

	pv := &corev1.PersistentVolume{
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				CSI: &corev1.CSIPersistentVolumeSource{
					VolumeAttributes: map[string]string{
						vcKeyServer:         testAgenticFsAPDomain,
						vcKeyAgenticSpaceId: testAgenticFsAgenticSpaceID,
						vcKeyAccesspointId:  testAgenticFsAccessPointID,
						filesystemIDKey:     testAgenticFsFilesystemID,
					},
				},
			},
		},
	}

	_, err := ctrl.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: testAgenticFsPVName}, pv)
	require.NoError(t, err)
	assert.NotEmpty(t, fake.deleteAccessPointIDs, "created-mode volumes must be cleaned up")
}
