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
	CreateSubUser(subUser user.SubUsers) error
	GetSubUsers(userId int) []user.SubUsers
	UpdateSubUserPhone(userID int, phone string) error
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

func (s *service) CreateSubUser(subUser user.SubUsers) error {
	return s.dao.CreateSubUser(subUser)
}

func (s *service) GetSubUsers(userId int) []user.SubUsers {
	return s.dao.GetSubUsers(userId)
}

func (s *service) UpdateSubUserPhone(userID int, phone string) error {
	return s.dao.UpdateSubUserPhone(userID, phone)
}
