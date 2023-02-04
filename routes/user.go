package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func userRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})

	users.GET("/:userId", func(c *gin.Context) {
		userId := c.Params.ByName("userId")
		if !core.IsValidUUID(userId) {
			c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
				Message: "user not found",
			})
			return
		}

		user, err := repository.GetUserById(userId)
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
					Message: "user not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, schemas.UserDetailResponse{
			Id:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			IsActive:    user.IsActive,
			IsSuperuser: user.IsSuperuser,
		})
	})

	users.POST("/", func(c *gin.Context) {
		var newUser schemas.UserCreateRequest
		err := c.BindJSON(&newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
				Message: err.Error(),
			})
			return
		}

		now := time.Now()
		createdUser, err := repository.CreateUser(newUser.Username, newUser.Email, newUser.Password, newUser.IsActive, newUser.IsSuperuser, now, &now)
		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, schemas.UserCreateResponse{
			Id:          createdUser.ID,
			Username:    createdUser.Username,
			Email:       createdUser.Email,
			IsActive:    createdUser.IsActive,
			IsSuperuser: createdUser.IsSuperuser,
		})
	})

	users.PUT("/:userId", func(c *gin.Context) {
		// get input user
		userId := c.Params.ByName("userId")
		if !core.IsValidUUID(userId) {
			c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
				Message: "user not found",
			})
			return
		}
		jsonRequest := schemas.UserUpdateRequest{}
		err := c.BindJSON(&jsonRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
				Message: err.Error(),
			})
			return
		}

		// get existing user
		user, err := repository.GetUserById(userId)
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
					Message: "user not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		// update user
		updatedUser, err := repository.UpdateUser(
			user,
			jsonRequest.Email,
			jsonRequest.Username,
			jsonRequest.Password,
			jsonRequest.IsActive,
			jsonRequest.IsSuperuser,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, schemas.UserUpdateResponse{
			Id:          updatedUser.ID,
			Username:    updatedUser.Username,
			Email:       updatedUser.Email,
			IsActive:    updatedUser.IsActive,
			IsSuperuser: updatedUser.IsSuperuser,
		})
	})

	users.DELETE("/:userId", func(c *gin.Context) {
		// get input user
		userId := c.Params.ByName("userId")
		if !core.IsValidUUID(userId) {
			c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
				Message: "user not found",
			})
			return
		}

		// get existing user
		user, err := repository.GetUserById(userId)
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, schemas.NotFoundResponse{
					Message: "user not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}

		_, err = repository.DeleteUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})
}
