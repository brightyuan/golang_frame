package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"

	/*使用键输出地图值 */
	for k, v := range countryCapitalMap {
		fmt.Println(k, "首都是", v)
	}

	var p = map[string]string{"aa": "bb", "cc": "dd"}
	fmt.Println(p["aa"])

}
