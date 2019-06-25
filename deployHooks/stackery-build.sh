#!/bin/sh

echo --------------------------------------------------  got here in stackery-build.sh

mkdir -p .aws-sam/build/src/
( cd src/Function && GOPATH=$PWD && make ) 
cp .stackery/template.yaml .aws-sam/build/template.yaml
cp -r src/Function .aws-sam/build/src/Function
