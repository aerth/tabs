tablint:
	go build -o tablint ./cmd/tablint

clean:
	rm -vf tablint

test:
	go test ./...

.PHONY: tablint clean
