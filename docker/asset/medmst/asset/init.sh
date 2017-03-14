cd /bootstrap
DATE=`date +%Y%m%d`
/go/bin/medmst -d $DATE -p $http_proxy -f
mv $DATE save
cd save/hot
jlha xif save/hot/*.lzh
cd /bootstrap
cd save/y
unzip -jo y.zip
