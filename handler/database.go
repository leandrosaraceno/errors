package handler

import (
	"errors"
	"fmt"
	"github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

// HandleDatabaseError database error handler
func HandleDatabaseError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {

		var errMsg string
		switch pqErr.Code {
		case "23505":
			// unique_violation
			errMsg = fmt.Sprintf("duplicate entry violates constraint: %s", pqErr.Constraint)
		case "23502":
			// NOT NULL violation
			errMsg = fmt.Sprintf("duplicate entry violates constraint: %s", pqErr.Column)
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
		log.Errorf("DATABASE Error:  <%s> ", errMsg)

		return mapConstraintError(pqErr)
	}
	return err
}
