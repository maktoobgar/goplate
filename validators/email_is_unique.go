package validators

import (
	"context"
	"database/sql"
	g "service/global"
	"service/repositories"
	"service/utils"
)

func EmailIsUnique(ctx context.Context, input interface{}) bool {
	db := ctx.Value(g.DbInstance).(*sql.DB)

	email, _ := input.(string)
	_, err := repositories.New(db).GetUserWithApprovedEmail(ctx, sql.NullString{String: email, Valid: true})
	return utils.IsErrorNotFound(err)
}
