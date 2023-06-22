# Install protobuf compiler (protoc)
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
rm protoc-3.15.8-linux-x86_64.zip

# Install go language binding for protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

# Install go rRPC server scaffolder 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
