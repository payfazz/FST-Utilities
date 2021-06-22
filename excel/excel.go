package excel

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	errSheetTabConst         = "0 Sheet Tab"
	errSheetDataConst        = "0 Sheet Data"
	errSheetTabNotMatchConst = "Sheet tab and content length are not match (number of tab and content length must be exact)"

	sheet1Const = "Sheet1"
)

// Note : sheetTabName length has to be exact same as content length(number of sheet)
func CreateExcel(sheetTabName []string, content [][][]interface{}) (*bytes.Buffer, error) {

	// Initial handling
	if len(sheetTabName) <= 0 {
		return nil, errors.New(errSheetTabConst)
	}
	if len(content) <= 0 {
		return nil, errors.New(errSheetDataConst)
	}
	if len(sheetTabName) != len(content) {
		return nil, errors.New(errSheetTabNotMatchConst)
	}

	f := excelize.NewFile()

	// Create a new sheet.
	excelMainIdx := f.NewSheet(sheetTabName[0])
	for i, value := range sheetTabName {
		if i == 0 {
			continue
		}

		f.NewSheet(value)
	}
	f.DeleteSheet(sheet1Const)

	// Write data to Sheet
	for index, value := range content {
		if len(value) == 0 {
			// Skip to next sheet tab
			continue
		}
		// Set values
		for idx, row := range value {
			// f.SetSheetRow(sheetTabName[index], fmt.Sprintf("A%d", idx+1), &row)
			for idy, cell := range row {
				cl := ""
				if idy > 25 {
					cl = fmt.Sprintf("%c%c", 64+(idy/25), 64+(idy%25))
				} else {
					cl = fmt.Sprintf("%c", 65+idy)
				}
				f.SetCellValue(sheetTabName[index], fmt.Sprintf("%s%d", cl, idx+1), cell)
			}
		}
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(excelMainIdx)

	// Save xlsx file by the given path.
	// err := f.SaveAs("./Book1.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	buffer, _ := f.WriteToBuffer()

	return buffer, nil
}
