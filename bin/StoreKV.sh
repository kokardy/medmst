#!/bin/bash
d=${PMDAKV:-yj}

key=$1
value=$2
mkdir -p $d
echo $value > $d/$key
