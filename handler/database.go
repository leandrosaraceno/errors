package handler

import (
	"errors"
	"fmt"
	"github.com/leandrosaraceno/errors/model"
	"github.com/lib/pq"
)

func HandleDatabaseError(err error) *model.CustomError {
	//if pqErr, ok := err.(*pq.Error); ok {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		var err *model.CustomError
		var errMsg string
		switch pqErr.Code {
		case "23505":
			// unique_violation
			errMsg = fmt.Sprintf("duplicate entry violates constraint: %s", pqErr.Constraint)
		case "23503":
			// foreign_key_violation
			errMsg = fmt.Sprintf("foreign key constraint violated: %s", pqErr.Constraint)
		case "23514":
			// check_violation
			errMsg = fmt.Sprintf("check constraint '%s' violated: %s", pqErr.Constraint, pqErr.Message)
		default:
			// generic PostgreSQL error
			errMsg = fmt.Sprintf("database error [%s]: %s", pqErr.Code, pqErr.Message)
		}
		err = mapConstraintError(pqErr)
		err.LogMessage = errMsg
		return err
	}
	return nil
}
