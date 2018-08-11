// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const outputDir = "output"
const contextTypeHeader = "Content-Type"
const contextTypeTextHtml = "text/html"

// TODO: url を getする
// TODO: 最初のpath（/）は自分でつくらないといけない
// TODO: htmlじゃなかったらpathで保存して終了
// TODO: htmlをパースする
// TODO: nodeがaでelementNodeだったらhrefがあるかチェックする
// TODO: hrefがあったらlinkをURL Parseして取得する
// TODO: hrefの対象がホストだったら頭にfile://を付与する
// TODO: hrefが / や # や hostURL（/）ではないかチェックする
// TODO: hrefが/で始まっている場合は host/... で保存する
// TODO: hrefが/で始まっていない場合は host/href で保存する
// TODO: nodeがaでelementNodeでimgでsrcがあるかチェックする
// TODO: srcのURLをパースする
// TODO: srcの内容を保存する
// TODO: nodeのFirstChildを取得する
// TODO: 上記のチェック操作を実行する
// TODO: nodeのnextSiblingを取得する
// TODO: 上記のチェック操作を実行する

// TODO: URLの指定
func main() {
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Fprintf(os.Stdout, "output directory exitsts. delete it")
			return
		}
	}
	if err := get("https://gopl.io"); err != nil {
		log.Fatal(err)
	}
}

func get(url string) error {
	fmt.Fprintf(os.Stdout, "url: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not get target. err: %v", err)
		return err
	}
	scheme := res.Request.URL.Scheme
	host := res.Request.URL.Host
	path := res.Request.URL.Path

	if err := os.Mkdir(outputDir+"/"+host, os.ModePerm); err != nil {
		return err
	}

	// TODO: 画像などの対応
	contentType, ok := res.Header[contextTypeHeader]
	if !ok {
		fmt.Println("content type could not be got")
		return nil
	}
	actualType := strings.Split(contentType[0], ";")[0]

	fmt.Fprintf(os.Stdout, "actualType: %v\n", actualType)

	if actualType == contextTypeTextHtml {
		dom, err := html.Parse(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "parse error: %v\n", err)
		}
		if path == "/" {
			path = "/index.html"
		}
		save(res, outputDir+"/"+host+path)
		fmt.Fprintf(os.Stdout, "dom: %v\n", dom)
		parse(dom, scheme+"://"+host)
	}
	return nil
}

func parse(node *html.Node, urlBase string) {
	fmt.Fprintf(os.Stdout, "node: %v\n", node)
	if node == nil {
		return
	}
	if node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				// アンカーリンクは無視する
				if strings.HasPrefix(attr.Val, "#") {
					return
				}
				if strings.HasPrefix(attr.Val, "/") {
					fmt.Fprintf(os.Stdout, "hasprefix / : %v\n", attr.Val)
					get(urlBase + attr.Val)
					return
				}
				fmt.Fprintf(os.Stdout, "does not have prefix / : %v\n", attr.Val)
				get(attr.Val)
			}
		}
	}
	// 次のnodeから実施する
	parse(node.NextSibling, urlBase)
	parse(node.FirstChild, urlBase)
}

// TODO: すぐに保存されない
func save(res *http.Response, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "creating path %s was failed\n", path)
	}
	io.Copy(f, res.Body)
	defer f.Close()
	defer res.Body.Close()
}
