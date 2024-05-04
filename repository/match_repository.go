package repository

import (
	"github.com/Adekabang/social-cat/model"
)

type MatchRepositoryInterface interface {
	RequestMatch(model.RequestMatch) bool
	GetMatchRequest(string) []model.GetMatch
	ApproveMatch(string) bool
	RejectMatch(string) bool
	DeleteRequestMatch(string, string) bool
}
