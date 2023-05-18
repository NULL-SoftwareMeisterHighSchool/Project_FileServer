package errors

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
)

func ErrStatusMatches(status codes.Code, err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("rpc error: code = %s", status.String()))
}
