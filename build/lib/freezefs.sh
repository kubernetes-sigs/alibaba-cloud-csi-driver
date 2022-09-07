#!/bin/bash

# sh fsfreeze.sh --type=freeze --path=/var/lib/kubelet/plugins/kubernetes.io/csi/pv/d-hp39lj3xo500dacme9hi/globalmount --timeout=10
# sh fsfreeze.sh --type=unfreeze --path=/var/lib/kubelet/plugins/kubernetes.io/csi/pv/d-hp39lj3xo500dacme9hi/globalmount

timeoutFunc() {
  errorTpl="unfreeze failed: Invalid argument"
  echo `date`" fsfreeze start to sleep: "$timeout  >> /var/log/alicloud/fsfreeze.log
  sleep $2
  echo `date`" fsfreeze-finish to path: "$1  >> /var/log/alicloud/fsfreeze.log
  returnStr=$(fsfreeze -u $1 2>&1)
  if [[ -n $returnStr ]]; then
    if [[ $returnStr == *$errorTpl* ]]; then
    echo `date`" the status of "$1" is already unfreeze" >> /var/log/alicloud/fsfreeze.log
    exit 0
    fi
    echo `date`" fsfreeze -u "$1" error: $returnStr" >> /var/log/alicloud/fsfreeze.log
    exit 0
  fi
}

if [[ "$1" = "--type=freeze" ]]; then
  type=freeze
elif [[ "$1" = "--type=unfreeze" ]]; then
  type=unfreeze
else
  echo `date`" input unsupport type: "$1 >> /var/log/alicloud/fsfreeze.log
  exit 1
fi

pathStr=$2
path=`echo ${pathStr:7}`

if [ "$type" = "unfreeze" ]; then
  fsfreeze -u $path &>> /var/log/alicloud/fsfreeze.log
  echo `date`" fsfreeze -u "$path" finished."  >> /var/log/alicloud/fsfreeze.log
  exit 0
fi

timeout=5
timeStr=$3
if [[ "$timeStr" != "" ]]; then
  timeout=`echo ${timeStr:10}`
fi

echo `date`" fsfreeze-start to path: "$path" with timeout: "$timeout  >> /var/log/alicloud/fsfreeze.log
fsfreeze -f $path &>> /var/log/alicloud/fsfreeze.log
if [ "$?" != "0" ]; then
  echo `date`" fsfreeze -f "$path" error." >> /var/log/alicloud/fsfreeze.log
  exit 1
fi

timeoutFunc $path $timeout &