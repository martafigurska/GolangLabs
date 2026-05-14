package main

import (
	"fmt"
	"encoding/json"
	"encoding/xml"
    // "io"
    "os"
	// "bufio"
)

type Filter interface {
	ApplyFilter(data Data) Data
}

type R interface {
	Load(filename string) Data
}

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
	Prof string `json:"profession" xml:"profession"`
}

type Data struct {
	Users []Person `json:"users" xml:"users"`
}

type Employees struct {
	users Data
	f []Filter
}

func getData(r R, filename string) Data{
	return r.Load(filename)
}

type JS struct {
	Data
}

func (j JS) Load(filename string) Data {
	jsonFile, err := os.Open(filename)
	if err != nil {
    	fmt.Println(err)
	}
	defer jsonFile.Close()

	var data Data
	json.NewDecoder(jsonFile).Decode(&data)

	return data
}

type XML struct {
	Data
}

func (x XML) Load(filename string) Data {
	xmlFile, err := os.Open(filename)
	if err != nil {
    	fmt.Println(err)
	}
	defer xmlFile.Close()

	var data Data
	xml.NewDecoder(xmlFile).Decode(&data)

	return data
}

type TXT struct {
	Data
}

// func (t TXT) Load(filename string) Data {
// 	txtFile, err := os.Open(filename)
// 	if err != nil {
//     	fmt.Println(err)
// 	}
// 	scanner := bufio.NewScanner(txtFile)

// 	return data
// }

type Age struct {
	minAge int
	maxAge int
}

func (a Age) ApplyFilter(data Data) Data {
	var result []Person
	for _, u := range data.Users {
		if u.Age >= a.minAge && u.Age <= a.maxAge {
			result = append(result, u)
		}
	}
	return Data{Users: result}
}

type Needed struct {
	professions []string
}

// func (n Needed) ApplyFilter(data Data) Data {

// }



func (e Employees) Select() Data{
	result := e.users
	for _, filter := range e.f {
		result = filter.ApplyFilter(result)
	}

	return result
}


func main(){
	dataJS := Employees{users: getData(JS{}, "data.json"), f: []Filter{Age{30, 40}}}.Select()
	dataXML := Employees{users: getData(XML{}, "data.xml"), f: []Filter{Age{20, 30}}}.Select()
	// dataTXT := Employees{users: getData(TXT{}, "data.txt"), f: []Filter{Needed{[]string{"economist", "judge"}}}}.Select()
	
	fmt.Println(dataJS)
	fmt.Println("****************************************************")
	fmt.Println(dataXML)
}