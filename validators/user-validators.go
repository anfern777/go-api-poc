package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/metadex-dao/we3-gaming/service"
)

func ValidateUserRole(field validator.FieldLevel) bool {
	roles := service.NewUserService().GetAllUserRoles()

	for _, role := range roles {
		if role == field.Field().String() {
			return true
		}
	}
	return false
}
