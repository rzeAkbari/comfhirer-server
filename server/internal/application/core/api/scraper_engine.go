package api

import "C"

import (
	"github.com/DataDog/go-python3"
)

func Scrape(file []byte) string {
	defer python3.Py_Finalize()

	python3.Py_Initialize()
	python3.PyRun_SimpleString("import sys")
	python3.PyRun_SimpleString("sys.path.insert(0, \"./lib\")")

	pdfModule := python3.PyImport_ImportModule("pdf_reader")
	pdfFile := python3.PyByteArray_FromStringAndSize(string(file))
	result := pdfModule.CallMethodArgs("Do", pdfFile)

	return python3.PyUnicode_AsUTF8(result)
}
