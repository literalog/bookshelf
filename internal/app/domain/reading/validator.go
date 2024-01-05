package reading

import "github.com/literalog/bookshelf/pkg/models"

type Validator struct{}

func (v *Validator) Validate(r *models.Reading) error {
	return nil
}
