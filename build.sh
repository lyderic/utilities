#!/bin/bash

export GO111MODULE=auto

DEST=/dev/shm
DEPLOY=false

main() {
	if [ -z ${1} ]; then
		usage
	fi
	if [ "${1}" == "-deploy" ] ; then
		DEPLOY=true
		shift
	fi
	utility=${1%.*}
	echo "Utility: ${utility}"
	echo '[building linux amd64]'
	build amd64 x86_64
	echo '[building linux arm]'
	build arm armv7l
	echo '[building linux arm64]'
	build arm64 aarch64
	if [ "${DEPLOY}" == "true" ] ; then
		echo '[ansible deployment]'
		copy-to-remote
	fi
}

copy-to-remote() {
	echo "ANSIBLE: utility=${utility}"
	cd ${HOME}/repositories/ansible
	ansible all -o -m setup
	ansible all -o -m copy -a "src=${DEST}/${utility}-{{ansible_architecture}} dest=~{{operator}}/bin/${utility} mode=0755"
}

build() {
	goarch=${1}
	ansiblearch=${2}
	GOARCH=${goarch} \
	go build -v -o "${DEST}/${utility}-${ansiblearch}" "${utility}.go"
	if [ ${goarch} == 'arm' ] ; then
		ln -sfv ${DEST}/${utility}-armv7l ${DEST}/${utility}-armv6l
	fi
}

usage() {
	echo "Usage: $(basename $0) [-deploy] <utility.go>"
	echo "build an executable for various architecture in ${DEST}"
	echo "  -deploy  do an ansible deployment on all machines in the inventory"
	exit 23
}

main $@
