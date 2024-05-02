package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
				imageUrlsJSON string
				createdAt     string
			)

			err := rows.Scan(&id, &name, &race, &sex, &ageInMonth, &description, &imageUrlsJSON, &createdAt)
			if err != nil {
				log.Println(err)
				continue
			}

			// Parse the custom format into a slice of strings
			imageUrls := parseCustomFormat(imageUrlsJSON)

			cat := model.Cat{Id: id, Name: name, Race: race, Sex: sex, AgeInMonth: int32(ageInMonth), Description: description, ImageUrls: imageUrls, CreatedAt: createdAt}
			cats = append(cats, cat)
		}
	}
	return cats
}

// InsertCat implements CatRepositoryInterface
func (m *CatRepository) InsertCat(post model.PostCat) bool {

	uuidCat := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO cats(id, name, race, sex, ageInMonth, description, imageUrls) VALUES ($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidCat, post.Name, post.Race, post.Sex, post.AgeInMonth, post.Description, pq.Array(post.ImageUrls))
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
