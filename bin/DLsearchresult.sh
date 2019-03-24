#!/bin/bash
cat effectvalue.txt | while read effectValue
do 
./POSTsearch.sh $effectValue
done
