#!/bin/bash

set -e

N_NODES=${N_NODES:-3}
ACK_REGION=${ACK_REGION:-cn-beijing}
ACK_ZONE=${ACK_ZONE:-cn-beijing-h cn-beijing-i cn-beijing-j}

function cluster-setup {
    local HERE
    local CASE_NAME=$1
    HERE=$(dirname "${BASH_SOURCE[0]}")

    echo "Creating VPC"
    VPC_ID=$(aliyun vpc CreateVpc --RegionId "$ACK_REGION" \
        --CidrBlock '172.16.0.0/12' \
        --VpcName "$CASE_NAME" \
        --ClientToken "$CASE_NAME-vpc" \
        --Tag.1.Key 'created-for' --Tag.1.Value "$CASE_NAME" | jq -r .VpcId)

    if [ -z "$VPC_ID" ]; then
        echo "failed to create VPC"
        exit 1
    fi
    echo "VPC_ID: $VPC_ID"

    echo "Create SSH key pair"
    if [ -f ~/.ssh/id_rsa ]; then
        echo "SSH key pair already exists"
    else
        ssh-keygen -t rsa -N '' -f ~/.ssh/id_rsa
    fi
    aliyun ecs ImportKeyPair --RegionId "$ACK_REGION" \
        --KeyPairName "$CASE_NAME" \
        --PublicKeyBody "$(cat ~/.ssh/id_rsa.pub)" \
        --Tag.1.Key 'created-for' --Tag.1.Value "$CASE_NAME"

    echo "waiting for VPC to be ready"
    while true; do
        STATE=$(aliyun vpc DescribeVpcAttribute --RegionId "$ACK_REGION" --VpcId "$VPC_ID" | jq -r .Status)
        echo "VPC state: $STATE"
        if [ "$STATE" = "Available" ]; then
            break
        fi
        sleep 3
    done

    local subnet=0
    local zone
    local -a vswitch_ids
    for zone in $ACK_ZONE; do
        echo "Creating VSwitch for zone $zone"
        local id
        id=$(aliyun vpc CreateVSwitch --RegionId "$ACK_REGION" \
            --ZoneId "$zone" \
            --VpcId "$VPC_ID" \
            --CidrBlock 172.16.$subnet.0/24 \
            --ClientToken "$CASE_NAME-vsw-$subnet" \
            --Tag.1.Key 'created-for' --Tag.1.Value "$CASE_NAME" | jq -r .VSwitchId)
        if [ -z "$id" ]; then
            echo "failed to create VSwitch"
            exit 1
        fi
        echo "VSwitch ID: $id"
        vswitch_ids+=("$id")
        subnet=$((subnet+1))
    done

    local cluster_params
    cluster_params=$(jsonnet "$HERE/cluster-template.jsonnet" \
        --ext-str region="$ACK_REGION" \
        --ext-str vpc_id="$VPC_ID" \
        --ext-str cluster_name="$CASE_NAME" \
        --ext-str os_image_alinux3="${OS_IMAGE_ALINUX3:-aliyun_3_x64_20G_alibase_20240819.vhd}" \
        --ext-str os_image_containeros3="${OS_IMAGE_CONTAINEROS3:-lifsea_3_x64_10G_alibase_20240918.qcow2}" \
        --ext-code n_nodes="$N_NODES" \
        --ext-code vswitch_ids="$(jq -n '$ARGS.positional' --args "${vswitch_ids[@]}")")
    CLUSTER_ID=$(aliyun --region "$ACK_REGION" cs POST /clusters --header "Content-Type=application/json" --body "$cluster_params" | jq -r .cluster_id)

    if [ -z "$CLUSTER_ID" ]; then
        echo "failed to create cluster"
        exit 1
    fi
    echo "CLUSTER_ID: $CLUSTER_ID"

    echo "waiting for cluster to be ready"
    while true; do
        local STATE
        STATE=$(aliyun --region "$ACK_REGION" cs GET "/clusters/${CLUSTER_ID}" | jq -r .state)
        echo "cluster state: $STATE"
        if [ "$STATE" = "running" ]; then
            break
        elif [ "$STATE" != "initial" ]; then
            echo "cluster state is not 'running' or 'initial'"
            exit 1
        fi
        sleep 15
    done
    get-kubeconfig
    ram-setup "$CASE_NAME"
}

