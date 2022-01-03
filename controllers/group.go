package controllers

import (
	"fmt"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/middlewares"
	"github.com/s-owl/sowl_manager_backend/models"
	"github.com/s-owl/sowl_manager_backend/utils"
)

// GroupController - /user 라우팅 설정
func GroupController(router *gin.RouterGroup) {
	r := router.Group("/group")
	r.Use(middlewares.AdminUserMiddleware())
	{
		r.POST("/regist", groupRegister)
	}
}

// groupRegist godoc
// @Summary 그룹 등록
// @Description 그룹 등록 API
// @ID group-regist
// @Accept json
// @Produce json
// @Param newGroupForm body models.GroupRegistInput true "그룹 등록을 위한 양식"
// @Success 200 {object} models.InfoDTO
// @Failure 400 {object} models.ErrorDTO
// @Router /group/regist [post]
func groupRegister(c *gin.Context) {
	var err error = nil
	var group *models.Group
	groupInput := models.GroupRegisterInput{}
	adminUser := c.MustGet("user").(*auth.UserRecord)

	if err = c.ShouldBindJSON(&groupInput); err == nil {
		group, err = groupRegisterLogic(c, &groupInput, adminUser.UserInfo.Email)
	} else {
		err = utils.GinJSONMarshalError(err)
	}

	if err != nil {
		err = fmt.Errorf("GroupRegister: %w", err)
		utils.AbortWithHTTPError(c, err)
		return
	}

	log.Printf("Successfully created group: %v", group)
}
