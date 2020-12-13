package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	PDF_TITLE              = 0
	ROW_SEQUENCE_NUMBER    = 1
	PAGE_NUMBER            = 2
	TOTAL_PAGES            = 3
	DOCUMENT_TITLE         = 4
	LANGUAGE               = 5
	THB                    = 6
	WORDING                = 7
	OPTION_SELECTED        = 8
	OPTION_NOT_SELECTED    = 9
	MISSPELLED_WORDS_FIXED = 10
)

type ReadExcelExtension interface {
	ReadXlsx() []ExceFormatData
	ReadXls() []ExceFormatData
	CountXlsx() int64
}

type ExtensionHandler struct {
	filePath string
}

func (exHandle *ExcelHandler) ExcelExtension() ReadExcelExtension {
	return exHandle
}

func (ex *ExcelHandler) ReadXlsx() []ExceFormatData {
	f, err := excelize.OpenFile(ex.filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	firstSheet := f.GetSheetName(0)
	rows, err := f.GetRows(firstSheet)

	var a = 0
	var result []ExceFormatData
	for index, row := range rows {
		a = a + 1
		isFirstRowHeader := index == 0
		if isFirstRowHeader {
			continue
		}

		var excelData ExceFormatData
		for colIndex, colCell := range row {
			switch colIndex {
			case PDF_TITLE:
				excelData.PdfTitle = colCell
			case ROW_SEQUENCE_NUMBER:
				excelData.RowSequenceNumber = colCell
			case PAGE_NUMBER:
				excelData.PageNumber = colCell
			case TOTAL_PAGES:
				excelData.PageNumber = colCell
			case DOCUMENT_TITLE:
				excelData.DocumentTitle = colCell
			case LANGUAGE:
				excelData.Language = colCell
			case THB:
				excelData.Type = colCell
			case WORDING:
				excelData.Wording = colCell
			case OPTION_SELECTED:
				excelData.OptionSelected = colCell
			case OPTION_NOT_SELECTED:
				excelData.OptionNotSelected = colCell
			case MISSPELLED_WORDS_FIXED:
				excelData.Misspelled = colCell
			}
		}
		result = append(result, excelData)
	}

	return result
}

func (ex *ExcelHandler) ReadXls() []ExceFormatData {
	var excelData []ExceFormatData
	return excelData
}

func (ex *ExcelHandler) CountXlsx() int64 {
	f, err := excelize.OpenFile(ex.filePath)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	firstSheet := f.GetSheetName(0)
	rows, err := f.GetRows(firstSheet)

	return int64(len(rows))
}
