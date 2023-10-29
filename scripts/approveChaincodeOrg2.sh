#!/bin/bash
export PATH=$PATH:~/fabric-samples/bin
export FABRIC_CFG_PATH=$HOME/fabric-samples/config

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

export ORDERER_CA=$HOME/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER_ADDRESS=localhost:7050
set -e

# Read the package ID
PACKAGE_ID=$(cat packageId.txt)

# Approve chaincode for Org2
peer lifecycle chaincode approveformyorg --channelID mychannel --name bphr --version 1.0 --package-id $PACKAGE_ID --sequence 1 --waitForEvent --tls --cafile $ORDERER_CA --orderer $ORDERER_ADDRESS