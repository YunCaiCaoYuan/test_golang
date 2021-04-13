package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
	static char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	static char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	static char* os = "linux";
#else
#	error(unknown os)
#endif
*/
import "C"

func main() {
	print(C.GoString(C.os))
}
