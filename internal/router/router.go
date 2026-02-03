package router

import (
	"file_storage/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Handler) *gin.Engine {
    r := gin.Default()
    
    // r.Use(middleware.Logger())
    // r.Use(middleware.Recovery())
	
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "OK"})
    })

	r.MaxMultipartMemory = 1 << 30

	v1 := r.Group("/api/v1")
	{
		// auth := v1.Group("/auth")
		// {
		// 	auth.POST("/login", handler.Login)
		// 	auth.POST("/register", handler.Register)
		// }

		v1.GET("/files", handler.ListFiles)
		v1.GET("/files/:id", handler.GetFile)
		v1.POST("/files", handler.UploadFile)
		v1.DELETE("/files/:id", handler.DeleteFile)
		v1.POST("/thumbnails/:id", handler.UploadThumbnail)

	}
	return r		
}