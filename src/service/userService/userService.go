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
	CreateUser(user user.User) (error, *user.User)
	GetUserByEmail(email string) user.User
}

func NewService() Service {
	return &service{dao: userDao.NewUserDao(database.GetMySqlEngine())}
}

type service struct {
	dao *userDao.Dao
}

func (s *service) CreateUser(user user.User) (error, *user.User) {
	return s.dao.CreateUser(user)
}

func (s *service) GetUserByEmail(email string) user.User {
	return s.dao.GetUserByEmail(email)
}

