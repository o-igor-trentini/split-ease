package validations

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"se-backend/config/selog"
	request "se-backend/controller/model/request/user"
	"strings"
)

// bindingWithJSONTag faz com que ao fazer o "binding" do JSON
// o campo de erro retornado seja a tag JSON ao invés do nome do campo da struct.
func bindingWithJSONTag() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	// Utilizado para que a validação "eqfield" da struct "request.UserRequest"
	// não retorne o nome do campo da struct ao invés do campo JSON na mensagem de validação.
	registerCustomUserRequestEqField := func(fieldName string) string {
		field, _ := reflect.TypeOf(request.UserRequest{}).FieldByName(fieldName)
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	}

	_ = validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
		return ut.Add("eqfield", "{0} deve ser igual a {1}.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field(), registerCustomUserRequestEqField(fe.Param()))
		return t
	})
}

// register Adiciona uma validação de "binding" customizada.
func register(tag, tagMsg string, validationFunc validator.Func) {
	// Registra a função de validação
	if err := validate.RegisterValidation(tag, validationFunc); err != nil {
		selog.Error("Erro ao registrar validação customizada", err)
	}

	// Registra a tradução
	err := validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, tagMsg, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
	if err != nil {
		selog.Error("Erro ao registrar mensagem customizada", err)
	}

}

// strongPass registra a validação customizada de senha forte.
func strongPass() {
	var v validator.Func = func(fl validator.FieldLevel) bool {
		if password, ok := fl.Field().Interface().(string); ok {
			passLen := len(password)

			if passLen < 8 || passLen > 100 {
				return false
			}

			rules := []*regexp.Regexp{
				// Mín. 1 caractere especial
				regexp.MustCompile("^(.*[!@#$%^&*])"),
				// Mín. 1 letra minúscula
				regexp.MustCompile("(.*[a-z])"),
				// Mín. 1 letra maiúscula
				regexp.MustCompile("(.*[A-Z])"),
				// Mín. 1 número
				regexp.MustCompile("(.*[0-9])"),
			}

			for _, v := range rules {
				if !v.MatchString(password) {
					return false
				}
			}
		}

		return true
	}

	register("strongpass", "{0} não atende aos requisitos de senha!", v)
}
