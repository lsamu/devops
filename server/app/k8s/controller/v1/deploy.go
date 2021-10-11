package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/lib/logger"
    "github.com/lsamu/ago/rest/handler"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/controller/response"
    "sass_devops/app/k8s/service"
)

var Deploy DeployController

func init() {
    Deploy = DeployController{}
}

type DeployController struct {
}

// @Description 列表
// @Tags k8s集群
// @Router	/api/k8s/Deploy/list [post]
func (u *DeployController) List(c *gin.Context) {
    req := new(request.DeployRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    list,total, err := service.DeployHandler.List(req)
    if err != nil {
        logger.Info("Deploy", "list", "err", err.Error())
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
// @Router	/api/k8s/Deploy/info [get]
func (u *DeployController) Info(c *gin.Context) {
    req := new(request.DeployInfoRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    info, err := service.DeployHandler.Info(req)
    if err != nil {
        logger.Info("Deploy", "info", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, info)
}

// @Description 创建
// @Tags k8s集群
// @Router	/api/k8s/Deploy/create [post]
func (u *DeployController) Create(c *gin.Context) {
    req := new(request.DeployCreateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    _, err := service.DeployHandler.Create(req)
    if err != nil {
        logger.Info("Deploy", "create", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 更新
// @Tags k8s集群
// @Router	/api/k8s/Deploy/update [post]
func (u *DeployController) Update(c *gin.Context) {
    req := new(request.DeployUpdateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.DeployHandler.Update(req)
    if err != nil {
        logger.Info("Deploy", "update", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 批量
// @Tags k8s集群
// @Router	/api/k8s/Deploy/bat [post]
func (u *DeployController) Bat(c *gin.Context) {
    req := new(request.DeployBatRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.DeployHandler.Bat(req)
    if err != nil {
        logger.Info("Deploy", "bat", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

