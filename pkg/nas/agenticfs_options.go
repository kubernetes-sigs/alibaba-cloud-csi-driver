//go:build !windows

package nas

import (
	"context"
	"strings"
	"unicode"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

const (
	// Shared contract with the CNFS controller and NAS StorageType.
	agenticFsVolumeAs     = cloud.StorageTypeAgentic
	cnfsSpecTypeAgenticFS = cloud.StorageTypeAgentic

	paramContainerNetworkFileSystem = "containerNetworkFileSystem"
	paramAgenticSpaceSizeLimit      = "agenticSpaceSizeLimit"
	paramAgenticSpaceFileCountLimit = "agenticSpaceFileCountLimit"

	vcKeyServer         = "server"
	vcKeyPath           = "path"
	vcKeyMountProtocol  = "mountProtocol"
	vcKeyOptions        = "options"
	vcKeyAccesspointId  = "accesspointId"
	vcKeyAgenticSpaceId = "agenticSpaceId"

	// Read-only compatibility for an earlier build. Write filesystemIDKey instead.
	vcKeyFilesystemIdLegacy = "filesystemId"
	mountProtocolAlinas     = "alinas"

	defaultAgenticFsMountOptions = "tls,vers=3,ram"
	maxVolumeNameLen             = 64
)

// Like nasVolumeArgs/diskVolumeArgs, this is a validated request snapshot, not
// mutable controller state. Cloud operations need not re-read untyped maps.
type agenticfsVolumeArgs struct {
	Name             string
	FileSystemID     string
	FileSystemPath   string
	ZoneID           string
	VpcID            string
	VSwitchID        string
	SizeLimit        int64
	FileCountLimit   int64
	MountOptions     string
	AuthType         string
	CredProviderName string
	SubstrateMode    bool
}

// Preserve validation order: name, filesystem/CNFS resolution, placement, then
// quota. In particular fileSystemId bypasses CNFS, and all validation finishes
// before the first billable NAS call.
func (c *agenticfsController) getAgenticfsVolumeOptions(ctx context.Context, req *csi.CreateVolumeRequest) (*agenticfsVolumeArgs, error) {
	if err := validateVolumeName(req.Name); err != nil {
		return nil, err
	}
	parameters := req.Parameters
	filesystemID := parameters[filesystemIDKey]
	if filesystemID == "" {
		var err error
		filesystemID, err = c.filesystemIDFromCNFS(ctx, parameters[paramContainerNetworkFileSystem])
		if err != nil {
			return nil, err
		}
	}

	zoneID := parameters[ZoneID]
	if zoneID == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAgenticSpace.Azone", ZoneID)
	}
	vpcID := parameters[VpcID]
	if vpcID == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAccessPoint.VpcId", VpcID)
	}
	vswID := parameters[VSwitchID]
	if vswID == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAccessPoint.VswId", VSwitchID)
	}

	sizeLimit, err := computeAgenticSpaceSizeLimit(req.GetCapacityRange(), parameters[paramAgenticSpaceSizeLimit])
	if err != nil {
		return nil, err
	}
	fileCountLimit, err := computeAgenticSpaceFileCountLimit(parameters[paramAgenticSpaceFileCountLimit])
	if err != nil {
		return nil, err
	}
	if err := validateAgenticSpaceQuota(sizeLimit, fileCountLimit); err != nil {
		return nil, err
	}
	return &agenticfsVolumeArgs{
		Name:             req.Name,
		FileSystemID:     filesystemID,
		FileSystemPath:   "/" + req.Name + "/",
		ZoneID:           zoneID,
		VpcID:            vpcID,
		VSwitchID:        vswID,
		SizeLimit:        sizeLimit,
		FileCountLimit:   fileCountLimit,
		MountOptions:     parameters[vcKeyOptions],
		AuthType:         parameters[jwtauth.OptAuthType],
		CredProviderName: parameters[jwtauth.OptSandboxCredProviderName],
		SubstrateMode:    common.IsSubstrateVolumeContext(parameters),
	}, nil
}

