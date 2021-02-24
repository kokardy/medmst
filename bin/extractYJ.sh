#!/bin/bash
input=$1
grep -o -e "[0-9]\{5\}[0-9A-Z][0-9A-Z][A-Z][0-9]\{4\}" $input | while read YJ; do
	filename=`basename $input`
	value=${filename%.*}
	./StoreKV.sh $YJ $value
done

