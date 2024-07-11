package req

import (
	"math"
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

func HandlerCreateProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products food.Product
		if err := c.ShouldBind(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Failed to create products",
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(products); err != nil {
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
		var num = toFixed(*products.Price, 2)
		currentTime := time.Now().UTC()
		product := food.Product{
			Product_ID:  id.String(),
			Title:       products.Title,
			Image:       products.Image,
			Description: products.Description,
			Price:       &num,
			Status:      products.Status,
			Menu_ID:     products.Menu_ID,
			CreatedAt:   &currentTime,
			UpdatedAt:   &currentTime,
		}

		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		if err := biz.NewCreateProducts(c.Request.Context(), &product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "Failed to create products",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Create food success!"))
	}
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func HandlerGetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodID := c.Param("product_id")
		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		pro, err := biz.NewGetProducts(c.Request.Context(), foodID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "GetFood failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(pro))
	}
}

func HandlerGetProducts(db *gorm.DB) gin.HandlerFunc {
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
		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		listData, err := biz.NewGetProductsList(c.Request.Context(), &pagging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "listProduct failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.MutiResponse(listData, pagging, nil, nil))
	}
}
func HandlerUpdateProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataUpdate food.Product
		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "dataUpdate failed",
			})
			return
		}
		dataID := c.Param("product_id")
		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		if err := biz.NewUpdateProducts(c.Request.Context(), &dataUpdate, dataID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "dataUpdate db failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Update success"))
	}
}
func HandlerDeletedProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dataID := c.Param("product_id")
		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		if err := biz.NewDeleteProducts(c.Request.Context(), dataID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "deleted db failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("deleted successfully"))
	}
}
func HandlerGetProductByName(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dataName := c.Param("title")
		biz := food_bussiness.NewProductsController(repoimpl.NewSql(db))
		data, err := biz.NewGetProductByName(c.Request.Context(), dataName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"comment": "getdatbyname db failed",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(data))
	}
}
