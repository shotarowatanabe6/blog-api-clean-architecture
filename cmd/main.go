// Description: Clean Architecture の Compotion Root として扱う
package main

import "blog-api-clean-architecture/internal"

func main() {
	port := "8080"

	err := internal.Run(port)
	if err != nil {
		panic(err)
	}
}
