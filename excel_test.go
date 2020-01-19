package cellsexporter

import (
	"testing"
)

func TestSheet1(t *testing.T) {
	excel := getExcel()

	a1 := excel.GetCellValue("Sheet1", "A1")
	if a1 != "" {
		t.Error("Sheet1!A1 is not empty")
	}

	b2 := excel.GetCellValue("Sheet1", "B2")
	if b2 != "a" {
		t.Error("Sheet!B2 is wrong: ", b2)
	}

	b4 := excel.GetCellValue("Sheet1", "B4")
	if b4 != "1" {
		t.Error("Sheet1!B4 is wrong: ", b4)
	}

	d3 := excel.GetCellValue("Sheet1", "D3")
	if d3 != "あ" {
		t.Error("Sheet1!D3 is wrong: ", d3)
	}

	d6 := excel.GetCellValue("Sheet1", "D6")
	expected := `aa	bb	cc
1
22`
	if d6 != expected {
		t.Error("Sheet1!D6 is wrong: ", d6)
	}
}

func TestSheet2(t *testing.T) {
	excel := getExcel()

	a1 := excel.GetCellValue("Sheet2", "A1")
	if a1 != "" {
		t.Error("Sheet2!A1 is not empty")
	}

	b6 := excel.GetCellValue("Sheet2", "B6")
	if b6 != "4444" {
		t.Error("Sheet2!B6 is wrong: ", b6)
	}

	c3 := excel.GetCellValue("Sheet2", "C3")
	if c3 != "eee" {
		t.Error("Sheet2!C3 is wrong, ", c3)
	}
}

func TestSheet3(t *testing.T) {
	excel := getExcel()

	a1 := excel.GetCellValue("Sheet3", "A1")
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
