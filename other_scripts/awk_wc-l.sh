#!/bin/bash

read -p "echo file name: " text
#echo $text
awk "BEGIN { i=0 } { i++ } END { print i }" $text
