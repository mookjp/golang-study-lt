package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "/dp/0131103628/uedata/nvp/unsticky/137-8278262-1262454/NoPageType/ntpoffrw?tepes=1&id=Y8HPDWEWKVHQ15CRJS92"
	fmt.Fprintf(os.Stdout, "hasprefix: %v", strings.HasPrefix(str, "/"))
}
