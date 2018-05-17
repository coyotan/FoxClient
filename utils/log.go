package utils

import (

  "log"
  "os"

)

var (
  Log *log.Logger
)

func NewLog(logpath string) {

  file, err := os.Crete(logpath)

  if err != nil {
      panic(err)
  }

  Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
