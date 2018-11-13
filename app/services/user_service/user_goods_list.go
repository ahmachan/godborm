package user_service

import (
	_ "godborm/app/models/user/view_model"
	"godborm/framework/db"
	"godborm/framework/models"
)

type UserGoodsListService struct {
	//vm  *models.ViewModel
	Dbh *db.DbHandle
}

func (m *UserGoodsListService) GetUserGoodsListServiceVm() map[int]map[string]string { //interface{}
	//ugVm := new(user.UserGoodsVm)
	//my_viewmodel := ugVm.GetUserGoodsListVm()
	//my_viewmodel := new(user.UserGoodsVm).GetUserGoodsListVm()
	t, _ := models.InstanceViewModel(m.Dbh)

	//my_viewmodel := make(map[string]map[string]interface{})
	// my_viewmodel := map[string]map[string]interface{}{}
	//my_viewmodel:=UserGoodsVm
	/*
		var mapWhere :=map[string]interface{} {
			"user.id":0,
			"user.gender":1,
			"_type":"or"
		}*/

	//return t.IocViewModel(my_viewmodel).Where("user.id>0 and user.gender=2").Limit(0, 10).OrderBy("user.id Desc,user.gender Asc").FindViewModel()

	return t.ViewModelIoc().Where("user.id>0 and user.gender=2").Limit(0, 10).OrderBy("user.id Desc,user.gender Asc").FindViewModel()

}
