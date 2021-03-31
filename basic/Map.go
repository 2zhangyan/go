package main

import "fmt"

func capitalMap() {
	var countryCapitalMap map[string]string //创建集合
	countryCapitalMap = make(map[string]string)

	//map 插入key - val对，各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["China"] = "北京"
	countryCapitalMap["Italy"] = "罗马"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	fmt.Println(capital) //false
	fmt.Println(ok)

	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

}

func deleteMap() {
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italt":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New delhi"}

	fmt.Println("原始值")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "India") //删除元素

	fmt.Println("删除后的值")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}

func main() {
	capitalMap()
	deleteMap()
}
