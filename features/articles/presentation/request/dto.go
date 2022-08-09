package request

import "myexample/go-gin/features/articles"

type Article struct {
	Title  string `json:"title" form:"title"`
	Detail string `json:"detail" form:"detail"`
	UserID int
}

func ToCore(artRequest Article) articles.Core {
	return articles.Core{
		Title:  artRequest.Title,
		Detail: artRequest.Detail,
		UserID: artRequest.UserID,
	}
}
