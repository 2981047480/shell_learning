#!/bin/bash

#可以用这个命令来对终端进行录制
script -t 2>timing.log -a output.session
#可以用scriptreplay对终端进行回放
scriptreplay timing.log output.session
