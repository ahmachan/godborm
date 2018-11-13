package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"godborm/app/services/user_service"
	"godborm/framework/db"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
 * 启动入口
 */
func Start() {
	initEnv()
	//router := new(framework.router)
	//router.dispatch();
	//log.flushLog();
}

func getConfigFile() string {
	return "/conf/conf.ini"
}

//获取当前目录路径
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	indexPos := strings.LastIndex(path, "/")
	if indexPos < 0 {
		indexPos = strings.LastIndex(path, "\\")
	}
	if indexPos < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	//如果最后有斜杆,则不返回,返回则 "indexPos+1"
	return string(path[0:indexPos]), nil
}

func initEnv() {
	currPath, pathErr := getCurrentPath()
	if pathErr != nil {
		fmt.Println(pathErr)
	}
	lastConfigFile := currPath + getConfigFile()
	dbh, err := db.SetConfig(lastConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	ugService := new(user_service.UserGoodsListService)
	ugService.Dbh = dbh
	mapRes := ugService.GetUserGoodsListServiceVm()
	jsonStr, err := json.Marshal(mapRes)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}
