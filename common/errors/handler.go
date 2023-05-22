package errors

import (
	"errors"

	article_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/articles/errors"
	comment_errors "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/domain/comments/errors"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Return status code with error message
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})

}

// err should not be nil
func StatusForError(err error) error {
	switch err {
	case article_errors.ErrArticleNotFound:
		return status.Error(codes.NotFound, err.Error())
	case article_errors.ErrPermissionDenied:
		return status.Error(codes.PermissionDenied, err.Error())
	case comment_errors.ErrCommentNotFound:
		return status.Error(codes.NotFound, err.Error())
	case comment_errors.ErrNoPermission:
		return status.Error(codes.PermissionDenied, err.Error())
	case comment_errors.ErrReplyingToReply:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
