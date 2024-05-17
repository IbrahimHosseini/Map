package main

// The line `var userProfile = map[string]string{}` is declaring a variable named `userProfile` in Go,
// which is a map data structure. This map is defined to have keys of type `string` and values of type
// `string`. Initially, the map is empty as it is initialized without any key-value pairs.
var userProfile = make(map[string]string)

// The main function adds, prints, deletes, and updates key-value pairs in a map data structure.
func main() {

	// The code snippet `addItem := func(key string, value string) {
	// 		userProfile[key] = value
	// 	}` is defining an anonymous function in Go that takes two parameters: `key` of type `string` and
	// `value` of type `string`. This function is then assigned to a variable `addItem` using the short
	// variable declaration `:=`.
	addItem := func(key string, value string) {
		userProfile[key] = value
	}

	// The code snippet `printMapData := func() {
	// 		println("==================================")
	// 		for key, value := range userProfile {
	// 			println(key, ": => ", value)
	// 		}
	// 	}` is defining an anonymous function in Go that iterates over the `userProfile` map and prints
	// each key-value pair.
	printMapData := func() {
		println("==================================")
		for key, value := range userProfile {
			println(key, ": => ", value)
		}
	}

	// The code snippet `deleteMapData := func(key string) {
	// 		delete(userProfile, key)
	// 	}` is defining an anonymous function in Go that takes a parameter `key` of type `string`. This
	// function is responsible for deleting a key-value pair from the `userProfile` map based on the
	// provided key.
	deleteMapData := func(key string) {
		delete(userProfile, key)
	}

	addItem("name", "Ibrahim")
	addItem("surename", "Hosseini")
	addItem("ocupation", "Software Engineer")

	printMapData()

	deleteMapData("surename")

	addItem("lastname", "Hosseini")
	addItem("age", "24")

	printMapData()

}
