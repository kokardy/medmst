#!/bin/bash
directory=$1
sh make_pdflist.sh $directory | xargs -n 1 sh wget.sh --pdf
