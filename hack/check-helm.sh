#!/bin/bash
# We need a clean Kubernetes cluster to run this script.

set -e

cd deploy/chart

kubectl create -f ./crds

helm lint
helm install alibaba-cloud-csi-driver . --namespace kube-system --dry-run=server

for v in values-*.yaml; do
    helm lint --values "$v"
    helm install alibaba-cloud-csi-driver . --namespace kube-system --dry-run=server --values "$v"
done

kubectl delete -f ./crds
