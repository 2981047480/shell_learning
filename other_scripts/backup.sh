#!/bin/bash

read -p "What floder should be backuped: " floder
read -p "What kind of files should be backuped: " suffix
echo $floder *.$suffix
find $floder -name '*.'$suffix -a ! -name  '~*' -print
find $floder -name '*.'$suffix -a ! -name  '~*' -exec cp {} /home/zephyr/backup \;
#   {}代表被处理文件
echo "Backuped successful!"
