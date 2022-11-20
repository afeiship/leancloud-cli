package misc

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

type LcResponseBody struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ObjectId  string `json:"objectId"`
}

func GetUrl(key string) string {
	return "https://8alqie7r.lc-cn-n1-shared.com/1.1/classes/options/" + key
}

func LcGet(optionKey string) LcResponseBody {
	url := GetUrl(optionKey)
	client := resty.New()
	// set header Lc_key, lc_id from os.get env
	lcId := os.Getenv("LC_ID")
	lcKey := os.Getenv("LC_KEY")
	client.SetHeaders(map[string]string{"X-LC-Id": lcId, "X-LC-Key": lcKey})
	resp, err := client.R().Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	var lcResponseBody LcResponseBody
	json.Unmarshal(resp.Body(), &lcResponseBody)
	return lcResponseBody
}

// create item by key/value
func LcPost(optionKey string, key, value string) LcResponseBody {
	url := GetUrl(optionKey)
	client := resty.New()
	resp, err := client.R().SetBody(map[string]string{
		"key":   key,
		"value": value,
	}).Post(url)
	if err != nil {
		log.Fatalln(err)
	}
	var lcResponseBody LcResponseBody
	json.Unmarshal(resp.Body(), &lcResponseBody)
	return lcResponseBody
}

func LcPut(optionKey string, key, value string) LcResponseBody {
	url := GetUrl(optionKey)
	client := resty.New()
	resp, err := client.R().SetBody(map[string]string{
		"key":   key,
		"value": value,
	}).Put(url)
	if err != nil {
		log.Fatalln(err)
	}
	var lcResponseBody LcResponseBody
	json.Unmarshal(resp.Body(), &lcResponseBody)
	return lcResponseBody
}

func LcDelete(optionKey string) {
	url := GetUrl(optionKey)
	client := resty.New()
	_, err := client.R().Delete(url)
	if err != nil {
		log.Fatalln(err)
	}
}
