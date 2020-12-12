package persist

import (
	"fmt"
	"gengycSrc/gycSearchEngine/model"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"strconv"
	"time"
)



type LianjiaES struct {
	index string
	mapping string
	client *elastic.Client
}

const (
	project = "lianjia_engine"
	author = "gengyc"
	mappingTpl = `{
	"mappings":{
		"properties":{
			"id": 				{ "type": "long" },
			"title": 		{ "type": "keyword" },
			"price":			{ "type": "text" },
			"room":			{ "type": "text" },
			"type":				{ "type": "text" },
			"area":		{ "type": "text" },
			"communityName":         { "type": "text" },
			"areaName":         { "type": "text" },
			"visitTime":         { "type": "text" },
			"update_time":		{ "type": "long" },
			"create_time":		{ "type": "long" }
			}
		}
	}`
	esRetryLimit = 3 //bulk 错误重试机制
)

func NewLianjiaES(client *elastic.Client) *LianjiaES {
	index := fmt.Sprintf("%s",project)
	lianjiaES := &LianjiaES{
		client: client,
		index: index,
		mapping: mappingTpl,
	}

	lianjiaES.init()
	return lianjiaES
}

func (es *LianjiaES) init()  {
	ctx := context.Background()

	exists,err := es.client.IndexExists(es.index).Do(ctx)
	if err !=nil{
		fmt.Printf("lianjia init exist failed err is %s\n", err)
	}

	if !exists{
		_,err := es.client.CreateIndex(es.index).Body(es.mapping).Do(ctx)
		if err != nil {
			fmt.Printf("lianjia init failed err is %s\n", err)

			return
		}
	}
}

func (es *LianjiaES) Add(ctx context.Context,profile model.Profile) error  {


	_, err := es.client.Index().
		Index(es.index). // 设置索引名称
		BodyJson(profile). // 指定前面声明的微博内容
		Do(ctx) // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		log.Printf("err happened when we trying to insert one data:%v",err)
		return err
	}
	return nil
}

func (es *LianjiaES) BatchAdd(ctx context.Context,profile []*model.Profile) error  {
	var err error
	for i :=0;i< esRetryLimit;i++{
		if err = es.batchAdd(ctx,profile);err !=nil{
			fmt.Println("batch add failed",err)
			continue
		}
		return err
	}
	return err
}

func (es *LianjiaES)batchAdd(ctx context.Context,profile []*model.Profile) error  {
	req := es.client.Bulk().Index(es.index)
	for _,u := range profile{
		u.UpdateTime = uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
		u.CreateTime = uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
		doc := elastic.NewBulkIndexRequest().Id(strconv.FormatUint(u.ID,10)).Doc(u)
		req.Add(doc)
	}
	if req.NumberOfActions() < 0{
		return nil
	}

	if _,err := req.Do(ctx);err !=nil{
		return err
	}
	return nil

}

func (es *LianjiaES)Search(ctx context.Context,filter *EsSearch) (int64,[]*model.Profile,error)  {

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(filter.MustQuery...)
	boolQuery.MustNot(filter.MustNotQuery...)
	boolQuery.Should(filter.ShouldQuery...)
	boolQuery.Filter(filter.Filters...)

	if len(filter.MustNotQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0{
		boolQuery.MinimumShouldMatch("1")
	}

	service := es.client.Search().Index(es.index).Query(boolQuery).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size)

	resp,err := service.Do(ctx)
	if err !=nil{
		return 0,nil, err
	}

	if resp.TotalHits() == 0{
		return 0,nil,nil
	}

	lianjiaInfo := make([]*model.Profile,0)

	for _,e:= range resp.Each(reflect.TypeOf(&model.Profile{})){
		lianjia := e.(*model.Profile)
		lianjiaInfo = append(lianjiaInfo,lianjia)
	}

	return resp.TotalHits(),lianjiaInfo,nil
}