#!/bin/bash
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
wget -c $url -O $output
if [ $xxx = "html" ]; then
	sh ./extractYJ.sh $output
fi
