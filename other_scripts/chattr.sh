#!/bin/bash

chattr +i a.txt
rm a.txt
chattr -i a.txt
