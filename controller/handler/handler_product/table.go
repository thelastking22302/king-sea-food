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

func HandlerCreateTables(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var table food.Table
		if err := c.ShouldBind(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Failed to create table",
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Failed to validate products",
			})
			return
		}

		id, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err,
				"comment": "Failed to generate UUID",
			})
			return
		}

		time := time.Now().UTC()
		tables := food.Table{
			Table_id:         id.String(),
			Number_of_guests: table.Number_of_guests,
			Table_number:     table.Table_number,
			Created_at:       &time,
			Updated_at:       &time,
		}

		biz := food_bussiness.NewTableController(repoimpl.NewSql(db))
		if err := biz.NewCreateTable(c.Request.Context(), &tables); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Failed to create table",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Create table success!"))
	}
}

func HandlerGetTable(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("table_id")
		biz := food_bussiness.NewTableController(repoimpl.NewSql(db))
		pro, err := biz.NewGetTable(c.Request.Context(), tableID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Get table failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(pro))
	}
}

func HandlerGetTables(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagging common.Paggings
		if err := c.ShouldBind(&pagging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "pagging failed",
			})
			return
		}
		pagging.Process()
		biz := food_bussiness.NewTableController(repoimpl.NewSql(db))
		listData, err := biz.NewGetListTable(c.Request.Context(), &pagging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "listTable failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.MutiResponse(listData, pagging, nil, nil))
	}
}
func HandlerUpdateTables(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataUpdate food.Table
		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "dataUpdate failed",
			})
			return
		}
		dataID := c.Param("table_id")
		biz := food_bussiness.NewTableController(repoimpl.NewSql(db))
		if err := biz.NewUpdateTable(c.Request.Context(), dataID, &dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "dataUpdate db failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Update success"))
	}
}
func HandlerDeletedTable(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dataID := c.Param("table_id")
		biz := food_bussiness.NewTableController(repoimpl.NewSql(db))
		if err := biz.NewDeleteTable(c.Request.Context(), dataID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "deleted db failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("deleted successfully"))
	}
}
