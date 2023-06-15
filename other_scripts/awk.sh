#!/bin/bash

awk -F: 'BEGIN {print "BEGIN"} {print $1} END {print "END"}' passwd
echo | awk '{print var}' var=$VAR
seq 5 | awk 'BEGIN{getline; print "Read ahead first line", $0} {print $0}'
seq 10|awk 'NR<5'
seq 10|awk 'NR==1,NR==4'
cat passwd|awk '/a/'
cat passwd|awk '/!a/'
