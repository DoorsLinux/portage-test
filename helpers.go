package main

import (
  "os"
  "errors"
)

func file_exists(dir string) bool {
  if _, err := os.Stat(dir); err == nil {
    return true
  } else if errors.Is(err, os.ErrNotExist) {
    return false
  }
  return false
}

func current_dir_can_build_ebs() bool {
  return file_exists("./ebuild")
}

func current_dir_can_build_pbs() bool {
  return file_exists("./pbuild")
}

