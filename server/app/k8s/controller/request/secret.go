package request

import "github.com/lsamu/ago/rest/handler"

//列表
type SecretRequest struct {
    PageSize int `json:"page_size" binding:"gte=0,lte=1000"`
    PageIndex int `json:"page_index" binding:"gte=0"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *SecretRequest) Check() *handler.CustomError {
    if p.PageSize == 0 {
        p.PageSize = 10
    }
    if p.PageIndex == 0 {
        p.PageIndex = 1
    }
    return nil
}

//信息
type SecretInfoRequest struct {
    Id uint32 `json:"id" binding:"gte=1"`
}

func (p *SecretInfoRequest) Check() *handler.CustomError {

    return nil
}

//创建
type SecretCreateRequest struct {
    Name       string    `json:"name" binding:"required"`
    Contact    string    `json:"contact" binding:"required"`
    Status     uint32    `json:"status"`
    Readme     string    `json:"readme"`
}

func (p *SecretCreateRequest) Check() *handler.CustomError {

    return nil
}

//更新
type SecretUpdateRequest struct {
    ID         uint32    `json:"id" binding:"required,gte=1"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *SecretUpdateRequest) Check() *handler.CustomError {

    return nil
}

//批量
type SecretBatRequest struct {
    OP string   `json:"op" binding:"required"`
    ID []uint32 `json:"id" binding:"gt=0,dive,required"`
}

func (p *SecretBatRequest) Check() *handler.CustomError {
    return nil
}
