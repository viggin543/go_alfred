package api

import (
	"example.com/banana/teamcity/internal/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func readCacheFile(route string) (string, os.FileInfo, error) {
	split := strings.Split(route, "/")
	cacheFileName := "./cache/" + split[len(split)-1]+".xml"
	file, err := os.Stat(cacheFileName)
	return cacheFileName, file, err
}


func cacheResponse(res *http.Response, cacheFileName string ) []byte  {
	body, _ := ioutil.ReadAll(res.Body)
	createCAcheDir(cacheFileName)
	writeCAcheFile(cacheFileName, body)
	return body
}

func writeCAcheFile(cacheFileName string, body []byte) {
	err := ioutil.WriteFile(cacheFileName, body, 777)
	if err != nil {
		logger.Log.Println("ERROR", err.Error())
	} else {
		logger.Log.Println("cached -> ", cacheFileName, len(body))
	}
}

func createCAcheDir(cacheFileName string) {
	split := strings.Split( cacheFileName,"/")
	cacheDir := strings.Join(split[:len(split)-1], "/")
	mkdirErr := os.Mkdir(cacheDir, os.FileMode(0777))
	if mkdirErr != nil {
		logger.Log.Println("ERROR", "failed to craete cache dir in: ", cacheDir, mkdirErr.Error())
	}
}

func notOldEnough(file os.FileInfo) bool {
	duration := time.Duration(-2629743000000000) // one month
	//duration := time.Duration(0) // one month
	oneMonth := time.Now().Add(duration)
	return file.ModTime().Before(oneMonth)
}