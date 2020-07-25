
# Setup

# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip

# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google

# Go:

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/golang/protobuf/proto

go get -u google.golang.org/grpc
