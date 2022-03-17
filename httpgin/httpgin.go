package httpgin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {

	// Define ambiente: DebugMode, ReleaseMode, TestMode:
	// export GIN_MODE=release
	gin.SetMode(gin.DebugMode)
	fmt.Println("Gin Mode:", gin.Mode())

	// Default with the Logger and Recovery middleware already attached
	// router := gin.Default()

	// without middleware by default
	router := gin.New()

	router.GET("/api/v1/names/:name", search)
	// router.POST("/api/v1/names/:name", insert)
	// router.PUT("/api/v1/names/:nameOld/:nameNew", update)
	// router.DELETE("/api/v1/delete/:name", delete)(
	router.Run(":5000")
	// router.Run()
}

// ****** HANDLERS

func search(ctx *gin.Context) {

	name := ctx.Param("name")
	fmt.Println("name:", name)

}

/*


 */
