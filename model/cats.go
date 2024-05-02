package model

type Cat struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int32    `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
	CreatedAt   string   `json:"createdAt"`
}

type PostCat struct {
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int32    `json:"ageInMonth"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
}

type GetCat struct {
	Id         string `json:"id"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Race       string `json:"race"`
	Sex        string `json:"sex"`
	HasMatched bool   `json:"hasMatched"`
	AgeInMonth int    `json:"ageInMonth"`
	Owned      bool   `json:"owned"`
	Search     string `json:"search"`
}
