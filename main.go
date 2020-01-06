package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello GitHub Actions!")
	fmt.Println("Why there is no way to manually run GitHub workflow?")
	json.Marshal("test")
}
