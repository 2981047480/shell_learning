#!/bin/bash

tput cols
#获取终端的行数
tput lines
#获取终端的列数
tput longname
#获取当前终端名
tput cup 2 2
#把光标移动到（2，2）处
tput setb n
#设置终端背景色（n在0-7之间）
tput setf n
#设置终端前景色（n在0-7之间）
tput bold
#加粗
tput smul
tput rmul
#下划线起止

