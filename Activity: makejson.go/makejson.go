package main

import(
	"encoding/json"
	"fmt"
)

func main() {
person := make(map[string]string)

var name string
fmt.Println("enter ur name: ")
fmt.Scanln(&name)

var address string
fmt.Println("enterur address: ")
fmt.Scanln(&address)

person["name"] = name
person["address"] = address
jsonData, err := json.Marshal(person)
if err!= nil{
	fmt.Println(err)
	return
}
fmt.Println("JSON Object: ")
fmt.Println(string(jsonData))
}