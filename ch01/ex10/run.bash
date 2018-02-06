#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)
go run ${SCRIPT_DIR}/fetchall.go https://golang.org http://gopl.io https://godoc.org
