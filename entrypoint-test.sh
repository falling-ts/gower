#!/bin/bash

go test -bench=Benchmark -tags test,tmpl,static
$1 run
