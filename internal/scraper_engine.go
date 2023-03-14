package internal

import "C"

import (
	"github.com/DataDog/go-python3"
)

func Do(file []byte) string {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	python3.PyRun_SimpleString("import sys")
	python3.PyRun_SimpleString("sys.path.insert(0, \"./lib\")")
	pyobj := python3.PyImport_ImportModule("simple")
	result := pyobj.CallMethodArgs("scrape")
	return python3.PyUnicode_AsUTF8(result)
}
