package handleruser

import (
	"fmt"
	"net/http"
	"strings"
	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/req_users"
	redisdb "thelastking/kingseafood/pkg/redisDB"
	"thelastking/kingseafood/pkg/security"
	"thelastking/kingseafood/repository/users_bussiness"
	repouserimpl "thelastking/kingseafood/repository/users_bussiness/repo_user_impl"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SignUpHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users model.Users
		if err := c.ShouldBind(&users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't request sign up",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validator",
			})
			return
		}

		idUsers, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "uuid fails",
			})
			return
		}
		time := time.Now().UTC()
		pwd := security.HashAndSalt([]byte(users.Password_user))
		role := model.MEMBERS.String()
		users1 := &model.Users{
			UserID:        idUsers.String(),
			FullName:      users.FullName,
			Email:         users.Email,
			Password_user: pwd,
			Male:          users.Male,
			Role:          role,
			CreatedAt:     &time,
			UpdatedAt:     &time,
		}
		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		dataUser, err := biz.NewSignUp(c.Request.Context(), users1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"comment": "Invalid database users",
			})
			return
		}
		accessToken, refreshToken, err := security.JwtToken(users1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild token",
			})
			return
		}
		//SAVE REFRESH TOKEN
		redisInstance := redisdb.GetInstanceRedis()
		Rediserr := redisInstance.SaveRefreshToken(refreshToken)
		if Rediserr != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"error":   "save refresh token faild",
				"details": Rediserr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseDataToken(dataUser, accessToken, refreshToken))
	}
}

// SIGNIN
func SignInHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//tao 1 access token moi neu refresh con han
		refreshToken := c.GetHeader("Authorization")
		if refreshToken != "" {
			refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")
		}

		// Nếu có refresh token, thử cập nhật access token
		if refreshToken != "" {
			newAccessToken, err := security.UpdateToken(refreshToken)
			if err == nil {
				// Nếu cập nhật thành công, trả về access token mới
				c.JSON(http.StatusOK, gin.H{
					"result":       "signin successful",
					"access token": newAccessToken,
				})
				return
			}
		}
		var reqSignIn req_users.RequestSignIn
		if err := c.ShouldBind(&reqSignIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "reqSignIn failed",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(reqSignIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validate",
			})
			return
		}

		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		data, err := biz.NewSignIn(c.Request.Context(), &reqSignIn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "reqSignIn database failed",
			})
			return
		}
		isValidPassword := security.ComparePasswords(data.Password_user, []byte(reqSignIn.Password))
		if !isValidPassword {
			c.JSON(http.StatusBadRequest, gin.H{
				"comment": "Email or password is incorrect",
			})
			return
		}
		//cap nhat lai token
		acToken, reqToken, _ := security.JwtToken(data)
		c.JSON(http.StatusOK, gin.H{
			"result":        "signin successful",
			"access token":  acToken,
			"refresh token": reqToken,
		})
	}
}

func ProfileUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userId")
		if !exists || userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"comment": "Missing user ID in token",
			})
			return
		}
		claims, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid user ID in token",
				"comment": "Failed to convert to string",
			})
			return
		}

		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		dataUser, err := biz.NewProfileUserByID(c.Request.Context(), claims)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "error dataUser",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(dataUser))
	}
}
func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req_update req_users.UpdateUsers
		if err := c.ShouldBind(&req_update); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"comment": "loi req_updated",
			})
			return
		}
		// hashedPassword := security.HashAndSalt([]byte(req_update.Password_user))
		// req_update.Password_user = hashedPassword
		// validate := validator.New()
		// if err := validate.Struct(req_update); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error":   err.Error(),
		// 		"comment": "Can't validator",
		// 	})
		// 	return
		// }
		userID, exists := c.Get("userId")
		if !exists || userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"comment": "Missing user ID in token",
			})
			return
		}
		claims, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid user ID in token",
				"comment": "Failed to convert to string",
			})
			return
		}
		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		if err := biz.NewUpdateUser(c.Request.Context(), &req_update, claims); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"comment": "loi tai biz",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Update thanh cong"))
	}
}

func DeletedUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Get("userId")
		if !ok || token == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"comment": "Missing user ID in token",
			})
			return
		}
		claims, ok := token.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid user ID in token",
				"comment": "Failed to convert to string",
			})
			return
		}
		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		fmt.Printf("biz: %v\n", biz)
		if err := biz.NewDeletedUserByID(c.Request.Context(), claims); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"comment": "Failed deleted",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Da xoa thanh cong nguoi dung"))
	}
}
func HistoryPurchasesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idUser := c.Param("user_id")
		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		dataUser, dataListProduct, err := biz.NewHistoryPurchases(c.Request.Context(), idUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"comment": "Failed seen history",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"User":           dataUser,
			"historyPurchas": dataListProduct,
		})
	}
}

func ChangePwdUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req_update req_users.ChangePwd
		if err := c.ShouldBind(&req_update); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"comment": "req_updated error",
			})
			return
		}
		hashedPassword := security.HashAndSalt([]byte(req_update.Password_user))
		req_update.Password_user = hashedPassword
		validate := validator.New()
		if err := validate.Struct(req_update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validator",
			})
			return
		}
		updUser := &req_users.ChangePwd{
			Email:         req_update.Email,
			Password_user: req_update.Password_user,
		}
		userID, exists := c.Get("userId")
		if !exists || userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"comment": "Missing user ID in token",
			})
			return
		}
		claims, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid user ID in token",
				"comment": "Failed to convert to string",
			})
			return
		}
		biz := users_bussiness.NewUserController(repouserimpl.NewSql(db))
		if err := biz.NewChangePwdUser(c.Request.Context(), claims, updUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors":  err.Error(),
				"comment": "biz error",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData("Change pwd success"))
	}
}
