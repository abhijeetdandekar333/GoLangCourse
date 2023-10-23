// Key value Pair.
// All the keys should be of same type and all the values should be of same type. 
package main

import (
	"fmt"
)

func main()  {
	// 1st way of eeclaring maps:
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",

	}
	fmt.Println(colors["red"])
	fmt.Println(colors)

	iterateOverMap(colors);

	// 2nd way of declaring a map:
	var colors1 map[string]string;
	fmt.Println(colors1)

	// 3rd way :
	colors2 := make(map[string]string)

	colors2["white"] = "#ffffff"

	delete(colors2, "white")  //Deletes the key white from the map

	fmt.Println(colors2)

}

// Iterating over the maps 
func iterateOverMap(colors map[string]string)  {
	for colorKey, colorValue  := range colors {
		fmt.Println(colorKey+ " has color code of " + colorValue)
	}
}