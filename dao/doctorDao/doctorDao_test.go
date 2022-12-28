/**
 * @author zhaiyuanji
 * @date 2022年02月18日 10:09 下午
 */
package doctorDao

import (
	"OpenSchedule/database"
	"OpenSchedule/model/doctorModel"
	"fmt"
	"testing"
)

var dao = NewDoctorDao(database.GetElasticSearchEngine(), database.GetMySqlEngine())

func TestDao_SearchDoctor(t *testing.T) {
	dao.SyncInsurance()
}

func TestDao_CreateUser(t *testing.T) {
	u := doctorModel.DoctorUser{
		FullName: "Jeffrey Glasser",
		FirstName: "Jeffrey",
		LastName: "Glasser",
		Email: "Jeffre.Glasser@zendoc.com",
		Password: "123456",
		Npi: 1902809254,
	}

	if err := dao.CreateUser(u); err != nil {
		t.Errorf("create user failed: %v\r\n", err)
	} else {
		fmt.Println("create user success.")
	}
}
