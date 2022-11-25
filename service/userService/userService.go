/**
 * @author zhaiyuanji
 * @date 2022年04月08日 10:48 上午
 */
package userService

import (
	userDao "OpenSchedule/dao/userDao"
	"OpenSchedule/database"
	"OpenSchedule/model/userModel"
)

type Service interface {
	CreateUser(user userModel.User) (error, *userModel.User)
	GetUserByEmail(email string) userModel.User
	CreateSubUser(subUser userModel.SubUsers) error
	GetSubUsers(userId int) []userModel.SubUsers
	UpdateSubUserPhone(userID int, phone string) error
	UpdateUserProfile(user userModel.UserProfile) error
	GetUserByID(userID int) (userModel.User, error)
	GetUserInsurance(userID int) ([]userModel.UserInsurance, error)
	UpdateUserInsurance(ins []userModel.UserInsurance) error
}

func NewService() Service {
	return &service{dao: userDao.NewUserDao(database.GetMySqlEngine())}
}

type service struct {
	dao *userDao.Dao
}

func (s *service) CreateUser(user userModel.User) (error, *userModel.User) {
	return s.dao.CreateUser(user)
}

func (s *service) GetUserByEmail(email string) userModel.User {
	return s.dao.GetUserByEmail(email)
}

func (s *service) CreateSubUser(subUser userModel.SubUsers) error {
	return s.dao.CreateSubUser(subUser)
}

func (s *service) GetSubUsers(userId int) []userModel.SubUsers {
	return s.dao.GetSubUsers(userId)
}

func (s *service) UpdateSubUserPhone(userID int, phone string) error {
	return s.dao.UpdateSubUserPhone(userID, phone)
}

func (s *service) UpdateUserProfile(user userModel.UserProfile) error {
	return s.dao.UpdateUserProfile(user)
}

func (s *service) GetUserByID(userID int) (userModel.User, error) {
	return s.dao.GetUserByID(userID)
}
