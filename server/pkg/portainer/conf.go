package portainer

import (
    portainer2 "github.com/leidruid/go-portainer"
    "log"
)

func InitPortainer()  {
    var err error
    portainer:=portainer2.NewPortainer(&portainer2.Config{
        Host:     "",
        Port:     0,
        Schema:   "",
        User:     "",
        Password: "",
        URL:      "",
    })
    err = portainer.Auth()
    if err!=nil{
        log.Print(err)
    }

}
