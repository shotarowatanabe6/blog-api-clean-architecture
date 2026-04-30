// Description: Clean Architecture の Compotion Root として扱う
package main

import "blog-api-clean-architecture/internal"

//	@title			Blog API
//	@version		1.0
//	@description	This is a blog server.

//	@contact.name	shotarowatanabe6
//	@contact.email	shotarowatanabe6@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	port := "8080"

	err := internal.Run(port)
	if err != nil {
		panic(err)
	}
}
