package response

type BasePage struct {
    List      interface{} `json:"list"`
    Total     int64         `json:"total"`
    PageSize  int         `json:"page_size"`
    PageIndex int         `json:"page_index"`
}

