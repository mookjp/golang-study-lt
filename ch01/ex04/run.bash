#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)
go run ${SCRIPT_DIR}/dup.go ${SCRIPT_DIR}/input1 ${SCRIPT_DIR}/input2