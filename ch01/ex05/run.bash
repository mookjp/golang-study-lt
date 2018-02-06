#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)
go run ${SCRIPT_DIR}/lissajous.go > /tmp/out.gif