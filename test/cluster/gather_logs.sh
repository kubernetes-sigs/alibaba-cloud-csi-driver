#!/bin/bash

set -e
OUTPUT_DIR=${OUTPUT_DIR:-/tmp}

function covdata() {
    nodes=$(shopt -s nullglob; echo $OUTPUT_DIR/gocover/* | tr ' ' ,)
    if [ -z "$nodes" ]; then
        echo "no coverage data found"
        return 0
    fi
    go tool covdata func -i="$nodes" | tee $OUTPUT_DIR/gocover/cover.txt | tail -n 1
    go tool covdata textfmt -i="$nodes" -o $OUTPUT_DIR/gocover/merged.txt
    go tool cover -html $OUTPUT_DIR/gocover/merged.txt -o $OUTPUT_DIR/cover.html

    rm $OUTPUT_DIR/gocover/merged.txt
    tar -C $OUTPUT_DIR -czf $OUTPUT_DIR/cover.tar.gz gocover
    rm -r $OUTPUT_DIR/gocover
}

if [ -f ~/.ssh/id_rsa ]; then
    mkdir -p $OUTPUT_DIR/gocover
    NODES=$(kubectl get nodes -o json)
    for node in $(echo $NODES | jq -r '.items[].metadata.name'); do
        PRIVATE_IP=$(echo $NODES | jq -r ".items[] | select(.metadata.name==\"$node\") | .status.addresses[] | select(.type==\"InternalIP\") | .address")
        echo "Gathering logs of node $node ($PRIVATE_IP)"
        ssh -o StrictHostKeyChecking=no root@$PRIVATE_IP 'journalctl -o export --since "8 hours ago" | gzip' > $OUTPUT_DIR/$node.journal.gz
        ssh -o StrictHostKeyChecking=no root@$PRIVATE_IP 'cd /var/log/containers && shopt -s nullglob && tar --dereference -czf - {csi,storage,terway}-*' > $OUTPUT_DIR/$node-containers.tar.gz
        scp -o StrictHostKeyChecking=no -r root@$PRIVATE_IP:/opt/csi/gocover $OUTPUT_DIR/gocover/$node || true
    done
    covdata
else
    echo "No SSH key found, skipping gathering logs of nodes"
fi
