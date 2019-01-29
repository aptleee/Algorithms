package main

import "net/http"
import "fmt"
import "io/ioutil"

func main(){

	resp, err := http.Get("http://google.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := fmt.Sprint(body)
	fmt.Printf("%T",s)
}