function get-kubeconfig {
    echo "getting kubeconfig"
    KUBECONFIG=${KUBECONFIG:-~/.kube/config.$CASE_NAME}
    aliyun cs GET "/k8s/${CLUSTER_ID}/user_config" --TemporaryDurationMinutes 480 | jq -r .config > "$KUBECONFIG"
    chmod 600 "$KUBECONFIG"
}

function ssh-forward-setup {
    echo "setup SLB ssh forwards"
    touch ~/.ssh/config
    chmod 600 ~/.ssh/config

    echo "getting SLB ID"
    local RESOURCES
    RESOURCES=$(aliyun cs GET "/clusters/${CLUSTER_ID}/resources")
    SLB_ID=$(echo "$RESOURCES" | jq -r '.[] | select(.resource_type == "SLB") | .instance_id')
    EIP=$(echo "$RESOURCES" | jq -r '.[] | select(.resource_type == "EIP") | .resource_info')
    echo "SLB_ID: $SLB_ID, EIP: $EIP"

    echo "waiting for $N_NODES nodes to appear"
    while true; do
        local CURRENT_N_NODES
        CURRENT_N_NODES=$(aliyun cs GET "/clusters/${CLUSTER_ID}/nodes" | jq -r '.nodes | length')
        echo "current number of nodes: $CURRENT_N_NODES"
        if [ "$CURRENT_N_NODES" -ge "$N_NODES" ]; then
            break
        fi
        sleep 10
    done

    local node_idx=0
    local node
    aliyun cs GET "/clusters/${CLUSTER_ID}/nodes" | jq -r '.nodes[]|[.instance_id, .ip_address[0]]|@tsv' | \
    while read -r node PRIVATE_IP; do
        echo "creating vServerGroup for node $node_idx $node ($PRIVATE_IP)"
        local VSG_ID
        VSG_ID=$(aliyun slb CreateVServerGroup --RegionId "$ACK_REGION" \
            --LoadBalancerId "$SLB_ID" \
            --VServerGroupName "vsg-$node_idx" \
            --BackendServers '[{"ServerId": "'"$node"'", "Port": 22, "Weight": 100}]' | jq -r .VServerGroupId)
        echo "VSG_ID: $VSG_ID"
        
        local PORT=$((22000+node_idx))
        aliyun slb CreateLoadBalancerTCPListener --RegionId "$ACK_REGION" \
            --LoadBalancerId "$SLB_ID" \
            --ListenerPort $PORT \
            --HealthCheckSwitch off \
            --Bandwidth="-1" \
            --VServerGroupId "$VSG_ID"
        echo "created listener for node $node on port $PORT"

        aliyun slb StartLoadBalancerListener --RegionId "$ACK_REGION" \
            --LoadBalancerId "$SLB_ID" \
            --ListenerPort $PORT
        echo "started listener"
        
        cat <<EOF >> ~/.ssh/config

Host $node $PRIVATE_IP
    HostName $EIP
    Port $PORT
EOF

        node_idx=$((node_idx+1))
    done
}

KEY=1234567890ABCDEF
IV=0000000000000000

function encrypt {
    echo -n $IV
    openssl enc -aes-128-cbc -e -K "$(xxd -p <<< $KEY)" -iv "$(xxd -p <<< $IV)"
}

function create-ram-role {
    aliyun ram CreateRole --region "$ACK_REGION" \
        --RoleName "$CASE_NAME-$1" \
        --MaxSessionDuration 28800 \
        --AssumeRolePolicyDocument "$(jq -n '{
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Effect": "Allow",
            "Principal": {
                "RAM": $arn
            },
            "Condition": {
                "StringEquals": {
                    "sts:ExternalId": $externalID
                }
            }
        }
    ],
    "Version": "1"
}' --arg arn "$arn" --arg externalID "$external_id")"
    echo "Created role $CASE_NAME-$1"
}

function ram-role-setup {
    create-ram-role "$1"

    aliyun ram CreatePolicy --region "$ACK_REGION" \
        --PolicyName "$CASE_NAME-$1" \
        --PolicyDocument "$(cat "$HERE/../../docs/ram-policies/disk/$1.json")"
    echo "Created policy $CASE_NAME-$1"

    aliyun ram AttachPolicyToRole --region "$ACK_REGION" \
        --PolicyType "Custom" \
        --PolicyName "$CASE_NAME-$1" \
        --RoleName "$CASE_NAME-$1"
    echo "Attached policy to role"
}

