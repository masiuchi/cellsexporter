build:
	go build cmd/excelpicker/excel-picker.go

clean:
	-rm excel-picker

test:
	go test

test-cmd:
	./excel-picker --file=test.xlsx --cell=Sheet1/B2
	./excel-picker --file=test.xlsx --cell=B2

.PHONY: build clean test test-cmd
