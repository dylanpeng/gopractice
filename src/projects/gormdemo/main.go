package main

import (
	"encoding/json"
	"fmt"
	"gopractice/projects/gormdemo/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dsn := "manager:123qweasd@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("init db fail. err: %s\n", err)
		return
	}

	filePath := "/Users/chenpeng/Downloads/address/{name}.json"

	for i := 1; i < 100; i++ {
		name := strconv.Itoa(i)

		if i < 10 {
			name = "0" + name
		}

		fileName := strings.Replace(filePath, "{name}", name, -1)

		// 打开json文件
		jsonFile, err := os.Open(fileName)

		// 最好要处理以下错误
		if err != nil {
			fmt.Println(err)
			jsonFile.Close()
			continue
		}

		// 要记得关闭
		byteValue, _ := ioutil.ReadAll(jsonFile)

		fmt.Printf("%s \n", string(byteValue))

		address := make(map[string]*Postcode)

		err = json.Unmarshal(byteValue, &address)

		if err != nil {
			fmt.Printf("%s \n", err)
		} else {
			for k, v := range address{
				AddPostcode(db, k, v)
			}
		}

		jsonFile.Close()
	}


}

func DBDemo() {
	dsn := "manager:123qweasd@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("init db fail. err: %s\n", err)
		return
	}

	AddDemo(db)

	demo1 := &table.Demo{}

	err = db.Model(&table.Demo{}).First(demo1).Error

	if err != nil {
		fmt.Printf("get demo fail. err: %s\n", err)
		return
	}

	demo2 := &table.Demo{}

	err = db.Table("demo").Where("id=?", 1).First(demo2).Error

	if err != nil {
		fmt.Printf("get demo fail. err: %s\n", err)
		return
	}
}

func AddDemo(db *gorm.DB) {
	now := time.Now()

	demo := &table.Demo{
		Name:      "test",
		Type:      1,
		Desc:      "描述",
		TimeStamp: now,
		TimeDate:  now,
		TimeDate2: &now,
	}

	err := db.Create(demo).Error

	if err != nil {
		fmt.Printf("create demo fail. err: %s\n", err)
		return
	}
}

func AddPostcode(db *gorm.DB, postcode string, postAddress *Postcode) {
	pa := &table.PostcodeAddress{
		Postcode: postcode,
		State:    postAddress.State,
		City:     postAddress.City,
		Street:   strings.Join(postAddress.List, ","),
		Ctime:    time.Now().UnixMilli(),
		Utime:    time.Now().UnixMilli(),
	}

	err := db.Create(pa).Error

	if err != nil {
		fmt.Printf("create pa fail. err: %s\n", err)
		return
	}
}

type Postcode struct {
	State string   `json:"state"`
	City  string   `json:"city"`
	List  []string `json:"list"`
}

func (e *Postcode) String() string {
	return fmt.Sprintf("%+v", *e)
}
