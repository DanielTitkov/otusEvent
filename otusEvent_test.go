package main


import (
    "testing"
    "time"
    "github.com/bouk/monkey"
)


// sets datetime to 1997.07.01 16:30:55
func setupMockTime() *monkey.PatchGuard {
    mockTime := time.Date(1997, time.July, 1, 16, 30, 55, 0, time.UTC)
    patch := monkey.Patch(time.Now, func() time.Time { return mockTime })
    return patch
}


func TestHwAcceptedLog(t *testing.T) {
    defer setupMockTime().Unpatch()
    hwa := HwAccepted{123999, 7}
    expected := "1997.07.01 16:30:55 | homework #123999 accepted with grade 7"
    if res := hwa.Log(); res != expected {
        t.Errorf("expected %v, got %v", expected, res)
    }
}


func TestHwSubmittedLog(t *testing.T) {
    defer setupMockTime().Unpatch()
    hws := HwSubmitted{123999, "print('Foobar!')", "Best code ever!!"}
    expected := "1997.07.01 16:30:55 | homework #123999 submitted, comment: 'Best code ever!!'"
    if res := hws.Log(); res != expected {
        t.Errorf("expected %v, got %v", expected, res)
    }
}
