package food

import "time"

type InvoiceFood struct {
	Invoice_ID       string    `json:"invoice_id" gorm:"column:invoice_id;"`
	Order_ID         string    `json:"order_id" validate:"required" gorm:"column:order_id;"`
	Payment_method   *string   `json:"payment_method" validate:"eq=CARD|eq=CASH|eq=" gorm:"column:payment_method;"`
	Payment_status   *string   `json:"payment_status" validate:"required,eq=PENDING|eq=PAID" gorm:"column:payment_status;"`
	Payment_due_date time.Time `json:"Payment_due_date" gorm:"column:payment_due_date;"`
	Created_at       time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at       time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
