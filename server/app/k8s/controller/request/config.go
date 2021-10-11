package request

import "github.com/lsamu/ago/rest/handler"

//列表
type ConfigRequest struct {
    PageSize int `json:"page_size" binding:"gte=0,lte=1000"`
    PageIndex int `json:"page_index" binding:"gte=0"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *ConfigRequest) Check() *handler.CustomError {
    if p.PageSize == 0 {
        p.PageSize = 10
    }
    if p.PageIndex == 0 {
        p.PageIndex = 1
    }
    return nil
}

//信息
type ConfigInfoRequest struct {
    Id uint32 `json:"id" binding:"gte=1"`
}

func (p *ConfigInfoRequest) Check() *handler.CustomError {

    return nil
}

//创建
type ConfigCreateRequest struct {
    Name       string    `json:"name" binding:"required"`
    Contact    string    `json:"contact" binding:"required"`
    Status     uint32    `json:"status"`
    Readme     string    `json:"readme"`
}

func (p *ConfigCreateRequest) Check() *handler.CustomError {

    return nil
}

//更新
type ConfigUpdateRequest struct {
    ID         uint32    `json:"id" binding:"required,gte=1"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *ConfigUpdateRequest) Check() *handler.CustomError {

    return nil
}

//批量
type ConfigBatRequest struct {
    OP string   `json:"op" binding:"required"`
    ID []uint32 `json:"id" binding:"gt=0,dive,required"`
}

func (p *ConfigBatRequest) Check() *handler.CustomError {
    return nil
}
