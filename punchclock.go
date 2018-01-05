package punchclock

import (
	"sort"
	"time"
)

// Record : 簽到記錄
//
// 到達時間應小於離開時間 (Arrival < Departure)
type Record struct {
	// 到達時間
	Arrival time.Time

	// 離開時間
	Departure time.Time
}

// int64 實作int64型態的排序方法
type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

// GetMaxNumDuringPeriod : 給定一群簽到記錄，與一段時間，取得該時間內最大同時人數(ccu)
//
// StartTime : 指定起始時間
// EndTime : 指定結束時間
// PunchRecords : 簽到記錄
func GetMaxNumDuringPeriod(StartTime, EndTime time.Time, PunchRecords []Record) int {
	// countMap : 用於紀錄時間區間內的人數變化
	//
	// key: time.Duraiion 有人數變化的時間點(有人離開或進入)
	// value: count 該時間點的變化數量(可能增加或減少)
	countMap := make(map[int64]int)

	for _, record := range PunchRecords {
		// 結束時間後才進入或在區間開始前已離開,不需記錄人數變化
		if record.Arrival.After(EndTime) ||
			record.Departure.Before(StartTime) ||
			record.Departure.Equal(StartTime) {
			continue
		}

		// 在區間開始時間之前已進入
		if record.Arrival.Equal(StartTime) || record.Arrival.Before(StartTime) {
			countMap[StartTime.Unix()]++

			// 在開始時間後~結束時間前進入
		} else if record.Arrival.Equal(EndTime) || record.Arrival.Before(EndTime) {
			countMap[record.Arrival.Unix()]++
		}

		// 在區間結束前離開則記錄 -1
		if record.Departure.Before(EndTime) {
			countMap[record.Departure.Unix()]--
		}

	}

	// 取出人數變化紀錄，並依照時間先後排序
	sortMapKey := []int64{}
	for k := range countMap {
		sortMapKey = append(sortMapKey, k)
	}
	sort.Sort(int64arr(sortMapKey))

	// 遍歷所有人數變化狀況，找出最大值
	max := 0
	current := 0
	for _, timeDuration := range sortMapKey {
		current += countMap[timeDuration]
		if max < current {
			max = current
		}
	}

	return max
}
