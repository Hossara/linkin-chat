package helpers

import (
	"context"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"

	ut "github.com/go-playground/universal-translator"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

type ServiceGetter[T any] func(context.Context) T

var validate = validator.New()

var uni = ut.New(en.New(), en.New())

var trans, _ = uni.GetTranslator("en")

func ValidateRequestBody[T any](body T) map[string]string {
	if errs := validate.Struct(body); errs != nil {
		validationErrors := errs.(validator.ValidationErrors).Translate(trans)

		cleanedErrors := make(map[string]string)

		// Remove the struct name prefix from each field name
		for field, errMsg := range validationErrors {
			// Extract the field name after the last dot
			if dotIndex := strings.LastIndex(field, "."); dotIndex != -1 {
				field = field[dotIndex+1:]
			}
			cleanedErrors[field] = errMsg
		}

		return cleanedErrors
	}

	return nil
}

func ParseRequestBody[T any](c *fiber.Ctx, body *T) fiber.Map {
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	errParse := c.BodyParser(body)

	msg := fiber.Map{"error": ErrRequiredBodyNotFound.Error()}

	if errParse != nil {
		msg["message"] = errParse.Error()
		return msg
	}

	errValidation := ValidateRequestBody[T](*body)

	if errValidation != nil {
		msg["details"] = errValidation

		return msg
	}

	return nil
}
