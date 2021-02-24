#!/bin/bash
d=$1

today=`date  +%Y%m%d`

lastdate=`ls -r $d/searchresult | grep -e "[0-9]\{8\}" | sed -n 1P`
lastdate=${lastdate%.*}

echo "lastdate: $lastdate .... today: $today"

if [ $lastdate = "00000000_" ] ;then
		lastdate="20200415"
fi

day=$lastdate
while [ $day -le $today ]
do
	./POSTsearchDay.sh $d $day
	./DLhtml.sh $d/searchresult/$day.html
	echo "done: $day "
	day=`date --date "$day 1 days" +%Y%m%d`
done
