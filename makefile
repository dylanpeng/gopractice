#!/bin/bash

dep:
	cd src; ${ENVARG} go get ./...; cd -

p:
	mkdir -p src/lib/proto
	rm -rf src/lib/proto/*

	cd src; protoc -I ./protocol --go_out=. --go-grpc_out=. common.proto; cd -
	cd src; protoc -I ./protocol --gofast_out=plugins=grpc:. protocol_demo.proto; cd -

	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s@\"lib/proto/@\"gopractice/lib/proto/@"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s/,omitempty//"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1\" form:\"\1\"/g"
	ls src/lib/proto/*/*.pb.go | xargs sed -i -e "/force omitempty/{n;s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1,omitempty\"/g;}"

	rm -f src/lib/proto/*/*.pb.go-e