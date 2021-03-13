package urls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type urlsSingleton struct {
	URLs map[string]interface{}
}

var instance *urlsSingleton
var once sync.Once

func ReadJsonFile(path string) []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	return byteValue
}

func ParseJson(jsonDate []byte) map[string]interface{} {
	var paths map[string]interface{}

	json.Unmarshal(jsonDate, &paths)

	urls := paths["paths"].(map[string]interface{})

	return urls
}

func InitShortUrlsData() map[string]interface{} {
	jsonDate := ReadJsonFile("urlsDB/shorten_urls.json")
	urls := ParseJson(jsonDate)
	//fmt.Println(urls["/go-gophers"].(string))
	return urls
}

func GetUrlDBInstance() *urlsSingleton {
	once.Do(func() {
		instance = &urlsSingleton{}
		instance.URLs = InitShortUrlsData()
	})
	return instance
}


