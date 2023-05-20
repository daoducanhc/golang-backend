package handler

import (
	"net/http"
	"os"
	"std/pkg/dto"
	"std/pkg/entity"
	"std/pkg/utils"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *handler) getUser(c *gin.Context) {
	username, err := ExtractUsernameFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEnt, err := h.UserService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDto := convertUserEntityToDto(userEnt)

	// show username, nickname, profile picture
	// c.HTML(http.StatusOK, "profile.html", userDto)
	c.JSON(http.StatusOK, userDto)
}

func (h *handler) login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if govalidator.IsNull(input.Username) || govalidator.IsNull(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data can not empty"})
		return
	}

	username := input.Username
	password := input.Password

	// str := c.Request.URL.Path

	// u, err := url.Parse(str)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// username := u.User.Username()
	// password, _ := u.User.Password()

	// username, err := ExtractUsernameFromToken(c)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	b, err := h.UserService.CheckLogin(c.Request.Context(), username, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if b {
		token, err := GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		// userEnt, err := h.UserService.GetUserByUsername(c.Request.Context(), username)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// userDto := convertUserEntityToDto(userEnt)
		// c.JSON(http.StatusOK, userDto)
	}

}

func (h *handler) editNickname(c *gin.Context) {
	username, err := ExtractUsernameFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req struct {
		Nickname string `form:"nickname" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if govalidator.IsNull(req.Nickname) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nickname can not empty"})
		return
	}

	b, err := h.UserService.ChangeNickname(c.Request.Context(), username, req.Nickname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if b {
		c.JSON(http.StatusOK, gin.H{"message": "Nickname changed successfully"})
	}
}

func (h *handler) editPicture(c *gin.Context) {
	username, err := ExtractUsernameFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	randString := utils.RandomString()

	path, _ := os.Getwd()
	filename := path + "/assets/pictures/" + randString + username + ".jpg"
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	b, err := h.UserService.ChangePicture(c.Request.Context(), username, filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		if err := os.Remove(filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	if b {
		c.JSON(http.StatusOK, gin.H{"message": "Profile picture changed successfully"})
	}

}

func convertUserEntityToDto(userEnt *entity.UserEntity) *dto.UserDto {
	picture_url := userEnt.Picture_url
	if len(strings.TrimSpace(picture_url)) == 0 {
		picture_url = "../../assets/pictures/default.jpg"
	}
	return &dto.UserDto{
		Username:    userEnt.Username,
		Nickname:    userEnt.Nickname,
		Picture_url: picture_url,
	}
}
