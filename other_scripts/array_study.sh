#!/bin/bash

array_var=(test1 test2 test3 test4)
echo ${array_var[*]}
echo ${#array_var[*]}


#创建关联数组
declare -A fruits_value
fruits_value=([apple]='100 dollars' [orange]='150 dollars')
echo ${fruits_value[apple]}
#列出数组的索引
echo ${!array_var[*]}
