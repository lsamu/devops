package conf

import (
    "encoding/json"
    "fmt"
    "github.com/BurntSushi/toml"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "strings"
)

type ServerConfig struct {
    Debug bool
    Jwt   struct {
        Secret string
    }
    K8s struct {
        Host string
        Port uint32
    }
    Api struct {
        Host string
        Port uint32
    }
    MySql struct {
        Host     string
        Port     uint32
        UserName string
        Password string
        Database string
    }
    Redis struct {
        Address  string
        Password string
        DB       int
    }
    GinSession struct {
        Name      string
        SecretKey string
    }
    Email struct {
    }
}

var _cfg ServerConfig

func InitConfig() *ServerConfig {
    err := Load("server.toml", &_cfg)
    Dump(&_cfg)
    if nil != err {
        panic(err)
    }
    return &_cfg
}

func Load(path string, v interface{}) error {
    root := os.Getenv("SERVER_ROOT")
    if "" == root {
        root = "."
    }
    confPath := root + "/" + path
    content, err := ioutil.ReadFile(confPath)
    if nil != err {
        panic("Can not found conf file")
    }
    if strings.HasSuffix(confPath, ".json") {
        err = json.Unmarshal(content, v)
    }
    if strings.HasSuffix(confPath, ".yml") {
        err = yaml.Unmarshal(content, v)
    }
    if strings.HasSuffix(confPath, ".toml") {
        err = toml.Unmarshal(content, v)
    }
    return err
}

func Dump(v interface{}) {
    fmt.Println("config begin:--------------------------------")
    data, _ := json.MarshalIndent(v, "", "  ")
    fmt.Println(string(data))
    fmt.Println("config end:--------------------------------")
}
