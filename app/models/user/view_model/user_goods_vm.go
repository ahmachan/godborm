package user

import (
	"fmt"
	"godborm/framework/models"
)

type UserGoodsVm struct {
	//@TODO
}

func init() {
	fmt.Println("init2")
	models.ViewModelInstance = map[string]map[string]interface{}{
		//"user": map[string]interface{} {}
		"user": {
			"_table":   "t_user",
			"id":       "user_id",
			"username": "username",
			"_type":    "LEFT",
		},
		"captcha": {
			"_table": "t_user_captcha",
			"code":   "captcha_code",
			"_on":    "user.id=captcha.user_id",
			"_type":  "LEFT",
		},
		"goods": {
			"_table":     "t_user_goods",
			"goods_code": "goods_code",
			"goods_name": "goods_name",
			"_on":        "user.id=goods.user_id",
			"_type":      "LEFT",
		},
	}
}

/*
func (m *UserGoodsVm) GetUserGoodsListVm() map[string]map[string]interface{} {
	my_viewmodel := models.ViewModelInstance
	return my_viewmodel

}
*/
/*
func (m *UserGoodsVm) GetAllBestGoodsList(offset, limit)
{
        where := [
            'StoreGoods.is_show' => 1,
            'Store.is_del'       => 0
        ];

        orderBy := 'StoreGoods.goods_id DESC';

        return $this->where(where)->limit(offset, limit)->order(orderBy)->select();
}
*/
