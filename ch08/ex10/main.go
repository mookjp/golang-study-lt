package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML %v", url, err)
	}

	var links []string
	// vistNode は a タグの href の値を links に追加する
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// node に pre, post 処理を実行する
// node の 子要素に対する同操作を再帰的に実行する。
// 子要素が終わったら次の要素を対象にする
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	// Nodeの子要素を forEachNode にかける
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// 引数を worklist へ送信
	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			// unseenLinks から crawl へ link を渡す
			// crawl した link を worklink 送信
			for link := range unseenLinks {
				foundlinks := crawl(link)
				go func() {
					worklist <- foundlinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	// worklist から受信
	for list := range worklist {
		// 受信した link が seen map になかったら追加する
		// unseenLinks へ追加して 上記の goroutine に処理させる
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

	}
}
