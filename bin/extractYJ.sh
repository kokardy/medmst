#!/bin/bash
input=$1
grep -o -e "[0-9]\{7\}[A-Z][0-9]\{4\}" $input | while read YJ; do
	filename=`basename $input`
	value=${filename%.*}
	./StoreKV.sh $YJ $value
done

