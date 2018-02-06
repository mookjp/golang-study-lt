#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)
echo "=== http://gopl.io"
go run ${SCRIPT_DIR}/fetch.go http://gopl.io
echo "=== http://bad.gopl.io"
go run ${SCRIPT_DIR}/fetch.go http://bad.gopl.io
echo "=== gopl.io"
go run ${SCRIPT_DIR}/fetch.go gopl.io
