#!/bin/bash
x=1
for id in $(docker history -q my_image|tac|grep -v missing)
do
	docker tag "${id}" "my_image:latest_step_${x}"
	((x++))
done
