package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func main() {
	map1 := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	map2 := map[string]int{
		"b": 25,
		"d": 40,
		"e": 50,
	}

	mergedMap := mergeMaps(map1, map2)

	fmt.Println("Mapa 1:", map1)
	fmt.Println("Mapa 2:", map2)
	fmt.Println("Mapa Mesclado:", mergedMap)
}
