#!/bin/bash

go test -tags tmpl,static
# go test -bench=Benchmark -tags tmpl,static
$1 run
