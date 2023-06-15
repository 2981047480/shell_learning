#!/bin/bash

echo Count:
tput sc

#循环10秒
for count in `seq 0 10`
do
	tput rc
	tput ed
	echo -n $count
	sleep 1
done
