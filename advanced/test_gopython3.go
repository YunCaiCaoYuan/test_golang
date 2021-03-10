package main

// !!! 未通过
// python3 
import (
	"fmt"
	"github.com/DataDog/go-python3"
	"os"
)

func main() {
	/*
	i, err := python3.Py_Main(os.Args)
	if err != nil {
		fmt.Printf("error launching the python interpreter: %s\n", err)
		os.Exit(1)
	}
	if i == 1 {
		fmt.Println("The interpreter exited due to an exception")
		os.Exit(1)
	}
	if i == 2 {
		fmt.Println("The parameter list does not represent a valid Python command line")
		os.Exit(1)
	}
	 */

	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}
	v := ImportModule("/Users/yourname/Desktop/lab/pppython", "value_pp")

	b := v.GetAttrString("estimate_pp")
	fmt.Printf("[FUNC] b = %#v\n", b)
	bArgs := python3.PyTuple_New(1)
	python3.PyTuple_SetItem(bArgs, 0, python3.PyUnicode_FromString("/Users/yourname/Desktop/lab/pppython/srcdata/67.csv"))
	re := b.Call(bArgs, python3.Py_None)
	fmt.Println("IsCallable: ", python3.PyCallable_Check(b))
	re1 := python3.PyTuple_GetItem(re, 0)
	re2 := python3.PyTuple_GetItem(re, 1)
	fmt.Println("re1:", python3.PyLong_AsLong(re1))
	fmt.Println("re2:", python3.PyLong_AsLong(re2))
	python3.Py_Finalize()
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python3.PyObject {
	sysModule := python3.PyImport_ImportModule("sys") // import sys
	fmt.Printf("sysModule=%v\n", sysModule)
	path := sysModule.GetAttrString("path")                    // path = sys.path
	fmt.Printf("path=%v\n", path)
	err := python3.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)
	fmt.Printf("err=%v, path=%v\n", err, path)
	return python3.PyImport_ImportModule(name)            // return __import__(name)
}
