package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkin(OPs []string, OP string) bool {

	var status bool
	status = false
	for _, op := range OPs {

		if op == OP {
			status = true
		}
	}
	return status

}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func response(www string) string {
	req, err := http.NewRequest("GET", www, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//print("Response Status:", resp.Status, "\n")
	pool, err := ioutil.ReadAll(resp.Body)
	text := string(pool[:])
	return text

}

func responseRPD(url string, data string) []byte {

	var username string = "mgtgup"
	var passwd string = "mgt2017"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+data, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	_ = s
	return bodyText
}
