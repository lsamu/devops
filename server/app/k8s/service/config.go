package service

import (
    "errors"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/model"
)

var ConfigHandler ConfigService

func init() {
    ConfigHandler = ConfigService{}
}

type ConfigService struct {
}

func (p *ConfigService) List(req *request.ConfigRequest) (list []model.K8sConfig, total int64, err error) {
    startIndex := req.PageSize * (req.PageIndex - 1)
    pageSize := req.PageSize
    Config := model.K8sConfig{}
    db := model.Conn().Model(&Config).Order("id desc")
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

func (p *ConfigService) Info(req *request.ConfigInfoRequest) (info model.K8sConfig, err error) {
    Config := model.K8sConfig{}
    err = model.Conn().Model(&Config).Find(&info).Error
    if err != nil {
        return info, err
    }
    return
}

func (p *ConfigService) Create(req *request.ConfigCreateRequest) (lastId uint32, err error) {
    Config := model.K8sConfig{}
    var recordCount int64
    err = model.Conn().Model(&Config).Where("name=?", req.Name).Count(&recordCount).Error
    if err != nil {
        return 0, err
    }
    if recordCount > 0 {
        return 0, errors.New("已存在该记录！")
    }
    err = model.Conn().Model(&Config).Create(&Config).Error
    if err != nil {
        return 0, err
    }
    return
}

func (p *ConfigService) Update(req *request.ConfigUpdateRequest) (err error) {
    Config := model.K8sConfig{}
    data := make(map[string]interface{}, 0)
    err = model.Conn().Model(&Config).Where("id=?", req.ID).Updates(&data).Error
    if err != nil {
        return err
    }
    return
}

func (p *ConfigService) Bat(req *request.ConfigBatRequest) (err error) {
    Config := model.K8sConfig{}
    switch req.OP {
    case "del":
        err = model.Conn().Model(&Config).Where("id in ? ", req.ID).Delete(Config).Error
        break
    }
    if err != nil {
        return err
    }
    return
}
