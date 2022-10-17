package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name      string
	Job       string
	Id        string
	CreatedAt string
}

func postreq(data map[string]string) {


	fmt.Println("\nPost req Function")
	details, _ := json.Marshal(data)

	response, err := http.Post(usersApi, contentType, bytes.NewBuffer(details))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res User

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(res.Id, res.Name, res.Job, res.CreatedAt)
}
