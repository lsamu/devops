package model

import "time"

type K8sDeploy struct {
    ID            uint32    `gorm:"type:int(11);auto_increment;not null;primary_key;comment:'自增id'" json:"id"`
    PodName       string    `gorm:"type:varchar(255);default null;comment:'Pod名称'" json:"pod_name"`
    AppName       string    `gorm:"type:varchar(50);default null;comment:'app_name'" json:"app_name"`
    Image         string    `gorm:"type:varchar(255);default 0;comment:'镜像'" json:"image"`
    Command       string    `gorm:"type:varchar(255);default 0;comment:'启动命令'" json:"command"`
    ContainerPort string    `gorm:"type:int(11);default null;comment:'容器Port'" json:"container_port"`
    ConfigList    string    `gorm:"type:varchar(255);default null;comment:'配置列表'" json:"config_list"`
    SecretList    string    `gorm:"type:varchar(255);default null;comment:'密钥列表'" json:"secret_list"`
    EnvList       string    `gorm:"type:varchar(255);default null;comment:'环境列表'" json:"env_list"`
    VolumeList    string    `gorm:"type:varchar(255);default:'';comment:'磁盘列表'" json:"volume_list"`
    CPU           string    `gorm:"type:varchar(255);default null;comment:'cpu资源'" json:"cpu"`
    Memory        string    `gorm:"type:varchar(255);default null;comment:'memory资源'" json:"memory"`
    Status        int       `gorm:"type:int(1);default null;comment:'结果0/1'" json:"status"`
    ClusterId     string    `gorm:"type:int(11);default 0;comment:'所属集群'" json:"cluster"`
    PodNum        int       `gorm:"type:int(11);default 1;comment:'pod_num 数量'" json:"pod_num"`
    NodeName      string    `gorm:"type:varchar(50);not null;default:'';comment:'node name'" json:"node_name"`
    CreatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_time"`
    UpdatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"updated_time"`
}

func (k *K8sDeploy) TableName() string {
    return "k8s_deploy_service"
}
