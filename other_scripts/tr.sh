#!/bin/bash

echo 12345 | tr '0-9' '9876543210'
echo 12345 | tr '09876543210' '0-9'
echo 12345 hello 6789 world | tr -d '0-9'
echo 12345 hello 6789 world | tr -c -d 'a-zA-Z'
echo 12345 hello 6789 world | tr -c 'a-zA-Z' ' '
echo 12345 hello 6789 world | tr -c -s 'a-zA-Z' ' '
