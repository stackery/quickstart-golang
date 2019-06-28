#!/bin/sh
mkdir -p .aws-sam/build/src/

if [ -f .stackery/.stackery.template.yaml ]; then
    cp .stackery/.stackery.template.yaml .aws-sam/build/template.yaml
else
    cp .stackery/template.yaml .aws-sam/build/template.yaml
fi

export GOPATH=$PWD

( cd src/getItems && make ) 
mkdir -p .aws-sam/build/src/getItems
cp src/getItems/main .aws-sam/build/src/getItems/main

( cd src/postItem && make ) 
mkdir -p .aws-sam/build/src/postItem
cp src/postItem/main .aws-sam/build/src/postItem/main
