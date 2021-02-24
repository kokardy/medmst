#!/bin/bash

if [ "$2" = "" ]; then
	echo "変更なし "
	exit
fi

url=$2
base=`basename $url`
while getopts ":-:" OPT
do 
		case $OPT in
				-) 
						case "$OPTARG" in
								pdf) 
									 	xxx=pdf
										;;
								html)
									 	xxx=html
										;;
								sgml)
									 	xxx=sgml
										;;
						esac
						;;
		esac
done
output="${xxx}/${base}.${xxx}"
mkdir -p ${xxx}
echo "download: $output"
wget -q -c $url -O $output
if [ $xxx = "html" ]; then
	sh ./extractYJ.sh $output
fi
