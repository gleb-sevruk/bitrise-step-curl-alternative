package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	aws_build_internal_id := os.Getenv("aws_build_internal_id")
	briteapps__build_id := os.Getenv("briteapps_build_id")
	api__url := os.Getenv("api_url")

	//aws_build__internal_id := "intuitive_web_solutions/2020-11-04_18-16-13_ee806e7a-cd50-4b8f-90fa-619440b775e8"
	println(briteapps__build_id)
	println(api__url)
	fmt.Println("This is the value specified for the input 'aws_build_internal_id':", aws_build_internal_id)
	fmt.Println("This is the value specified for the input 'api__url':", api__url)
	fmt.Println("This is the value specified for the input 'briteapps__build_id':", briteapps__build_id)

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:

	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}

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

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}
