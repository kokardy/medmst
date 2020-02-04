#!/bin/bash
d=$1
./DLsearchresult.sh $d
./DLhtml.sh $d/searchresult/00000000_

"0 2 * * * bash /asset/bin/Routine.sh /asset/data" | crontab -
