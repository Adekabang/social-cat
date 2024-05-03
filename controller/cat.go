package controller

import (
	"database/sql"
	"strconv"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/repository"
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
	reqQuery := c.Request.URL.Query()

	var id string
	if val, ok := reqQuery["id"]; ok && len(val) > 0 {
		id = val[0]
	} else {
		id = ""
	}

	var limit int
	if val, ok := reqQuery["limit"]; ok && len(val) > 0 {
		limit, _ = strconv.Atoi(val[0])
	} else {
		limit = 0
	}

	var offset int
	if val, ok := reqQuery["offset"]; ok && len(val) > 0 {
		offset, _ = strconv.Atoi(val[0])
	} else {
		offset = 0
	}

	var race string
	if val, ok := reqQuery["race"]; ok && len(val) > 0 {
		race = val[0]
	} else {
		race = ""
	}

	var sex string
	if val, ok := reqQuery["sex"]; ok && len(val) > 0 {
		sex = val[0]
	} else {
		sex = ""
	}

	var hasMatched bool
	if val, ok := reqQuery["hasMatched"]; ok && len(val) > 0 {
		hasMatched, _ = strconv.ParseBool(val[0])
	} else {
		hasMatched = true
	}

	var ageInMonth int
	if val, ok := reqQuery["ageInMonth"]; ok && len(val) > 0 {
		ageInMonth, _ = strconv.Atoi(val[0])
	} else {
		ageInMonth = 0
	}

	var owned bool
	if val, ok := reqQuery["owned"]; ok && len(val) > 0 {
		owned, _ = strconv.ParseBool(val[0])
	} else {
		owned = true
	}

	var search string
	if val, ok := reqQuery["search"]; ok && len(val) > 0 {
		search = val[0]
	} else {
		search = ""
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
	})
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get cats successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "cats not found"})
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
	repository := repository.NewCatRepository(DB)
	insert := repository.InsertCat(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert cat successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert cat failed"})
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
		c.JSON(500, gin.H{"status": "failed", "msg": "delete cat failed"})
		return
	}
}