function create-token-secret {
    local resp
    resp=$(aliyun sts AssumeRole --region "$ACK_REGION" \
        --DurationSeconds 28800 \
        --RoleArn "acs:ram::$uid:role/$CASE_NAME-$1" \
        --ExternalId "$external_id" \
        --RoleSessionName "$CASE_NAME")
    echo "Assumed role ID $(jq -r .AssumedRoleUser.AssumedRoleId <<< "$resp")"

    secret=$(jq -n '{"access.key.id": $id, "access.key.secret": $secret, "security.token": $token, "expiration": $expiration, "keyring": $keyring}' \
        --arg id "$(jq -j .Credentials.AccessKeyId <<< "$resp" | encrypt | base64)" \
        --arg secret "$(jq -j .Credentials.AccessKeySecret <<< "$resp" | encrypt | base64)" \
        --arg token "$(jq -j .Credentials.SecurityToken <<< "$resp" | encrypt | base64)" \
        --arg expiration "$(jq -j .Credentials.Expiration <<< "$resp")" \
        --arg keyring $KEY \
    )
    kubectl create secret generic "$2" -n kube-system --from-literal "addon.token.config=$secret"
}

function ram-setup {
    local HERE
    local CASE_NAME=$1
    HERE=$(dirname "${BASH_SOURCE[0]}")

    read -r arn uid < <(aliyun sts GetCallerIdentity | jq -r '[.Arn, .AccountId] | @tsv')
    echo "Current identity ARN: $arn"

    external_id=$(openssl rand 12 | basenc --base64url)

    ram-role-setup node
    create-token-secret node addon.aliyuncsmanagedcsipluginrole.token
    ram-role-setup controller
    create-token-secret controller addon.aliyuncsmanagedcsiprovisionerrole.token

    create-ram-role sanity
    aliyun ram AttachPolicyToRole --region "$ACK_REGION" \
        --PolicyType "Custom" \
        --PolicyName "$CASE_NAME-node" \
        --RoleName "$CASE_NAME-sanity"
    echo "Attached node policy to sanity role"

    aliyun ram AttachPolicyToRole --region "$ACK_REGION" \
        --PolicyType "Custom" \
        --PolicyName "$CASE_NAME-controller" \
        --RoleName "$CASE_NAME-sanity"
    echo "Attached controller policy to sanity role"

    create-token-secret sanity sanity-ram-token
}

function cluster-teardown {
    ram-teardown "$CASE_NAME"

    echo "deleting cluster"
    aliyun cs DELETE "/clusters/${CLUSTER_ID}"

    echo "waiting for cluster to be deleted"
    while true; do
        local STATE
        STATE=$(aliyun cs GET "/clusters/${CLUSTER_ID}" | jq -r .state)
        echo "cluster state: $STATE"
        if [ "$STATE" != "deleting" ]; then
            break
        fi
        sleep 30
    done

    echo "deleting VPC"
    aliyun vpc DeleteVpc --RegionId "$ACK_REGION" --ForceDelete true --VpcId "${VPC_ID}"

    echo "deleting SSH key pair"
    aliyun ecs DeleteKeyPairs --RegionId "$ACK_REGION" --KeyPairNames "$(jq -n '$ARGS.positional' --args "$CASE_NAME")"
}

function ram-role-teardown {
    echo "detaching RAM policy"
    aliyun ram DetachPolicyFromRole --region "$ACK_REGION" --PolicyName "$CASE_NAME-$1" --RoleName "$CASE_NAME-$1" --PolicyType "Custom"

    echo "deleting RAM policy"
    aliyun ram DeletePolicy --region "$ACK_REGION" --PolicyName "$CASE_NAME-$1"

    echo "deleting RAM role"
    aliyun ram DeleteRole --region "$ACK_REGION" --RoleName "$CASE_NAME-$1"
}

function ram-teardown {
    local CASE_NAME=$1

    echo "cleaning up sanity role"
    aliyun ram DetachPolicyFromRole --region "$ACK_REGION" --PolicyName "$CASE_NAME-node" --RoleName "$CASE_NAME-sanity" --PolicyType "Custom"
    aliyun ram DetachPolicyFromRole --region "$ACK_REGION" --PolicyName "$CASE_NAME-controller" --RoleName "$CASE_NAME-sanity" --PolicyType "Custom"
    aliyun ram DeleteRole --region "$ACK_REGION" --RoleName "$CASE_NAME-sanity"
    kubectl delete secret -n kube-system sanity-ram-token

    ram-role-teardown controller
    kubectl delete secret -n kube-system addon.aliyuncsmanagedcsiprovisionerrole.token

    ram-role-teardown node
    kubectl delete secret -n kube-system addon.aliyuncsmanagedcsipluginrole.token
}
