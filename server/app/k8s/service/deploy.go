package service

import (
    "errors"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/model"
)

var DeployHandler DeployService

func init() {
    DeployHandler = DeployService{}
}

type DeployService struct {
}

func (p *DeployService) List(req *request.DeployRequest) (list []model.K8sDeploy, total int64, err error) {
    startIndex := req.PageSize * (req.PageIndex - 1)
    pageSize := req.PageSize
    Deploy := model.K8sDeploy{}
    db := model.Conn().Model(&Deploy).Order("id desc")
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

func (p *DeployService) Info(req *request.DeployInfoRequest) (info model.K8sDeploy, err error) {
    Deploy := model.K8sDeploy{}
    err = model.Conn().Model(&Deploy).Find(&info).Error
    if err != nil {
        return info, err
    }
    return
}

func (p *DeployService) Create(req *request.DeployCreateRequest) (lastId uint32, err error) {
    Deploy := model.K8sDeploy{}
    var recordCount int64
    err = model.Conn().Model(&Deploy).Where("name=?", req.Name).Count(&recordCount).Error
    if err != nil {
        return 0, err
    }
    if recordCount > 0 {
        return 0, errors.New("已存在该记录！")
    }
    err = model.Conn().Model(&Deploy).Create(&Deploy).Error
    if err != nil {
        return 0, err
    }
    return
}

func (p *DeployService) Update(req *request.DeployUpdateRequest) (err error) {
    Deploy := model.K8sDeploy{}
    data := make(map[string]interface{}, 0)
    err = model.Conn().Model(&Deploy).Where("id=?", req.ID).Updates(&data).Error
    if err != nil {
        return err
    }
    return
}

func (p *DeployService) Bat(req *request.DeployBatRequest) (err error) {
    Deploy := model.K8sDeploy{}
    switch req.OP {
    case "del":
        err = model.Conn().Model(&Deploy).Where("id in ? ", req.ID).Delete(Deploy).Error
        break
    }
    if err != nil {
        return err
    }
    return
}
