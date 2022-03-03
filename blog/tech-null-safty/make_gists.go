package main

import "fmt"

type gist_list struct {
	url      string
	filename string
}

func main() {
	list := []gist_list{
		gist_list{"https://gist.github.com/woinary/e8d808489326626a03cbf3e17f2cecce", "hello_scan.go"},
		gist_list{"https://gist.github.com/woinary/8932e8a4431e53378237b7c856a200a7", "hello.go"},
		gist_list{"https://gist.github.com/woinary/96a5aab65a1ea4a9e26fef6c6dc661d6", "num_add_refine_fir.swift"},
		gist_list{"https://gist.github.com/woinary/5d6fc521a7ab677cf758dd9e59f4074e", "num_add_refine.swift"},
		gist_list{"https://gist.github.com/woinary/dd180f456deb9fb5402b9ac238b36125", "num_add.swift"},
		gist_list{"https://gist.github.com/woinary/51eaca6d7495e60d9bf07260af7f112d", "hello_refine.swift"},
		gist_list{"https://gist.github.com/woinary/d8437fa4f5306b970a92a0b10e0139a1", "hello.swift"},
	}

	for _, item := range list {
		fmt.Printf("git submodule add %s %s\n", item.url, item.filename)
	}

}
