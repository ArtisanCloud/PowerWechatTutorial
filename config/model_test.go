package config

import (
  "fmt"
  "github.com/go-playground/assert/v2"
  "github.com/jinzhu/configor"
  "log"
  "os"
  "testing"
)

func TestGet(t *testing.T)  {
  //os.Setenv("app_id", "wx_xxxxx")
  _ = os.Setenv("mch_id", "mch_xxx")
  var c Payment
  err := configor.Load(&c, "./model_test.yaml")
  if err != nil {
    log.Println(c)
    panic(err)
  }
  fmt.Printf("config: %#v\n", c)
  assert.Equal(t, c.AppID, "hello-app")
  assert.Equal(t, c.MchID, "mch_xxx")
}