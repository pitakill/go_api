#!/usr/bin/env sh

find . ! -path '*.git*' ! -path '*backend-config*' ! -path './config' -type d -exec go test {} \;
