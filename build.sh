#!/bin/sh

pwd=$(pwd)
export GOPATH=$pwd/../../../..
go install github.com/volcano822/journey

rm -rf output
mkdir output
cp -r resources output/
cp config.json output/
cp control.sh output/
cp $GOPATH/bin/journey output/

cd output
tar -czf journey.tar.gz *
