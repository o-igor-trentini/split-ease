package validations

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	ptBRLocale "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"se-backend/config/seerror"
	"se-backend/config/selog"

	ptBRTranslation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ptBR := ptBRLocale.New()
		unt := ut.New(ptBR, ptBR)
		trans, _ = unt.GetTranslator("pt_BR")
		_ = ptBRTranslation.RegisterDefaultTranslations(v, trans)

		validate = v

		bindingWithJSONTag()
	}

	strongPass()
}

func Validate(validationErr error) seerror.SEError {
	var res seerror.SEError
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		res = seerror.NewBadRequestErr("Tipo de campo inválido")
	} else if errors.As(validationErr, &jsonValidationError) {
		causes := make([]seerror.Cause, 0)

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := seerror.Cause{
				Message: e.Translate(trans),
				Field:   e.Field(),
			}

			causes = append(causes, cause)
		}

		res = seerror.NewUnprocessableEntityErr("Campos inválidos", causes)
	} else {
		res = seerror.NewBadRequestErr("Não foi possível converter os campos")
	}

	selog.Error(res.GetMessage(), res)
	return res
}
