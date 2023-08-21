package uenv

import (
	"fmt"
	"os"
	"se-backend/utils/uslice"

	"strings"
)

// GetStr retorna o valor da variável de ambiente como string ou erro caso o valor esteja vazio.
func GetStr(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("variável de ambiente não encontrada: %s", name)
	}

	return v, nil
}

// CheckEnvs verifica se os valores de todas as variáveis de ambiente foram informadas.
func CheckEnvs(table []CheckEnvTable) (map[string]string, error) {
	var (
		missing []string
		wrong   []string
		values  = make(map[string]string, len(table))
		errMsgs []string
	)

	for _, item := range table {
		v, err := GetStr(item.Name)
		if err != nil && (v == "" && !item.IgnoreEmpty && len(item.Expected) == 0) {
			missing = append(missing, item.Name)
		}

		if len(item.Expected) > 0 {
			if !uslice.Contains(item.Expected, v) {
				wrong = append(wrong, fmt.Sprintf(`%s="%s"`, item.Name, v))
			}
		}

		values[item.Name] = v
	}

	if len(missing) > 0 {
		errMsgs = append(errMsgs, fmt.Sprintf("-> Não carregadas: [%s]", strings.Join(missing, ", ")))
	}

	if len(wrong) > 0 {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"-> As seguintes variáveis de ambiente não receberam os valores esperados:\n%s",
			strings.Join(wrong, "\n"),
		))
	}

	var err error
	if len(errMsgs) > 0 {
		err = fmt.Errorf(
			"[ENV] Ocorrem o seguintes erros ao carregar as variáveis de ambiente:\n%s",
			strings.Join(errMsgs, "\n\n"),
		)
	}

	return values, err
}
