package handlers

import (
	"gengycSrc/gycSearchEngine/model"
	"gengycSrc/gycSearchEngine/persist"
	"gengycSrc/gycSearchEngine/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LianjiaHandler struct {
	engine *gin.Engine
	service *service.LianjiaService
}

func NewLianjiaHandler(engine *gin.Engine,lianjiaService *service.LianjiaService) *LianjiaHandler  {
	return &LianjiaHandler{
		engine: engine,
		service: lianjiaService,
	}
}

func (h *LianjiaHandler) Run() {
	gin.ForceConsoleColor()
	h.engine.Use(gin.Logger())
	h.engine.Use(gin.Recovery())
	h.registerRouter()

	err := h.engine.Run()
	if err !=nil{
		log.Fatalln("server start failed")
	}
}


func (h *LianjiaHandler) registerRouter()  {
	lianjia := h.engine.Group("api/houseInfo")
	{
		lianjia.POST("/batchAdd",h.Create)
		lianjia.POST("/add",h.Add)
		lianjia.POST("/search", h.Search)
	}
}

func (h *LianjiaHandler)Create(c *gin.Context)  {
	var err error
	houseInfoList := make([]*model.Profile,0)
	houseInfo := model.Profile{}

	if err = c.ShouldBindJSON(&houseInfo);err !=nil{
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}

	houseInfoList = append(houseInfoList,&houseInfo)
	if err = h.service.BatchAdd(c,houseInfoList);err !=nil{
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success"})

}

func (h *LianjiaHandler)Add(c *gin.Context)  {
	var err error
	houseInfo := model.Profile{}
	if err = c.ShouldBindJSON(&houseInfo);err !=nil{
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}

	if err = h.service.Add(c,houseInfo);err !=nil{
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":0,"msg":"success"})

}
func (h *LianjiaHandler)Search(c *gin.Context)  {
	var err error
	req := persist.SearchRequest{}
	err = c.ShouldBindJSON(&req)
	if err !=nil{
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}

	total,res,e := h.service.Search(c,&req)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": e.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
		"total": total,
	})
}