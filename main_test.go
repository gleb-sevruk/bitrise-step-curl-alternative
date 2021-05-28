package main

import (
	"bytes"
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	println("x")
	url := "http://0.0.0.0:8081/event-sink/"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(`
	{
		"api_key": "cat-with-dog-head",
		"event_type":"builds.androidBundleBuilt",
		"body": {
				"build_id": "30f64de1-8734-47de-ba80-bbe79af65d22",
				"aab_location": null
				}
	}
`)
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
}

//
//def remove_transparency_from_png(img_path, background_color):
//	with Image.open(img_path).convert('RGBA') as png:
//	with Image.new('RGBA', png.size, background_color) as background:
//	# draw logo on background with given color
//	alpha_composite = Image.alpha_composite(background, png)
//	with alpha_composite.convert('RGB') as converted:
//	converted.save(img_path)
