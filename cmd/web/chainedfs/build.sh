#install go-bindata
go get -u github.com/jteeuwen/go-bindata/...
#run go-bindata and then build
${GOPATH:-~/go}/bin/go-bindata ./static && go build