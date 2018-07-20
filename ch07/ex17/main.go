package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
)

type item struct {
	name      string
	startItem xml.StartElement
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	stack := make([]item, 10)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stackItem := item{tok.Name.Local, tok}
			stack = append(stack, stackItem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAllItem(stack, os.Args[1:]) {
				fmt.Printf("%v: %s\n", stack, tok)
			}
			//names := getElementNames(stack)
			//if containsAll(names, os.Args[1:]) {
			//	fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			//}
		}
	}
}

// containsAll reports whether x contains the item of y, in order.
func containsAll(itemNames, args []string) bool {
	for len(args) <= len(itemNames) {
		if len(args) == 0 {
			return true
		}
		if itemNames[0] == args[0] {
			args = args[1:]
		}
		itemNames = itemNames[1:]
	}
	return false
}

func containsAllItem(items []item, args []string) bool {
	for len(args) <= len(items) {
		if len(args) == 0 {
			return true
		}
		// マッチしたら次の要素へ
		name, value, err := ParseAttrArg(args[0])
		if err != nil {
			fmt.Fprintf(os.Stdout, "error: %v", err)
			os.Exit(1)
		}
		// 要素のみの場合
		if value == "" {
			if name == args[0] {
				args = args[1:]
			}
			// class="hoge" のように attr である場合
		} else {
			for _, attr := range items[0].startItem.Attr {
				if name == attr.Name.Local && value == attr.Value {
					args = args[1:]
				}
			}
		}
		items = items[1:]
	}
	return false
}

func ParseAttrArg(arg string) (name string, value string, err error) {
	r, err := regexp.Compile(`(.+)="(.+)"`)
	if err != nil {
		return "", "", err
	}
	res := r.FindStringSubmatch(arg)
	// class="hoge" の場合はname,valueをセットで返す
	if len(res) > 0 {
		return res[1], res[2], nil
	}
	// div などの要素名単体の場合は name のみ返す
	return arg, "", nil
}

func containsAttr(items []item, attrName string, attrValue string) bool {
	for _, item := range items {
		for _, attr := range item.startItem.Attr {
			if attrName == attr.Name.Local && attrValue == attr.Value {
				return true
			}
		}
	}
	return false
}

func getElementNames(elms []item) []string {
	var names []string
	for _, v := range elms {
		names = append(names, v.name)
	}
	return names
}

//!-
