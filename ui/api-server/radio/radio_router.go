package radio

import "github.com/pkg/errors"

var (
	ErrorRadio      = errors.New("not found record")
	ErrorRadioId    = errors.New("id is not allow")
	ErrorRadioParam = errors.New("list team param error")
)

type CreateRadioParam struct {

}