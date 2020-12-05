package usecase

import "github.com/kindai-csg/d-blog-engine/domain"

type ArticleRepository interface {
    Post(domain.Article) error
    GetRawArticle(string) (domain.Article, error)
}