// volume builds the CSI response after backend provisioning; it makes no API calls.
func (args *agenticfsVolumeArgs) volume(agenticSpaceID, accesspointID, server string) *csi.Volume {
	vc := map[string]string{
		vcKeyServer:         server,
		vcKeyPath:           "/",
		vcKeyMountProtocol:  mountProtocolAlinas,
		vcKeyOptions:        mergeAgenticFsMountOptions(args.MountOptions, defaultAgenticFsMountOptions),
		filesystemIDKey:     args.FileSystemID,
		vcKeyAccesspointId:  accesspointID,
		vcKeyAgenticSpaceId: agenticSpaceID,
	}
	if args.AuthType != "" {
		// NodePublishVolume needs the identity selection; atelet supplies the actor UID later.
		vc[jwtauth.OptAuthType] = args.AuthType
	}
	if args.CredProviderName != "" {
		vc[jwtauth.OptSandboxCredProviderName] = args.CredProviderName
	}
	volumeID := args.Name
	if args.SubstrateMode {
		volumeID = substrateAgenticVolumeID(args.FileSystemID, agenticSpaceID)
	}
	return &csi.Volume{
		VolumeId:      volumeID,
		CapacityBytes: args.SizeLimit,
		VolumeContext: vc,
	}
}

// Resolves the filesystem a CNFS created, for a StorageClass that does not name one with fileSystemId.
func (c *agenticfsController) filesystemIDFromCNFS(ctx context.Context, cnfsName string) (string, error) {
	if cnfsName == "" {
		return "", status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s and parameters.%s are both empty; one of them must name the AgenticFS filesystem",
			filesystemIDKey, paramContainerNetworkFileSystem)
	}
	cnfs, err := c.cnfsGetter.GetCNFS(ctx, cnfsName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return "", status.Errorf(codes.InvalidArgument, "CNFS not found: %s", cnfsName)
		}
		return "", status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
	}

	// AgenticFS is distinguished by StorageType only; FilesystemType is "standard".
	if cnfs.Spec.StorageType != cnfsSpecTypeAgenticFS {
		return "", status.Errorf(codes.InvalidArgument,
			"CNFS %s spec.type is %q, expected %q", cnfsName, cnfs.Spec.StorageType, cnfsSpecTypeAgenticFS)
	}
	if cnfs.Status.FsAttributes.StorageType != cloud.StorageTypeAgentic {
		return "", status.Errorf(codes.InvalidArgument,
			"CNFS %s status.fsAttributes.storageType is %q, expected %q (AgenticFS not ready or mismatched)",
			cnfsName, cnfs.Status.FsAttributes.StorageType, cloud.StorageTypeAgentic)
	}
	filesystemID := cnfs.Status.FsAttributes.FilesystemID
	if filesystemID == "" {
		return "", status.Errorf(codes.InvalidArgument, "CNFS %s status.fsAttributes.filesystemId is empty", cnfsName)
	}
	return filesystemID, nil
}

// Completes missing options after controllerserver.go has overwritten "options".
// Caller-supplied values always win. Other volume modes are unchanged.
func enforceAgenticFsMountOptions(volumeAs string, volumeContext map[string]string) {
	if volumeAs != agenticFsVolumeAs || volumeContext == nil {
		return
	}
	current := volumeContext[vcKeyOptions]
	merged := mergeAgenticFsMountOptions(current, defaultAgenticFsMountOptions)
	if merged != current {
		volumeContext[vcKeyOptions] = merged
	}
}

func mergeAgenticFsMountOptions(options, forced string) string {
	base := mounterutils.SplitMountOptions(options)
	// "none" suppresses optional defaults, not the AgenticFS required options.
	if len(base) == 1 && strings.EqualFold(base[0], "none") {
		base = nil
	}
	merged := mounterutils.MergeMountOptions(base, mounterutils.SplitMountOptions(forced))
	return strings.Join(merged, ",")
}

// Guards the name shared by FileSystemPath, ClientToken and AccessPointName.
func validateVolumeName(name string) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "volume name is empty")
	}
	if strings.TrimSpace(name) != name {
		return status.Errorf(codes.InvalidArgument,
			"volume name %q must not have leading or trailing whitespace", name)
	}
	if len(name) > maxVolumeNameLen {
		return status.Errorf(codes.InvalidArgument,
			"volume name %q is longer than %d bytes", name, maxVolumeNameLen)
	}
	if strings.ContainsAny(name, `/\`) {
		return status.Errorf(codes.InvalidArgument, "volume name %q must not contain '/' or '\\'", name)
	}
	if name == "." || name == ".." {
		return status.Errorf(codes.InvalidArgument, "volume name %q must not be '.' or '..'", name)
	}
	for _, r := range name {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return status.Errorf(codes.InvalidArgument, "volume name %q must only contain printable ASCII characters", name)
		}
	}
	return nil
}

func agenticfsFilesystemID(attributes map[string]string) string {
	if id := attributes[filesystemIDKey]; id != "" {
		return id
	}
	return attributes[vcKeyFilesystemIdLegacy]
}
