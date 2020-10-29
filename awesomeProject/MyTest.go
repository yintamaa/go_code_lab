package main

import (
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"
	//"unsafe"

)
var mp map[string] int
//mp := make(map[string] int)
func checkProblem(pos int, str string) bool {
	if str[pos] == 'c' && str[pos+1] == 'o' && str[pos+2] == 'n' && str[pos+3] == 't' && str[pos+4] == 'e' && str[pos+5] == 's' && str[pos+6] == 't' && str[pos+7] == '/'{
		return true
	}else {
		return false
	}
}

func checkAccept(pos int, str string) bool {
	if str[pos] == 'A' && str[pos+1] == 'c' && str[pos+2] == 'c' && str[pos+3] == 'e' &&str[pos+4] == 'p'{
		return true
	}else {
		return false
	}
}

func connect(id int, name string) string{
	var url string
	if id == 1 { // codeforces
		url = "https://codeforces.com/submissions/" + name
	} else {

	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http error",err)
		return ""
	}

	//data := make([]byte, 1025)
	//reader := bufio.NewReader(res.Body)
	//reader.Read(data)
	////stringPointer := (*string)(unsafe.Pointer(&data)) // pointer to save memory
	//str := string(data)
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c",str[i])
	//}


	// waste memory
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("read error",err)
	}
	str := string(body)
	return str
}
func codeForces(name string) int{
	str := connect(1, name)
	//var lastProblem string
	var tmpStr string
	var sum int
	for i := 0; i < len(str); i++ {
		if checkProblem(i, str) { //
			tmpStr = ""
			for j := i + 8; ; j++ {
				tmpStr += string(str[j])
				if str[j] == '/' {
					tmpStr += string(str[j+9])
					break;
				}
			}
		}else if(checkAccept(i, str) && tmpStr != "") {
			if mp[tmpStr] == 0 {
				mp[tmpStr] = 1
				sum++
			}
		}
	}
	return sum
	//}
	//fmt.Println(str)
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

	mp = make(map[string] int)
	fmt.Printf("codeforces: %d",codeForces("yintama"))
	//mp["123"] = 456
	//fmt.Println(mp["123"])
}