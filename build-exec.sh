#!/bin/sh

# ------------------------------------
# Purpose:
# - Builds executables/binaries (Certificate Client).
#
# Releases:
# - v1.0.0 - 2024/09/26: initial release
#
# Remarks:
# - go tool dist list
# ------------------------------------

# set -v -o xtrace
set -v -o verbose

# renew vendor content
go mod vendor

# lint
golangci-lint run --no-config --enable gocritic

# compile 'aix'
env GOOS=aix GOARCH=ppc64 go build -v -o build/aix-ppc64/gpxdist

# compile 'darwin'
env GOOS=darwin GOARCH=amd64 go build -v -o build/darwin-amd64/gpxdist
env GOOS=darwin GOARCH=arm64 go build -v -o build/darwin-arm64/gpxdist

# compile 'dragonfly'
env GOOS=dragonfly GOARCH=amd64 go build -v -o build/dragonfly-amd64/gpxdist

# compile 'freebsd'
env GOOS=freebsd GOARCH=amd64 go build -v -o build/freebsd-amd64/gpxdist
env GOOS=freebsd GOARCH=arm64 go build -v -o build/freebsd-arm64/gpxdist

# compile 'illumos'
env GOOS=illumos GOARCH=amd64 go build -v -o build/illumos-amd64/gpxdist

# compile 'linux'
env GOOS=linux GOARCH=amd64 go build -v -o build/linux-amd64/gpxdist
env GOOS=linux GOARCH=arm64 go build -v -o build/linux-arm64/gpxdist
env GOOS=linux GOARCH=mips64 go build -v -o build/linux-mips64/gpxdist
env GOOS=linux GOARCH=mips64le go build -v -o build/linux-mips64le/gpxdist
env GOOS=linux GOARCH=ppc64 go build -v -o build/linux-ppc64/gpxdist
env GOOS=linux GOARCH=ppc64le go build -v -o build/linux-ppc64le/gpxdist
env GOOS=linux GOARCH=riscv64 go build -v -o build/linux-riscv64/gpxdist
env GOOS=linux GOARCH=s390x go build -v -o build/linux-s390x/gpxdist

# compile 'netbsd'
env GOOS=netbsd GOARCH=amd64 go build -v -o build/netbsd-amd64/gpxdist
env GOOS=netbsd GOARCH=arm64 go build -v -o build/netbsd-arm64/gpxdist

# compile 'openbsd'
env GOOS=openbsd GOARCH=amd64 go build -v -o build/openbsd-amd64/gpxdist
env GOOS=openbsd GOARCH=arm64 go build -v -o build/openbsd-arm64/gpxdist

# compile 'solaris'
env GOOS=solaris GOARCH=amd64 go build -v -o build/solaris-amd64/gpxdist

# compile 'windows'
env GOOS=windows GOARCH=amd64 go build -v -o build/windows-amd64/gpxdist.exe
env GOOS=windows GOARCH=386 go build -v -o build/windows-386/gpxdist.exe
env GOOS=windows GOARCH=arm go build -v -o build/windows-arm/gpxdist.exe
