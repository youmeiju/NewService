package main

import (
	//"net/http"
	//"io/ioutil"
	//"encoding/json"
	"fmt"
)

//
//import (
//	"fmt"
//	//"net/http"
//	//"io/ioutil"
//	//"encoding/json"
//)
//
func main() {
	//m := make(map[string]interface{})
	url:= fmt.Sprintf("http://localhost:8080/v1/object/GetFirstMenu/%s","1")
	fmt.Printf(url)
	//res, _ := http.Get(strconv.Itoa(url))
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//err := json.Unmarshal(body, &m)
	//if err == nil {
	//	fmt.Println(m)
	//} else {
	//	fmt.Println(body, m, err)
	//}
}
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered in f", r)
//		}
//	}()
//f()
//}