#!/bin/bash
swagger generate client -f swagger.json
go build -o refactor ./cmd/refactor
mkdir tmp
mv refactor tmp/
mv -f tmp/* client/repository_management/*
