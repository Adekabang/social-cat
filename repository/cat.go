package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Adekabang/social-cat/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CatRepository struct {
	Db *sql.DB
}

// Function to parse custom format into a slice of strings
func parseCustomFormat(input string) []string {
	// Remove curly braces and split by comma
	parts := strings.Split(strings.Trim(input, "{}"), ",")
	// Trim whitespace from each part
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

func NewCatRepository(db *sql.DB) CatRepositoryInterface {
	return &CatRepository{Db: db}
}

// GetAllCats implements CatRepositoryInterface
func (m *CatRepository) GetAllCats(cat model.GetCat) []model.Cat {

	var query string
	var conditions []string

	// Add condition for id
	if cat.Id != "" {
		conditions = append(conditions, fmt.Sprintf("id = '%s'", cat.Id))
	}

	// Add condition for search
	if cat.Search != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", cat.Search))
	}

	// Add condition for sex (enum male or female)
	if cat.Sex == "male" || cat.Sex == "female" {
		conditions = append(conditions, fmt.Sprintf("sex = '%s'", cat.Sex))
	}

	// Add condition for race (enums)
	validRaces := map[string]bool{
		"Persian":           true,
		"Maine Coon":        true,
		"Siamese":           true,
		"Ragdoll":           true,
		"Bengal":            true,
		"Sphynx":            true,
		"British Shorthair": true,
		"Abyssinian":        true,
		"Scottish Fold":     true,
		"Birman":            true,
	}
	if validRaces[cat.Race] {
		conditions = append(conditions, fmt.Sprintf("race = '%s'", cat.Race))
	}

	if cat.AgeInMonth != "" {
		conditions = append(conditions, fmt.Sprintf(cat.AgeInMonth))
	}
	if cat.HasMatched {
		conditions = append(conditions, fmt.Sprintf("hasmatched = '%t'", cat.HasMatched))
	}

	if cat.Owned {
		conditions = append(conditions, fmt.Sprintf("ownerid = '%s'", cat.OwnerId))
	}
	conditions = append(conditions, "softdelete = false")

	// Construct the WHERE clause
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Construct the final query with WHERE clause
	query = fmt.Sprintf("SELECT * FROM cats %s", whereClause)

	// Add LIMIT condition
	if cat.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", cat.Limit)
	}

	// Add OFFSET condition
	if cat.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", cat.Offset)
	}

	// Execute the query
	log.Println(query)
	rows, err := m.Db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	// defer rows.Close()

	var cats []model.Cat
	if rows != nil {
		for rows.Next() {
			var (
				id            string
				name          string
				race          string
				sex           string
				ageInMonth    int
				description   string
				hasMatched    bool
				imageUrlsJSON string
				createdAt     string
				ownerId       string
				softDelete    bool
			)

			err := rows.Scan(&id, &name, &race, &sex, &ageInMonth, &description, &hasMatched, &imageUrlsJSON, &createdAt, &ownerId, &softDelete)
			if err != nil {
				log.Println(err)
				continue
			}

			// Parse the custom format into a slice of strings
			imageUrls := parseCustomFormat(imageUrlsJSON)
			createdAtFormated, err := time.Parse(time.RFC3339, createdAt)
			if err != nil {
				fmt.Println(err)
			}

			cat := model.Cat{Id: id, Name: name, Race: race, Sex: sex, AgeInMonth: int32(ageInMonth), Description: description, ImageUrls: imageUrls, HasMatched: hasMatched, CreatedAt: createdAtFormated.String()}
			cats = append(cats, cat)
		}
	}
	return cats
}

