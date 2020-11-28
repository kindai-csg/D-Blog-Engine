package infrastructure

import (
    "fmt"
    "net/http"
    "crypto/tls"
    "io/ioutil"
)

type GrowiHandler struct {
    endpoint string
    token string
}

func NewGrowiHandler(endpoint string, token string) *GrowiHandler {
    return &GrowiHandler {
        endpoint,
        token,
    }
}

func (handler *GrowiHandler) GetPage(path string) (string, error) {
    url := handler.endpoint + "pages.get"
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer " + handler.token)

    params := req.URL.Query();
    params.Add("access_token", handler.token)
    params.Add("path", path)
    req.URL.RawQuery = params.Encode()

    fmt.Println(req.URL.String())
    client := &http.Client{
        Transport: &http.Transport {
            TLSClientConfig: &tls.Config{ InsecureSkipVerify: true },
        },
    }
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    fmt.Println(resp.StatusCode)
    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byteArray))
    return "", nil
}
