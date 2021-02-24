#!/bin/bash
d=$1
root="https://www.pmda.go.jp"
grep "PmdaSearch/iyakuDetail/ResultDataSetPDF/" -rn $d | sed -E "s@^.*?'(/PmdaSearch/iyakuDetail/ResultDataSetPDF/.*?)'>PDF.*@${root}\1@g" 
