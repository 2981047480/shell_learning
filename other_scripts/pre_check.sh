#!/bin/bash

color_theme() {
    RED='\E[1;31m'
    GREEN='\E[1;32m'
    YELLOW='\E[1;33m'
    RES='\E[0m'
}
color_theme

system_version() {
	echo -e "$YELLOW system version $RES\n"
	cat /etc/redhat-release
	echo -e "$YELLOW kernel $RES\n"
	uname -a
}
cpu() {
	echo -e "$YELLOW cpu info $RES\n"
	lscpu|grep -i '^cpu(s)'|awk '{print $1,$2}'
	lscpu|grep -o 'avx2'
}
memory() {
	echo -e "$YELLOW memory $RES\n"
	free -h|grep -v '^ '|awk 'BEGIN {print "type total used"} {print $1,$2,$3}'
}
disk() {
	echo -e "$YELLOW disk $RES\n"
	df -h|grep -v "tmpfs\|overlay\|boot\|run\|mnt"
	echo -e "$YELLOW lsblk $RES\n"
	lsblk
}
ntp() {
	echo -e "$YELLOW ntp $RES\n"
	ntpstat | grep -i 'synchronised'
}
yum() {
	echo -e "$YELLOW yum $RES\n"
	yum makecache
}

print_help() {
	echo -e "$YELLOW This script is used to check the environment $RES"
	echo -e "$YELLOW   -h $RES print this help"
	echo -e "$YELLOW   -c $RES checkall"
}

if [ ! $# ];then
	print_help && exit -1
fi
while [ $# -gt 0 ]
do
	case $1 in 
		-c) system_version && cpu && memory && disk && ntp && yum ; shift;;
		-h) print_help;shift;;
	esac
done
echo -e "$YELLOW end $RES\n"
