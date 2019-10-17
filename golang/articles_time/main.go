package articles_time

import (
	"fmt"
	"time"
)

//--------------------------------------------------------
// ローカルタイムと指定した時刻を扱う方法(Time.now)
//--------------------------------------------------------
func inInTimeNow() {
	// time.Now()
	// 何も指定していない場合システムの時刻が返ります（私の場合は日本時刻が返ってきています
	t1 := time.Now()
	fmt.Println(t1)
	// => 2019-10-16 21:52:41.994421 +0900 JST m=+0.000315501

	// UTCの時刻を取得したい場合
	t2 := t1.UTC()
	fmt.Println(t2)
	// => 2019-10-16 13:16:49.428464 +0000 UTC

	// 自分で任意の国の時間にしたい場合
	// time.FixedZone("タイムゾーンに付ける名前", UTCとの差分の秒数)
	pt := time.FixedZone("PT", -8*60*60)
	fmt.Println(time.Now().In(pt))
	// => 2019-10-16 05:32:34.011723 -0800 PT
}

//--------------------------------------------------------
// 時刻から曜日を取得する方法(time.Weekday)
//--------------------------------------------------------

func inInWeekday() {
	// time.Weekday()
	// 曜日の数値を返します

	// 日曜 0
	// 月曜 1
	// 火曜 2
	// 水曜 3
	// 木曜 4
	// 金曜 5
	// 土曜 6

	t1 := time.Now()
	fmt.Println(t1.Weekday())
	// => Wednesday

	// 日本語にしたい場合
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	jpWeekday := fmt.Sprintf("%s曜日", wdays[t1.Weekday()])
	fmt.Println(jpWeekday)
	// => 水曜日
}

//--------------------------------------------------------
// 時刻から秒、分、時、日、月、年をそれぞれ取得するやり方
//--------------------------------------------------------

func inInSecMinEtc() {
	t := time.Now()
	fmt.Println(t)
	// => 2019-10-16 23:06:40.569593 +0900 JST m=+0.000338457

	// 秒
	fmt.Println(t.Second())
	// => 40

	// 分
	fmt.Println(t.Minute())
	// => 6

	// 時
	fmt.Println(t.Hour())
	// => 23

	// 日
	fmt.Println(t.Day())
	// => 6

	// 月
	fmt.Println(t.Month())
	// => October

	// 年
	fmt.Println(t.Year())
	// => 2019
}

//--------------------------------------------------------
// 指定した時刻を文字列から作成する(time.ParseInLocation)
//--------------------------------------------------------

func inInTimeParseInLocation() {
	// ParseInLocation(layout, value string, loc *Location) (Time, error)
	// 指定した時刻を文字列から作成するにはParseもしくはParseInLocationを使います
	// Parseの場合ローカルタイムが適用されるので指定したロケーションで時刻が取得できるParseInLocationを使います

	// time.Formatと同様の使い方をします

	// Formatに与えられる引数は既に決められていて

	// 秒 05
	// 分 04
	// 時 15
	// 日 02
	// 月 01
	// 年 2006

	jst, _ := time.LoadLocation("Asia/Tokyo")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-10-17 00:12:15", jst)
	fmt.Println(t1)
	// => 2019-10-17 00:12:15 +0900 JST
}

//--------------------------------------------------------
// 指定した時刻を作成する(time.Date)
//--------------------------------------------------------

func inInTimeDate() {
	// Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location)

	t := time.Date(2019, 10, 16, 23, 6, 40, 0, time.UTC)
	fmt.Println(t)
	// => 2019-10-16 23:06:40 +0000 UTC
}

//--------------------------------------------------------
// 時刻オブジェクトから文字列を整形する(time.Format)
//--------------------------------------------------------

