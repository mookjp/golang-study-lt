#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)

# TODO: daemonize
go run ${SCRIPT_DIR}/main.go