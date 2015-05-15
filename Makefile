all:
#	go build src/dtrace/dtrace.go
	rm -rf pkg
	env GOPATH=${PWD} GOBIN=${PWD} go install test.go
