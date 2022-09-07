#!/bin/bash

# sh fsfreeze.sh --type=freeze --path=/var/lib/kubelet/plugins/kubernetes.io/csi/pv/d-hp39lj3xo500dacme9hi/globalmount --timeout=10
# sh fsfreeze.sh --type=unfreeze --path=/var/lib/kubelet/plugins/kubernetes.io/csi/pv/d-hp39lj3xo500dacme9hi/globalmount

timeoutFunc() {
  pathTmp=$1
  echo `date`" fsfreeze start to sleep: "$timeout  >> /var/log/alicloud/fsfreeze.log
  sleep $2
  if ["$1" != pathTmp ]; then
    echo `date`" skip fsfreeze-finish as already unfreeze"
    exit 0
  fi
  echo `date`" fsfreeze-finish to path: "$1  >> /var/log/alicloud/fsfreeze.log
  fsfreeze -u $1 &>> /var/log/alicloud/fsfreeze.log
  if [ "$?" != "0" ]; then
    echo `date`" fsfreeze -u "$1" error." >> /var/log/alicloud/fsfreeze.log
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

if [ "$type" = "freeze" ]; then
  timeoutFunc $path $timeout &
fi