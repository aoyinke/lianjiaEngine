package service

import (
	"gengycSrc/gycSearchEngine/model"
	"gengycSrc/gycSearchEngine/persist"
	"golang.org/x/net/context"
)

type LianjiaService struct {
	es *persist.LianjiaES
}

func NewLianjiaService(es *persist.LianjiaES) *LianjiaService  {
	return &LianjiaService{
		es: es,
	}
}

func (l *LianjiaService) Add(ctx context.Context,profile model.Profile) error  {
	return l.es.Add(ctx,profile)
}
func (l *LianjiaService)BatchAdd(ctx context.Context,lianjia []*model.Profile)error  {
	return l.es.BatchAdd(ctx,lianjia)
}


func (l *LianjiaService) Search(ctx context.Context, req *persist.SearchRequest) (int64,[]*model.Profile, error) {
	return l.es.Search(ctx, req.ToFilter())
}