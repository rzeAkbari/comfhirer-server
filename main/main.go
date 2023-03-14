package main

import (
	"fmt"
	"github.com/DataDog/go-python3"
)

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	python3.PyRun_SimpleString("import sys")
	python3.PyRun_SimpleString("sys.path.insert(0, \"./lib\")")
	pyobj := python3.PyImport_ImportModule("simple")
	result := pyobj.CallMethodArgs("scrape")
	fmt.Print(python3.PyUnicode_AsUTF8(result))
}
