package req

import (
	"fmt"
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

func HandlerCreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataOrder1 food.Order
		if err := c.ShouldBind(&dataOrder1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bind data faild",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(dataOrder1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "validate data faild",
			})
			return
		}
		idOrders, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "uuid faild",
			})
			return
		}
		fmt.Printf("idOrders: %v\n", idOrders)
		time := time.Now().UTC()
		dataOrder := food.Order{
			Order_id:   idOrders.String(),
			Table_id:   dataOrder1.Table_id,
			Order_date: dataOrder1.Order_date,
			Created_at: time,
			Updated_at: time,
		}

		biz := food_bussiness.NewOrderController(repoimpl.NewSql(db))
		if err := biz.NewCreateOrderTable(c.Request.Context(), &dataOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "uuid faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("create success"))
	}
}
func HandlerGetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idOrder := c.Param("order_id")

		biz := food_bussiness.NewOrderController(repoimpl.NewSql(db))
		data, err := biz.NewGetOrderTable(c.Request.Context(), idOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "getOrder faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(data))
	}
}
func HandlerUpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idOrder := c.Param("order_id")
		var data food.Order
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bind data faild",
			})
			return
		}
		biz := food_bussiness.NewOrderController(repoimpl.NewSql(db))
		if err := biz.NewUpdateOrder(c.Request.Context(), idOrder, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": " data faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("update success"))
	}
}
func HandlerDeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idOrder := c.Param("order_id")
		biz := food_bussiness.NewOrderController(repoimpl.NewSql(db))
		if err := biz.NewDeleteOrderTable(c.Request.Context(), idOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": " delete order table faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("delete success"))
	}
}
