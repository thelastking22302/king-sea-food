package router

import (
	req "thelastking/kingseafood/controller/handler/handler_product"
	handleruser "thelastking/kingseafood/controller/handler/handler_user"
	"thelastking/kingseafood/middleware"
	"thelastking/kingseafood/server"

	"github.com/gin-gonic/gin"
)

// FACADE
func KingRouters(incomingRoutes *gin.Engine) {
	conn := server.GetInstance()
	r := incomingRoutes.Group("/kingseafood")

	setupAuthRoutes(r.Group("/auth"), conn)
	setupMenuRoutes(r.Group("/menu"), conn)
	setupTableRoutes(r.Group("/table"), conn)
	setupOrderRoutes(r.Group("/order"), conn)
	setupInvoiceRoutes(r.Group("/invoice"), conn)
	setupFoodRoutes(r.Group("/food"), conn)
	setupOrderItemsRoutes(r.Group("/orderItems"), conn)
	setupAdminRoutes(r.Group("/admin"), conn)
	setupUserRoutes(r.Group("/users"), conn)
}

func setupAuthRoutes(auth *gin.RouterGroup, conn *server.Singleton) {
	auth.POST("/sign-up", handleruser.SignUpHandler(conn.Run()))
	auth.POST("/sign-in/id", handleruser.SignInHandler(conn.Run()))
}

func setupMenuRoutes(menu *gin.RouterGroup, conn *server.Singleton) {
	menu.POST("/", req.CreateMenuHandler(conn.Run()))
	menu.GET("/:menu_id", req.HandlerGetMenu(conn.Run()))
	menu.GET("/", req.HandlerGetListMenu(conn.Run()))
	menu.PATCH("/:menu_id", req.HandlerUpdateMenus(conn.Run()))
	menu.DELETE("/:menu_id", req.HandlerDeleteMenu(conn.Run()))
	menu.GET("viewproduct/:menu_id", req.ViewProductHandler(conn.Run()))
}

func setupTableRoutes(table *gin.RouterGroup, conn *server.Singleton) {
	table.POST("/", req.HandlerCreateTables(conn.Run()))
	table.GET("/:table_id", req.HandlerGetTable(conn.Run()))
	table.GET("/", req.HandlerGetTables(conn.Run()))
	table.PATCH("/:table_id", req.HandlerUpdateTables(conn.Run()))
	table.DELETE("/:table_id", req.HandlerDeletedTable(conn.Run()))
}

func setupOrderRoutes(order *gin.RouterGroup, conn *server.Singleton) {
	order.POST("/", req.HandlerCreateOrder(conn.Run()))
	order.GET("/:order_id", req.HandlerGetOrder(conn.Run()))
	order.PATCH("/:order_id", req.HandlerUpdateOrder(conn.Run()))
	order.DELETE("/:order_id", req.HandlerDeleteOrder(conn.Run()))
}

func setupInvoiceRoutes(invoice *gin.RouterGroup, conn *server.Singleton) {
	invoice.POST("/", req.HandlerCreateInvoice(conn.Run()))
	invoice.GET("/:invoice_id", req.HandlerGetInvoice(conn.Run()))
	invoice.PATCH("/:invoice_id", req.HandlerUpdateInvoices(conn.Run()))
	invoice.DELETE("/:invoice_id", req.HandlerDeleteInvoice(conn.Run()))
}

func setupFoodRoutes(food *gin.RouterGroup, conn *server.Singleton) {
	food.POST("/", req.HandlerCreateProducts(conn.Run()))
	food.GET("/product/:product_id", req.HandlerGetProduct(conn.Run()))
	food.GET("/product/", req.HandlerGetProducts(conn.Run()))
	food.PATCH("/product/:product_id", req.HandlerUpdateProducts(conn.Run()))
	food.PATCH("/product/deleted/:product_id", req.HandlerDeletedProduct(conn.Run()))
	food.GET("/product/searchName/:title", req.HandlerGetProductByName(conn.Run()))
}

func setupOrderItemsRoutes(orderItems *gin.RouterGroup, conn *server.Singleton) {
	orderItems.POST("/", req.HandlerCreateOrderItems(conn.Run()))
	orderItems.GET("/:order_item_id", req.HandlerGetOrderItems(conn.Run()))
	orderItems.GET("/order-items-by-product/:product_id", req.HandlerGetOrderItemsByProduct(conn.Run()))
	orderItems.GET("/order-items-by-order/:order_id", req.HandlerGetOrderItemsByOder(conn.Run()))
	orderItems.PATCH("/:order_item_id", req.HandlerUpdateOrderItems(conn.Run()))
}

func setupAdminRoutes(admin *gin.RouterGroup, conn *server.Singleton) {
	admin.Use(middleware.IsAdmin(), middleware.JwtMiddleware())
	admin.GET("/profile-admin/id", handleruser.ProfileUser(conn.Run()))
	admin.PATCH("/update-admin/id", handleruser.UpdateUserHandler(conn.Run()))
	admin.DELETE("/delete-admin/id", handleruser.DeletedUserHandler(conn.Run()))
	admin.GET("/historyPurchase/:user_id", handleruser.HistoryPurchasesHandler(conn.Run()))
}
func setupUserRoutes(users *gin.RouterGroup, conn *server.Singleton) {
	users.Use(middleware.JwtMiddleware())

	users.GET("/profile/id", handleruser.ProfileUser(conn.Run()))
	users.PATCH("/update/id", handleruser.UpdateUserHandler(conn.Run()))
	users.PATCH("/change/id", handleruser.ChangePwdUserHandler(conn.Run()))
	users.DELETE("/delete/id", handleruser.DeletedUserHandler(conn.Run()))
}

