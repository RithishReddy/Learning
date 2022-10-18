package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseData struct {
	Data    map[string]interface{}
	Support map[string]string
}

func getreq() {

	fmt.Println("\nGet req Function")


	resp, err := http.Get("https://reqres.in/api/users/1")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data ResponseData

	json.Unmarshal(body, &data)

	// fmt.Println(data)

	//iterating Data
	for k,v:=range data.Data{
		fmt.Printf("%v : %v \n",k,v)
	}

	//iterating support
	for k,v:=range data.Support{
		fmt.Printf("%v : %v \n",k,v)
	}

}
