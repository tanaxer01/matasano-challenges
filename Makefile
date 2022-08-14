
all:
	go test -v ./...


set1:
	go test -v ./set1

set2:
	go test -v ./set2



.PHONY: set1 set2
