package validators

import (
	"context"
	"database/sql"
	g "service/global"
	"service/repositories"
	"service/utils"
)

func PhoneNumberIsUnique(ctx context.Context, input interface{}) bool {
	db := ctx.Value(g.DbInstance).(*sql.DB)
	defer db.Close()

	phoneNumber, _ := input.(string)
	_, err := repositories.New(db).GetUserWithApprovedPhoneNumber(ctx, phoneNumber)
	return utils.IsErrorNotFound(err)
}
