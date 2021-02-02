package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	dao "github.com/jonanata/tempoil2/dao"
	dto "github.com/jonanata/tempoil2/dto"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/getData", func(c *gin.Context) {

		dataMt := dao.GetMTData()

		dataOp := dao.GetOPData()

		fmt.Println("dataMt: ", len(dataMt.MarstempNodes))
		fmt.Println("dataOp: ", len(dataOp.OilpriceNodes))

		var tData []dto.Tabledata

		for i := 0; i < len(dataMt.MarstempNodes); i++ {

			d := new(dto.Tabledata)

			d.Date = dataMt.MarstempNodes[i].Date
			d.Degree = dataMt.MarstempNodes[i].Degree
			d.Price = dataOp.OilpriceNodes[i].Price

			tData = append(tData, *d)
		}

		//check last 7 days data , if mars temperature and oil price rise togather more than 3 days ,
		//next mars temperature rise , oil price will rise
		var s int = 0

		var maxi int = len(dataMt.MarstempNodes) - 7

		if maxi < 0 {

			maxi = 0
		}

		for i := len(dataMt.MarstempNodes) - 1; i > 0; i-- {

			if dataMt.MarstempNodes[i].Degree > dataMt.MarstempNodes[i-1].Degree {

				if (dataOp.OilpriceNodes[i].Price - dataOp.OilpriceNodes[i-1].Price) > 0 {

					s++

				}
			}

		}

		var opr string = "fall"

		if s > 3 {

			opr = "rise"
		}

		c.JSON(200, gin.H{
			"message": "pong",
			"tData":   tData,
			"opr":     opr,
		})
	})

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
