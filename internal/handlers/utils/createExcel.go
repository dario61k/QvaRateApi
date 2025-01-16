package utils

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateExcel(data [][]interface{}) ([]byte, error) {

	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("error closing excel")
		}
	}()

	index, err := f.NewSheet("Tasa de cambio")
	if err != nil {
		return nil, errors.New("error creating a new sheet")
	}

	if err := f.DeleteSheet("Sheet1"); err != nil {
		return nil, errors.New("error creating excel file")
	}

	for idx, row := range data {

		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			return nil, fmt.Errorf("coordination with cell %d failed", idx+1)
		}
		f.SetSheetRow("Tasa de cambio", cell, &row)
	}

	if err := f.SetColWidth("Tasa de cambio", "A", "A", 15); err != nil {
		return nil, fmt.Errorf("error setting column width: %v", err)
	}

	f.SetActiveSheet(index)

	var file bytes.Buffer

	if err := f.Write(&file); err != nil {
		return nil, errors.New("error creating buffer")
	}

	return file.Bytes(), nil
}
