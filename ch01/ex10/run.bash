#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0);pwd)

# Directory to save output
OUTPUT_DIR="${SCRIPT_DIR}/tmp"
mkdir -p ${OUTPUT_DIR}

# Url
URL_01=https://cookpad.com/search/%E3%83%90%E3%83%AC%E3%83%B3%E3%82%BF%E3%82%A4%E3%83%B3
URL_02=https://cookpad.com/search/%E3%82%AC%E3%83%88%E3%83%BC%E3%82%B7%E3%83%A7%E3%82%B3%E3%83%A9

echo "=== First fetchall"
go run ${SCRIPT_DIR}/fetchall.go ${OUTPUT_DIR} ${URL_01} ${URL_02}
echo "=== Second fetchall"
go run ${SCRIPT_DIR}/fetchall.go ${OUTPUT_DIR} ${URL_01} ${URL_02}
