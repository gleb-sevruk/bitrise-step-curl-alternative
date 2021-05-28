package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	briteapps__build_id := os.Getenv("briteapps_build_id")
	aws_aab_location := os.Getenv("aws_aab_location")
	api__url := os.Getenv("api_url")

	//aws_build__internal_id := "intuitive_web_solutions/2020-11-04_18-16-13_ee806e7a-cd50-4b8f-90fa-619440b775e8"
	println(briteapps__build_id)
	println(api__url)
	fmt.Println("This is the value specified for the input 'api__url':", api__url)
	fmt.Println("This is the value specified for the input 'aws_aab_location':", aws_aab_location)

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	// "http://0.0.0.0:8081/event-sink/"
	url := api__url
	fmt.Println("URL:>", url)
	var jsonStr = []byte(get_request_string(briteapps__build_id, aws_aab_location))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}

func get_aws_build_bucket() string {
	aws_build_bucket := "briteapps-builds-output"
	return aws_build_bucket
}

func get_request_string(build_id string, aab_location string) string {

	return fmt.Sprintf(`
	{
		"api_key": "cat-with-dog-head",
		"event_type":"builds.androidBundleBuilt",
		"body": {
				"build_id": "%s",
				"aab_location": "%s"
				}
	}
`, build_id, aab_location)
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}
