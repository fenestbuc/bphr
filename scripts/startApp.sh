#!/bin/bash

set -e

# Navigate to app directory
cd $HOME/fabric-samples/bphr/application

# Start app
npm start
npm go get
