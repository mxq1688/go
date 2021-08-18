package main

import "fmt"

func main() {
	//map  如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	var mMap map[string]string //定义空map  不能存放键值对
	var nMap = make(map[string]string)//创建map并初始化
	nMap ["title"] = "ajeif"
	nMap ["content"] = "faeoifjqje"
	nMap ["name"] = "mxq"
	fmt.Println(mMap,nMap)

	for key,value := range nMap {
		fmt.Println(key, value)
	}

	//判断map中元素是否存在
	mkey, mvalue := nMap["uid"]
	fmt.Println(mkey, mvalue)

	map delete()
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图", countryCapitalMap)
	for key, value := range countryCapitalMap {
		fmt.Println(key, value, countryCapitalMap[key])
	}
	//fmt.Println("\n")
	delete(countryCapitalMap, "France")
	for key, value := range countryCapitalMap {
		fmt.Println(key, value, countryCapitalMap[key])
	}
}
