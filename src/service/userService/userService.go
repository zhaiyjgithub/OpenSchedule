/**
 * @author zhaiyuanji
 * @date 2022年04月08日 10:48 上午
 */
package userService

import (
	userDao "OpenSchedule/src/dao/user"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/user"
)

type Service interface {
	CreateUser(user user.Users) (error, *user.Users)
	GetUserByEmail(email string) user.Users
}

func NewService() Service {
	return &service{dao: userDao.NewUserDao(database.GetMySqlEngine())}
}

type service struct {
	dao *userDao.Dao
}

func (s *service) CreateUser(user user.Users) (error, *user.Users) {
	return s.dao.CreateUser(user)
}

func (s *service) GetUserByEmail(email string) user.Users {
	return s.dao.GetUserByEmail(email)
}

