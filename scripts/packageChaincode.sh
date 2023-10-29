#!/bin/bash

set -e
export FABRIC_CFG_PATH=$HOME/fabric-samples/config
export PATH=$PATH:~/fabric-samples/bin

# Navigate to chaincode directory
cd $HOME/fabric-samples/bphr/chaincode

# Package chaincode
peer lifecycle chaincode package bphr.tar.gz --path . --lang golang --label bphr_1.0
