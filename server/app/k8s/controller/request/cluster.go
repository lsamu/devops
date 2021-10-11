package request

import "github.com/lsamu/ago/rest/handler"

//列表
type ClusterRequest struct {
    PageSize int `json:"page_size" binding:"gte=0,lte=1000"`
    PageIndex int `json:"page_index" binding:"gte=0"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *ClusterRequest) Check() *handler.CustomError {
    if p.PageSize == 0 {
        p.PageSize = 10
    }
    if p.PageIndex == 0 {
        p.PageIndex = 1
    }
    return nil
}

//信息
type ClusterInfoRequest struct {
    Id uint32 `json:"id" binding:"gte=1"`
}

func (p *ClusterInfoRequest) Check() *handler.CustomError {

    return nil
}

//创建
type ClusterCreateRequest struct {
    Name       string    `json:"name" binding:"required"`
    Contact    string    `json:"contact" binding:"required"`
    Status     uint32    `json:"status"`
    Readme     string    `json:"readme"`
}

func (p *ClusterCreateRequest) Check() *handler.CustomError {

    return nil
}

//更新
type ClusterUpdateRequest struct {
    ID         uint32    `json:"id" binding:"required,gte=1"`
    Name       *string    `json:"name" binding:"omitempty"`
    Contact    *string    `json:"contact" binding:"omitempty"`
    Status     *uint32    `json:"status" binding:"omitempty"`
    Readme     *string    `json:"readme" binding:"omitempty"`
}

func (p *ClusterUpdateRequest) Check() *handler.CustomError {

    return nil
}

//批量
type ClusterBatRequest struct {
    OP string   `json:"op" binding:"required"`
    ID []uint32 `json:"id" binding:"gt=0,dive,required"`
}

func (p *ClusterBatRequest) Check() *handler.CustomError {
    return nil
}