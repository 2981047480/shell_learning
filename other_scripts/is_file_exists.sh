#!/bin/bash

fpath=$1
echo $fpath
if [ -e $fpath ]; then
	echo File exists;
else
	echo Does not exist;
fi
