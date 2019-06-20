package main


import (
    "io"
    "fmt"
    "time"
    "os"
    "log"
)


type HwAccepted struct {
    Id int
    Grade int
}


func (self HwAccepted) Log() string {
    dtFmt, logFmt := "2006.01.02 15:04:05", `%s | homework #%d accepted with grade %d`
    return fmt.Sprintf(logFmt, time.Now().Format(dtFmt), self.Id, self.Grade)
}


type HwSubmitted struct {
    Id int
    Code string
    Comment string
}


func (self HwSubmitted) Log() string {
    dtFmt, logFmt := "2006.01.02 15:04:05", `%s | homework #%d submitted, comment: '%s'`
    return fmt.Sprintf(logFmt, time.Now().Format(dtFmt), self.Id, self.Comment)
}


type OtusEvent interface {
    Log() string
}


func LogOtusEvent(e OtusEvent, w io.Writer) error {
    _, err := io.WriteString(w, e.Log()+"\n")
    return err
}


func main() {
    events := []OtusEvent{
        HwSubmitted{162342, "print stupid computer!", "2nd try"},
        HwSubmitted{162347, "Ook. Ook.", "Ook!"},
        HwAccepted{162342, 9},
        HwSubmitted{123445, "Bar <- Foo", "Do you want to go out?"},
        HwSubmitted{167652, "// js is just the best", "True story"},
        HwAccepted{167652, 10},
        HwSubmitted{166543, "Foo:=bar", "Me though"},
        HwAccepted{162347, 8},
        HwAccepted{166543, 4},
        HwAccepted{123445, 8},
    }

    for _, event := range(events) {
        if err := LogOtusEvent(event, os.Stdout); err != nil {
            log.Fatal(err)
        }
        time.Sleep(413 * time.Millisecond)
    }
}
