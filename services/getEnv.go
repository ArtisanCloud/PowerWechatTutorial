package services

import (
  "errors"
  "os"
  "strconv"
)

var ErrEnvVarEmpty = errors.New("getEnv: environment variable empty")

func getEnvStr(key string) (string, error) {
  v := os.Getenv(key)
  if v == "" {
    return v, ErrEnvVarEmpty
  }
  return v, nil
}

func getEnvInt(key string) (int, error) {
  s, err := getEnvStr(key)
  if err != nil {
    return 0, err
  }
  v, err := strconv.Atoi(s)
  if err != nil {
    return 0, err
  }
  return v, nil
}

func getEnvBool(key string) (bool, error) {
  s, err := getEnvStr(key)
  if err != nil {
    return false, err
  }
  v, err := strconv.ParseBool(s)
  if err != nil {
    return false, err
  }
  return v, nil
}
