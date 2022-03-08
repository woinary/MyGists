package main

import "fmt"

type gist_list struct {
	url      string
	filename string
}

func main() {
	list := []gist_list{
		gist_list{"https://gist.github.com/woinary/f8d9f093c72b362ba2f58f58f74315b4", "cascade.dart"},
		gist_list{"https://gist.github.com/woinary/18b6ea5be0c1debcace76ee5c076a6ea", "other_operators.dart"},
	}

	for _, item := range list {
		fmt.Printf("git submodule add %s %s\n", item.url, item.filename)
	}

}
