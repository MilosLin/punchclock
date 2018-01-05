package punchclock_test

import (
	"punchclock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var GMT8 *time.Location

func init() {
	GMT8 = time.FixedZone("GMT+8", 8*60*60)
}

// Test_TimeParse : 測試time取Unix後再還原
func Test_TimeParse(t *testing.T) {
	now := time.Now()
	unitTime := now.Unix()
	anow := time.Unix(unitTime, 0)

	assert.Equal(t,
		now.Format(time.RFC3339),
		anow.Format(time.RFC3339),
	)
}

// Test_getMaxNumDuringPeriod : 取得指定區間(12:00~13:00)最大連線數
//
// 測試資料入下圖:
// c1   ↑-------↑
//    11:00   11:30
// c2   ↑---------------↑
//    11:00           12:00
// c3   ↑-------------------------↑
//    11:00                     12:30
// c4                   ↑---------↑
//                    12:00     12:30
// c5                          ↑--------------↑
//                           12:20          12:50
// c6                          ↑--------------------↑
//                           12:20                13:00
// c7                                     ↑---------------↑
//                                      12:40           13:10
// c8                                               ↑-----↑
//                                                13:00 13:10
// c9  ↑--------------------------------------------------↑
//   11:00                                              13:10
//
// period               ↑---------------------------↑
//                    12:00                       13:00
//
func Test_getMaxNumDuringPeriod(t *testing.T) {
	StartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:00:00", GMT8)
	EndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:00:00", GMT8)

	assert.Equal(
		t,
		5,
		punchclock.GetMaxNumDuringPeriod(StartTime, EndTime, getTestCase()),
	)
}

// Test_getMaxNumDuringPeriod_2 : 移動指定區間(11:00~12:00)，檢驗結果
func Test_getMaxNumDuringPeriod_2(t *testing.T) {
	StartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:00:00", GMT8)
	EndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:00:00", GMT8)

	assert.Equal(
		t,
		4,
		punchclock.GetMaxNumDuringPeriod(StartTime, EndTime, getTestCase()),
	)
}

// getTestCase : 回傳測試用簽到記錄
func getTestCase() []punchclock.Record {
	ret := []punchclock.Record{}

	ret = append(
		ret,
		case1(),
		case2(),
		case3(),
		case4(),
		case5(),
		case6(),
		case7(),
		case8(),
		case9(),
	)

	return ret
}

func case1() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:30:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case2() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:00:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case3() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:30:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case4() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:30:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case5() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:20:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:50:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case6() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:20:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:00:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case7() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 12:40:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:10:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case8() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:10:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}

func case9() punchclock.Record {
	Arrival, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 11:00:00", GMT8)
	Departure, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-05 13:10:00", GMT8)
	return punchclock.Record{
		Arrival:   Arrival,
		Departure: Departure,
	}
}
