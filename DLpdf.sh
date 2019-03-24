#!/bin/bash
directory=$1
sh make_pdflist.sh $directory | xargs -L 1 sh wget.sh --pdf
