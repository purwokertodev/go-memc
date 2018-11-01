package main

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p *Person) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func main() {
	mc := memcache.New("127.0.0.1:11211")
	setSimpleValue(mc)
	setMultipleValue(mc)
	setJSONValue(mc)
}

func setSimpleValue(mc *memcache.Client) {

	err := mc.Set(&memcache.Item{Key: "a", Value: []byte("wury")})

	if err != nil {
		fmt.Println(err)
	}

	val, err := mc.Get("a")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(val.Value))
}

func setMultipleValue(mc *memcache.Client) {

	var err error
	err = mc.Set(&memcache.Item{Key: "1", Value: []byte("wury")})
	err = mc.Set(&memcache.Item{Key: "2", Value: []byte("yanto")})

	if err != nil {
		fmt.Println(err)
	}

	vals, err := mc.GetMulti([]string{"1", "2"})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range vals {
		fmt.Println(string(v.Value))
	}
}

func setJSONValue(mc *memcache.Client) {

	person := &Person{ID: "U1", Name: "Wuriyanto"}
	personJSON, _ := person.JSON()

	err := mc.Set(&memcache.Item{Key: person.ID, Value: personJSON})

	if err != nil {
		fmt.Println(err)
	}

	val, err := mc.Get("U1")
	if err != nil {
		fmt.Println(err)
	}

	var personResult Person
	err = json.Unmarshal(val.Value, &personResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(personResult)
}
