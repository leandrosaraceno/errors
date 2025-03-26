package handler

import (
	"github.com/leandrosaraceno/errors/constants"
	"github.com/leandrosaraceno/errors/model"
	"github.com/lib/pq"
	"net/http"
)

var constraintErrorMap = map[string]*model.CustomError{
	"ActivePrograms.check_sequence_id": model.NewCustomError(constants.ActiveProgramsInvalidSequence, http.StatusConflict),
}

func mapConstraintError(pgErr *pq.Error) *model.CustomError {
	if customErr, exists := constraintErrorMap[pgErr.Table+"."+pgErr.Constraint]; exists {
		return customErr
	}
	return model.NewCustomError(
		constants.InternalServerError,
		http.StatusInternalServerError)
}
