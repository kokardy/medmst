#!/bin/bash
d=$1
root="https://www.pmda.go.jp"
path="/PmdaSearch/iyakuDetail/"
grep "PmdaSearch/iyakuDetail/ResultDataSetPDF/" -rn $d | sed -E "s@^.*?'/PmdaSearch/iyakuDetail/ResultDataSetPDF/(.*?)'>PDF.*@${root}${path}\1@g" 
