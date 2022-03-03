package main

import "fmt"

type gist_list struct {
	url      string
	filename string
}

func main() {
	list := []gist_list{
		gist_list{"https://gist.github.com/woinary/c8e64d265727b84df3a878effc65740d", "basic_dart_program.dart"},
		gist_list{"https://gist.github.com/woinary/af2f2c4a93b813305996eb6fa6f3afd8", "integer_object.dart"},
		gist_list{"https://gist.github.com/woinary/af04f212693172d7c2df3fd41557628a", "type_anotation.dart"},
		gist_list{"https://gist.github.com/woinary/b30f3f1bc4149e7e040ee1c7c14caf44", "null_safety.dart"},
		gist_list{"https://gist.github.com/woinary/d8fcbc3445bf911b69ecc7a2a9074195", "object.dart"},
		gist_list{"https://gist.github.com/woinary/1328503cb4f61258a92bcf29e9c2eb15", "dynamic.dart"},
		gist_list{"https://gist.github.com/woinary/c7f970f327448679d29aabce5567dc84", "dynamic_sample.dart"},
		gist_list{"https://gist.github.com/woinary/b01c1bf2ac0cb5807f5dd2268948b04d", "generic_list.dart"},
		gist_list{"https://gist.github.com/woinary/53ad2738c6eec58f83bb2c4f4f5e1dee", "class.dart"},
		gist_list{"https://gist.github.com/woinary/9dcea003d3039ea6d3168f0c3a1e9c46", "inner_func.dart"},
		gist_list{"https://gist.github.com/woinary/40f8fa9af40b5740d6b05a291f512746", "variables.dart"},
		gist_list{"https://gist.github.com/woinary/ec44f7340a20817d8ae75e59419f5c70", "default_value.dart"},
		gist_list{"https://gist.github.com/woinary/a90341e36caee91640d734a5c08b3ae1", "late_variables.dart"},
		gist_list{"https://gist.github.com/woinary/4f085f4156a180e92d130750488ac3ad", "final_and_const.dart"},
		gist_list{"https://gist.github.com/woinary/e7c6f2d0fe618fe4af54c59f9cfcc4bd", "list_spread.dart"},
		gist_list{"https://gist.github.com/woinary/4f7ff09c59e6b54a914a37433af8b78e", "list_collection_if_for.dart"},
		gist_list{"https://gist.github.com/woinary/606a613da72fbc1e8016639fbd644249", "set_collection.dart"},
		gist_list{"https://gist.github.com/woinary/b3fdac8463cf1b293ec7973b3653a9e0", "maps.dart"},
	}

	for _, item := range list {
		fmt.Printf("git submodule add %s %s\n", item.url, item.filename)
	}

}
