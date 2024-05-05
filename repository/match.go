package repository

import (
	"database/sql"
	"log"

	"github.com/Adekabang/social-cat/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		log.Println(err)
	}
	if isValidCatId != 2 {
		log.Println(isValidCatId)
		return model.CreateMatchResponse{StatusCode: 404, Message: "Match Cat ID or User Cat ID not found"}
	}

	// check if issuerCatId is not belong issuer
	var isValidIssuerCat int
	err = m.Db.QueryRow("SELECT count(*) FROM cats WHERE ownerid = $1 and id =$2", requestMatch.IssuedBy, requestMatch.UserCatId).Scan(&isValidIssuerCat)
	if err != nil {
		log.Println(err)
	}
	if isValidIssuerCat != 1 {
		log.Println(isValidIssuerCat)
		return model.CreateMatchResponse{StatusCode: 404, Message: "User Cat ID is not belong to the Issuer"}
	}

	var userCatSex string
	var matchCatSex string
	var userCatMatchedStatus bool
	var matchCatMatchedStatus bool
	var userCatOwner string
	var matchCatOwner string
	// check if gender is the same
	err = m.Db.QueryRow("SELECT sex, hasmatched, ownerid FROM cats WHERE ownerid = $1 AND id =$2", requestMatch.IssuedBy, requestMatch.UserCatId).Scan(&userCatSex, &userCatMatchedStatus, &userCatOwner)
	if err != nil {
		log.Println(err)
	}
	err = m.Db.QueryRow("SELECT sex, hasmatched, ownerid FROM cats WHERE id =$1", requestMatch.MatchCatId).Scan(&matchCatSex, &matchCatMatchedStatus, &matchCatOwner)
	if err != nil {
		log.Println(err)
	}
	if userCatSex == matchCatSex {
		return model.CreateMatchResponse{StatusCode: 400, Message: "Requested cats have same gender"}
	}

	// check issuerCatId and receiverCatId already matched
	if userCatMatchedStatus {
		return model.CreateMatchResponse{StatusCode: 400, Message: "Requested user cat already matched"}
	}
	if matchCatMatchedStatus {
		return model.CreateMatchResponse{StatusCode: 400, Message: "Requested match cat already matched"}
	}

	// check issuerCatId and receiverCatId same owner
	if userCatOwner == matchCatOwner {
		return model.CreateMatchResponse{StatusCode: 400, Message: "Requested cats belong to the same owner"}
	}

	uuidMatch := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO matches(id, issuedBy, receiverid, issuerCatId, receiverCatId, message, status) VALUES ($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Println(err)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidMatch, requestMatch.IssuedBy, matchCatOwner, requestMatch.MatchCatId, requestMatch.UserCatId, requestMatch.Message, "pending")
	if err2 != nil {
		log.Println(err2)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

	}
	query, err := m.Db.Query("SELECT * FROM matches WHERE id = $1", uuidMatch)
	if err != nil {
		log.Println(err)
		return model.CreateMatchResponse{StatusCode: 500, Message: "server Error", IdMatch: "", CreatedAt: ""}

	}
	defer query.Close()
	if query != nil {
		for query.Next() {
			var (
				id            string
				createdAt     string
				issuedBy      string
				receiverid    string
				issuerCatId   string
				receiverCatId string
				message       string
				status        string
			)
			err := query.Scan(&id, &createdAt, &issuedBy, &receiverid, &issuerCatId, &receiverCatId, &message, &status)

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

	query, err := m.Db.Query("SELECT * FROM matches WHERE issuedby = $1 or receiverId=$1", userId)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer query.Close()
	var matchs []model.GetMatch
	if query != nil {
		for query.Next() {
			var (
				id            string
				createdAt     string
				issuedBy      string
				receiverId    string
				issuerCatId   string
				receiverCatId string
				message       string
				status        string
			)
			err := query.Scan(&id, &createdAt, &issuedBy, &receiverId, &issuerCatId, &receiverCatId, &message, &status)

			if err != nil {
				log.Println(err)
				return nil
			}

			var issuerName, issuerEmail, issuerCreatedAt string
			// get data issuer by userid
			err = m.Db.QueryRow("SELECT name, email, created_at FROM users WHERE id = $1", issuedBy).Scan(&issuerName, &issuerEmail, &issuerCreatedAt)
			if err != nil {
				log.Println("get issuer")
				log.Println(err)
			}
			issuerData := model.IssuedBy{Name: issuerName, Email: issuerEmail, CreatedAt: issuerCreatedAt}

			var userCat, matchCat model.Cat
			// get data userCatDetail by issuerCatId
			err = m.Db.QueryRow("SELECT id, name, race, sex, ageinmonth, description, hasmatched, created_at, imageurls FROM cats WHERE id = $1", issuerCatId).Scan(&userCat.Id, &userCat.Name, &userCat.Race, &userCat.Sex, &userCat.AgeInMonth, &userCat.Description, &userCat.HasMatched, &userCat.CreatedAt, pq.Array(&userCat.ImageUrls))
			if err != nil {
				log.Println("get issuer cat")
				log.Println(err)
			}

			// get data matchCatDetail by receiverCatId
			err = m.Db.QueryRow("SELECT id, name, race, sex, ageinmonth, description, hasmatched, created_at, imageurls FROM cats WHERE id = $1", receiverCatId).Scan(&matchCat.Id, &matchCat.Name, &matchCat.Race, &matchCat.Sex, &matchCat.AgeInMonth, &matchCat.Description, &matchCat.HasMatched, &matchCat.CreatedAt, pq.Array(&matchCat.ImageUrls))
			if err != nil {
				log.Println("get match cat")
				log.Println(err)
			}

			match := model.GetMatch{Id: id, IssuedBy: issuerData, Message: message, UserCatDetail: userCat, MatchCatDetail: matchCat, CreatedAt: createdAt}
			matchs = append(matchs, match)
		}
	} else {
		return nil
	}
	return matchs
}

// DeleteRequestMatch implements CatRepositoryInterface
func (m *MatchRepository) DeleteRequestMatch(matchId string, userId string) model.DeleteMatchResponse {

	var status string
	// check if already approved / reject
	err := m.Db.QueryRow("SELECT status FROM matches WHERE id = $1 AND issuedby = $2", matchId, userId).Scan(&status)
	if status == "" {
		return model.DeleteMatchResponse{StatusCode: 404, Message: "Match Id not found"}
	}
	if err != nil {
		log.Println(err)
		return model.DeleteMatchResponse{StatusCode: 500, Message: "Server Error"}
	}
	if status == "pending" {
		_, err := m.Db.Exec("DELETE FROM matches WHERE id = $1 AND issuedby = $2", matchId, userId)
		// num, _ := delete.RowsAffected()
		// log.Println(num)
		if err != nil {
			log.Println(err)
			return model.DeleteMatchResponse{StatusCode: 500, Message: "Server Error"}
		}
	} else {
		return model.DeleteMatchResponse{StatusCode: 400, Message: "Failed to delete. Match status is Approved or Rejected"}
	}
	return model.DeleteMatchResponse{StatusCode: 200, Message: "Delete request successfully"}
}

// ApproveMatch implements CatRepositoryInterface
func (m *MatchRepository) ApproveMatch(idMatch string, userIdFromToken string) bool {

	dataMatches, err := m.Db.Query("SELECT receiverId, status FROM matches WHERE id = $1", idMatch)
	if err != nil {
		log.Println(err)
		return false

	}
	defer dataMatches.Close()

	if dataMatches != nil {
		for dataMatches.Next() {
			var (
				receiverId string
				status     string
			)
			err := dataMatches.Scan(&receiverId, &status)

			if err != nil {
				log.Println(err)
				return false

			}

			// check userid si receiver bukan
			if userIdFromToken != receiverId {
				log.Println("beda userIdFromToken ama receiverId")
				return false
			}

			// check statusnya pending apa bukan
			if status != "pending" {
				log.Println("status bukan pending")
				return false
			}

		}
	}

	// change status di table matches
	approve, err := m.Db.Exec("UPDATE matches SET status = $1 WHERE id = $2", "approve", idMatch)
	num, _ := approve.RowsAffected()
	if num == 0 {
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}

	// hasMatched tiap kucing jadi true
	// select id kucing issuer dan receiver
	idsCat, err := m.Db.Query("SELECT issuercatid, receivercatid FROM matches WHERE id = $1", idMatch)
	if err != nil {
		log.Println(err)
		return false

	}
	defer idsCat.Close()

	if idsCat != nil {
		for idsCat.Next() {
			var (
				issuerCatId   string
				receiverCatId string
			)
			err := idsCat.Scan(&issuerCatId, &receiverCatId)

			if err != nil {
				log.Println(err)
				return false

			}

			// update hasMatched issuer cat jadi true
			updateIssuer, err := m.Db.Exec("UPDATE cats SET hasMatched = $1 WHERE id = $2", true, issuerCatId)
			numIssuer, _ := updateIssuer.RowsAffected()
			if numIssuer == 0 {
				log.Println(err)
				return false
			}
			if err != nil {
				log.Println(err)
				return false
			}

			// update hasMatched receiver cat jadi true
			updateReceiver, err := m.Db.Exec("UPDATE cats SET hasMatched = $1 WHERE id = $2", true, receiverCatId)
			numReceiver, _ := updateReceiver.RowsAffected()
			if numReceiver == 0 {
				log.Println(err)
				return false
			}
			if err != nil {
				log.Println(err)
				return false
			}

			// other match request that matches both the issuer and the receiver cat’s, will get removed
			delete, err := m.Db.Exec("DELETE FROM matches WHERE (issuerCatId = $1 OR receiverCatId = $2 OR issuerCatId = $2 OR receiverCatId = $1) AND status = 'pending'", issuerCatId, receiverCatId)
			num, _ := delete.RowsAffected()
			log.Println(num)
			if err != nil {
				log.Println(err)
				return false
			}

		}
	}

	return true
}

func (m *MatchRepository) RejectMatch(idMatch string, userIdFromToken string) bool {

	dataMatches, err := m.Db.Query("SELECT receiverId, status FROM matches WHERE id = $1", idMatch)
	if err != nil {
		log.Println(err)
		return false

	}
	defer dataMatches.Close()

	if dataMatches != nil {
		for dataMatches.Next() {
			var (
				receiverId string
				status     string
			)
			err := dataMatches.Scan(&receiverId, &status)

			if err != nil {
				log.Println(err)
				return false

			}

			// check userid si receiver bukan
			if userIdFromToken != receiverId {
				log.Println("beda userIdFromToken ama receiverId")
				return false
			}

			// check statusnya pending apa bukan
			if status != "pending" {
				log.Println("status bukan pending")
				return false
			}

		}
	}

	// change status di table matches
	approve, err := m.Db.Exec("UPDATE matches SET status = $1 WHERE id = $2", "reject", idMatch)
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
