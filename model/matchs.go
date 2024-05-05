package model

type IssuedBy struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type RequestMatch struct {
	MatchCatId string `json:"matchCatId"`
	UserCatId  string `json:"userCatId"`
	Message    string `json:"message"`
	IssuedBy   string `json:"issuedBy"`
}

type GetMatch struct {
	Id             string   `json:"id"`
	IssuedBy       IssuedBy `json:"issuedBy"`
	MatchCatDetail Cat      `json:"matchCatDetail"`
	UserCatDetail  Cat      `json:"userCatDetail"`
	Message        string   `json:"message"`
	CreatedAt      string   `json:"createdAt"`
}

type MatchUri struct {
	ID string `uri:"id" binding:"required"`
}

type PostApproveReject struct {
	MatchId string `json:"matchId"`
}

type CreateMatchResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	IdMatch    string `json:"idMatch"`
	CreatedAt  string `json:"createdAt"`
}
