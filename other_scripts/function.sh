#!/bin/bash

fname()
{
	echo $1, $2;
	echo "$@";
	echo "$*";
	return 0;
}
fname echo hello world
export -f fname
