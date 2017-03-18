package main

import (
	"fmt"
	"github.com/lonli078/psnapi"
)

func main() {
	fmt.Println("start")
	api := &psnapi.Api{Username: "username", Password: "password"}
	api.Auth()
	info, _ := api.Get_my_info()
	fmt.Println(info)
	friends, _ := api.Get_friends("36", "0")
	fmt.Println(friends)
	data2, _ := api.GetGameTrophies("NPWR06221_00", "all", "en", "")
	fmt.Println(data2)
	data3, _ := api.GetGameGroupList("NPWR06221_00", "en")
	fmt.Println(data3)
	data4, _ := api.GetTrophyInfo("NPWR06221_00", "all", "60", "en")
	fmt.Println(data4)
}
