package repository

import (
	"github.com/Adekabang/social-cat/model"
)

type MatchRepositoryInterface interface {
	RequestMatch(model.RequestMatch) model.CreateMatchResponse
	GetMatchRequest(string) []model.GetMatch
	ApproveMatch(string, string) bool
	RejectMatch(string, string) bool
	DeleteRequestMatch(string, string) model.DeleteMatchResponse
}
