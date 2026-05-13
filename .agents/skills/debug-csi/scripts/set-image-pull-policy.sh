#!/bin/bash
set -e

kubectl -n kube-system patch deploy/csi-provisioner -p '
spec:
  template:
    spec:
      containers:
      - name: csi-provisioner
        imagePullPolicy: Always
'

kubectl -n kube-system patch daemonset/csi-plugin -p '
spec:
  template:
    spec:
      containers:
      - name: csi-plugin
        imagePullPolicy: Always
      initContainers:
      - name: init
        imagePullPolicy: Always
'
