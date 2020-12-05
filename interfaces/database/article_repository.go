package database

import (
    "regexp"

    "github.com/kindai-csg/d-blog-engine/domain"
)

type ArticleRepository struct {
    fileHandler FileHandler
    growiHandler GrowiHandler
    hugoHandler HugoHandler
    downloadDir string
    articleDir string
}

func NewArticleRepository(fileHandler FileHandler, growiHandler GrowiHandler, hugoHandler HugoHandler, downloadDir string, articleDir string) *ArticleRepository {
    articleRepository := ArticleRepository{
        fileHandler,
        growiHandler,
        hugoHandler, downloadDir,
        articleDir,
    }
    return &articleRepository
}

var regFilePath = regexp.MustCompile(`/attachment/(([A-Z]|[a-z]|[0-9])+)`)
func (repository *ArticleRepository) GetRawArticle(path string) (domain.Article, error) {
    article, err := repository.growiHandler.GetPage(path)
    if err != nil {
        return domain.Article{}, err
    }
    imagePathGroups := regFilePath.FindAllStringSubmatch(article.Body, -1)
    if imagePathGroups == nil || len(imagePathGroups) == 0 {
        return article, nil
    }
    var imagePaths []string
    for _, imagePathGroup := range imagePathGroups {
        imagePaths = append(imagePaths, imagePathGroup[1])
    }
    err = repository.growiHandler.DownloadFile(imagePaths, repository.downloadDir)
    return article, err
}

func (repository *ArticleRepository) Post(article domain.Article) error {
    err := repository.hugoHandler.Update()
    if err != nil {
        return err
    }
    err = repository.fileHandler.Write(repository.articleDir + article.Id + ".md", article.Body)
    if err != nil {
        return err
    }
    err = repository.hugoHandler.Deploy()
    if err != nil {
        return err
    }
    return nil
}
