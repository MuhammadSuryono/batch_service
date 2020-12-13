package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

type WriteHandle struct {
	FileName  string           `json:"file_name" binding:"required"`
	Split     int              `json:"split"`
	Header    []string         `json:"header" binding:"required"`
	IndexData []string         `json:"index_data" binding:"required"`
	Data      []ExceFormatData `json:"data" binding:"required"`
}

func (*ExcelHandler) WriteToExcel(writeHandle *WriteHandle) string {
	xlsx := excelize.NewFile()
	sheetName := "Sheet1"

	// Set sheetname
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	// Set Header
	header := writeHandle.Header
	if len(writeHandle.Header) <= 0 {
		return "Header not found!"
	}
	for i, head := range header {
		xlsx.SetCellValue(sheetName, fmt.Sprintf(toCharStr(i+1)+"%d", 1), head)
	}

	iData := writeHandle.IndexData
	data := writeHandle.Data
	var split = 0
	for index, objData := range data {
		split = split + 1
		if split == writeHandle.Split {
			err := xlsx.SaveAs("./upload/" + writeHandle.FileName + "_" + strconv.Itoa(index) + ".xlsx")
			if err != nil {
				return err.Error()
			}
			split = 0
		}
		for i, _ := range iData {
			xlsx.SetCellValue(sheetName, fmt.Sprintf(toCharStr(i+1)+"%d", index+2), objData.PdfTitle)
		}
	}

	return ""
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}
