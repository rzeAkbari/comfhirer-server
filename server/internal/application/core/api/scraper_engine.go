package api

import "C"

import (
	"errors"
	"github.com/DataDog/go-python3"
	"os"
)

func Scrape(file []byte) (string, error) {
	defer python3.Py_Finalize()
	directory := os.Getenv("DIRECTORY")
	if directory == "" {
		return "", errors.New("DIRECTORY env variable is not passed")
	}
	internalPath := "/server/internal/application/core/api/lib"
	modulePath := directory + internalPath

	python3.Py_Initialize()
	python3.PyRun_SimpleString("import sys")
	python3.PyRun_SimpleString("import os")
	python3.PyRun_SimpleString("os.environ[\"DIRECTORY\"] = \"" + modulePath + "\"")
	python3.PyRun_SimpleString("sys.path.append(\"" + modulePath + "\")")
	pdfModule := python3.PyImport_ImportModule("pdf_reader")
	pdfFile := python3.PyByteArray_FromStringAndSize(string(file))
	result := pdfModule.CallMethodArgs("Do", pdfFile)

	return python3.PyUnicode_AsUTF8(result), nil
}
