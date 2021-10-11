package router

import (
    "github.com/gin-gonic/gin"
    v1 "sass_devops/app/k8s/controller/v1"
)

func InitK8sSecretRouter(Router *gin.RouterGroup) {
    K8sClusterRouter := Router.Group("k8s/secret")
    {
        K8sClusterRouter.POST("create", v1.Secret.Create) // 新建
        K8sClusterRouter.POST("bat", v1.Secret.Bat)       // 批量
        K8sClusterRouter.POST("update", v1.Secret.Update) // 更新
        K8sClusterRouter.GET("info", v1.Secret.Info)      // 根据ID获取
        K8sClusterRouter.POST("list", v1.Secret.List)     // 列表
    }
}
