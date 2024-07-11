package req

import (
	"net/http"
	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/repository/food_bussiness"
	repoimpl "thelastking/kingseafood/repository/food_bussiness/repo_impl"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMenuHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var menus food.MenuFood
		if err := c.ShouldBind(&menus); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bind menus faild",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(menus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "validator faild",
			})
			return
		}
		idMenu, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "uuid faild",
			})
			return
		}

		time := time.Now().UTC()
		menu := food.MenuFood{
			Menu_ID:    idMenu.String(),
			Name:       menus.Name,
			Category:   menus.Category,
			Created_At: &time,
			Updated_At: &time,
			Products:   []*food.Product{},
		}
		biz := food_bussiness.NewMenuController(repoimpl.NewSql(db))
		if err := biz.NewCreateMenu(c.Request.Context(), &menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "create menu database faild",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("create menu success"))
	}
}

func HandlerGetMenu(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idMenu := c.Param("menu_id")

		biz := food_bussiness.NewMenuController(repoimpl.NewSql(db))
		dataMenu, err := biz.NewGetMenu(c.Request.Context(), idMenu)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "get menu database faild",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(dataMenu))
	}
}

func HandlerGetListMenu(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagging common.Paggings
		if err := c.ShouldBind(&pagging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "pagging faild",
			})
			return
		}
		pagging.Process()
		biz := food_bussiness.NewMenuController(repoimpl.NewSql(db))
		dataList, err := biz.NewGetListMenu(c.Request.Context(), &pagging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "getList menu database faild",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.MutiResponse(dataList, pagging, nil, nil))
	}
}

func HandlerUpdateMenus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var upMenu food.MenuFood
		if err := c.ShouldBind(&upMenu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "update faild",
			})
			return
		}
		idMenu := c.Param("menu_id")
		biz := food_bussiness.NewMenuController(repoimpl.NewSql(db))
		if err := biz.NewUpdateFoodMenu(c.Request.Context(), idMenu, &upMenu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "update db faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("update success"))
	}
}

func HandlerDeleteMenu(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idMenu := c.Param("menu_id")

		biz := food_bussiness.NewMenuController(repoimpl.NewSql(db))
		if err := biz.NewDeleteFoodMenu(c.Request.Context(), idMenu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "get menu database faild",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("delete success"))
	}
}
