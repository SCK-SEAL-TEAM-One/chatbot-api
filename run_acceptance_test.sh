#!/bin/bash
GOPACKAGES=chatbot

export GOPATH=$(pwd)
export PATH=$PATH:./node_modules/.bin

rm -rf ./bin/*
rm -rf ./log/*

go fmt ${GOPACKAGES}
go test ${GOPACKAGES} | tee ./log/chatbot_test.log
go install ${GOPACKAGES}
go build -o ./bin/main_chatbot_api main

./bin/main_chatbot_api &
newman run acceptance_test.postman_collection.json | tee ./log/acceptance_test.log
killall -9 main_chatbot_api
