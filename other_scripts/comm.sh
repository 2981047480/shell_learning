#!/bin/bash

comm a.txt b.txt
#第一列：a中特有的，第二列：b中特有的，第三列：c中特有的
comm a.txt b.txt -1 -2
#忽略第一第二列
comm a.txt b.txt -3 | tr -d '\t' 
