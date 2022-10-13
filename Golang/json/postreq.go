package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type postData struct {
	Name      string
	Job       string
	Id        string
	CreatedAt string
}

func postreq() {

	fmt.Println("\nPost req Function")
	details, _ := json.Marshal(map[string]string{
		"name": "rithish",
		"job":  "developer",
	})

	response, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(details))

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(response.Body)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res postData

	json.Unmarshal(body, &res)

	// fmt.Println(string(body))

	fmt.Println(res.Id, res.Name, res.Job,res.CreatedAt)
}
