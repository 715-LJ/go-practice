package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", err)
		return
	}

	// 从文件中读取 JSON 数据
	file, err := os.Open(wd + "/json/data.json")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	// 解码 JSON 数据到结构体
	var personMap map[string]Person
	err = json.Unmarshal(data, &personMap)
	if err != nil {
		fmt.Println("Failed to decode JSON:", err)
		return
	}

	personList := make([]Person, 0, len(personMap))
	for _, person := range personMap {
		personList = append(personList, person)
	}

	fmt.Println(personList)
	for _, person := range personList {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("Email:", person.Email)
	}
}
