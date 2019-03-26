#!/bin/bash
saveDir=$1

#POST param
updateDocDt=$2 #指定時は8桁数字19801028

#入力値チェック
if [ $saveDir = "" ] || [ $updateDocDt = "" ]; then
	echo "num of command option must be 2."
	echo "POSTsearchDay.sh [saveDir] [update_date(YYYYMMDD)]"
	exit
fi	
	

#output param
filename=$2

mkdir -p $saveDir

http \
--ignore-stdin \
--form \
--output $saveDir/${filename}.html \
POST \
https://www.pmda.go.jp/PmdaSearch/iyakuSearch/ \
User-Agent:"Mozilla/5.0 (Linux; Android 4.2.2; ja-jp; SC-04E Build/JDQ39) AppleWebkit/535.19 (KHTML,like Gecko) Version/1.0 Chrome/18.0.1025.308 Mobile Safari/535.19" \
Host:www.pmda.go.jp \
Referer:https://www.pmda.go.jp/PmdaSearch/iyakuSearch/ \
btnA.x=$((66 + $RANDOM % 5)) \
btnA.y=$((5 + $RANDOM % 5)) \
nameWord= \
iyakuHowtoNameSearchRadioValue=1 \
howtoMatchRadioValue=1 \
effectValue=${effectValue} \
infoindicationsorefficacy= \
infoindicationsorefficacyHowtoSearch=and \
warnings= \
warningsHowtoSearch=and \
contraindicationsAvoidedadministration= \
contraindicationsAvoidedadministrationHowtoSearch=and \
contraindicatedcombinationPrecautionsforcombination= \
contraindicatedcombinationPrecautionsforcombinationHowtoSearch=and \
dispColumnsList[0]=1 \
dispColumnsList[1]=2 \
dispColumnsList[2]=3 \
dispColumnsList[3]= \
dispColumnsList[4]= \
dispColumnsList[5]= \
dispColumnsList[6]= \
dispColumnsList[7]= \
dispColumnsList[8]= \
dispColumnsList[9]= \
dispColumnsList[10]= \
dispColumnsList[11]= \
dispColumnsList[12]= \
dispColumnsList[13]= \
dispColumnsList[14]= \
dispColumnsList[15]= \
changeColumnsList= \
updateDocFrDt="${updateDocFrDt}" \
updateDocToDt="${updateDocToDt}" \
tglOpFlg= \
compNameWord= \
koumoku1Value= \
koumoku2Value= \
koumoku3Value= \
koumoku1Word= \
koumoku2Word= \
koumoku3Word= \
koumoku1HowtoSearch=and \
koumoku2HowtoSearch=and \
koumoku3HowtoSearch=and \
howtoRdSearchSel=or \
relationDoc1Sel= \
relationDoc2Sel= \
relationDoc3Sel= \
relationDoc1check1=on \
relationDoc1check2=on \
relationDoc2check1=on \
relationDoc2check2=on \
relationDoc3check1=on \
relationDoc3check2=on \
relationDoc1Word=検索語を入力 \
relationDoc2Word=検索語を入力 \
relationDoc3Word=検索語を入力 \
relationDoc1HowtoSearch=and \
relationDoc2HowtoSearch=and \
relationDoc3HowtoSearch=and \
relationDoc1FrDt="年月 [YYYYMM]" \
relationDoc1ToDt="年月 [YYYYMM]" \
relationDoc2FrDt="年月 [YYYYMM]" \
relationDoc2ToDt="年月 [YYYYMM]" \
relationDoc3FrDt="年月 [YYYYMM]" \
relationDoc3ToDt="年月 [YYYYMM]" \
relationDocHowtoSearchBetween12=and \
relationDocHowtoSearchBetween23=and \
ListRows=1000 \
listCategory=BOOK \
personFlg=false \
gs1code= \
iyakuKoumokuSelectSwitchRadio=2
