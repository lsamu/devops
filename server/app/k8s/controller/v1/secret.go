package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/lib/logger"
    "github.com/lsamu/ago/rest/handler"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/controller/response"
    "sass_devops/app/k8s/service"
)

var Secret SecretController

func init() {
    Secret = SecretController{}
}

type SecretController struct {
}

// @Description 列表
// @Tags k8s集群
// @Router	/api/k8s/Secret/list [post]
func (u *SecretController) List(c *gin.Context) {
    req := new(request.SecretRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    list,total, err := service.SecretHandler.List(req)
    if err != nil {
        logger.Info("Secret", "list", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, response.BasePage{
        List:      list,
        Total:     total,
    })
}

// @Description 详细
// @Tags k8s集群
// @Router	/api/k8s/Secret/info [get]
func (u *SecretController) Info(c *gin.Context) {
    req := new(request.SecretInfoRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    info, err := service.SecretHandler.Info(req)
    if err != nil {
        logger.Info("Secret", "info", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, info)
}

// @Description 创建
// @Tags k8s集群
// @Router	/api/k8s/Secret/create [post]
func (u *SecretController) Create(c *gin.Context) {
    req := new(request.SecretCreateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    _, err := service.SecretHandler.Create(req)
    if err != nil {
        logger.Info("Secret", "create", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 更新
// @Tags k8s集群
// @Router	/api/k8s/Secret/update [post]
func (u *SecretController) Update(c *gin.Context) {
    req := new(request.SecretUpdateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.SecretHandler.Update(req)
    if err != nil {
        logger.Info("Secret", "update", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 批量
// @Tags k8s集群
// @Router	/api/k8s/Secret/bat [post]
func (u *SecretController) Bat(c *gin.Context) {
    req := new(request.SecretBatRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.SecretHandler.Bat(req)
    if err != nil {
        logger.Info("Secret", "bat", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

