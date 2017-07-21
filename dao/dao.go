package dao

import "time"

var GlogMap map[string]Blog

func init() {
	GlogMap = map[string]Blog{
		"1": Blog{"1", "james", "what a fucking day!!", int64(time.Now().Unix())},
		"2": Blog{"2", "jeff", "what a good day!!", int64(time.Now().Unix())},
		"3": Blog{"3", "leff", "what a wanderfull day!!", int64(time.Now().Unix())},
		"4": Blog{"4", "Ann", "what a day!!", int64(time.Now().Unix())},
	}
}
