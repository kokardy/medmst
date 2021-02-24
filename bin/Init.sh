#!/bin/bash
d=$1
./DLsearchresult.sh $d
./DLhtml.sh $d/searchresult/00000000_

echo "15 2,6,17 * * * cd /asset/bin && bash Routine.sh /asset/data >> /asset/cron.log" | crontab -
