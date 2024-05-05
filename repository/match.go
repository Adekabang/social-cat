package repository

import (
	"database/sql"
	"log"

	"github.com/Adekabang/social-cat/model"
	"github.com/google/uuid"
)

type MatchRepository struct {
	Db *sql.DB
}

func NewMatchRepository(db *sql.DB) MatchRepositoryInterface {
	return &MatchRepository{Db: db}
}

// RequestMatch implements MatchRepositoryInterface
func (m *MatchRepository) RequestMatch(requestMatch model.RequestMatch) model.CreateMatchResponse {

	// check if issuerCatId or receiverCatId not exist
	var isValidCatId int
	err := m.Db.QueryRow("SELECT COUNT(*) FROM cats WHERE id = $1 OR id =$2", requestMatch.MatchCatId, requestMatch.UserCatId).Scan(&isValidCatId)
	if err != nil {
		log.Fatal(err)
	}
	if isValidCatId != 2 {
		log.Println(isValidCatId)
		return model.CreateMatchResponse{StatusCode: 404, Message: "Match Cat ID or User Cat ID not found"}
	}

	// check if issuerCatId is not belong issuer
	var isValidIssuerCat int
	err = m.Db.QueryRow("SELECT count(*) FROM cats WHERE ownerid = $1 OR id =$2", requestMatch.IssuedBy, requestMatch.UserCatId).Scan(&isValidIssuerCat)
	if err != nil {
		log.Fatal(err)
	}
	if isValidIssuerCat != 1 {
		log.Println(isValidIssuerCat)
		return model.CreateMatchResponse{StatusCode: 404, Message: "User Cat ID is not belong to the Issuer"}
	}

	// check if gender is the same

	// check issuerCatId and receiverCatId already matched

	// check issuerCatId and receiverCatId same owner

	uuidMatch := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO matches(id, issuedBy, issuerCatId, receiverCatId, message, status) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		log.Println(err)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidMatch, requestMatch.IssuedBy, requestMatch.MatchCatId, requestMatch.UserCatId, requestMatch.Message, "pending")
	if err2 != nil {
		log.Println(err2)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

	}
	query, err := m.Db.Query("SELECT * FROM matches WHERE id = $1", uuidMatch)
	if err != nil {
		log.Println(err)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

	}
	if query != nil {
		for query.Next() {
			var (
				id            string
				createdAt     string
				issuedBy      string
				issuerCatId   string
				receiverCatId string
				message       string
				status        string
			)
			err := query.Scan(&id, &createdAt, &issuedBy, &issuerCatId, &receiverCatId, &message, &status)

			if err != nil {
				log.Println(err)
				return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

			}
			if id != "" {
				return model.CreateMatchResponse{StatusCode: 201, Message: "Match sucessfully created", IdMatch: id, CreatedAt: createdAt}
			}

		}
	} else {
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

	}
	return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

}

// GetAllMatchs implements CatRepositoryInterface
func (m *MatchRepository) GetMatchRequest(userId string) []model.GetMatch {

	// gimana caranya ngeget kalo kita sebagai receiver cok

	query, err := m.Db.Query("SELECT * FROM matches WHERE issuedby = $1 ", userId)
	if err != nil {
		log.Println(err)
		return nil
	}

	var matchs []model.GetMatch
	if query != nil {
		for query.Next() {
			var (
				id            string
				createdAt     string
				issuedBy      string
				issuerCatId   string
				receiverCatId string
				message       string
				status        string
			)
			err := query.Scan(&id, &createdAt, &issuedBy, &issuerCatId, &receiverCatId, &message, &status)

			if err != nil {
				log.Println(err)
				return nil
			}

			// get data issuer by userid

			// get data matchCatDetail by receiverCatId

			// get data userCatDetail by issuerCatId

			match := model.GetMatch{Id: id, Message: message, CreatedAt: createdAt}
			matchs = append(matchs, match)
		}
	} else {
		return nil
	}
	return matchs
}

// DeleteRequestMatch implements CatRepositoryInterface
func (m *MatchRepository) DeleteRequestMatch(matchId string, userId string) bool {

	// check if already approved / reject

	delete, err := m.Db.Exec("DELETE FROM matches WHERE id = $1 AND id IN ( SELECT id FROM match WHERE id = $1) AND issuedby = $2", matchId, userId)
	num, _ := delete.RowsAffected()
	if num == 0 {
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// ApproveMatch implements CatRepositoryInterface
func (m *MatchRepository) ApproveMatch(id string) bool {

	// other match request that matches both the issuer and the receiver cat’s, will get removed

	approve, err := m.Db.Exec("UPDATE matches SET status = $1 WHERE id = $2", "approve", id)
	num, _ := approve.RowsAffected()
	if num == 0 {
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (m *MatchRepository) RejectMatch(id string) bool {

	// other match request that matches both the issuer and the receiver cat’s, will get removed

	approve, err := m.Db.Exec("UPDATE matches SET status = $1 WHERE id = $2", "reject", id)
	num, _ := approve.RowsAffected()
	if num == 0 {
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
