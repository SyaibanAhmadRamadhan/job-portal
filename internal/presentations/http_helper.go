package presentations

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strings"
)

type HTTPHelper struct {
	trans     ut.Translator
	validator *validator.Validate
}

func NewHTTPHelper(t ut.Translator, v *validator.Validate) *HTTPHelper {
	return &HTTPHelper{
		trans:     t,
		validator: v,
	}
}

func (h *HTTPHelper) ErrResp(c *fiber.Ctx, err error) error {
	var (
		validationErrors validator.ValidationErrors
		errTrace         *tracer.ErrTrace
		errFiber         *fiber.Error
		httpCode         = http.StatusInternalServerError
	)

	var resp any

	if errors.As(err, &validationErrors) {
		httpCode = http.StatusBadRequest
		resp = parseValidationErrors(validationErrors, h.trans)

	} else if errors.As(err, &errFiber) {
		httpCode = errFiber.Code
		resp = api.Error{
			Message: errFiber.Error(),
		}

	} else if errors.As(err, &errTrace) {
		httpCode, resp = handleTraceError(c, errTrace)

	} else if errors.Is(err, context.DeadlineExceeded) {
		httpCode = http.StatusRequestTimeout
		resp = api.Error{
			Message: "request timeout",
		}
	} else {
		resp = api.Error{
			Message: err.Error(),
		}
	}
	return c.Status(httpCode).JSON(resp)
}

func (h *HTTPHelper) Bind(c *fiber.Ctx, v interface{}) error {
	err := c.BodyParser(&v)
	if err != nil {
		return err
	}

	err = h.validator.Struct(v)
	if err != nil {
		return err
	}

	return nil
}

func parseValidationErrors(validationErrors validator.ValidationErrors, trans ut.Translator) api.Error400 {
	err400 := api.Error400{
		Errors:  map[string][]string{},
		Message: "invalid your request",
	}
	for _, validationError := range validationErrors {
		fieldName := ""
		namespaceParts := strings.Split(validationError.Namespace(), ".")
		if len(namespaceParts) > 1 {
			fieldName = strings.Join(namespaceParts[1:], ".")
		}
		err400.Errors[fieldName] = append(err400.Errors[fieldName], validationError.Translate(trans))
	}
	return err400
}

func handleTraceError(c *fiber.Ctx, errTrace *tracer.ErrTrace) (int, api.Error) {
	resp := api.Error{
		Message: errTrace.Msg,
	}
	span := trace.SpanFromContext(c.UserContext())
	if span.IsRecording() {
		span.SetAttributes(attribute.String("error_server", errTrace.Error()))
	}
	switch {
	case strings.Contains(errTrace.Msg, tracer.PrefixBadRequest):
		return http.StatusBadRequest, resp
	case strings.Contains(errTrace.Msg, tracer.PrefixNotFound):
		return http.StatusNotFound, resp
	case strings.Contains(errTrace.Msg, tracer.PrefixBadGateway):
		return http.StatusBadGateway, resp
	case strings.Contains(errTrace.Msg, tracer.PrefixUnAuthorization):
		return http.StatusUnauthorized, resp
	case strings.Contains(errTrace.Msg, tracer.PrefixTimeOut):
		return http.StatusRequestTimeout, resp
	default:
		return http.StatusInternalServerError, resp
	}
}

func (h *HTTPHelper) BindToPaginationInput(c *fiber.Ctx) (primitive.PaginationInput, error) {
	page := c.QueryInt("page")
	pageSize := c.QueryInt("page_size")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 50 {
		pageSize = 50
	}

	return primitive.PaginationInput{
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (h *HTTPHelper) BindToPaginationResponse(input primitive.PaginationOutput) *api.PaginationResponse {
	return &api.PaginationResponse{
		Page:      input.Page,
		PageCount: input.PageCount,
		PageSize:  input.PageSize,
		TotalData: input.TotalData,
	}
}
