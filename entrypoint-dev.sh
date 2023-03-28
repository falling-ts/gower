#!/bin/bash

go test -bench=Benchmark -tags tmpl,static
$1 run
