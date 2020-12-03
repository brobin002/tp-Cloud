//a concurrent web crawler who explore  the link graph of the web in breadth-first order
//To test, go run crawler.go https://www.catie.fr
package main

import (
	"fmt"
	"net/http"
    "os"
	"log"
	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	var links []string
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	defer resp.Body.Close()
	doc, _ := html.Parse(resp.Body)
	
	var visitNode func(*html.Node)
	visitNode = func (n *html.Node)  {
		if n.Type == html.ElementNode && n.Data == "a"{
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Printf("%v\n", a.Val)
					links = append(links, a.Val)
				}
			}
		}
	}


	forEachNode(doc, visitNode, nil)
	return links, nil
}




//!-Extract

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//recursive call
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}



//!+
func main() {

	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	seen := make(map[string]bool)

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				worklist <- crawl(link)
			}
		} ()
	}

	for links := range worklist {
		for _,link := range links{
			if (seen[link] == false){
				unseenLinks <- link
				seen[link] = true
			}
		}
	} 
}
