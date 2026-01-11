package disk

var (
	// DiskXattrName xattr is applied on the block device file to indicate that it is managed by the CSI driver.
	// Value is the disk ID.
	// Linux only support trusted namespace xattr in tmpfs before kernel v6.6,
	// but setting trusted xaattr requires CAP_SYS_ADMIN capability, we may use user namespace instead in unit tests.
	DiskXattrName = "trusted.csi-managed-disk"

	// DiskXattrVirtioBlkName xattr is applied on the block device file to indicate that it is managed by the CSI driver in PVM ways.
	DiskXattrVirtioBlkName = "trusted.virtio-blk"
)
