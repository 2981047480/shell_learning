#!/bin/bash

read -p 'please enter depth: ' depth
find . -maxdepth $depth -exec sh -c 'echo -n {} | tr -d "[:alnum:]-_." | tr "/" "---";basename {}' \;
#find . -exec sh -c 'echo -n {}|tr -d "[:alnum:]-._"|tr "/" "-";basename {}' \;
