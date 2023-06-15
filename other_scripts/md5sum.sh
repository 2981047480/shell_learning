#!/bin/bash

md5sum example.txt > example.md5
md5sum -c example.md5
find . -name '*' -print0 | xargs -0 md5sum