// GetOneCat implements CatRepositoryInterface
func (m *CatRepository) GetOneCat(id string) bool {
	query, err := m.Db.Query("SELECT * FROM cats WHERE id = $1 AND softdelete = false", id)
	if err != nil {
		log.Println(err)
		return false
	}
	if query != nil {
		for query.Next() {
			var (
				id            string
				name          string
				race          string
				sex           string
				ageInMonth    int
				description   string
				hasMatched    bool
				imageUrlsJSON string
				createdAt     string
				ownerId       string
				softDelete    bool
			)
			err := query.Scan(&id, &name, &race, &sex, &ageInMonth, &description, &hasMatched, &imageUrlsJSON, &createdAt, &ownerId, &softDelete)

			if err != nil {
				log.Println(err)
				return false
			}

			if id != "" {
				return true
			}

		}
	} else {
		return false
	}
	return false
}

// InsertCat implements CatRepositoryInterface
func (m *CatRepository) InsertCat(post model.PostCat) model.CatResponseMessage {

	uuidCat := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO cats(id, name, race, sex, ageInMonth, description, imageUrls, hasMatched, ownerId) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		log.Println(err)
		return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidCat, post.Name, post.Race, post.Sex, post.AgeInMonth, post.Description, pq.Array(post.ImageUrls), false, post.OwnerId)
	if err2 != nil {
		log.Println(err2)
		return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}
	}
	query, err := m.Db.Query("SELECT * FROM cats WHERE id = $1", uuidCat)
	if err != nil {
		log.Println(err)
		return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}
	}
	if query != nil {
		for query.Next() {
			var (
				id            string
				name          string
				race          string
				sex           string
				ageInMonth    int
				description   string
				hasMatched    bool
				imageUrlsJSON string
				createdAt     string
				ownerId       string
				softDelete    bool
			)
			err := query.Scan(&id, &name, &race, &sex, &ageInMonth, &description, &hasMatched, &imageUrlsJSON, &createdAt, &ownerId, &softDelete)

			if err != nil {
				log.Println(err)
				return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}

			}
			if id != "" {
				return model.CatResponseMessage{Status: true, Id: id, CreatedAt: createdAt}
			}

		}
	} else {
		return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}
	}
	return model.CatResponseMessage{Status: false, Id: "", CreatedAt: ""}
}

// UpdateCat implements CatRepositoryInterface
func (m *CatRepository) UpdateCat(id string, post model.PostCat) int {

	var count int

	err := m.Db.QueryRow("SELECT COUNT(*) FROM matches WHERE receivercatid = $1 OR issuercatid = $1", id).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if count >= 1 {
		var sex string
		err := m.Db.QueryRow("SELECT sex FROM cats WHERE id = $1", id).Scan(&sex)
		if err != nil {
			log.Println(err)
		}
		if sex != post.Sex {
			return 400
		}
	}
	update, err := m.Db.Exec("UPDATE cats SET name = $1, race = $2, sex = $3, ageInMonth = $4, description = $5, imageUrls = $6 WHERE id = $7 AND ownerid = $8", post.Name, post.Race, post.Sex, post.AgeInMonth, post.Description, pq.Array(post.ImageUrls), id, post.OwnerId)
	num, _ := update.RowsAffected()
	if num == 0 {
		log.Println(err)
		return 404
	}
	if err != nil {
		log.Println(err)
		return 500
	}

	getCat := m.GetOneCat(id)
	if getCat {
		return 200
	}

	return 500
}

// DeleteCat implements CatRepositoryInterface
func (m *CatRepository) DeleteCat(id string, userId string) bool {
	// delete, err := m.Db.Exec("DELETE FROM cats WHERE id = $1", id)
	// num, _ := delete.RowsAffected()
	// if num == 0 {
	// 	return false
	// }
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }
	// return true
	var deleteStatus bool
	err := m.Db.QueryRow("SELECT softdelete FROM cats WHERE id = $1", id).Scan(&deleteStatus)
	if err != nil {
		log.Println(err)
	}
	if deleteStatus {
		return false
	}

	update, err := m.Db.Exec("UPDATE cats SET softdelete = true WHERE id = $1 AND ownerid = $2", id, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	num, _ := update.RowsAffected()
	log.Println(num)
	if num == 0 {
		log.Println(err)
		return false
	}
	return true

}
