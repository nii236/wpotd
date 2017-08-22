package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response is the response from the Bing server
type Response struct {
	Images []struct {
		Startdate     string        `json:"startdate"`
		Fullstartdate string        `json:"fullstartdate"`
		Enddate       string        `json:"enddate"`
		URL           string        `json:"url"`
		Urlbase       string        `json:"urlbase"`
		Copyright     string        `json:"copyright"`
		Copyrightlink string        `json:"copyrightlink"`
		Quiz          string        `json:"quiz"`
		Wp            bool          `json:"wp"`
		Hsh           string        `json:"hsh"`
		Drk           int           `json:"drk"`
		Top           int           `json:"top"`
		Bot           int           `json:"bot"`
		Hs            []interface{} `json:"hs"`
	} `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

const getURL = "http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"

func main() {
	url, err := getWPOTD()
	if err != nil {
		log.Println(err)
		return
	}

	data, err := getImage(url)
	if err != nil {
		log.Println(err)
		return
	}
	err = run(data)
	if err != nil {
		log.Println(err)
		return
	}
}

func getWPOTD() (string, error) {
	resp, err := http.Get(getURL)
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result := &Response{}
	json.Unmarshal(buf, result)
	return "https://www.bing.com" + result.Images[0].URL, nil
}

func getImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return buf, nil
}

func run(b []byte) error {
	f, err := os.Create("/usr/share/wallpapers/wpotd.jpg")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
