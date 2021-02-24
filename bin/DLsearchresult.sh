#!/bin/bash
d=$1
cat effectvalue.txt | while read effectValue
do 
./POSTsearch.sh $d $effectValue
done
