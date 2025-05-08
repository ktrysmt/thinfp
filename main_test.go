package thinfp

import (
	"fmt"
	"strconv"
)

func main() {
	// --- Map の使用例 ---
	fmt.Println("--- Map Examples ---")
	nums := []int{1, 2, 3, 4, 5}

	// int スライスを string スライスに変換
	strings := Map(nums, func(i int) string {
		return fmt.Sprintf("Num_%d", i)
	})
	fmt.Println("Map int to string:", strings) // Output: [Num_1 Num_2 Num_3 Num_4 Num_5]

	// int スライスの各要素を2倍にする
	doubledNums := Map(nums, func(i int) int {
		return i * 2
	})
	fmt.Println("Map int to int (double):", doubledNums) // Output: [2 4 6 8 10]

	// string スライスの長さを計算する
	words := []string{"apple", "banana", "cherry"}
	lengths := Map(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Map string to int (length):", lengths) // Output: [5 6 6]

	// --- Filter の使用例 ---
	fmt.Println("\n--- Filter Examples ---")

	// int スライスから偶数だけをフィルタリング
	evenNums := Filter(nums, func(i int) bool {
		return i%2 == 0 // 偶数ならtrue
	})
	fmt.Println("Filter even numbers:", evenNums) // Output: [2 4]

	// string スライスから長さが5以上の単語をフィルタリング
	longWords := Filter(words, func(s string) bool {
		return len(s) >= 5 // 長さが5以上ならtrue
	})
	fmt.Println("Filter long words:", longWords) // Output: [apple banana cherry]

	// --- Reduce の使用例 ---
	fmt.Println("\n--- Reduce Examples ---")

	// int スライスの合計を計算 (intをintにReduce)
	sum := Reduce(nums, 0, func(acc int, val int) int {
		return acc + val // アキュムレータに現在の要素を加算
	})
	fmt.Println("Reduce sum of numbers:", sum) // Output: 15 (1+2+3+4+5)

	// int スライスをカンマ区切りの string に変換 (intをstringにReduce)
	commaSeparatedNums := Reduce(nums, "", func(acc string, val int) string {
		if acc == "" {
			return strconv.Itoa(val) // 最初の要素
		}
		return acc + "," + strconv.Itoa(val) // カンマを追加して結合
	})
	fmt.Println("Reduce numbers to comma-separated string:", commaSeparatedNums) // Output: 1,2,3,4,5

	// string スライスをスペース区切りの単一の string に結合 (stringをstringにReduce)
	sentence := Reduce(words, "", func(acc string, val string) string {
		if acc == "" {
			return val // 最初の要素
		}
		return acc + " " + val // スペースを追加して結合
	})
	fmt.Println("Reduce words to sentence:", sentence) // Output: apple banana cherry
}
