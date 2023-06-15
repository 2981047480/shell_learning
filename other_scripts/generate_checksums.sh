#!/bin/bash

PIDARRAY=()
for file in `find . -name '*.sh'`
do
	md5sum $file &
	PIDARRAY+=("$!")
	#$! 最近一个进程的pid
done
echo ${PIDARRAY[@]}
wait ${PIDARRAY[@]}
