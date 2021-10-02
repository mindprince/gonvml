#!/bin/bash

FILES=$(go list ./...  | grep -v /vendor/)
go build $FILES
