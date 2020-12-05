package usecase

import (
    "time"

    "github.com/kindai-csg/d-blog-engine/domain"
)

type ArticleInteractor struct {
    articleRepository ArticleRepository
}

func NewArticleInteractor(articleRepository ArticleRepository) *ArticleInteractor {
    articleInteractor := ArticleInteractor{
        articleRepository,
    }
    return &articleInteractor
}

func (interactor *ArticleInteractor) createHeader(article domain.Article) string {
    header := "---\n"
    header += "title: \"" + article.Title + "\"\n"
    header += "date: " + time.Now().Format("2006-01-02T15:04:05") + "+09:00\n"
    header += "draft: false\n"
    header += "tags: ["
    for _, tag := range article.Tags {
        header += "\"" + tag + "\","
    }
    if len(article.Tags) > 0 {
        header = header[:len(header) - 1]
    }
    header += "]\n"
    header += "---\n\n"
    return header
}

func (interactor *ArticleInteractor) Post(path string, title string) error {
    article, err := interactor.articleRepository.GetRawArticle(path)
    if err != nil {
        return err
    }
    article.Title = title
    article.Body = interactor.createHeader(article) + article.Body
    err = interactor.articleRepository.Post(article)
    return err
}
