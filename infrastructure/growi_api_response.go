package infrastructure

import "time"

// auto generate https://mholt.github.io/json-to-go/
type GrowiPageGetResponse struct {
	Page struct {
		Status       string        `json:"status"`
		Grant        int           `json:"grant"`
		GrantedUsers []interface{} `json:"grantedUsers"`
		Liker        []interface{} `json:"liker"`
		SeenUsers    []string      `json:"seenUsers"`
		CommentCount int           `json:"commentCount"`
		PageID           string        `json:"_id"`
		CreatedAt    time.Time     `json:"createdAt"`
		UpdatedAt    time.Time     `json:"updatedAt"`
		Path         string        `json:"path"`
		Creator      struct {
			IsGravatarEnabled bool      `json:"isGravatarEnabled"`
			IsEmailPublished  bool      `json:"isEmailPublished"`
			Lang              string    `json:"lang"`
			Status            int       `json:"status"`
			Admin             bool      `json:"admin"`
			ID                string    `json:"_id"`
			CreatedAt         time.Time `json:"createdAt"`
			Name              string    `json:"name"`
			Username          string    `json:"username"`
			Email             string    `json:"email"`
			LastLoginAt       time.Time `json:"lastLoginAt"`
			ImageAttachment   struct {
				CreaterImage_ID                  string `json:"_id"`
				FilePathProxied     string `json:"filePathProxied"`
				DownloadPathProxied string `json:"downloadPathProxied"`
				CreaterImageID                  string `json:"id"`
			} `json:"imageAttachment"`
		} `json:"creator"`
		LastUpdateUser struct {
			IsGravatarEnabled bool      `json:"isGravatarEnabled"`
			IsEmailPublished  bool      `json:"isEmailPublished"`
			Lang              string    `json:"lang"`
			Status            int       `json:"status"`
			Admin             bool      `json:"admin"`
			ID                string    `json:"_id"`
			CreatedAt         time.Time `json:"createdAt"`
			Name              string    `json:"name"`
			Username          string    `json:"username"`
			Email             string    `json:"email"`
			LastLoginAt       time.Time `json:"lastLoginAt"`
			ImageAttachment   struct {
				LastUpdateImage_ID                  string `json:"_id"`
				FilePathProxied     string `json:"filePathProxied"`
				DownloadPathProxied string `json:"downloadPathProxied"`
				LastUpdateImageID                  string `json:"id"`
			} `json:"imageAttachment"`
		} `json:"lastUpdateUser"`
		RedirectTo   interface{} `json:"redirectTo"`
		GrantedGroup interface{} `json:"grantedGroup"`
		V            int         `json:"__v"`
		Revision     struct {
			Format    string    `json:"format"`
			ID        string    `json:"_id"`
			CreatedAt time.Time `json:"createdAt"`
			Path      string    `json:"path"`
			Body      string    `json:"body"`
			Author    struct {
				IsGravatarEnabled bool      `json:"isGravatarEnabled"`
				IsEmailPublished  bool      `json:"isEmailPublished"`
				Lang              string    `json:"lang"`
				Status            int       `json:"status"`
				Admin             bool      `json:"admin"`
				ID                string    `json:"_id"`
				CreatedAt         time.Time `json:"createdAt"`
				Name              string    `json:"name"`
				Username          string    `json:"username"`
				Email             string    `json:"email"`
				LastLoginAt       time.Time `json:"lastLoginAt"`
				ImageAttachment   struct {
					RevisionImage_ID                  string `json:"_id"`
					FilePathProxied     string `json:"filePathProxied"`
					DownloadPathProxied string `json:"downloadPathProxied"`
					RevisionImageID                  string `json:"id"`
				} `json:"imageAttachment"`
			} `json:"author"`
			V int `json:"__v"`
		} `json:"revision"`
		ID string `json:"id"`
	} `json:"page"`
	Ok bool `json:"ok"`
}

type GrowiGetPageTagResponse struct {
    Tags []string `json:"tags"`
    Ok bool `json:"ok"`
}
