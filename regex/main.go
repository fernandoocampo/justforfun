package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r, err := regexp.Compile(`zero-to-one/labs/0([a-z0-9\-]+)-java/(.?)`)
	if err != nil {
		panic(err)
	}
	match := r.MatchString(
		"zero-to-one/labs/06-concurrency-java/path/path",
	)
	fmt.Println(match)

	match2 := r.MatchString(
		"zero-to-one/labs/06-concurrency-java/",
	)
	fmt.Println(match2)

	language := "java"
	// suffixLanguage := "-"+language
	pattern := `(.?)zero-to-one/labs/0([a-z0-9\-]+)-%s/(.?)`

	regexPattern := fmt.Sprintf(pattern, language)

	r2, err2 := regexp.Compile(regexPattern)
	if err2 != nil {
		panic(err2)
	}
	path := "/tmp/epic/resources/zero-to-one/labs/06-observability-java/deploy/chart/charts"
	fmt.Println("match?", r2.MatchString(path))
	path = "/tmp/epic/resources/zero-to-one/labs/03-database-java/deploy/chart/charts"
	fmt.Println("match?", r2.MatchString(path))
	path = "/tmp/epic/resources/zero-to-one-base/scripts"
	fmt.Println("match?", r2.MatchString(path))

	pattern2 := `-%s`
	regexPattern2 := fmt.Sprintf(pattern2, language)
	r3, err3 := regexp.Compile(regexPattern2)
	if err3 != nil {
		panic(err3)
	}

	path = "/tmp/epic/resources/zero-to-one/labs/06-observability-java/deploy/chart/charts"
	fmt.Println("match -java?", r3.MatchString(path))
	fmt.Println("let's replace", r3.ReplaceAllString(path, ""))
	path = "/tmp/epic/resources/zero-to-one/labs/03-database-java/deploy/chart/charts"
	fmt.Println("match -java?", r3.MatchString(path))
	fmt.Println("let's replace", r3.ReplaceAllString(path, ""))
	path = "/tmp/epic/resources/zero-to-one-base/scripts"
	fmt.Println("match -java?", r3.MatchString(path))
	fmt.Println("let's replace", r3.ReplaceAllString(path, ""))
	path = "/tmp/epic/resources/zero-to-one/labs/03-database-java/deploy/chart/charts"
	fmt.Println("let's replace using strings", strings.Replace(path, "-java", "", 1))
	path = "/tmp/epic/resources/zero-to-one-base/scripts"
	fmt.Println("let's replace using strings", strings.Replace(path, "-java", "", 1))

	pattern3 := `zero-to-one/labs/0([a-z0-9\-]+)(.?)`

	r4, err4 := regexp.Compile(pattern3)
	if err4 != nil {
		panic(err4)
	}
	path = "zero-to-one/labs/01-introduction/"
	fmt.Println("match?", r4.MatchString(path))
	path = "zero-to-one/labs/01-introduction"
	fmt.Println("match?", r4.MatchString(path))
}
