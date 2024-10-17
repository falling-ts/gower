#!/bin/bash

go test -tags prod,tmpl,static
# go test -bench=Benchmark -tags prod,tmpl,static
$1 run
