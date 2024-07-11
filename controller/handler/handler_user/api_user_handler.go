package handleruser

import (
	"fmt"
	"net/http"
	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/req_users"
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
		var reqSignUp req_users.RequestSignUp
		if err := c.ShouldBind(&reqSignUp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't request sign up",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(reqSignUp); err != nil {
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
		pwd := security.HashAndSalt([]byte(reqSignUp.Password))
		role := model.MEMBERS.String()
		users = model.Users{
			UserID:    idUsers.String(),
			FullName:  reqSignUp.FullName,
			Email:     reqSignUp.Email,
			Password:  pwd,
			Male:      reqSignUp.Male,
			Role:      role,
			CreatedAt: &time,
			UpdatedAt: &time,
		}
		biz := users_bussiness.NewSignUpController(repouserimpl.NewSql(db))
		dataUser, err := biz.NewSignUp(c.Request.Context(), &users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"comment": "Invalid database users",
			})
			return
		}
		accessToken, refreshToken, err := security.JwtToken(&users)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "faild token",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseDataToken(dataUser, accessToken, refreshToken))
	}
}

// SIGNIN
func SignInHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		biz := users_bussiness.NewSignInController(repouserimpl.NewSql(db))
		data, err := biz.NewSignIn(c.Request.Context(), &reqSignIn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "reqSignIn database failed",
			})
			return
		}
		isValidPassword := security.ComparePasswords(data.Password, []byte(reqSignIn.Password))
		if !isValidPassword {
			c.JSON(http.StatusBadRequest, gin.H{
				"comment": "Email or password is incorrect",
			})
			return
		}
		//updated token
		acToken, reqToken, _ := security.UpdateToken(data)
		c.JSON(http.StatusOK, gin.H{
			"result":        "signin successful",
			"access token":  acToken,
			"refresh token": reqToken,
		})

		c.JSON(http.StatusOK, common.ReponseDataToken("Đăng nhập thành công", acToken, reqToken))
	}
}

func ProfileUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
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

		biz := users_bussiness.NewProfileController(repouserimpl.NewSql(db))
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
		hashedPassword := security.HashAndSalt([]byte(req_update.Password))
		req_update.Password = hashedPassword
		validate := validator.New()
		if err := validate.Struct(req_update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validator",
			})
			return
		}
		userID, exists := c.Get("user_id")
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
		biz := users_bussiness.NewUpdateUserController(repouserimpl.NewSql(db))
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
		token, ok := c.Get("user_id")
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
		biz := users_bussiness.NewDeletedUserService(repouserimpl.NewSql(db))
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
