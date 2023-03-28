#!/bin/bash

go test -bench=Benchmark -tags prod,tmpl,static
$1 run