// func KingRouters(incomingRoutes *gin.Engine) {
// 	conn := server.GetInstance()
// 	r := incomingRoutes.Group("/kingseafood")
// 	{
// 		auth := r.Group("/auth")
// 		{
// 			auth.POST("/sign-up", handleruser.SignUpHandler(conn.Run()))
// 			auth.POST("/sign-in/id", handleruser.SignInHandler(conn.Run()))
// 		}

// 		menu := r.Group("/menu")
// 		{
// 			menu.POST("/", req.CreateMenuHandler(conn.Run()))
// 			menu.GET("/:menu_id", req.HandlerGetMenu(conn.Run()))
// 			menu.GET("/", req.HandlerGetListMenu(conn.Run()))
// 			menu.PATCH("/:menu_id", req.HandlerUpdateMenus(conn.Run()))
// 			menu.DELETE("/:menu_id", req.HandlerDeleteMenu(conn.Run()))
// 			menu.GET("viewproduct/:menu_id", req.ViewProductHandler(conn.Run()))
// 		}
// 		table := r.Group("/table")
// 		{
// 			table.POST("/", req.HandlerCreateTables(conn.Run()))
// 			table.GET("/:table_id", req.HandlerGetTable(conn.Run()))
// 			table.GET("/", req.HandlerGetTables(conn.Run()))
// 			table.PATCH("/:table_id", req.HandlerUpdateTables(conn.Run()))
// 			table.DELETE("/:table_id", req.HandlerDeletedTable(conn.Run()))
// 		}
// 		order := r.Group("/order")
// 		{
// 			order.POST("/", req.HandlerCreateOrder(conn.Run()))
// 			order.GET("/:order_id", req.HandlerGetOrder(conn.Run()))
// 			order.PATCH("/:order_id", req.HandlerUpdateOrder(conn.Run()))
// 			order.DELETE("/:order_id", req.HandlerDeleteOrder(conn.Run()))
// 		}
// 		invoice := r.Group("/invoice")
// 		{
// 			invoice.POST("/", req.HandlerCreateInvoice(conn.Run()))
// 			invoice.GET("/:invoice_id", req.HandlerGetInvoice(conn.Run()))

// 			invoice.PATCH("/:invoice_id", req.HandlerUpdateInvoices(conn.Run()))
// 			invoice.DELETE("/:invoice_id", req.HandlerDeleteInvoice(conn.Run()))
// 		}
// 		food := r.Group("/food")
// 		{
// 			food.POST("/", req.HandlerCreateProducts(conn.Run()))
// 			food.GET("/product/:product_id", req.HandlerGetProduct(conn.Run()))
// 			food.GET("/product/", req.HandlerGetProducts(conn.Run()))
// 			food.PATCH("/product/:product_id", req.HandlerUpdateProducts(conn.Run()))
// 			food.PATCH("/product/deleted/:product_id", req.HandlerDeletedProduct(conn.Run()))
// 			food.GET("/product/searchName/:title", req.HandlerGetProductByName(conn.Run()))
// 		}
// 		orderItems := r.Group("/orderItems")
// 		{
// 			orderItems.POST("/", req.HandlerCreateOrderItems(conn.Run()))
// 			orderItems.GET("/:order_item_id", req.HandlerGetOrderItems(conn.Run()))
// 			orderItems.GET("/order-items-by-product/:product_id", req.HandlerGetOrderItemsByProduct(conn.Run()))
// 			orderItems.GET("/order-items-by-order/:order_id", req.HandlerGetOrderItemsByOder(conn.Run()))
// 			orderItems.PATCH("/:order_item_id", req.HandlerUpdateOrderItems(conn.Run()))
// 		}

// 		//ADMIN
// 		admin := r.Group("/admin", middleware.IsAdmin(), middleware.JwtMiddleware())
// 		{
// 			admin.GET("/profile-admin/id", handleruser.ProfileUser(conn.Run()))
// 			admin.PATCH("/update-admin/id", handleruser.UpdateUserHandler(conn.Run()))
// 			admin.DELETE("/delete-admin/id", handleruser.DeletedUserHandler(conn.Run()))
// 			admin.GET("/historyPurchase/:user_id", handleruser.HistoryPurchasesHandler(conn.Run()))
// 		}

// 	}
// }

// users := r.Group("/users", middleware.JwtMiddleware())
// {
// 	users.GET("/profile/id", handleruser.ProfileUser(server.Run()))
// 	users.PATCH("/update/id", handleruser.UpdateUserHandler(server.Run()))
// 	users.DELETE("/delete/id", handleruser.DeletedUserHandler(server.Run()))

// 	menu := r.Group("/menu")
// 	{
// 		menu.GET("menufood/:menu_id", req.HandlerGetMenu(server.Run()))
// 		menu.GET("menufood/", req.HandlerGetListMenu(server.Run()))
// 	}
// 	table := r.Group("/table")
// 	{
// 		table.GET("tablefood/:table_id", req.HandlerGetTable(server.Run()))
// 		table.GET("tablefood/", req.HandlerGetTables(server.Run()))
// 	}
// 	food := r.Group("/food")
// 	{
// 		food.GET("productfood/:product_id", req.HandlerGetProduct(server.Run()))
// 		food.GET("productfood/", req.HandlerGetProducts(server.Run()))
// 		food.GET("productfood/searchName/:title", req.HandlerGetProductByName(server.Run()))
// 	}
// }
