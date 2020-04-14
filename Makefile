build:
	go build cmd/excelpicker/excel-picker.go

clean:
	-rm excel-picker

test:
	go test

.PHONY: build clean test
