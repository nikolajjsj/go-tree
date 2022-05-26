package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
  args := []string{"."}
  if len(os.Args) > 1 {
    args = os.Args[1:]
  }

  for _, arg := range args {
    err := tree(arg, "")
    if err != nil {
      log.Printf("tree %s: %v\n", arg, err)
    }
  }
}

func tree(root, indent string) error {
  info, err := os.Stat(root)
  if err != nil {
    return fmt.Errorf("could not stat %s: %v", root, err)
  }

  fmt.Println(info.Name())
  if !info.IsDir() {
    return nil 
  }

  dirInfo, err := ioutil.ReadDir(root)
  if err != nil {
    return fmt.Errorf("could not read dir %s: %v", root, err)
  }

  for i, info := range dirInfo {
    add := "|   "
    if i == len(dirInfo) - 1 {
      fmt.Printf(indent + "└--") 
      add = "    "
    } else {
      fmt.Printf(indent + "|--") 
    }

    if err := tree(filepath.Join(root, info.Name()), indent + add); err != nil {
      return err
    }
  }

  return nil
}
