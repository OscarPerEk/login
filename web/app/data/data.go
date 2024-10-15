package data

import (
	"01-Login/web/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Handler(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("app_db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	var users []types.User
	res := db.Find(&users)
	if res.Error != nil {
		panic("failed to get all items from database")
	}

	ctx.HTML(http.StatusOK, "data.html", gin.H{"Users": users})
}
