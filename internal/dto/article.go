package dto

type RequestArticle struct {
	AuthorId string `json:"author_id" validate:"required"`
	Title    string `json:"title"validate:"required"`
	Body     string `json:"body" validate:"required"`
}

type GetAllArticleReq struct {
	Page   int    `json:"page"`
	Query  string `json:"query"`
	Author string `json:"author"`
}
