package comment_errors

import "errors"

var ErrCommentNotFound = errors.New("no matching comment not found")
var ErrNoPermission = errors.New("you don't have permission to manipulate this comment")
