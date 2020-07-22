package varios

import (
	"aplicacaoExemplo/domain"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func SalvarPerson(w http.ResponseWriter, r *http.Request) {
	{
		w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
		var person1 domain.Person
		validarPerson(person1, w, r)
		err := json.NewDecoder(r.Body).Decode(&person1) // storing in person   //variable of type user
		if err != nil {
			fmt.Print(err)
		}
	}
	var person domain.Person
	json.NewDecoder(r.Body).Decode(&person)

	insertResult:= saveMongo(person)

	json.NewEncoder(w).Encode(insertResult) // return the //mongodb ID of generated document

}
