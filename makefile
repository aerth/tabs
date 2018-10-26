tablint:
	go build -o tablint ./cmd/tablint

clean:
	rm -vf tablint

test:
	go test ./...

install:
	go install ./cmd/tablint

.PHONY: tablint clean
