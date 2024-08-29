// 자료형 : 맵(1)
package main

import "fmt"

func main() {
	//맵(Map)
	//Key, Value 쌍으로 자료 저장하는 자료구조로 key를 통해 식별하고 순서가 정해져있지 않음

	//map의 생성방법
	var map1 map[int]string = make(map[int]string) //정석
	map2 := make(map[int]string)                   //축약형

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println()

	//map에 데이터를 저장하는 방법
	map1[10] = "최지원"
	map1[15] = "김수민"
	map1[17] = "박수혁"

	fmt.Println(map1)
	fmt.Println()

	//map 데이터를 사용하는 방법

	fmt.Println("map1[10] : ", map1[10])
	fmt.Println("map1[17] : ", map1[17])

	//생성과 동시에 초기화 가능
	map3 := map[string]int{
		"sport": 35,
		"ani":   40,
		"game":  52,
	}
	fmt.Println(map3)

	//용량을 정해서 사용할 수 있다
	map4 := make(map[string]int, 2)
	map4["sport"] = 5
	map4["ani"] = 10
	map4["game"] = 15

	fmt.Println(map4)

	//key로 식별하기때문에  동일한 key에 값 대입시 덮어씌워짐
	map4["sport"] = 20
	fmt.Println(map4)

	//map에 데이터를 삭제할때는delete함수
	delete(map4, "sport")
	fmt.Println(map4)

	//map은 리턴값이 2개  -> 값, 키존재유무
	value, isKey := map4["sport"]
	fmt.Println(value, isKey)

	//맵의 모든값에 순차적 접근가능 -> 순서는 랜덤
	for k, v := range map4 {
		fmt.Println(k, v)
	}
}
