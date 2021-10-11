package service

import (
    "errors"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/model"
)

var SecretHandler SecretService

func init() {
    SecretHandler = SecretService{}
}

type SecretService struct {
}

func (p *SecretService) List(req *request.SecretRequest) (list []model.K8sSecret, total int64, err error) {
    startIndex := req.PageSize * (req.PageIndex - 1)
    pageSize := req.PageSize
    Secret := model.K8sSecret{}
    db := model.Conn().Model(&Secret).Order("id desc")
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

func (p *SecretService) Info(req *request.SecretInfoRequest) (info model.K8sSecret, err error) {
    Secret := model.K8sSecret{}
    err = model.Conn().Model(&Secret).Find(&info).Error
    if err != nil {
        return info, err
    }
    return
}

func (p *SecretService) Create(req *request.SecretCreateRequest) (lastId uint32, err error) {
    Secret := model.K8sSecret{}
    var recordCount int64
    err = model.Conn().Model(&Secret).Where("name=?", req.Name).Count(&recordCount).Error
    if err != nil {
        return 0, err
    }
    if recordCount > 0 {
        return 0, errors.New("已存在该记录！")
    }
    err = model.Conn().Model(&Secret).Create(&Secret).Error
    if err != nil {
        return 0, err
    }
    return
}

func (p *SecretService) Update(req *request.SecretUpdateRequest) (err error) {
    Secret := model.K8sSecret{}
    data := make(map[string]interface{}, 0)
    err = model.Conn().Model(&Secret).Where("id=?", req.ID).Updates(&data).Error
    if err != nil {
        return err
    }
    return
}

func (p *SecretService) Bat(req *request.SecretBatRequest) (err error) {
    Secret := model.K8sSecret{}
    switch req.OP {
    case "del":
        err = model.Conn().Model(&Secret).Where("id in ? ", req.ID).Delete(Secret).Error
        break
    }
    if err != nil {
        return err
    }
    return
}
