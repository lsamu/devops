package router

import (
    "github.com/gin-gonic/gin"
    v1 "sass_devops/app/k8s/controller/v1"
)

func InitK8sVolumeRouter(Router *gin.RouterGroup) {
    K8sClusterRouter := Router.Group("k8s/volume")
    {
        K8sClusterRouter.POST("create", v1.Volume.Create) // 新建
        K8sClusterRouter.POST("bat", v1.Volume.Bat)       // 批量
        K8sClusterRouter.POST("update", v1.Volume.Update) // 更新
        K8sClusterRouter.GET("info", v1.Volume.Info)      // 根据ID获取
        K8sClusterRouter.POST("list", v1.Volume.List)     // 列表
    }
}
