package controller

import (
    "fmt"

    "github.com/kindai-csg/d-blog-engine/interfaces/gateway"
    "github.com/kindai-csg/d-blog-engine/usecase"
)

type ArticleController struct {
    interactor *usecase.ArticleInteractor
}

func NewArticleController(repository usecase.ArticleRepository) *ArticleController {
    articleController := ArticleController{
        interactor: usecase.NewArticleInteractor(repository),
    }
    return &articleController
}

func (controller *ArticleController) PostArticle(c Context) {
    input := gateway.ArticleInput{}
    err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(400, gateway.ResultOutput{false, "Parameter is invalid."})
        return
    }

    err = controller.interactor.Post(input.Path, input.Title)
    if err != nil {
        fmt.Println(err.Error())
        c.JSON(500, gateway.ResultOutput{false, "faild to post article."})
        return
    }
    c.JSON(200, gateway.ResultOutput{true, "success!"})
}


