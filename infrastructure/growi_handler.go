package infrastructure

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
        "fmt"

	"github.com/kindai-csg/d-blog-engine/domain"
	"gopkg.in/xmlpath.v2"
)


type GrowiHandler struct {
	endpoint string
	token    string
        userId string
        password string
}

func NewGrowiHandler(endpoint string, token string, userId string, password string) *GrowiHandler {
	return &GrowiHandler{
		endpoint,
		token,
                userId,
                password,
	}
}

func (handler *GrowiHandler) GetPage(path string) (domain.Article, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

        // get page body
	url := handler.endpoint + "/_api/pages.get"
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	params.Add("access_token", handler.token)
	params.Add("path", path)
	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return domain.Article{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
        jsonBytes := ([]byte)(byteArray)
        pageData := new(GrowiPageGetResponse)
        if err := json.Unmarshal(jsonBytes, pageData); err != nil {
            return domain.Article{}, err
        }

        article := domain.Article {
            Id: pageData.Page.PageID,
            Body:pageData.Page.Revision.Body,
        }

        // get page tags
	url = handler.endpoint + "/_api/pages.getPageTag"
	req, _ = http.NewRequest("GET", url, nil)

	params = req.URL.Query()
	params.Add("access_token", handler.token)
	params.Add("pageId", article.Id)
	req.URL.RawQuery = params.Encode()

	resp, err = client.Do(req)
	if err != nil {
		return domain.Article{}, err
	}

	byteArray, _ = ioutil.ReadAll(resp.Body)
        jsonBytes = ([]byte)(byteArray)
        tagData := new(GrowiGetPageTagResponse)
        if err := json.Unmarshal(jsonBytes, tagData); err != nil {
            return domain.Article{}, err
        }
        article.Tags = tagData.Tags
        fmt.Println("testtags")
        fmt.Println(article.Tags)

	defer resp.Body.Close()
	return article, nil
}

func (handler *GrowiHandler) DownloadFile(paths []string, downloadDir string) error {
        // generate client & cookie
        jar, err := cookiejar.New(nil)
        if err != nil {
            return err
        }
	client := &http.Client{
                Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

        // get csrf from login page
        req, _ := http.NewRequest("GET", handler.endpoint + "/login", nil)
        resp, err := client.Do(req)
        csrfPath := xmlpath.MustCompile(`/html/body/@data-csrftoken`)
        root, err := xmlpath.ParseHTML(resp.Body)
        if err != nil {
            return err
        }
        iter := csrfPath.Iter(root)
        if !iter.Next() {
            return errors.New("get csrf")
        }
        csrf := iter.Node().String()
        resp.Body.Close()

        // post login
        loginData := url.Values{}
        loginData.Add("loginForm[username]", handler.userId)
        loginData.Add("loginForm[password]", handler.password)
        loginData.Add("_csrf", csrf)
	req, _ = http.NewRequest("POST", handler.endpoint + "/login", strings.NewReader(loginData.Encode()))
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
        resp, err = client.Do(req)
	if err != nil {
		return err
	}

        // download file
        for _, path := range paths {
            req, _ = http.NewRequest("GET", handler.endpoint + "/download/" + path, nil)
            resp, err = client.Do(req)
            if err != nil {
                    return err
            }

            file, err := os.OpenFile(downloadDir + path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
            if err != nil {
                return err
            }
            io.Copy(file, resp.Body)
            file.Close()
        }

        return nil
}
