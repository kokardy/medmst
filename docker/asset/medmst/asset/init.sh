cd /bootstrap
DATE=`date +%Y%m%d`
medmst -d $DATE -p $HTTP_PROXY -f
mv $DATE save
cd save/hot
jlha xif save/hot/*.lzh
cd /bootstrap
cd save/y
unzip -jo y.zip
