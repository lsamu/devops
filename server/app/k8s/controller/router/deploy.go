package router

import (
    "github.com/gin-gonic/gin"
    v1 "sass_devops/app/k8s/controller/v1"
)

func InitK8sDeployRouter(Router *gin.RouterGroup) {
    K8sClusterRouter := Router.Group("k8s/deploy")
    {
        K8sClusterRouter.POST("create", v1.Deploy.Create) // 新建
        K8sClusterRouter.POST("bat", v1.Deploy.Bat)       // 批量
        K8sClusterRouter.POST("update", v1.Deploy.Update) // 更新
        K8sClusterRouter.GET("info", v1.Deploy.Info)      // 根据ID获取
        K8sClusterRouter.POST("list", v1.Deploy.List)     // 列表
    }
}
