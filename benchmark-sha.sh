#!/usr/bin/env bash
# gotesting_test.sh
# Described at https://wilsonmar.github.io/golang
# Based on code from https://app.pluralsight.com/library/courses/go-create-test-applications/exercise-files
#    by Michael Van Sickle (@vansimke)

# Install https://graphviz.org/download

cd benchmark-sha
go test -bench . -benchtime 10s -benchmem
#              . to run all test functions.
# EXPECTED RESPONSE: ns/op = nanosections per operation.
# goos: darwin
# goarch: amd64
# pkg: gotesting
# BenchmarkSHA1-12           	88175071	       140 ns/op	       0 B/op	       0 allocs/op
# BenchmarkSHA256-12         	57695949	       206 ns/op	       0 B/op	       0 allocs/op
# BenchmarkSHA512-12         	45616371	       269 ns/op	       0 B/op	       0 allocs/op
# BenchmarkSHA512Alloc-12    	31149062	       389 ns/op	     320 B/op	       2 allocs/op
# PASS
# ok  	gotesting	47.868s

# Profile run:
if ! command -v graphviz >/dev/null; then  # command not found, so:
    brew install graphviz   # os version specific!
       # dependencies include gd, pcre, sqlite, python 3.8-3.8.5, glib, jasper, netpbm, gts, libtool, pixman, ciaro, fribidi, graphite2, icu4c, harfbuzz, pango, graphviz
fi

go test -bench Alloc -memprofile profile.out
go tool pprof profile.out <<EOF
help
svg
EOF

# go test -typeblockprofile .
   # types: block, cover, spu, mem, mutex

# NOTE: gotesting.test and profile.out are in .gitignore
