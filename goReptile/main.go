package main

import (
	"fmt"
	"time"

	"./stockUtil"
	"./util"
)

func main() {
	fmt.Println("我的程序goReptile start...")
	//股票代码
	stockList := []string{"002351", "000673", "300315", "002463", "300201"}
	//获取股票信息
	// var resList list.List
	// resList = stockUtil.GetStockInfo(&stockList)
	// for i := resList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }
	//获取股票详情
	for {
		resList := stockUtil.GetStockDetailInfo(&stockList)
		stockUtil.PrintStockDetailInfo(&resList)
		time.Sleep(5 * time.Second) //延时
	}

	util.Pause()
}
