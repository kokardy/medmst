cd /bootstrap
DATE=`date +%Y%m%d`
medmst -d $DATE -f
mv $DATE save
cd save/hot
jlha xif save/hot/*.lzh
cd /bootstrap
cd save/y
unzip -jo y.zip
