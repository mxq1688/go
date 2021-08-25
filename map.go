package main

import "fmt"

// Map 类型是引用类型
/*
	Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
	申明map: 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
		var map_variable map[key_type]value_type
	make():
		初始化map：
			map_variable := make(map[key_data_type]value_data_type)
		初始化并赋值：
			map_variable := make(map[key_data_type]value_data_type){}
*/

func main() {
	//申明map，nil map 不能用来存放键值对
	var mMap map[string]string

	//创建map并初始化
	var nMap = make(map[string]string)
	mMap = nMap //Map 类型是引用类型
	nMap["title"] = "title"
	nMap["name"] = "mxq"
	fmt.Println(mMap, nMap)

	//使用for迭代map
	for key, value := range nMap {
		fmt.Println(key, value)
	}

	//判断map中元素是否存在
	key1, value1 := nMap["uid"]
	if value1 {
		fmt.Println("存在", key1)
	} else {
		fmt.Println("不存在", key1)
	}

	// delete()
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图", countryCapitalMap)
	delete(countryCapitalMap, "France")
	fmt.Println("删除后", countryCapitalMap)

}
