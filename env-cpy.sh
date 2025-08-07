#!/usr/bin/bash

ENVS=$(find . -name ".env.example")

for ENV in $ENVS
do
	DIR="${ENV%%.env.example}"
	FILE="${ENV##.*/}"
	FILE="${FILE%%.example}"

	echo "Copying $ENV to $DIR$FILE"
	cp $ENV $DIR$FILE
done

