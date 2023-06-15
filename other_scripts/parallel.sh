#!/bin/bash

ls -l *.sh | parallel echo {.}
