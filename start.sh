#!/bin/bash

#cd cmd && go build -o GWorkerPool
make
./build/GWorkerPool/GWorkerPool -v=5 -logtostderr=true -config ./examples/simple.yml >> ./GWorkerPool.log 2>&1 &
