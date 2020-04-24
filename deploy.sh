#!/bin/bash

main() {
	if [ -z ${1} ]; then
		usage
	fi
	utility=${1%.*}
	echo "Utility: ${utility}"
	echo '[building linux amd64]'
	build amd64 x86_64
	echo '[building linux arm]'
	build arm armv7l
	echo '[ansible deployment]'
	copy-to-remote
}

copy-to-remote() {
	echo "ANSIBLE: utility=${utility}"
	cd ${HOME}/repositories/ansible
	ansible all -o -m setup
	ansible all -o -m copy -a "src=/dev/shm/${utility}-{{ansible_architecture}} dest=~{{operator}}/bin/${utility} mode=0755"
}

build() {
	goarch=${1}
	ansiblearch=${2}
	GOARCH=${goarch} \
	GOBIN=$HOME/go/bin \
	go build -v -o "/dev/shm/${utility}-${ansiblearch}" "${utility}.go"
	if [ ${goarch} == 'arm' ] ; then
		ln -sfv /dev/shm/${utility}-armv7l /dev/shm/${utility}-armv6l
	fi
}

usage() {
	echo "Usage: deploy.sh <utility.go>"
	exit 23
}

main $@
