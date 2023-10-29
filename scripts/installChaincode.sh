#!/bin/bash

set -e

export PATH=$PATH:~/fabric-samples/bin
export FABRIC_CFG_PATH=$HOME/fabric-samples/config

# Set peer environment variables for Org1 (change as per your network config)
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Install chaincode on Org1's peer
peer lifecycle chaincode install $HOME/fabric-samples/bphr/chaincode/bphr.tar.gz

# Environment variables for Org2MSP peer
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=$HOME/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

# Install chaincode on Org2MSP peer
peer lifecycle chaincode install $HOME/fabric-samples/bphr/chaincode/bphr.tar.gz

# Fetch the package ID
PACKAGE_ID=$(peer lifecycle chaincode queryinstalled --output json | jq -r '.installed_chaincodes[0].package_id')

echo $PACKAGE_ID > packageId.txt

# Install chaincode on other peers in other organizations (if any)
# ... (similar to above with updated environment variables for other peers)
