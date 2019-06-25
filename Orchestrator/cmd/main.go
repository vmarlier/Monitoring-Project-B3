package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"monitoring/orchestrator/internal/app/config"
	"net/http"
	"net/url"

	"github.com/jmcvetta/napping"
)

func main() {
	/*	c := cron.New()
		c.AddFunc("@every 5s", readConf)
		c.Start()

		// For testing don't forget to uncomment this line
		// For deployment, don't forget to comment this line
		//defer c.Stop()
		time.Sleep(time.Second * 30)
	*/

	testKill()
}

func testKill() {
	url := "http://localhost:8001/process/kill"
	fmt.Println("URL:>", url)

	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "test")
	s.Header = h

	var jsonStr = []byte(`
{
	"pid": "35141"
}`)

	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Status:", resp.Status())
	fmt.Println("response Headers:", resp.HttpResponse().Header)
	fmt.Println("response Body:", resp.RawText())
}

func readConf() {
	conf := config.ParseConf("./configs/config.json")

	resp, _ := http.PostForm(conf.Host,
		url.Values{"key": {"Value"}, "disk": {"/"}})

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
