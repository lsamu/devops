package model

import "time"

type K8sConfig struct {
    ID            uint32    `gorm:"type:int(11);auto_increment;not null;primary_key;comment:'自增id'" json:"id"`
    ConfigName    string    `gorm:"type:varchar(255);default null;" json:"config_name"`
    ConfigContent string    `gorm:"type:text;default null;comment:'configmap内容'" json:"config_content"`
    Content       string    `gorm:"type:text;default null;comment:'Secret json/yaml'" json:"content"`
    Status        int       `gorm:"type:int(1);default 0;comment:'结果0/1'" json:"status"`
    K8sEnv        string    `gorm:"type:varchar(10);default null;comment:'dev prod'" json:"k8s_env"`
    CreatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_time"`
    UpdatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"updated_time"`
}

func (*K8sConfig) TableName() string {
    return "k8s_config"
}