func inInTimeFormat() {
	// Formatに与えられる引数は既に決められていて

	// 秒 05
	// 分 04
	// 時 15
	// 日 02
	// 月 01
	// 年 2006
	// これをFormatに与えていく

	t := time.Now()
	fmt.Println(t.Format("2006年01月02日 15時04分05秒"))
	// => 2019年10月16日 23時51分15秒
}

//--------------------------------------------------------
// 時刻の比較( time.Before time.After)
//--------------------------------------------------------

func inInTimeComparition() {
	// 16日と17日で比較します
	t16 := time.Date(2019, 10, 16, 23, 6, 40, 0, time.UTC)
	t17:= time.Date(2019, 10, 17, 23, 6, 40, 0, time.UTC)

	// 17日は16日の前なのでtrueが返ります
	fmt.Println(t16.Before(t17))
	// => true

	// 16日は17日の後なのでtrueが返ります
	fmt.Println(t17.After(t16))
	// => true
}

//--------------------------------------------------------
// 現在の日付を文字列で取得する
//--------------------------------------------------------

func inInDate() {
	// そんな関数存在しないので自分で作る
	day := time.Now()
	fmt.Println(day.Format("2006-01-02"))
	// => 2019-10-17
}

//--------------------------------------------------------
// 昨日、○日前、○週間前、○月前、明日、○週間後、○日後、○ヶ月後などを取得する(time.AddDate)
//--------------------------------------------------------

func inInTimeAddDate() {
	// 前系

	// time.AddDate(years, months, days)
	// -1を渡すと前の日にも戻れる

	// 昨日
	t1 := time.Date(2019, 12, 1, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t1.AddDate(0, 0, -1))
	// ちゃんと月をまたげる
	// => 2019-11-30 23:59:59.999999999 +0000 UTC

	// 10日前
	t2 := time.Date(2019, 12, 1, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t2.AddDate(0, 0, -10))
	// => 2019-11-21 23:59:59.999999999 +0000 UTC

	// 1週間前
	t3 := time.Date(2019, 12, 1, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t3.AddDate(0, 0, -7))
	// => 2019-11-24 23:59:59.999999999 +0000 UTC

	// 1ヶ月前
	t4 := time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t4.AddDate(0, -1, 0))
	// => 2019-12-01 23:59:59.999999999 +0000 UTC
	// 30日を引いてるだけに見える

	// 後系

	// 明日
	t5 := time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t5.AddDate(0, 0, 1))
	// => 2020-01-01 23:59:59.999999999 +0000 UTC

	// 10日後
	t6 := time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t6.AddDate(0, 0, 10))
	// => 2020-01-10 23:59:59.999999999 +0000 UTC

	// 1週間後
	t7 := time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t7.AddDate(0,0,7))
	// => 2020-01-07 23:59:59.999999999 +0000 UTC

	// 1ヶ月後
	t8 := time.Date(2019, 12, 31, 23, 59, 59, 999999999, time.UTC)
	fmt.Println(t8.AddDate(0,1,0))
	// => 2020-01-31 23:59:59.999999999 +0000 UTC
}

//--------------------------------------------------------
// 月末月初の時刻を取得する
//--------------------------------------------------------

func inInFirstOrEndMonth() {
	// 月初の取得は簡単 Dateで求めてあげればいいです
	t1 := time.Now()
	t2 := time.Date(t1.Year(), t1.Month(), 1, 0, 0, 0,0, time.UTC)
	fmt.Println(t2)
	// => 2019-10-01 00:00:00 +0000 UTC

	// 月末は月初から-1日すれば求まります
	t3 := t2.AddDate(0, 0, -1)
	fmt.Println(t3)
	// => 2019-09-30 00:00:00 +0000 UTC
}

func InInTime() {
	inInTimeNow()
	inInWeekday()
	inInSecMinEtc()
	inInTimeParseInLocation()
	inInTimeDate()
	inInTimeFormat()
	inInTimeComparition()
	inInDate()
	inInTimeAddDate()
	inInFirstOrEndMonth()
}