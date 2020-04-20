#!/bin/bash

main() {
	if [ -z ${1} ]; then
		usage
	fi
	utility=${1%.*}
	echo "Utility: ${utility}"
	go build -v -o "${HOME}/bin/${utility}" "${utility}.go"
}

usage() {
	echo "Usage: install-local.sh <utility.go>"
	exit 23
}

main $@
