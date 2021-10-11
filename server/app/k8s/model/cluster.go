package model

type K8sCluster struct {
    BaseModel
    ID             uint    `json:"id" gorm:"primarykey" form:"id"`
    ClusterName    string  `json:"clusterName" gorm:"comment:集群名称"`
    KubeConfig     string  `json:"kubeConfig" gorm:"comment:Kubeconfig内容;type:varchar(12800)"`
    ClusterVersion float32 `json:"clusterVersion" gorm:"comment:集群版本"`
}

func (*K8sCluster) TableName()	string  {
    return "k8s_clusters"
}
