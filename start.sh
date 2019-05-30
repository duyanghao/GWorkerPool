#!/bin/bash

cd cmd && go build -o GWorkerPools
./GWorkerPools -v=5 -logtostderr=true -config ../examples/simple.yml >> ../GWorkerPools.log 2>&1 &

