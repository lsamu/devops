package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/lib/logger"
    "github.com/lsamu/ago/rest/handler"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/controller/response"
    "sass_devops/app/k8s/service"
)

var Volume VolumeController

func init() {
    Volume = VolumeController{}
}

type VolumeController struct {
}

// @Description 列表
// @Tags k8s集群
// @Router	/api/k8s/Volume/list [post]
func (u *VolumeController) List(c *gin.Context) {
    req := new(request.VolumeRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    list,total, err := service.VolumeHandler.List(req)
    if err != nil {
        logger.Info("Volume", "list", "err", err.Error())
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
// @Router	/api/k8s/Volume/info [get]
func (u *VolumeController) Info(c *gin.Context) {
    req := new(request.VolumeInfoRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    info, err := service.VolumeHandler.Info(req)
    if err != nil {
        logger.Info("Volume", "info", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, info)
}

// @Description 创建
// @Tags k8s集群
// @Router	/api/k8s/Volume/create [post]
func (u *VolumeController) Create(c *gin.Context) {
    req := new(request.VolumeCreateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    _, err := service.VolumeHandler.Create(req)
    if err != nil {
        logger.Info("Volume", "create", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 更新
// @Tags k8s集群
// @Router	/api/k8s/Volume/update [post]
func (u *VolumeController) Update(c *gin.Context) {
    req := new(request.VolumeUpdateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.VolumeHandler.Update(req)
    if err != nil {
        logger.Info("Volume", "update", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 批量
// @Tags k8s集群
// @Router	/api/k8s/Volume/bat [post]
func (u *VolumeController) Bat(c *gin.Context) {
    req := new(request.VolumeBatRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.VolumeHandler.Bat(req)
    if err != nil {
        logger.Info("Volume", "bat", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

