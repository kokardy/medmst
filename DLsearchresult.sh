#!/bin/bash
cat effectvalue.txt | while read effectValue
do 
./post.sh $effectValue
done
