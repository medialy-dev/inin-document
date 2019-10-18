package article_slice

import "fmt"

//--------------------------------------------------------
// 配列の作り方
//--------------------------------------------------------

func inInSliceNew() {
	// 空配列の作り方
	// []型{}
	// ここでは文字列(string)の配列を扱います
	hoge1 := []string{}
	fmt.Println(hoge1)
	// => []

	// 初期値こみの配列の作り方
	hoge2 := []string{"a", "i", "u", "e", "o"}
	fmt.Println(hoge2)
	// => [a i u e o]

	// 配列の大きさがわかっている場合は makeを使います
	// make(型, サイズ)
	hoge3 := make([]string, 5)
	fmt.Println(hoge3)

	// 当然ながら範囲外を参照するとエラーで落ちます
	// fmt.Println(hoge3[5])
	// panic: runtime error: index out of range
}

//--------------------------------------------------------
// 配列への追加
//--------------------------------------------------------

func inInSliceAdd() {
	// append(追加元, 追加内容)
	hoge := []string{}
	fmt.Println(hoge)

	hoge = append(hoge, "a")
	fmt.Println(hoge)
}

//--------------------------------------------------------
// 配列(slice)から指定位置の要素の取り出し方
//--------------------------------------------------------

func inInSliceGetByIndex() {
	hoge := []string{"a", "i", "u", "e", "o"}

	// 配列(slice)の要素は0から開始
	// hoge[0] が aに
	// hoge[4] が oに概要する

	fmt.Println(hoge[0])
	// => a
	fmt.Println(hoge[4])
	// => o

	// 順番に抜き出す場合for rangeを使うのが楽
	for _, value := range hoge {
		fmt.Println(value)
		// => a
		//    i
		//    u
		//    e
		//    o
	}
}

//--------------------------------------------------------
// 配列(slice)から要素を削除する
//--------------------------------------------------------

func inInSliceDelete() {
	// ないので自分で作る
	hoge := []string{"a","i","u","e","o"}
	fmt.Println(hoge)
	// => [a i u e o]
	hoge = deleteByIndex(hoge, 2)
	// => [a i e o]
	fmt.Println(hoge)
}

func deleteByIndex(s []string, index int) []string{
	if index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

//--------------------------------------------------------
// 配列(slice)から長さを求める
//--------------------------------------------------------

func inInSliceLength() {
	hoge := []string{"a","i","u","e","o"}
	fmt.Println(len(hoge))
	// => 5
}

func InInSlice() {
	inInSliceNew()
	inInSliceAdd()
	inInSliceGetByIndex()
	inInSliceDelete()
	inInSliceLength()
}