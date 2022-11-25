package userService

import (
	"OpenSchedule/model/userModel"
	"errors"
)

func (s *service) GetUserInsurance(userID int) ([]userModel.UserInsurance, error) {
	return s.dao.GetUserInsurance(userID)
}

func (s *service) UpdateUserInsurance(ins []userModel.UserInsurance) error {
	if len(ins) == 0 {
		return errors.New("ins is required")
	}
	return s.dao.UpdateUserInsurance(ins)
}
