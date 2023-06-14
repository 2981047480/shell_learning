#!/bin/bash

#color
color_theme() {
    RED='\E[1;31m'
    GREEN='\E[1;32m'
    YELLOW='\E[1;33m'
    RES='\E[0m'
}
color_theme

#help
printHelp() {
	echo -e '$GREEN -h/--help $RES 查看帮助'
	echo -e '$GREEN -f/--host_file $RES 存放ip的文件'
	echo -e '$GREEN -p $RES 指定ssh密码'
	exit 2
}

#vars
password='zsh1101.'
user=`whoami`
host_file='./hosts'
while [ $# -gt 0 ]
do
	case $1 in
		-h) printHelp && exit ;;
		--help)  printHelp && exit ;;
		--host_file) shift; host_file=$1; shift;;
		-f) shift; host_file=$1; shift;;
		-p) shift; password=$1; shift;;
	esac
done

#init
init() {
	echo -e '$GREEN 您正在使用$user 用户执行该脚本$RES'
}
#hosts
hosts_set() {
	count=0
	for ip in `cat $host_file`;do 
		response=`grep $ip aaa|awk '{print $1}'`
		if [ ! $response ]; then
			echo "$ip node$count" >> aaa && let count++;
		fi
	done
}
hosts_set
