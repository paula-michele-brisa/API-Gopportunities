package router

import "github.com/gin-gonic/gin"

// Initialize inicializa o router utilizando as configurações Default do gin
func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	// Run the server
	r.Run()
	// listen and serve on 0.0.0.0:8080
}
