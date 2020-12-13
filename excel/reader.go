package excel

import (
	"path/filepath"
	"strings"
)

type ExcelEngine interface {
	ReadPfLegal() []ExceFormatData
	CountRows() int64
	WriteToExcel(writeHandle *WriteHandle) string
	ExcelExtension() ReadExcelExtension
}

type ExcelHandler struct {
	filePath string
}

func ExcelReadWriteEngine(filePath string) ExcelEngine {
	return &ExcelHandler{
		filePath: filePath,
	}
}

func (ex *ExcelHandler) ReadPfLegal() []ExceFormatData {
	ext := strings.ToLower(filepath.Ext(ex.filePath))

	handleExtension := ExcelReadWriteEngine(ex.filePath).ExcelExtension()
	switch ext {
	case ".xlsx":
		return handleExtension.ReadXlsx()
	case ".xls":
		return handleExtension.ReadXls()
	default:
		return nil
	}
}

func (ex *ExcelHandler) CountRows() int64 {
	ext := strings.ToLower(filepath.Ext(ex.filePath))

	handleExtension := ExcelReadWriteEngine(ex.filePath).ExcelExtension()
	switch ext {
	case ".xlsx":
		return handleExtension.CountXlsx()
	case ".xls":
		return handleExtension.CountXlsx()
	default:
		return 0
	}
}
