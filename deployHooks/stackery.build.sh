#!/bin/sh
set -e
mkdir -p .aws-sam/build/src/

for function in src/*; do
  echo compiling $function
  ( cd $function && make )
  rm -rf .aws-sam/build/$function
  cp -r $function .aws-sam/build/$function
done
