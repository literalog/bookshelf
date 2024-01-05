package reading

import (
	"net/http"

	"github.com/literalog/cerrors"
)

var (
	ErrEmptyId = cerrors.New("id is empty", http.StatusBadRequest)
)
