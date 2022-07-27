/**
 * @author zhaiyuanji
 * @date 2022年06月11日 8:46 上午
 */
package user

import (
	"OpenSchedule/database"
	"OpenSchedule/model/user"
	"fmt"
	"testing"
)

func TestDao_CreateSubUser(t *testing.T) {
	d := NewUserDao(database.GetMySqlEngine())
	subUser := user.SubUsers{
		FirstName: "Yuanji",
		LastName:  "Test",
		Phone:     "1234567890",
		Email:     "ajflsd@sfjslf.net",
		Birthday:  "2020-10-10",
		UserID:    4,
	}
	err := d.CreateSubUser(subUser)
	if err != nil {
		t.Errorf("create uer err: %v\r\n", err.Error())
	}
	fmt.Println("Create user success")
}

func TestDao_UpdateSubUserPhone(t *testing.T) {
	d := NewUserDao(database.GetMySqlEngine())
	err := d.UpdateSubUserPhone(1, "8888888888")
	if err != nil {
		t.Errorf("Update failed %s\r\n", err.Error())
	}
	fmt.Println("Update phone success.")
}

func TestDao_GetSubUsers(t *testing.T) {
	d := NewUserDao(database.GetMySqlEngine())
	users := d.GetSubUsers(4)
	fmt.Println(users)
}
