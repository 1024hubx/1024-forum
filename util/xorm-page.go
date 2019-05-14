package util

var (
	PAGE_NUM int = 10
)

type Page struct {
	Per_page int    `json:"Per_page"  form:"Per_page"` //每页多少条 默认10条
	Page     int    `json:"Page"  form:"Page"`         //当前页数 默认第1页
	Sortby   string `json:"Sortby"  form:"Sortby"`     // 排序字段 默认 updateAt
	Order    string `json:"Order"  form:"Order"`       //排序方式： asc 正序/desc  倒叙(默认)
}

//设置有多少页 参数为 总条数
func SetPage(TotalNum int) int {
	var Total int
	if TotalNum/PAGE_NUM < 1 {
		Total = 1
		return Total
	} else {
		if TotalNum%PAGE_NUM == 0 {
			Total = TotalNum / PAGE_NUM
			return Total
		} else {
			Total = TotalNum/PAGE_NUM + 1
		}
	}
	return Total
}

type Pagination struct {
	Page       int // 当前第几页
	Per_page   int // 每页条数
	Total      int // 总条数
	TotalPages int // 总页数

}

//根据传进来的页数 返回从第n条数据查询
func SetSelectNum(p int) int {
	if p == 1 {
		p = 0
		return p - 1
	}
	//如果是第二页 page =2 从第11条开始查
	p = PAGE_NUM*(p-1) + 1
	return p - 1
}
