package main

import (
	"godborm/framework"
)

func main() {
	app.Start()

	/*
		dbh, err := db.SetConfig("./conf/conf.ini")
		if err != nil {
			fmt.Println(err)
		}
	*/

	/*
		t, _ := models.InstanceViewModel(dbh)
		my_viewmodel := make(map[string]map[string]interface{})
		ugVm:=new(user.UserGoodsVm)
		my_viewmodel=ugVm.GetUserGoodsListVm()
		fmt.Println(t)
		fmt.Println(my_viewmodel)
	*/
	//fmt.Println(t.IocViewModel(my_viewmodel).Where("user.id>0").Limit(1, 4).OrderBy("user.id Desc").FindViewModel())

	/*
		t.Query("INSERT INTO t_user (`username`,`passwd`) VALUES ('xiaoweitest','xiaowei')")
		d := t.Query("select * from t_user")
		fmt.Println(d)

		n := t.Query("INSERT INTO t_user (`username`,`passwd`) VALUES ('xiaoweitest','xiaowei')")
		fmt.Println(n)
	*/
	//fmt.Println(t.Fileds("id,username").FindAll())
	//d := t.Query("select * from t_user")
	//fmt.Println(t.IocViewModel().Fileds("user.id as user_id,user.username,captcha.`code`").FindViewModel())

	//fmt.Println(t.SetTable("t_user").Fileds("id,username").FindAll())
	/*

		data := t.Fileds("user.id", "data.keywords", "user.username", "user.passwd").Join("data", "user.id = data.id").FindAll()
		fmt.Println(data)

		n = t.Query("update t_user set username='ceshishenma' where id =17 ")
		fmt.Println(n)

		n = t.Query("delete from t_user where id=16 ")
		fmt.Println(n)

		data = t.Query("select username,passwd from t_user")
		fmt.Println(data)

		value := make(map[string]interface{})
		value["username"] = "widuu"
		value["passwd"] = "widuu"
		_, err = t.Insert(value)
		fmt.Println(err)

		n, err = t.SetTable("user").Delete("id = 16")
		fmt.Println(n, err)

		sss := make(map[string]interface{})
		sss["username"] = "widuuweb"
		r, err := t.Where("username = 'widuu'").Update(sss)
		fmt.Println(r, err)
	*/
	/*
		data = c.SetTable("user").SetParam([]string{"*"}).Where("id>1").Limit(1, 5).OrderBy("id Desc").FindAll()
		for _, v := range data {
			for k, value := range v {
				fmt.Println(k, value)
			}
		}
	*/

}
