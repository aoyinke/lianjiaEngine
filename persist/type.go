package persist

import "github.com/olivere/elastic/v7"

type Item struct {
	Id string
	Type string
	Payload interface{}
	Url string
}

type SearchRequest struct {
	Title  string `json:"title"`
	Price     string `json:"price"`
	Num       int    `json:"num"`
	Size      int    `json:"size"`
}

//bool query 条件
type EsSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	From         int //分页
	Size         int
}

func (r *SearchRequest) ToFilter() *EsSearch {
	var search EsSearch
	if len(r.Title) !=0{
		search.ShouldQuery = append(search.ShouldQuery,elastic.NewMatchQuery("Title",r.Title))

	}
	if len(r.Price) !=0{
		search.ShouldQuery = append(search.ShouldQuery,elastic.NewMatchQuery("Price",r.Price))
	}

	search.From = (r.Num - 1) * r.Size
	search.Size = r.Size
	return &search
}