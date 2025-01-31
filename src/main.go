package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://malapi.io/winapi/"
const red = "\033[31m"
const green = "\033[32m"
const reset = "\033[0m"

func findAPICall(call string) {
	resp, err := http.Get(url + call)
	if err != nil {
		fmt.Println(red + "[!]" + reset + " Error fetching the URL: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(red + "[!]" + reset + " Error: Status code ", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(red + "[!]" + reset + " Error loading HTTP response body: ", err)
	}

	contents := doc.Find("div.content")
	if contents.Length() == 0 {
		fmt.Println(red + "[!]" + reset + " Not found " + call + " in malAPI")
		return
	}

	function := contents.Eq(0).Text()
	description := contents.Eq(1).Text()
	library := contents.Eq(2).Text()

	var tags string
	doc.Find("span.attack-container").Each(func(i int, s *goquery.Selection) {
		tags += s.Text() + " "		
	})

	documentation, _ := doc.Find("a").Last().Attr("href")

	fmt.Println(green + "\n\n[+]" + reset + " API CALL:  ", function)
	fmt.Println(green + "[+]" + reset + " DESCRIPTION:  ", description)
	fmt.Println(green + "[+]" + reset + " LIBRARY(DLL):  ", library)
	fmt.Println(green + "[+]" + reset + " TAGS:  ",  tags)
	fmt.Println(green + "[+]" + reset + "DOCUMENTATION:  ", documentation)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\tmalAPI-go - " + red + "Get information about a Windows API calls simply!" + reset)
		fmt.Println("\t\tUsage: ./malapi-go <API CALL>")
		return
	}

	call := os.Args[1]
	findAPICall(call)
}
