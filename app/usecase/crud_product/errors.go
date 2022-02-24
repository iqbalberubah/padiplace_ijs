package crud_product

import "errors"

var ErrUnexpected = errors.New("unexpected internal error")
var ErrProductNotFound = errors.New("product not found")
