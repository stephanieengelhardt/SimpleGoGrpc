build-osx: compile-proto
	cd main && \
	go build -o bank-osx

run-dev:
	./main/bank-osx -logLevel=1

run:
	./main/bank-osx -logLevel=0

build-linux: compile-proto
	cd main && \
	env GOOS=linux GOARCH=arm go build -o bank-linux

run-linux:
	./main/bank-linux -logLevel=0

run-linux-dev:
	./main/bank-linux -logLevel=1

clean:
	lsof -t -i tcp:9080 | xargs kill
	lsof -t -i tcp:9081 | xargs kill
	lsof -t -i tcp:9082 | xargs kill
	lsof -t -i tcp:9083 | xargs kill
	lsof -t -i tcp:9084 | xargs kill
	lsof -t -i tcp:9085 | xargs kill
	lsof -t -i tcp:9086 | xargs kill
	lsof -t -i tcp:9087 | xargs kill
	lsof -t -i tcp:9088 | xargs kill
	lsof -t -i tcp:9089 | xargs kill
	lsof -t -i tcp:9090 | xargs kill
	rm main/bank-osx
	rm main/bank-linux
	rm main/bank.pb.go

compile-proto:
	protoc -I. --go_out=plugins=grpc:. protos/bank.proto

