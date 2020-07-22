package varios

import (
	"aplicacaoExemplo/domain"
	"encoding/json"
	"net/http"
)

func validarPerson(person domain.Person, w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Erro no formato JSON", http.StatusBadRequest)
		return
	}
	v := validator.New()
	err1 := v.Struct(person)
	if err1 != nil {
		var channels = []*domain.Error{}
		for _, e := range err1.(validator.ValidationErrors) {
			var error domain.Error
			error.Message = "Implementar arquivos para ler essa mensagem"
			error.Field = e.Field()
			channels = append(channels, &error)
		}
		errorHandler(w, r, channels)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(person)
	w.Write(res)
}


func errorHandler(w http.ResponseWriter, r *http.Request, l []*domain.Error) {
	jData, _ := json.Marshal(l)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jData)
}

