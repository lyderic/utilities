#!/bin/bash

main() {
	if [ -z ${1} ] ; then
		usage
	fi
	gofile="${1}"
  # utility is gofile without the '.go' extension
	utility="${gofile%.*}"
	go build -v "${gofile}" && mv -v "${utility}" ~/bin
	if [ $? -eq 0 ] ; then
		echo "${utility} deployed"
	fi
}

usage() {
	echo "Usage: $(basename ${0}) <utility.go>"
	exit 23
}

main $@
