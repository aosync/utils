#!/bin/sh

[ ! -d build ] && mkdir build

cd build

for d in ../commands/*; do
	echo $d
	go build $d
done
