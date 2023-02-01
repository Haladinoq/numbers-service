package model

//NumbersRequest fee put request
type NumbersRequest struct {
	Client string `json:"client"`
	Number int64  `json:"number"`
}

//NumbersResponse fee put request
type NumbersResponse struct {
	ID        int64  `json:"id"`
	Client    string `json:"client"`
	Number    int64  `json:"number"`
	CreatedAt int64  `json:"CreatedAt"`
	UpdatedAt int64  `json:"UpdatedAt"`
}
