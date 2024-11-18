package structlab

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type artice struct {
	ArticleName string 	`json:"article_name"`
	AuthorName  string	`json:"author_name"` 
	PublishedOn string	`json:"published_on"`
	CreatedAt   time.Time	`json:"created_at"`
}

func ExportStruct() error{
	artice := New()
	fileName := "structLab/my_article.json"
	
	json, err := json.Marshal(artice)
	if err != nil {
		return errors.New("Error while converting to Json")
	}

	return os.WriteFile(fileName, json, 0644)
}

func New() *artice {
	return &artice{
		"Life Amazing secrete",
		"Gaur Gopal Das",
		"10/10/2001",
		time.Now(),
	}
}