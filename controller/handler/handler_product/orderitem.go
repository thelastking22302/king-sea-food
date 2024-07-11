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

func HandlerCreateOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataOrderItems1 food.OrderItem
		if err := c.ShouldBind(&dataOrderItems1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild request dataOI",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(dataOrderItems1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild validate dataOI",
			})
			return
		}
		dataId, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild id dataOI",
			})
			return
		}
		time := time.Now().UTC()

		dataOrderItems := food.OrderItem{
			Order_item_id: dataId.String(),
			Order_id:      dataOrderItems1.Order_id,
			Quantity:      dataOrderItems1.Quantity,
			Unit_price:    dataOrderItems1.Unit_price,
			Created_at:    time,
			Updated_at:    time,
			Food_id:       dataOrderItems1.Food_id,
		}
		biz := food_bussiness.NewOrderItemController(repoimpl.NewSql(db))
		if err := biz.NewCreateOrderItem(c.Request.Context(), &dataOrderItems); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild db",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("create orderitems success!"))
	}
}
func HandlerUpdateOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataOI food.OrderItem
		if err := c.ShouldBind(&dataOI); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild request dataOI",
			})
			return
		}
		idDataOI := c.Param("order_item_id")
		biz := food_bussiness.NewOrderItemController(repoimpl.NewSql(db))
		if err := biz.NewUpdateOrderItem(c.Request.Context(), idDataOI, &dataOI); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild db dataOI",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("update orderitems success!"))
	}
}
func HandlerGetOrderItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idDataOI := c.Param("order_item_id")
		biz := food_bussiness.NewOrderItemController(repoimpl.NewSql(db))
		DataOI, err := biz.NewGetOrderItems(c.Request.Context(), idDataOI)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild db dataOI",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(DataOI))
	}
}
func HandlerGetOrderItemsByOder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idOrder := c.Param("order_id")
		biz := food_bussiness.NewOrderItemController(repoimpl.NewSql(db))
		DataOI, err := biz.NewGetOrderItemsByOder(c.Request.Context(), idOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild db dataOI",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(DataOI))
	}
}
func HandlerGetOrderItemsByProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagging common.Paggings
		if err := c.ShouldBind(&pagging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild request pagging",
			})
			return
		}
		pagging.Process()
		idProduct := c.Param("product_id")
		biz := food_bussiness.NewOrderItemController(repoimpl.NewSql(db))
		dataList, err := biz.NewGetOrderItemsByProduct(c.Request.Context(), idProduct, &pagging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild request db dataListOI",
			})
			return
		}
		c.JSON(http.StatusOK, common.MutiResponse(dataList, pagging, nil, nil))
	}
}
