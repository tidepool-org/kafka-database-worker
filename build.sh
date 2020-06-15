#!/bin/sh -eu

rm -rf dist
mkdir dist
go build -o dist/kafka-database-worker kafka-database-worker.go
