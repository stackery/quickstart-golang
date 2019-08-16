#!/bin/sh

set -e
mkdir -p .aws-sam/build/src/

cp .stackery/template.yaml .aws-sam/build/template.yaml

( cd src/getItems && GOBIN=$PWD GOPATH=$PWD make )

rm -rf .aws-sam/build/src/getItems
cp -r src/getItems .aws-sam/build/src/getItems

( cd src/newItem && GOBIN=$PWD GOPATH=$PWD make )

rm -rf .aws-sam/build/src/newItem
cp -r src/newItem .aws-sam/build/src/newItem