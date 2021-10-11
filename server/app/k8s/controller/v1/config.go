package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/lib/logger"
    "github.com/lsamu/ago/rest/handler"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/controller/response"
    "sass_devops/app/k8s/service"
)

var Config ConfigController

func init() {
    Config = ConfigController{}
}

type ConfigController struct {
}

// @Description 列表
// @Tags k8s配置
// @Router	/api/k8s/Config/list [post]
func (u *ConfigController) List(c *gin.Context) {
    req := new(request.ConfigRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    list,total, err := service.ConfigHandler.List(req)
    if err != nil {
        logger.Info("Config", "list", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, response.BasePage{
        List:      list,
        Total:     total,
    })
}

// @Description 详细
// @Tags k8s配置
// @Router	/api/k8s/Config/info [get]
func (u *ConfigController) Info(c *gin.Context) {
    req := new(request.ConfigInfoRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    info, err := service.ConfigHandler.Info(req)
    if err != nil {
        logger.Info("Config", "info", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, info)
}

// @Description 创建
// @Tags k8s配置
// @Router	/api/k8s/Config/create [post]
func (u *ConfigController) Create(c *gin.Context) {
    req := new(request.ConfigCreateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    _, err := service.ConfigHandler.Create(req)
    if err != nil {
        logger.Info("Config", "create", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 更新
// @Tags k8s配置
// @Router	/api/k8s/Config/update [post]
func (u *ConfigController) Update(c *gin.Context) {
    req := new(request.ConfigUpdateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.ConfigHandler.Update(req)
    if err != nil {
        logger.Info("Config", "update", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 批量
// @Tags k8s配置
// @Router	/api/k8s/Config/bat [post]
func (u *ConfigController) Bat(c *gin.Context) {
    req := new(request.ConfigBatRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.ConfigHandler.Bat(req)
    if err != nil {
        logger.Info("Config", "bat", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

