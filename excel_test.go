package excelexporter

import (
	"testing"
)

func TestToGetCellValuesFromSheet(t *testing.T) {
	excel := getExcel()

	a1, err := excel.GetCellValue("Sheet1", "A1")
	if err != nil {
		panic(err)
	}
	if a1 != "" {
		t.Error("Sheet1!A1 is not empty")
	}

	b2, err := excel.GetCellValue("Sheet1", "B2")
	if err != nil {
		panic(err)
	}
	if b2 != "a" {
		t.Error("Sheet!B2 is wrong: ", b2)
	}

	b4, err := excel.GetCellValue("Sheet1", "B4")
	if err != nil {
		panic(err)
	}
	if b4 != "1" {
		t.Error("Sheet1!B4 is wrong: ", b4)
	}

	d3, err := excel.GetCellValue("Sheet1", "D3")
	if err != nil {
		panic(err)
	}
	if d3 != "あ" {
		t.Error("Sheet1!D3 is wrong: ", d3)
	}

	d6, err := excel.GetCellValue("Sheet1", "D6")
	if err != nil {
		panic(err)
	}
	expected := `aa	bb	cc
1
22`
	if d6 != expected {
		t.Error("Sheet1!D6 is wrong: ", d6)
	}
}

func TestToGetCellValuesFromOtherSheet(t *testing.T) {
	excel := getExcel()

	a1, err := excel.GetCellValue("Sheet2", "A1")
	if err != nil {
		panic(err)
	}
	if a1 != "" {
		t.Error("Sheet2!A1 is not empty")
	}

	b6, err := excel.GetCellValue("Sheet2", "B6")
	if err != nil {
		panic(err)
	}
	if b6 != "4444" {
		t.Error("Sheet2!B6 is wrong: ", b6)
	}

	c3, err := excel.GetCellValue("Sheet2", "C3")
	if err != nil {
		panic(err)
	}
	if c3 != "eee" {
		t.Error("Sheet2!C3 is wrong, ", c3)
	}
}

func TestWithNotExistingSheet(t *testing.T) {
	excel := getExcel()

	a1, err := excel.GetCellValue("Sheet3", "A1")
	if err == nil {
		t.Error("no error occurred")
	}
	if a1 != "" {
		t.Error("Sheet3!A1 is not empty")
	}
}

func getExcel() *Excel {
	excel, err := LoadExcel("test.xlsx")
	if err != nil {
		panic(err)
	}
	return excel
}
