package crud_po

import "errors"

var ErrUnexpected = errors.New("unexpected internal error")
var ErrProductNotFound = errors.New("po not found")
