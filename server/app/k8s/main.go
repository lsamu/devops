package main

import (
    "fmt"
    "github.com/lsamu/ago/lib/jwt"
    "sass_devops/app/k8s/controller"
    "sass_devops/app/k8s/model"
    "sass_devops/conf"
)

func main() {
    cfg := conf.InitConfig()
    jwt.InitJwt(cfg.Jwt.Secret)
    model.InitMySql(cfg.MySql.UserName,
        cfg.MySql.Password,
        cfg.MySql.Host,
        fmt.Sprintf("%d", cfg.MySql.Port),
        cfg.MySql.Database, cfg.Debug)
    controller.InitRouter(cfg.K8s.Host, int(cfg.K8s.Port))
}
