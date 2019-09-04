#!/usr/bin/env bash

protoc --go_out=plugins=micro:. rpc.proto