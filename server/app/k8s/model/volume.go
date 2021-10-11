package model

import "time"

type K8sVolume struct {
    ID          uint32    `gorm:"type:int(11);auto_increment;not null;primary_key;comment:'自增id'" json:"id"`
    Name        string    `gorm:"type:varchar(255);default null;comment:'挂载名称'" json:"name"`
    VolumeType  string    `gorm:"type:varchar(255);default null;comment:'挂载类型'" json:"volume_type"`
    Path        string    `gorm:"type:varchar(255);default null;comment:'path'" json:"path"`
    Storage     string    `gorm:"type:varchar(255);default null;comment:'挂载大小'" json:"storage"`
    K8sEnv      string    `gorm:"type:varchar(10);default null;comment:'dev prod'" json:"k8s_env"`
    Status      int       `gorm:"type:int(1);default null;comment:'结果0/1'" json:"status"`
    CreatedTime time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_time"`
    UpdatedTime time.Time `gorm:"type:datetime;default:current_timestamp" json:"updated_time"`
}

func (*K8sVolume) TableName() string {
    return "k8s_volume"
}
