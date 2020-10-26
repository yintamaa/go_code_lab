package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func codeForces(id string)  {
	url := "https://codeforces.com/submissions/" + id
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http error",err)
		return
	}
	data := make([]byte, 1025)
	reader := bufio.NewReader(res.Body)
	reader.Read(data)
	stringPointer := (*string)(unsafe.Pointer(&data)) // pointer to save memory
	for i := 0; i < len(*stringPointer); i++ {
		if (*stringPointer[i])
	}
	// waste memory
	//body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("read error",err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	//fmt.Println("Hello world")
	//res, err := http.Get("https://codeforces.com/submissions/yintama/page/1")
	//if err != nil {
	//	fmt.Println("http error",err)
	//	return
	//}
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("read error",err)
	//	return
	//}
	//fmt.Println(string(body))
}