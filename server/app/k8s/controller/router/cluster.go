package router

import (
    "github.com/gin-gonic/gin"
    v1 "sass_devops/app/k8s/controller/v1"
)

func InitK8sClusterRouter(Router *gin.RouterGroup) {
    K8sClusterRouter := Router.Group("k8s/cluster")
    {
        K8sClusterRouter.POST("create", v1.Cluster.Create) // 新建
        K8sClusterRouter.POST("bat", v1.Cluster.Bat)       // 批量
        K8sClusterRouter.POST("update", v1.Cluster.Update) // 更新
        K8sClusterRouter.GET("info", v1.Cluster.Info)      // 根据ID获取
        K8sClusterRouter.POST("list", v1.Cluster.List)     // 列表
    }
}
