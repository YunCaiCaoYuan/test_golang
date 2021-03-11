package test

/*
#cgo pkg-config: python3-embed
#include <Python/Python.h>

void Test() {
   Py_Initialize();
   PyRun_SimpleString("import sys");
   PyRun_SimpleString("sys.path.append('./')");

   PyObject *pmodule = PyImport_ImportModule("hello");
   if (pmodule == NULL) {
       perror("pmodule is null");
   }

   PyObject *pfunc = PyObject_GetAttrString(pmodule, "get_response");
   if (pfunc == NULL) {
       perror("pmodule is null");
   }
   printf("pfunc = %p\n", pfunc);

   PyObject_CallFunction(pfunc, NULL);
}
*/
import "C"

//#cgo LDFLAGS: -L/usr/local/Cellar/python@3.9/3.9.1_6/Frameworks/Python.framework/Versions/3.9/lib -lpython3.9

func Test() {
	C.Test()
}

//func main()  {
//	Test()
//}
