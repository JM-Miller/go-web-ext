//Routes/Routes.go
package Routes

import (
	"home-api/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/home", Controllers.GetHome)
	return r
}
