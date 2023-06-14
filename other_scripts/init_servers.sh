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
	echo -e `$GREEN 您正在使用$user用户执行该脚本$RES`
}
#hosts
hosts_set() {
	count=0
	for host in `cat $host_file`;do let count++ && echo "$ip node$count">>/etc/hosts;done
}
#ssh
ssh_without_password() {
	ssh-keygen -t rsa -q -f ~/.ssh/id_rsa -P ''
	yum install sshpass
	for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh-copy-id $user@$host;done
	localhost_ip=`ifconfig|grep 'inet '|grep -v '127'|awk '{print $2}'`
	localhost_ips=`sed 's/ /|/g'`
	for host in `grep -Ev $localhost_ips $host_file`;do ssh $host "mkdir -p ~/backup&&cp /etc/hosts ~/backup";done
	for host in `grep -Ev $localhost_ips $host_file`;do scp -o stricthostkeychecking=no /etc/hosts $user@$host:/etc;done
	for host in `grep -Ev $localhost_ips $host_file`;do ssh -o stricthostkeychecking=no $host "cat /etc/hosts|grep node|awk '{print $1}'|xargs -i ssh-pass -p $password ssh-copy-id $user@{}";done
}
#docker
curl -sSL https://get.daocloud.io/docker | sh
#daemon
sudo mkdir -p /etc/docker
sudo cat > /etc/docker/daemon.json << EOF
{
    "registry-mirrors": [
        "https://mirror.ccs.tencentyun.com"
    ]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

