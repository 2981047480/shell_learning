#!/bin/bash

printHelp () {
	echo "本脚本用来将客户的版本信息提取出来"
	echo "使用方法："
	echo "1. 提取版本信息sh info_to_text.sh -f 客户信息info文件名"
	echo "2. 比较两个版本信息 sh info_to_text.sh -f1 info文件1 -f2 info文件2"
	exit 2
}

function log() {
    log_color=$1
    log_context=$2
    echo -e "\033[${log_color}m$log_context \033[0m"
}

function red() {
	log "1;31" "$@"
}

if [ $# -ne 2 ] && [ $# -ne 4 ];
then
	printHelp
	exit -1
fi

info_file=''
file1=''
file2=''

while  [ $# -gt 0 ]
do
	case $1 in
		-f) shift ; info_file=$1; shift;;
		-h) printHelp; exit;;
		-f1) shift ; file1=$1; shift;;
		-f2) shift ; file2=$1; shift;;
	esac
done

[ ! -e ${info_file} ] && echo "${info_file}: No such file" && printHelp && exit -1;
[ ! -e ${file1} ] && echo "${file1}: No such file" && printHelp && exit -1;
[ ! -e ${file2} ] && echo "${file2}: No such file" && printHelp && exit -1;

if [  ${info_file} ];then
	cat $info_file| egrep '  .*{.*version:.*}' --color|sed 's/  \(.*\): {.*version: \(.*\)}/\1 \2/g'
	exit 0
fi

if [ ${file1} ] && [ ${file2} ];then
	#echo $file1,$file2
	declare -A production1
	declare -A production2
	eval `cat $file1| egrep '  .*{.*version:.*}' --color|sed 's/  \(.*\): {.*version: \(.*\)}/\1 \2/g'|awk '{production1[$1]=$2} END {for(i in production1) print "production1["i"]="production1[i]}'` 
	#cat $file2| egrep '  .*{.*version:.*}' --color|sed 's/  \(.*\): {.*version: \(.*\)}/\1 \2/g'|awk -v pro=$production '{${pro[$1]}!=$2) print $1,${pro[$1]},"1111",$1,$2}'
	eval `cat $file2| egrep '  .*{.*version:.*}' --color|sed 's/  \(.*\): {.*version: \(.*\)}/\1 \2/g'|awk '{production2[$1]=$2} END {for(i in production2) print "production2["i"]="production2[i]}'`
	for i in `echo ${!production1[*]} ${!production2[*]}|sed 's/ /\n/g'|sort|uniq`;
	do 
		if [[ -v production1[$i] ]] && [[  -v production2[$i] ]] ; then
			#if [[ $production1[$i] != $production2[$i] ]] ; then
		       	#	echo $i,${production1[$i]},${production2[$i]} 
			#else
			#	echo 
			#fi
			[[ ${production1[$i]} != ${production2[$i]} ]] && echo -e $i '\t' ${production1[$i]} '\t' ${production2[$i]}
		elif [[ -v production1[$i] ]] && [[ ! -v production2[$i] ]] ; then
			echo -e $i '\t' ${production1[$i]}
		else
			echo -e $i '\t' ${production2[$i]}
		fi
	done
	#echo ${production1[@]}
	#echo ${production2[@]}
	exit 0
fi

printHelp
exit 0
