#!/bin/bash

for dir in $(find . -type dir -depth 1 | grep -v \.git); do
	cd $dir
	go test
	cd ..
done
