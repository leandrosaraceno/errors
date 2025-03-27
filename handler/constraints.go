package handler

import (
	"github.com/leandrosaraceno/errors/constants"
	"github.com/leandrosaraceno/errors/model"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var constraintErrorMap = map[string]error{
	"ActivePrograms.check_sequence_id": model.NewConflictError(constants.ActiveProgramsInvalidSequence),
}

func mapConstraintError(pgErr *pq.Error) error {
	key := pgErr.Table + "." + pgErr.Constraint
	if customErr, exists := constraintErrorMap[key]; exists {
		return customErr
	}
	log.Errorf("Constraint key <%s> DOES NOT EXIST", key)
	return model.NewInternalServerError(constants.InternalServerError)
}
