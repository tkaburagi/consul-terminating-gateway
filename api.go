package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

type User struct {
    ID      string `gorm:"primary_key" json:"id"`
    Username  string `gorm:"primary_key" json:"username"`
}

type Version struct {
    Version string `json:"version"`
}

func (s *User) TableName() string {
    return "auth_user"
}

func gormConnect() *gorm.DB {
    DBMS     := "mysql"
    USER     := "admin"
    PASS     := "g9x68LZtYpG39IY717XE"
    PROTOCOL := "tcp(:3306)"
    DBNAME   := "handson"

    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
      panic(err.Error())
    }
    return db
}

func main() {

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Get("/users", GetAllUsers),
        rest.Get("/version", GetVersion),
    )
    if err != nil {
        log.Fatal(err)
    }

    api.SetApp(router)

    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

    fmt.Println("Exit!!!")
}

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
    db := gormConnect()
    defer db.Close()

    // 全件取得
    var allUsers []User
    db.Find(&allUsers)
    fmt.Println(allUsers)

    w.WriteHeader(http.StatusOK)
    w.WriteJson(&allUsers)
}

func GetVersion(w rest.ResponseWriter, r *rest.Request) {
    version := Version{"v1.0.0"}

    w.WriteJson(version)
}
