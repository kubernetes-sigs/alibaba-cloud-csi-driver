# CSI Metric

## Overview

CSI persistent volume metric is enabled by default in latest version.

Only EBS/NAS is supported to get metric.

## Requirements

Kubernetes should be 1.15 or newer.

## Usage:

You can use "curl" command to test the metrics feature.

Exec the command on the volume mounted node. For example:

```
# curl -s localhost:10255/metrics | grep disk-pvc
kubelet_volume_stats_available_bytes{namespace="default",persistentvolumeclaim="disk-pvc"} 2.6225094656e+10
kubelet_volume_stats_capacity_bytes{namespace="default",persistentvolumeclaim="disk-pvc"} 2.6288033792e+10
kubelet_volume_stats_inodes{namespace="default",persistentvolumeclaim="disk-pvc"} 1.6384e+06
kubelet_volume_stats_inodes_free{namespace="default",persistentvolumeclaim="disk-pvc"} 1.638389e+06
kubelet_volume_stats_inodes_used{namespace="default",persistentvolumeclaim="disk-pvc"} 11
kubelet_volume_stats_used_bytes{namespace="default",persistentvolumeclaim="disk-pvc"} 4.616192e+07
```

Also, you can check volume metric on prometheus.