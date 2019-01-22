package algo

import "fmt"

/// Find Pairs Game
/// Author Duncwinn
/// Date Jan 19
var Count int = 0
var array = [4]int{1, 3, 5, 7}
var bottom, top, tup int = 0, 0, 0

//var mymap map[int]int

/// oN dearch
func FindPairsCompareEverything() {
	/**
	for v := range array {

		for i := 1; i < len(array); i++ {
			Count++

			if v+array[i] == 8 {
				fmt.Println("Matchs on :", v, " + ", i)

			}

		}

	}
	**/

	for i := 0; i < len(array); i++ {

		for j := i + 1; j < len(array); j++ {
			Count++

			if array[j]+array[i] == 8 {
				fmt.Println("Matchs on :", j, " + ", i)
			}

		}

	}

}

func HasPairWithSum(sum int) bool {
	//create an unordered set/array
	length := len(array)
	var comp = make([]int, length)
	Count = 0

	//loop over the array
	for v := range array {
		for cv := range comp {
			Count++
			if v == cv {
				// Found!
				//break
				return true
			}
			comp = append(comp, sum-v)

		}
	}
	return false
}

func MyForLoop() {

	sum := 0

	for sum = 0; sum < 1000; sum++ {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

}
