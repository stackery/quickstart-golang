#!/bin/sh
mkdir -p .aws-sam/build/src/
cp .stackery/template.yaml .aws-sam/build/template.yaml

export GOPATH=$PWD

( cd src/getItems && make ) 
mkdir -p .aws-sam/build/src/getItems
cp src/getItems/main .aws-sam/build/src/getItems/main

( cd src/postItem && make ) 
mkdir -p .aws-sam/build/src/postItem
cp src/postItem/main .aws-sam/build/src/postItem/main
