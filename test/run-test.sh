#! /bin/sh

## Disk Test Case: use dynamic storage
sh disk/disk-dynamic.sh
sh disk/disk-aware.sh


## NAS Case test
nasserver=""
sh nas/nas-static.sh $nasserver


## OSS Case test: Input your access key;
akid=""
aksec=""
sh oss/oss-static.sh $akid $aksec
