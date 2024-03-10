package handler

import (
	"OCsemmerApp/pkg/domain"
	"OCsemmerApp/pkg/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// GET /data
func datahandler(ctx *gin.Context) {
	ctx.HTML(200, "data.html", nil)
}

// GET /search
func searchhandler(ctx *gin.Context) {
	prodname := ctx.Query("search")
	if prodname == "" {
		log.Println("製品名空文字")
	}
	fmt.Println("検索製品名", prodname)
	data, err := model.SearchData(prodname)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	ctx.JSON(200, data)
}

// GET /test
func testsearchhandler(ctx *gin.Context) {
	prodname := ctx.Query("search")
	if prodname == "" {
		log.Println("not prodname request")
		ctx.JSON(400, "Bad Request")
		return
	}
	resp := domain.DataResponse{
		{
			ProductID:    "1",
			ProductName:  "パイソン掃除機Mk5",
			UserName:     "John Yamashita",
			PurchaseDate: "2021-03-02",
		},
		{
			ProductID:    "2",
			ProductName:  "DELL PowerEdge",
			UserName:     "Yuta kiguchi",
			PurchaseDate: "2021-04-21",
		},
		{
			ProductID:    "3",
			ProductName:  "超電磁砲",
			UserName:     "Syu Kim",
			PurchaseDate: "2021-12-15",
		},
		{
			ProductID:    "4",
			ProductName:  "ガンダニウム合金",
			UserName:     "Kota Sagawa",
			PurchaseDate: "2022-03-03",
		},
		{
			ProductID:    "5",
			ProductName:  "ウラニウム",
			UserName:     "Yuki Kitayama",
			PurchaseDate: "20XX-09-08",
		},
	}
	ctx.JSON(200, resp)
}
