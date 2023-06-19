#!/bin/bash

#color
function color_theme() {
    RED='\E[1;31m'
    GREEN='\E[1;32m'
    YELLOW='\E[1;33m'
    RES='\E[0m'
}
color_theme

#help
function printHelp() {
	echo -e '$GREEN -h/--help $RES 查看帮助'
	echo -e '$GREEN -f/--host_file $RES 存放ip的文件'
	echo -e '$GREEN -p $RES 指定ssh密码'
	exit 2
}

#vars
password='zsh1101.'
user=`whoami`
host_file='./hosts'


#init
function init() {
	#echo -e `$GREEN 您正在使用$user用户执行该脚本$RES`
	echo -e '$GREEN You are using $user to executing�$RES'
}
#hosts
function hosts_set() {
	count=0
	#for ip in `cat $host_file`;do echo "$ip node$count" >> /etc/hosts && let count++;done
	for ip in `cat $host_file`;do
		response=`cat /etc/hosts|grep -v local|grep $ip|awk '{print $1}'`
		if [ ! $response ]; then
			echo "$ip node$count" >> /etc/hosts && let count++;
		fi
	done
}
#docker
function docker() {
curl -sSL https://get.daocloud.io/docker | sh
yum install docker-ce-rootless-extras -qy
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
}
#kind
function kind() {
	yum install kubernetes-client
	wget --http-user=pandag --http-passwd=z5lZiumYp http://ufile.vip.sensorsdata.cn/dl/kind-linux-amd64
	mv kind-linux-amd64 kind && chmod 755 kind && cp kind /bin
}
function minikube() {
	curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
	sudo install minikube-linux-amd64 /usr/local/bin/minikube
}
#ssh
function ssh_without_password() {
	ssh-keygen -t rsa -q -f ~/.ssh/id_rsa -P ''
	yum install sshpass -qy
	for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh-copy-id $user@$host;done
	#localhost_ip=`ifconfig|grep 'inet '|grep -v '127'|awk '{print $2}'`
	#localhost_ips=`echo $localhost_ip|sed 's/ /|/g'`
	#cat /etc/hosts|grep node|awk '{print $1}'|xargs -i sshpass -p $password ssh $user@{} "for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh-copy-id $user@$host;done"
	for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh -o stricthostkeychecking=no $host "cat /etc/hosts|grep node|awk '{print $1}'|xargs -i sshpass -p $password ssh-copy-id $user@{}";done
	#for host in `grep -Ev $localhost_ips $host_file`;do ssh $host "mkdir -p ~/backup&&cp /etc/hosts ~/backup";done
	#for host in `grep -Ev $localhost_ips $host_file`;do scp -o stricthostkeychecking=no /etc/hosts $user@$host:/etc;done
	#for host in `grep -Ev $localhost_ips $host_file`;do ssh -o stricthostkeychecking=no $host "cat /etc/hosts|grep node|awk '{print $1}'|xargs -i sshpass -p $password ssh-copy-id $user@{}";done
	for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh $host docker;done
	for host in `cat /etc/hosts|grep node|awk '{print $1}'`;do sshpass -p $password ssh $host kind && minikube;done
}
function run_init_ssh() {
	echo '$GREEN init$RES'
	init
	echo '$GREEN hosts$RES'
	hosts_set
	echo '$GREEN ssh$RES'
	ssh_without_password
}

function main() {
	while [ $# -gt 0 ]
	do
		case $1 in
			-h) printHelp && exit ;;
			--help)  printHelp && exit ;;
			--host_file) shift; host_file=$1; shift;;
			-f) shift; host_file=$1; shift;;
			-p) shift; password=$1; shift;;
			-y) shift; run_init_ssh; shift ;;
		esac
	done
}