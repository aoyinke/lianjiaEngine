package model

type Profile struct {
	Price string
	Title string
	ID         uint64 `json:"id,omitempty" mapstructure:"id"`
	UpdateTime uint64 `json:"update_time,omitempty" mapstructure:"update_time"`
	CreateTime uint64 `json:"create_time,omitempty" mapstructure:"create_time"`
	HouseInfo
	AroundInfo
}

type HouseInfo struct {
	Room string
	Type string
	Area string
}

type AroundInfo struct {
	CommunityName string
	AreaName string
	VisitTime string
}