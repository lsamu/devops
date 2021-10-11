package model

import "time"

type K8sSecret struct {
    ID            uint32    `gorm:"type:int(11);auto_increment;not null;primary_key;comment:'自增id'" json:"id"`
    SecretName    string    `gorm:"type:varchar(255);default null;" json:"secret_name"`
    SecretContent string    `gorm:"type:text;default null;comment:'secret内容'" json:"secret_content"`
    Status        int       `gorm:"type:int(1);default null;comment:'结果0/1'" json:"status"`
    Content       string    `gorm:"type:text;default null;comment:'Secret json/yaml'" json:"content"`
    K8sEnv        string    `gorm:"type:varchar(10);default null;comment:'dev prod'" json:"k8s_env"`
    CreatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_time"`
    UpdatedTime   time.Time `gorm:"type:datetime;default:current_timestamp" json:"updated_time"`
}

func (*K8sSecret) TableName() string {
    return "k8s_secret"
}

