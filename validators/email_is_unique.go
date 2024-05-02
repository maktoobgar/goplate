package validators

import (
	"context"
	"database/sql"
	g "service/global"
	"service/pkg/errors"
	"service/repositories"
	"service/utils"
)

func EmailIsUnique(input interface{}) bool {
	db, err := g.DB()
	if err != nil {
		panic(errors.New(errors.ServiceUnavailable, "DbNotFound", err.Error(), nil))
	}
	defer db.Close()

	email, _ := input.(string)
	_, err = repositories.New(db).GetUserByEmail(context.TODO(), sql.NullString{String: email, Valid: true})
	return utils.IsErrorNotFound(err)
}
