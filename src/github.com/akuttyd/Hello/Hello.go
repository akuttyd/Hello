package main

import (
"fmt"
"github.com/huandu/xstrings"
)


func main() {

	fmt.Printf("Hello World1 ")

    var count int = 0
	count = xstrings.Count("This is arun and akankgha . This has lots of problems","a-b")
	fmt.Println(count)

}
