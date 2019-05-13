#!/bin/sh

while getopts "d:n:" opt; do
  case $opt in
    d)
      devs=$$OPTARG
      echo "The target device: $OPTARG"
      ;;
    n)
      vgName=$OPTARG
      echo "Volume group name: $OPTARG"
      ;;
    \?)
      echo "Invalid option: -$OPTARG"
      ;;
  esac
done

if [ "$devs" = "" ] || [ "$vgName" = "" ]; then
    echo "Must input device and volume group name"
    exit 1
fi

vgcreate --version > /dev/null
if [ "$?" != "0" ]; then
    echo "Vgcreate is unavailable"
    exit 1
fi

echo "Create Volume Group on $devs"
vgcreate $vgName $devs
