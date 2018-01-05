# Punch Clock

golang 練習題目，若有錯誤麻煩多多指教。

題目:

    假設每位訪客進入(Arrival)與離開(Departure)系統都會留下時間紀錄，
    試寫一function 可從計算指定時間區間內的最大同時上線人數

## 思維

1. timestamp 為1970年1月1日0時0分0秒起至現在的總秒數
2. 人數變化只會在有人離開、或進入時會改變
3. 檢查時間區間內的所有人數變化點，即可知道最大線上人數

## 筆記

#### ▶ 排序陣列元素

官方的sort包沒有每一種變數型態的排序方法，但有提供介面可以實作，例如以下方法實作int64陣列的排序

```go
import "sort"

// int64 實作int64型態的排序方法
type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

sortMapKey := []int64{5,8,1,6,898,457,2}

sort.Sort(int64arr(sortMapKey))
```
