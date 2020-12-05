package database

import "github.com/kindai-csg/d-blog-engine/domain"

type GrowiHandler interface {
    GetPage(string) (domain.Article, error)
    DownloadFile([]string, string) error
}
