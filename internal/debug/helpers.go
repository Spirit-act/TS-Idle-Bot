package debug

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func DumpStruct(obj any) {
	fmt.Printf("%+v", obj)
}

func DdStruct(obj any) {
	fmt.Printf("%+v", obj)
	os.Exit(1)
}

func DumpStructFancy(obj any) {
	s, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Println(string(s))
}

func DdStructFancy(obj any) {
	s, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Println(string(s))
	os.Exit(1)
}

func DumpResBody(res *http.Response) {
	x, _ := io.ReadAll(res.Body)
	println(string(x))
}

func DdResBody(res *http.Response) {
	x, _ := io.ReadAll(res.Body)
	println(string(x))
	os.Exit(1)
}

func DumpRequest(req *http.Request) {
	res, err := httputil.DumpRequest(req, true)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(res))
}

func DdRequest(req *http.Request) {
	res, err := httputil.DumpRequest(req, true)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(res))

	os.Exit(1)
}
