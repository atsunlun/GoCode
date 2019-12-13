package stockUtil

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"../util"
)

const (
	TX_QUERY_STOCK_URL        = "http://qt.gtimg.cn/q=s_"
	TX_QUERY_STOCK_DETAIL_URL = "http://qt.gtimg.cn/q="
)

//获取股票简略消息
func GetStockInfo(arr *[]string) list.List {
	var resList list.List
	for _, code := range *arr {
		str := util.GetHtmlStr(TX_QUERY_STOCK_URL + "sh" + code) //上海
		if !strings.Contains(str, code) {
			str = util.GetHtmlStr(TX_QUERY_STOCK_URL + "sz" + code) //深圳
			if !strings.Contains(str, code) {
				continue
			}
		}
		resList.PushBack(strings.TrimSpace(str)) //去除左或右空格等特殊字符
	}
	return parsingStockInfo(&resList)
}

func parsingStockInfo(inList *list.List) list.List {
	var resList list.List
	//列表头
	kongStr := "\t" //空格间隔
	resList.PushBack("股票名称(代码)" + kongStr + "当前价格" + kongStr + "涨跌额" + kongStr + "涨跌幅" + kongStr + "成交量(万股)" + kongStr + "成交额(万)" + kongStr + "总市值(亿)")
	for i := inList.Front(); i != nil; i = i.Next() {
		oriStr, ok := i.Value.(string)
		if !ok {
			resList.PushBack("parsing Error")
			continue
		}
		var resStr string
		arr := strings.Split(oriStr, "~")
		for j, arrStr := range arr {
			switch j {
			case 0: //无
			case 1: //股票名称
				resStr += arrStr
			case 2: //股票代码
				resStr += "(" + arrStr + ")" + kongStr
			case 3: //当前价格
				resStr += arrStr + kongStr
			case 4: //涨跌
				resStr += arrStr + kongStr
			case 5: //涨跌%
				resStr += arrStr + "%" + kongStr
			case 6: //成交量(万股)
				num, _ := strconv.Atoi(arrStr)
				num /= 100
				resStr += strconv.Itoa(num) + kongStr + kongStr
			case 7: //成交额(万)
				resStr += arrStr + kongStr + kongStr
			case 8: //无
			case 9: //总市值
				resStr += arrStr + kongStr
			default:
			}
		}
		resList.PushBack(resStr)
	}
	return resList
}

//获取股票详情
func GetStockDetailInfo(arr *[]string) list.List {
	var resList list.List
	for _, code := range *arr {
		str := util.GetHtmlStr(TX_QUERY_STOCK_DETAIL_URL + "sh" + code) //上海
		if !strings.Contains(str, code) {
			str = util.GetHtmlStr(TX_QUERY_STOCK_DETAIL_URL + "sz" + code) //深圳
			if !strings.Contains(str, code) {
				continue
			}
		}
		resList.PushBack(strings.TrimSpace(str)) //去除左或右空格等特殊字符
	}
	return parsingStockDetailInfo(&resList)
}

func parsingStockDetailInfo(inList *list.List) list.List {
	var resList list.List
	titleMap := map[int]string{
		// 0:  "未知",
		1: "股票名字",
		// 2:  "股票代码",
		3: "当前价格",
		// 4:  "昨收",
		// 5:  "今开",
		// 6:  "成交量", //成交量（手）
		7: "外盘(万)",
		8: "内盘(万)",
		// 9:  "买一",
		// 10: "买量", //买一量（手）
		// 11: "买二",
		// 12: "买量", //买二量（手）
		// 13: "买三",
		// 14: "买量", //买三量（手）
		// 15: "买四",
		// 16: "买量", //买四量（手）
		// 17: "买五",
		// 18: "买量", //买五量（手）
		// 19: "卖一",
		// 20: "卖量", //卖一量（手）
		// 21: "卖二",
		// 22: "卖量", //卖二量（手）
		// 23: "卖三",
		// 24: "卖量", //卖三量（手）
		// 25: "卖四",
		// 26: "卖量", //卖四量（手）
		// 27: "卖五",
		// 28: "卖量", //卖五量（手）
		// 29: "最近逐笔成交",
		// 30: "时间",
		31: "涨跌",
		32: "涨跌%",
		// 33: "最高",
		// 34: "最低",
		// 35: "价格/成交量（手）/成交额",
		36: "成交量(万股)", //成交量（手）
		37: "成交额(亿)",  //成交额（万）
		38: "换手率",
		39: "市盈率",
		// 40: "未知",
		// 41: "最高",
		// 42: "最低",
		43: "振幅",
		44: "流通市值(亿)",
		45: "总市值(亿)",
		// 46: "市净率",
		// 47: "涨停价",
		// 48: "跌停价",
	}
	resList.PushBack(titleMap)
	for i := inList.Front(); i != nil; i = i.Next() {
		valueMap := make(map[int]string)
		oriStr, ok := i.Value.(string)
		if ok {
			arr := strings.Split(oriStr, "~")
			for j, arrStr := range arr {
				_, ok := titleMap[j]
				if !ok {
					continue //忽略，不打印
				}
				if "" == arrStr {
					arrStr = "null"
				}
				switch j {
				case 7:
					fallthrough
				case 8:
					f := 0.0000
					num, err := strconv.Atoi(arrStr)
					if nil == err {
						f = float64(num) / 10000
					}
					arrStr = strconv.FormatFloat(f, 'f', 2, 64)
				case 44:
					arrStr += "\t"
				case 36:
					num, err := strconv.Atoi(arrStr)
					if nil == err {
						num /= 100
					}
					arrStr = strconv.Itoa(num) + "\t"
				case 37:
					f := 0.0000
					num, err := strconv.Atoi(arrStr)
					if nil == err {
						f = float64(num) / 10000
					}
					arrStr = strconv.FormatFloat(f, 'f', 4, 64) + "\t"
				case 32:
					fallthrough
				case 38:
					fallthrough
				case 39:
					fallthrough
				case 43:
					arrStr += "%"
				}
				valueMap[j] = arrStr
			}
		}
		resList.PushBack(valueMap)
	}
	return resList
}

func PrintStockDetailInfo(inList *list.List) {
	//计算打印的最大键值
	size := 0
	{
		i := inList.Front()
		if i != nil {
			valueMap, ok := i.Value.(map[int]string)
			if !ok {
				fmt.Println("PrintStockDetailInfo Error")
				return
			}
			for k, _ := range valueMap {
				if k >= size {
					size = k
				}
			}
		}
	}
	for i := inList.Front(); i != nil; i = i.Next() {
		valueMap, ok := i.Value.(map[int]string)
		if !ok {
			fmt.Println("PrintStockDetailInfo Error")
			continue
		}
		for j := 0; j <= size; j++ {
			v, ok := valueMap[j]
			if !ok {
				continue
			}
			fmt.Printf("%v\t", v)
		}
		fmt.Print("\n")
	}
	fmt.Print("---------------------------------------------------------------------------------------------------------\n")
}
