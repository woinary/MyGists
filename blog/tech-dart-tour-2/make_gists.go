package main

import "fmt"

type gist_list struct {
	url      string
	filename string
}

func main() {
	list := []gist_list{
		gist_list{"https://gist.github.com/woinary/b4b86401706199a6852caa649db0dd5c", "functions.dart"},
		gist_list{"https://gist.github.com/woinary/f362eeb60d0571c63a709506345dd5a9", "named_parameters.dart"},
		gist_list{"https://gist.github.com/woinary/db2bf5d0b62b0bcdea6ccb8cfe150b24", "optional_position_parameters.dart"},
		gist_list{"https://gist.github.com/woinary/14b859fa81032fba77c5ac9afc8823af", "default_parameters.dart"},
		gist_list{"https://gist.github.com/woinary//672c9d414a15ca1841aa2e9c01fbfa9d", "first_class_object.dart"},
		gist_list{"https://gist.github.com/woinary/060606f31e73a10b3e24280ea9f6f30d", "anonymous_functions.dart"},
		gist_list{"https://gist.github.com/woinary/0aee99b1f0e13e6d511c6c9ba43e0d0b", "lexical_scope.dart"},
		gist_list{"https://gist.github.com/woinary/5e2656fafcb1b8acb634dae3f59e499b", "lexical_scope2.dart"},
		gist_list{"https://gist.github.com/woinary/5d427051051392608832745e6b89ec5e", "lexical_scope3.dart"},
		gist_list{"https://gist.github.com/woinary/aa01efd9d0829dd8f86ab56a3ae3f083", "lexical_closure.dart"},
	}

	for _, item := range list {
		fmt.Printf("git submodule add %s %s\n", item.url, item.filename)
	}

}
