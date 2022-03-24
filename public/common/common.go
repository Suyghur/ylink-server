//@Author   : KaiShin
//@Time     : 2021/11/1

package logic

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gopkg.in/fsnotify.v1"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func GetSyncMapKeys(mm *sync.Map) []interface{} {
	var keyList []interface{}

	mm.Range(func(key, value interface{}) bool {
		keyList = append(keyList, key)
		return true
	})
	return keyList
}

func GetSyncMapValues(mm *sync.Map) []interface{} {
	var valList []interface{}

	mm.Range(func(key, value interface{}) bool {
		valList = append(valList, value)
		return true
	})
	return valList
}

func PreCheckConfig(path string) {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("<PreCheckConfig> err:", err)
		log.Println("<PreCheckConfig> load conf failed, wait file create...")
		watch, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		defer watch.Close()

		// 获取当前目录的上一级
		prePath := strings.Split(path, "/")

		err = watch.Add(prePath[0])
		if err != nil {
			log.Fatal(err)
		}

		select {
		case ev := <-watch.Events:
			{
				log.Println("<PreCheckConfig> ev.Op: ", ev.Op)
				if ev.Op&fsnotify.Create == fsnotify.Create {
					break
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					break
				}
			}
		}

		log.Printf("<PreCheckConfig> %s has bean created", path)
	}
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func CleanContent(chatStr string) string {
	cleanStr := strings.Replace(chatStr, " ", "", -1)
	cleanStr = strings.Replace(cleanStr, "\n", "", -1)
	cleanStr = strings.Replace(cleanStr, "\x0B", "", -1)
	cleanStr = strings.Replace(cleanStr, "\t", "", -1)
	return cleanStr
}

func StructToMap(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	data, _ := json.Marshal(v)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil
	}
	return m
}
