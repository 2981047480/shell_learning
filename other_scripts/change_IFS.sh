#!/bin/bash

data="name,gender,rollno,location"

oldIFS=$IFS
IFS=,
for item in $data;
do
	echo Item:$item
done

IFS=$oldIFS

oldIFS=$IFS
IFS=:
line=`cat /etc/passwd|grep zephyr`
for i in $line;
do
	echo $i
done

