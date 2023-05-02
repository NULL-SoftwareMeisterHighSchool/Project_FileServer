package article_errors

import "errors"

var ErrPermissionDenied = errors.New("you don't have permission for it")
var ErrArticleNotFound = errors.New("matching article not found")
