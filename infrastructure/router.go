package infrastructure

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/kindai-csg/d-blog-engine/interfaces/controller"
    "github.com/kindai-csg/d-blog-engine/interfaces/database"
)

var Router *gin.Engine

func init() {
    router := gin.Default()

    endpoint := os.Getenv("D_BLOG_ENGINE_ENDPOINT")
    accessToken := os.Getenv("D_BLOG_ENGINE_ACCESSTOKEN")
    userId := os.Getenv("D_BLOG_ENGINE_USERID")
    password := os.Getenv("D_BLOG_ENGINE_PASSWORD")
    hugoDir := os.Getenv("D_BLOG_ENGINE_HUGODIR")

    fileHandler := NewFileHandler()
    growiHandler := NewGrowiHandler(endpoint, accessToken, userId, password)
    hugoHandler := NewHugoHandler(hugoDir)

    articleRepository := database.NewArticleRepository(fileHandler, growiHandler, hugoHandler, hugoDir + "/static/attachment/", hugoDir + "/content/post/")
    articleController := controller.NewArticleController(articleRepository)

    router.POST("/post", func(c *gin.Context) { articleController.PostArticle(c) })

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "msg": "hello world",
        })
    })

    Router = router
}
