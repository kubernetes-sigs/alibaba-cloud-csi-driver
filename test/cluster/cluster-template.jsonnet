local clusterName = std.extVar("cluster_name");
local vswitchIDs = std.extVar("vswitch_ids");
local nNodes = std.extVar("n_nodes");
local nodePool = {
    nodepool_info: {
        name: "default-nodepool"
    },
    scaling_group: {
        vswitch_ids: vswitchIDs,
        system_disk_category: "cloud_essd",
        system_disk_size: 60,
        system_disk_performance_level: "PL0",
        system_disk_encrypted: false,
        instance_types: [
            "ecs.u1-c1m2.xlarge"
        ],
        instance_charge_type: "PostPaid",
        key_pair: clusterName,
        soc_enabled: false,
        cis_enabled: false,
        login_as_non_root: false,
        platform: "AliyunLinux",
        image_id: std.extVar("os_image_alinux3"),
        image_type: "AliyunLinux3"
    },
    kubernetes_config: {
        cpu_policy: "none",
        cms_enabled: false,
        unschedulable: false,
        runtime: "containerd",
        runtime_version: "2.1.3",
    },
    management: {
        enable: false
    },
    count: 1
};

{
    name: clusterName,
    cluster_type: "ManagedKubernetes",
    kubernetes_version: "1.33.3-aliyun.1",
    region_id: std.extVar("region"),
    snat_entry: true,
    cloud_monitor_flags: false,
    endpoint_public_access: true,
    deletion_protection: false,
    node_cidr_mask: "26",
    proxy_mode: "ipvs",
    cis_enable_risk_check: false,
    tags: [
        {
            key: "created-for",
            value: "csi-daily"
        }
    ],
    timezone: "Asia/Shanghai",
    addons: [
        {
            name: "csi-plugin",
            disabled: true
        },
        {
            name: "csi-provisioner",
            disabled: true
        },
        {
            name: "storage-operator",
            disabled: true
        },
        {
            name: "nginx-ingress-controller",
            disabled: true
        }
    ],
    cluster_spec: "ack.pro.small",
    load_balancer_spec: "slb.s1.small",
    charge_type: "PostPaid",
    vpcid: std.extVar("vpc_id"),
    container_cidr: "10.236.0.0/16",
    service_cidr: "192.168.0.0/16",
    vswitch_ids: vswitchIDs,
    ip_stack: "ipv4",
    maintenance_window: {
        enable: false
    },
    nodepools: [
        nodePool {
            count: nNodes - 3
        },
        nodePool {
            nodepool_info+: {
                name: "ssd_only"
            },
            scaling_group+: {
                vswitch_ids: vswitchIDs[1:] + vswitchIDs[:1], // try different available zones
                system_disk_category: "cloud_efficiency",
                system_disk_performance_level: "",
                instance_types: [
                    "ecs.mn4.xlarge",
                ]
            }
        },
        nodePool {
            count: 2,
            nodepool_info+: {
                name: "eed"
            },
            scaling_group+: {
                instance_types: [
                    "ecs.g8y.xlarge",
                ],
                image_id: std.extVar("os_image_alinux3_arm64"),
                image_type: "AliyunLinux3Arm64",
                data_disks: [ // add an EED data disk to ensure the zone supports EED
                    {
                        category: "elastic_ephemeral_disk_standard",
                        size: 64,
                        auto_format: true,
                        file_system: "ext4",
                        mount_target: "/mnt"
                    }
                ],
            }
        },
    ]
}
