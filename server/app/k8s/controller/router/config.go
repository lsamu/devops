package router

import (
    "github.com/gin-gonic/gin"
    v1 "sass_devops/app/k8s/controller/v1"
)

func InitK8sConfigRouter(Router *gin.RouterGroup) {
    K8sClusterRouter := Router.Group("k8s/config")
    {
        K8sClusterRouter.POST("create", v1.Config.Create) // 新建
        K8sClusterRouter.POST("bat", v1.Config.Bat)       // 批量
        K8sClusterRouter.POST("update", v1.Config.Update) // 更新
        K8sClusterRouter.GET("info", v1.Config.Info)      // 根据ID获取
        K8sClusterRouter.POST("list", v1.Config.List)     // 列表
    }
}
