#!/bin/bash

echo thisthisthis | sed 's/this/THIS/2g'
sed 's/\b[0-9]\{3\}\b/NUMBER/g' sed_data.txt
echo this is an example | sed 's/\w\+/[&]/g' #这里&可以指代之前匹配到的值
echo abc | sed -e 's/a/A/' -e 's/c/C'
text=hello
echo hello world | sed "s/$text/HELLO"

