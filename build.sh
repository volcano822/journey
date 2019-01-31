#!/bin/sh

pwd=$(pwd)
export GOPATH=$pwd/../../../..
go install github.com/volcano822/journey

if [ ! -d "output" ]
then
    mkdir output
fi
cp -r resources output/
cp config.json output/
cp control.sh output/
cp $GOPATH/bin/journey output/

cd output
rm journey.tar.gz
tar -czf journey.tar.gz *
cd ..