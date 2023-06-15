#!/bin/bash

if [ $# -ne 3 ];
then
	echo "Usage: $0 URL -d DIRECTORY"
	exit -1
fi

while [ $# -gt 0 ]
do
	case $1 in 
		-d) shift; directory=$1; echo $1; shift ;;
		*) base_url=$1; echo $1; shift;;
	esac
done

echo $directory
echo $base_url
#curl -s $base_url | grep 'img' | sed 's/.*<img[^>]*src=\/\(.*\)>.*/title: \1/g'|sed 's/\(.\)>.*/\1/g'
ex_url=`curl -s $base_url | grep 'img' | sed 's/.*<img[^>]*src=\/\(.*\)>.*/\1/g'|sed 's/\(.\)>.*/\1/g'`
echo $ex_url
