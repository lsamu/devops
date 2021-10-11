package service

import (
    "errors"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/model"
)

var ClusterHandler ClusterService

func init() {
    ClusterHandler = ClusterService{}
}

type ClusterService struct {
}

func (p *ClusterService) List(req *request.ClusterRequest) (list []model.K8sCluster, total int64, err error) {
    startIndex := req.PageSize * (req.PageIndex - 1)
    pageSize := req.PageSize
    Cluster := model.K8sCluster{}
    db := model.Conn().Model(&Cluster).Order("id desc")
    err = db.Offset(startIndex).Limit(pageSize).Find(&list).Error
    if err != nil {
        return nil, 0, err
    }
    err = db.Count(&total).Error
    if err != nil {
        return nil, 0, err
    }
    return
}

func (p *ClusterService) Info(req *request.ClusterInfoRequest) (info model.K8sCluster, err error) {
    Cluster := model.K8sCluster{}
    err = model.Conn().Model(&Cluster).Find(&info).Error
    if err != nil {
        return info, err
    }
    return
}

func (p *ClusterService) Create(req *request.ClusterCreateRequest) (lastId uint32, err error) {
    Cluster := model.K8sCluster{}
    var recordCount int64
    err = model.Conn().Model(&Cluster).Where("name=?", req.Name).Count(&recordCount).Error
    if err != nil {
        return 0, err
    }
    if recordCount > 0 {
        return 0, errors.New("已存在该记录！")
    }
    err = model.Conn().Model(&Cluster).Create(&Cluster).Error
    if err != nil {
        return 0, err
    }
    return
}

func (p *ClusterService) Update(req *request.ClusterUpdateRequest) (err error) {
    Cluster := model.K8sCluster{}
    data := make(map[string]interface{}, 0)
    err = model.Conn().Model(&Cluster).Where("id=?", req.ID).Updates(&data).Error
    if err != nil {
        return err
    }
    return
}

func (p *ClusterService) Bat(req *request.ClusterBatRequest) (err error) {
    Cluster := model.K8sCluster{}
    switch req.OP {
    case "del":
        err = model.Conn().Model(&Cluster).Where("id in ? ", req.ID).Delete(Cluster).Error
        break
    }
    if err != nil {
        return err
    }
    return
}