#!/bin/bash

# Build the module
go build *.go 

# Copy it to the subcommands
cp ./dirtest ~/.config/portage/subcommands/test


