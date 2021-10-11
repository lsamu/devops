package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/lsamu/ago/rest"
    "sass_devops/app/k8s/controller/router"
)

func InitRouter(host string, port int) {
    server := rest.NewServer(rest.RestConf{
        Host: host,
        Port: port,
    })
    defer server.Stop()
    engine := server.GetEngine()
    err := RunRouter(engine)
    if err != nil {
        return
    }
    server.Start()
}

func RunRouter(eng *gin.Engine) (err error) {
    k8s := eng.Group("api")
    router.InitK8sClusterRouter(k8s)
    router.InitK8sConfigRouter(k8s)
    router.InitK8sDeployRouter(k8s)
    router.InitK8sSecretRouter(k8s)
    router.InitK8sVolumeRouter(k8s)
    return
}
