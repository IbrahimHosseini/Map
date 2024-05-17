package main

// The line `var userProfile = map[string]string{}` is declaring a variable named `userProfile` in Go,
// which is a map data structure. This map is defined to have keys of type `string` and values of type
// `string`. Initially, the map is empty as it is initialized without any key-value pairs.
var userProfile = map[string]string{}

// The main function adds, prints, deletes, and updates key-value pairs in a map data structure.
func main() {

	addItem("name", "Ibrahim")
	addItem("surename", "Hosseini")
	addItem("ocupation", "Software Engineer")

	printMapData()

	deleteMapData("surename")

	addItem("lastname", "Hosseini")
	addItem("age", "24")

	printMapData()
}

// The function `printMapData` iterates over a map called `userProfile` and prints out key-value pairs.
func printMapData() {
	println("==================================")
	for key, value := range userProfile {
		println(key, ": => ", value)
	}
}

// The function `deleteMapData` deletes a key-value pair from a map called `userProfile` using the
// specified key.
func deleteMapData(key string) {
	delete(userProfile, key)
}

// The function `addItem` adds a key-value pair to the `userProfile` map in Go.
func addItem(key string, value string) {
	userProfile[key] = value
}
