package users_handlers

import (
	g "service/global"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"
	"time"

	"github.com/kataras/iris/v12"
)

type MeRes struct {
	ID          int32     `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Avatar      string    `json:"avatar"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DisplayName string    `json:"display_name"`
	Gender      int32     `json:"gender"`
	IsActive    bool      `json:"is_active"`
	IsAdmin     bool      `json:"is_admin"`
	IsSuperuser bool      `json:"is_superuser"`
	CreatedAt   time.Time `json:"created_at"`
}

func Me(ctx iris.Context) {
	user := ctx.Value(g.UserKey).(*repositories.User)
	utils.SendJson(ctx, copier.Copy(&MeRes{}, user))
}
