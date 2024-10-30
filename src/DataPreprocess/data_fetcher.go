package DataPreprocess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DataFetcher(url string, jsonBody interface{}) (err error) {
	method := "GET"
	client := &http.Client{}

	if req, err := http.NewRequest(method, url, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if err = json.Unmarshal(body, &jsonBody); err == nil {
					return nil
				}
			}
		}
	}
	fmt.Println(err)
	return err
}
