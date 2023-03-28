#!/bin/bash

go test -tags test,tmpl,static
go test -bench=Benchmark -tags test,tmpl,static
$1 run
