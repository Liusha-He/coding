package coding

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type MyStruct struct {
	A, B, C string
	D       float32
}

type MyFunc func(a string) string

func anonymous_call() MyFunc {
	return func(a string) string {
		return fmt.Sprintf("anonymous function called with arg %s", a)
	}
}

func main() {
	text := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(text)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	mydata := MyStruct{
		A: "a",
	}
	mydata.B = "b"

	fmt.Println(mydata)

	greet := anonymous_call()
	fmt.Println(greet("hello"))

	fmt.Println("========== Concurrency =========")
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://yahoo.com",
		"https://baidu.com",
		"https://linkedin.com",
	}

	startTime := time.Now()

	c := make(chan string)

	for _, url := range links {
		go pingUrl(url, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	exTime := time.Now().Sub(startTime).Seconds()

	fmt.Println("==============================")
	fmt.Println(fmt.Sprintf("the whole process took %f sec", exTime))
}

func pingUrl(url string, c chan string) {
	startTime := time.Now()
	_, err := http.Get(url)
	exTime := time.Now().Sub(startTime).Seconds()

	if err != nil {
		c <- fmt.Sprintf("%s is down", url)
		return
	}
	c <- fmt.Sprintf("%s is up, it took %f sec to ping it", url, exTime)
}
