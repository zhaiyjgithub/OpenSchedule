package userModel

type UserProfile struct {
	UserID        int     `json:"userID" validate:"required"`
	Email         *string `json:"email" validate:"omitempty,email"`
	Phone         *string `json:"phone" validate:"omitempty,e164"`
	Birthday      *string `json:"birthday" validate:"omitempty,datetime=2006-01-02"`
	Gender        *string `json:"gender" validate:"omitempty,oneof=F M"`
	StreetAddress *string `json:"streetAddress" validate:"omitempty"`
	Suit          *string `json:"suit" validate:"omitempty"`
	City          *string `json:"city" validate:"omitempty"`
	State         *string `json:"state" validate:"omitempty"`
	Zip           *string `json:"zip" validate:"omitempty"`
}
