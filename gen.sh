#!/bin/bash

python3 codegen.py
go fmt query.g.go
