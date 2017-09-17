package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// contains
func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// 与えられたファイルまたはディレクトリがあったらあったらtrue，なかったらfalse
func dir_exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func main() {
	// カレントディレクトリ取得
	current_path, _ := os.Getwd()

	// ファイルリストを取得
	files, err := ioutil.ReadDir(current_path)
	if err != nil {
		log.Fatal(err)
	}

	// 隠しファイルを除去してファイルリストを回す
	for _, f := range files {
		if f.Name()[:1] != "." && f.IsDir() == false {
			// 拡張子取得
			pos := strings.LastIndex(f.Name(), ".")

			// 引数で指定されたもののみを仕分けする
			if contains(os.Args, f.Name()[pos+1:]) {
				// 拡張子名のディレクトリを作ってそこにファイルをぶち込む
				if dir_exists(f.Name()[pos+1:]) {
					os.Rename(f.Name(), f.Name()[pos+1:]+"/"+f.Name())
				} else {
					os.Mkdir(f.Name()[pos+1:], 0777)
					os.Rename(f.Name(), f.Name()[pos+1:]+"/"+f.Name())
				}
			}
		}
	}
}
