package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/lib/logger"
    "github.com/lsamu/ago/rest/handler"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/controller/response"
    "sass_devops/app/k8s/service"
)

var Cluster ClusterController

func init() {
    Cluster = ClusterController{}
}

type ClusterController struct {
}

// @Description 列表
// @Tags k8s集群
// @Router	/api/k8s/cluster/list [post]
func (u *ClusterController) List(c *gin.Context) {
    req := new(request.ClusterRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    list,total, err := service.ClusterHandler.List(req)
    if err != nil {
        logger.Info("Cluster", "list", "err", err.Error())
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
// @Router	/api/k8s/cluster/info [get]
func (u *ClusterController) Info(c *gin.Context) {
    req := new(request.ClusterInfoRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    info, err := service.ClusterHandler.Info(req)
    if err != nil {
        logger.Info("Cluster", "info", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, info)
}

// @Description 创建
// @Tags k8s集群
// @Router	/api/k8s/cluster/create [post]
func (u *ClusterController) Create(c *gin.Context) {
    req := new(request.ClusterCreateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    _, err := service.ClusterHandler.Create(req)
    if err != nil {
        logger.Info("Cluster", "create", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 更新
// @Tags k8s集群
// @Router	/api/k8s/cluster/update [post]
func (u *ClusterController) Update(c *gin.Context) {
    req := new(request.ClusterUpdateRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.ClusterHandler.Update(req)
    if err != nil {
        logger.Info("Cluster", "update", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}

// @Description 批量
// @Tags k8s集群
// @Router	/api/k8s/cluster/bat [post]
func (u *ClusterController) Bat(c *gin.Context) {
    req := new(request.ClusterBatRequest)
    cusErr := handler.Parse(c, req)
    if cusErr != nil {
        handler.JSON(c, handler.CodeErr,cusErr.Error())
        return
    }
    err := service.ClusterHandler.Bat(req)
    if err != nil {
        logger.Info("Cluster", "bat", "err", err.Error())
        handler.JSON(c, handler.CodeErr, err.Error())
        return
    }
    handler.JSON(c, handler.CodeOK, handler.MsgOK, nil)
}
