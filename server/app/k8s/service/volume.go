package service

import (
    "errors"
    "sass_devops/app/k8s/controller/request"
    "sass_devops/app/k8s/model"
)

var VolumeHandler VolumeService

func init() {
    VolumeHandler = VolumeService{}
}

type VolumeService struct {
}

func (p *VolumeService) List(req *request.VolumeRequest) (list []model.K8sVolume, total int64, err error) {
    startIndex := req.PageSize * (req.PageIndex - 1)
    pageSize := req.PageSize
    Volume := model.K8sVolume{}
    db := model.Conn().Model(&Volume).Order("id desc")
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

func (p *VolumeService) Info(req *request.VolumeInfoRequest) (info model.K8sVolume, err error) {
    Volume := model.K8sVolume{}
    err = model.Conn().Model(&Volume).Find(&info).Error
    if err != nil {
        return info, err
    }
    return
}

func (p *VolumeService) Create(req *request.VolumeCreateRequest) (lastId uint32, err error) {
    Volume := model.K8sVolume{}
    var recordCount int64
    err = model.Conn().Model(&Volume).Where("name=?", req.Name).Count(&recordCount).Error
    if err != nil {
        return 0, err
    }
    if recordCount > 0 {
        return 0, errors.New("已存在该记录！")
    }
    err = model.Conn().Model(&Volume).Create(&Volume).Error
    if err != nil {
        return 0, err
    }
    return
}

func (p *VolumeService) Update(req *request.VolumeUpdateRequest) (err error) {
    Volume := model.K8sVolume{}
    data := make(map[string]interface{}, 0)
    err = model.Conn().Model(&Volume).Where("id=?", req.ID).Updates(&data).Error
    if err != nil {
        return err
    }
    return
}

func (p *VolumeService) Bat(req *request.VolumeBatRequest) (err error) {
    Volume := model.K8sVolume{}
    switch req.OP {
    case "del":
        err = model.Conn().Model(&Volume).Where("id in ? ", req.ID).Delete(Volume).Error
        break
    }
    if err != nil {
        return err
    }
    return
}
