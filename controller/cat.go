package controller

import (
	"database/sql"
	"strconv"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
	"github.com/Adekabang/social-cat/utils"
	"github.com/gin-gonic/gin"
)

type CatController struct {
	Db *sql.DB
}

func NewCatController(db *sql.DB) CatControllerInterface {
	return &CatController{Db: db}
}

// GetAllCats implements CatControllerInterface
func (m *CatController) GetAllCats(c *gin.Context) {

	var id string
	var limit int
	var offset int
	var race string
	var sex string
	var hasMatched bool
	var ageInMonth string
	var owned bool
	var search string
	var ownerId string

	reqQuery := c.Request.URL.Query()

	var uri model.CatUri
	if err := c.ShouldBindUri(&uri); err != nil && uri.ID == "" {
		if val, ok := reqQuery["id"]; ok && len(val) > 0 {
			id = val[0]
		} else {
			id = ""
		}

		if val, ok := reqQuery["limit"]; ok && len(val) > 0 {
			limit, _ = strconv.Atoi(val[0])
		} else {
			limit = 0
		}

		if val, ok := reqQuery["offset"]; ok && len(val) > 0 {
			offset, _ = strconv.Atoi(val[0])
		} else {
			offset = 0
		}

		if val, ok := reqQuery["race"]; ok && len(val) > 0 {
			race = val[0]
		} else {
			race = ""
		}

		if val, ok := reqQuery["sex"]; ok && len(val) > 0 {
			sex = val[0]
		} else {
			sex = ""
		}

		if val, ok := reqQuery["hasMatched"]; ok && len(val) > 0 {
			hasMatched, _ = strconv.ParseBool(val[0])
		} else {
			hasMatched = true
		}

		if val, ok := reqQuery["ageInMonth"]; ok && len(val) > 0 {
			ageInMonth = "ageInMonth" + val[0]
			// ageInMonth, _ = strconv.Atoi(val[0])
		} else {
			ageInMonth = ""
		}

		if val, ok := reqQuery["owned"]; ok && len(val) > 0 {
			owned, _ = strconv.ParseBool(val[0])

			userId, err := utils.GetUserId(c.GetHeader(("Authorization")))
			if err != nil {
				return
			}
			ownerId = userId

		} else {
			owned = false
			ownerId = ""
		}

		if val, ok := reqQuery["search"]; ok && len(val) > 0 {
			search = val[0]
		} else {
			search = ""
		}
	} else if uri.ID != "" {
		id = uri.ID
	} else {
		if err := c.ShouldBind(&uri); err != nil {
			c.JSON(400, gin.H{"status": "failed", "msg": err})
			return
		}
	}

	DB := m.Db
	repository := repository.NewCatRepository(DB)
	get := repository.GetAllCats(model.GetCat{
		Id:         id,
		Limit:      limit,
		Offset:     offset,
		Race:       race,
		Sex:        sex,
		HasMatched: hasMatched,
		AgeInMonth: ageInMonth,
		Owned:      owned,
		Search:     search,
		OwnerId:    ownerId,
	})
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get cats successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": make([]string, 0), "msg": "cats not found"})
		return
	}
}

// InsertCat implements CatControllerInterface
func (m *CatController) InsertCat(c *gin.Context) {
	DB := m.Db
	var post model.PostCat
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	userId, err := utils.GetUserId(c.GetHeader(("Authorization")))
	if err != nil {
		c.JSON(500, gin.H{"message": "failed", "msg": "failed to get user id"})
		return
	}

	post.OwnerId = userId

	repository := repository.NewCatRepository(DB)
	insert := repository.InsertCat(post)
	if insert.Status {
		c.JSON(201, gin.H{"message": "success", "data": gin.H{"id": insert.Id, "createdAt": insert.CreatedAt}})
		return
	} else {
		c.JSON(500, gin.H{"message": "failed", "msg": "insert cat failed"})
		return
	}
}

// UpdateCat implements CatControllerInterface
func (m *CatController) UpdateCat(c *gin.Context) {
	DB := m.Db
	var post model.PostCat
	var uri model.CatUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCatRepository(DB)
	update := repository.UpdateCat(uri.ID, post)
	if update {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update cat successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update cat failed"})
		return
	}
}

// DeleteCat implements CatControllerInterface
func (m *CatController) DeleteCat(c *gin.Context) {
	DB := m.Db
	var uri model.CatUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCatRepository(DB)
	delete := repository.DeleteCat(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete cat successfully"})
		return
	} else {
		c.JSON(404, gin.H{"data": make([]string, 0)})
		return
	}
}
