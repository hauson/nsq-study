#!/bin/sh

find . -name "*.go" | xargs -n1 -I {} awk -f ifmt.awk {}