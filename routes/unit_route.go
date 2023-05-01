package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/goods/unit/app"
	repository "github.com/mrokoo/goERP/internal/goods/unit/infra"
	"gorm.io/gorm"
)

func NewUnitRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewUnitRepository(db)
	s := app.NewUnitServiceImpl(m)
	h := app.NewUnitHandler(s)
	group.GET("/units", h.GetUnitList)
	group.GET("/units/:id", h.GetUnit)
	group.POST("/units", h.AddUnit)
	group.PUT("/units/:id", h.ReplaceUnit)
	group.PATCH("/units/:id", h.ReplaceUnit)
	group.DELETE("/units/:id", h.DeleteUnit)

}
