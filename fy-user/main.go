package main

import (
	"fmt"
	"fy-user/domain/repository"
	"fy-user/domain/service"
	"fy-user/handler"
	user "fy-user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
)

func main() {
	//服务器参数配置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库
	db, err := gorm.Open("mysql", "root:root@/micro?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)
	//创建服务实例
	userDataService := service.NewUserDataService(repository.NewUserRepository(db))
	//注册Handler
	err = user.RegisterUserHandler(srv.Server(),&handler.User{UserDataService:userDataService})

	if err != nil {
		fmt.Println(err)
	}

	//运行服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
