build:
	go build cmd/excelpicker/excel-picker.go

clean:
	-rm excel-picker

test:
	go test

test-cmd:
	./excel-picker --file=test.xlsx --cell=Sheet1/B2
	./excel-picker --file=test.xlsx --cell=B2
	./excel-picker --file=test.xlsx --cell=B2 --cell=D3 --cell=B4 --cell=D6
	./excel-picker --file=test.xlsx --file=test2.xlsx --cell=B2 --cell=D3 --cell=B4 --cell=D6
	./excel-picker --file=test.xlsx --cell=B2  --type=csv

.PHONY: build clean test test-cmd
