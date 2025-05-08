package thinfp

// --- ジェネリクスを使った一般的な関数 ---

// Map はスライスの各要素に関数を適用し、新しいスライスを生成します。
// T: 入力スライスの要素型
// U: 出力スライスの要素型
func Map[T, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v) // T型のvを関数fでU型に変換
	}
	return result
}

// Filter はスライスの要素に述語関数を適用し、条件を満たす要素のみで新しいスライスを生成します。
// T: スライスの要素型
func Filter[T any](list []T, predicate func(T) bool) []T {
	result := make([]T, 0) // 容量ゼロで開始
	for _, v := range list {
		if predicate(v) { // 述語がtrueを返したら
			result = append(result, v) // 結果に追加
		}
	}
	return result
}

// Reduce はスライスの要素を結合関数を使って単一の値に畳み込みます（または Fold）。
// T: 入力スライスの要素型
// U: 結果の型 (アキュムレータの型)
func Reduce[T, U any](list []T, initial U, accumulator func(U, T) U) U {
	result := initial // 初期値から開始
	for _, v := range list {
		result = accumulator(result, v) // 現在の結果と要素を結合して、新しい結果を生成
	}
	return result
}
