#!/bin/bash
directory=$1
sh make_htmllist.sh $directory | xargs -n 1 sh wget.sh --html
