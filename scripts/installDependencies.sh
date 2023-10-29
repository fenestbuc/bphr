cd $HOME/fabric-samples/bphr/chaincode
go mod init
go get github.com/hyperledger/fabric-chaincode-go/shim
go get github.com/hyperledger/fabric-protos-go/peer
export GO111MODULE=on