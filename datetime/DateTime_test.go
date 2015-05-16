package datetime

import (
	"fmt"
	"testing"
)

//func TestNew(t *testing.T){
//    if (NewDate(2000,1,1))
//}

func TestNewDate(t *testing.T) {
	dt := NewDate(2000, 1, 1).Time
	if dt.Year() != 2000 || dt.Month() != 1 || dt.Day() != 1 {
		t.Error("test NewDate faild")
	} else {
		t.Log("test NewDate pass")
	}
}

func TestGetDateString(t *testing.T) {
	dt := NewDate(2015, 8, 18)
	if dt.GetDateString() != "2015-08-18" {
		t.Error("DateTime:GetDateString faild")
	} else {
		t.Log("test DateTime:GetDateString pass")
	}
}

func TestGetTimeString(t *testing.T) {
	dt := NewDate(2015, 8, 18)
	if dt.GetTimeString() != "00:00:00" {
		t.Error("DateTime:GetTimeString faild")
	} else {
		t.Log("test DateTime:GetTimeString pass")
	}
}

func TestSetTime(t *testing.T) {
	dt := NewDate(2015, 8, 18)
	dt.SetTime("17:32:00")
	if dt.GetTimeString() != "17:32:00" {
		t.Error("DateTime:SetTime faild")
	} else {
		t.Log("test DateTime:SetTime pass")
	}
}

//func TestAdd(t *testing.T) {
//    if (Add(1, 2) != 3) {
//        t.Error("test foo:Addr failed")
//    } else {
//        t.Log("test foo:Addr pass")
//    }
//}

//func TestAdd(t *testing.T){
//    if (Add(1,3)!=5){
//        t.Error("!=5")
//    }else{
//        t.Log("addr pass")
//    }
//}
