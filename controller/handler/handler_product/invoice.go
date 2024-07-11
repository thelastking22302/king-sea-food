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

func HandlerCreateInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataInvoice food.InvoiceFood
		if err := c.ShouldBind(&dataInvoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bind dataInvoice faild",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(dataInvoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "validate dataInvoice faild",
			})
			return
		}
		idInvoice1, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "uuid faild",
			})
			return
		}
		status := "PENDING"
		if dataInvoice.Payment_status == nil {
			dataInvoice.Payment_status = &status
		}
		data := food.InvoiceFood{
			Invoice_ID:       idInvoice1.String(),
			Order_ID:         dataInvoice.Order_ID,
			Payment_method:   dataInvoice.Payment_method,
			Payment_status:   dataInvoice.Payment_status,
			Payment_due_date: dataInvoice.Payment_due_date,
			Created_at:       time.Now().UTC(),
			Updated_at:       time.Now().UTC(),
		}
		biz := food_bussiness.NewInvoiceController(repoimpl.NewSql(db))
		if err := biz.NewCreateInvoice(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "db dataInvoice faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("create success"))
	}
}

func HandlerGetInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idInvoice := c.Param("invoice_id")

		biz := food_bussiness.NewInvoiceController(repoimpl.NewSql(db))
		data, err := biz.NewGetInvoiceTable(c.Request.Context(), idInvoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "getInvoice faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(data))
	}
}
func HandlerUpdateInvoices(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idInvoice := c.Param("invoice_id")
		var data food.InvoiceFood
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bind data faild",
			})
			return
		}
		data.Updated_at = time.Now()

		biz := food_bussiness.NewInvoiceController(repoimpl.NewSql(db))
		if err := biz.NewUpdateInvoice(c.Request.Context(), idInvoice, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": " data faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("update success"))
	}
}
func HandlerDeleteInvoice(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idInvoice := c.Param("invoice_id")
		biz := food_bussiness.NewInvoiceController(repoimpl.NewSql(db))
		if err := biz.NewDeleteInvoiceTable(c.Request.Context(), idInvoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": " delete invocie faild",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("delete success"))
	}
}
