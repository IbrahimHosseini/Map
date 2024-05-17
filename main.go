package main

var userProfile = map[string]string{}

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

func printMapData() {
	println("==================================")
	for key, value := range userProfile {
		println(key, ": => ", value)
	}
}

func deleteMapData(key string) {
	delete(userProfile, key)
}

func addItem(key string, value string) {
	userProfile[key] = value
}
