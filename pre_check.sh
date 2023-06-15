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
	lscpu|grep -o 'avx2\|avx \|ssse3\|sse4\|sse4_2\|popcnt'
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
yum_check() {
	echo -e "$YELLOW yum $RES\n"
	yum makecache 2>&1|grep "Metadata cache created."
}
selinux() {
	echo -e "$YELLOW selinux $RES\n"
	getenforce
}

print_help() {
	echo -e "$YELLOW This script is used to check the environment $RES"
	echo -e "$YELLOW   -h $RES print this help"
	echo -e "$YELLOW   -c $RES checkall"
	echo -e "$YELLOW   -v $RES print version"
}

print_version() {
	echo -e "$YELLOW version: 0.1.1 $RES"
}

iptables_check() {
	echo -e "$YELLOW iptables $RES\n"
        iptables -nvL
}

ip_check() {
	echo -e "$YELLOW ip_check $RES\n"
	#ifconfig|egrep '([0-9]*\.){3}([0-9]){1}'|grep -v 127.0.0.1
	ips=$(ip a|egrep '([0-9]*\.){3}([0-9]){1}'|grep -v 127.0.0.1|awk '{print $2}' || ifconfig|egrep '([0-9]*\.){3}([0-9]){1}'|grep -v 127.0.0.1|awk '{print $2}')
	echo $ips | sed 's/ /\n/g'
}

root_login_check() {
	echo -e "$YELLOW root_login_check $RES\n"
	cat /etc/ssh/sshd_config|grep ^PermitRootLogin
}

end() {
	echo -e "$YELLOW end $RES\n"
}

if [ ! $# ];then
	print_help && exit -1
fi
while [ $# -gt 0 ]
do
	case $1 in 
		-c) system_version ; cpu ; memory ; disk ; ntp ; yum_check ; iptables_check; ip_check; selinux; root_login_check; end; exit 0; shift;;
		-h) print_help; end; exit 1;shift;;
		-v) print_version; end; exit 1;shift;;
	esac
done
print_help;
end;
