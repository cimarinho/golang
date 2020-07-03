package golang

import (
	"aplicacaoExemplo/domain"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DataSave(person domain.Person) {
	saveMysql(person)
}

func saveMongo() {
}

func saveMysql(person domain.Person) {
	db, err := gorm.Open("mysql", "user:password@/person") // colocar usuario e senha num arquivo
	if err != nil {
		panic(err)
	}
	defer db.Close()

	langs := []Language{{ Name: "Portugues"}, {Name: "Espanhol"}}

	p := PersonMysql{
		Age:             person.Age,
		Name:            person.Name,
		Email:           person.Email,
		Birthday:        person.Birthday,
		PersonLanguages: langs,
		Contacts: []ContactMysql{
			{CountryCode: 91, MobileNo: "956112"},
			{CountryCode: 91, MobileNo: "997555"}},
	}

	address := AddressMysql{
		City:   person.Address.City,
		Street: person.Address.Street,
		Person: p,
	}
	db.SingularTable(true)
	db.Model(PersonLanguage{}).AddForeignKey("person_mysql_id", "Person(id)", "CASCADE", "CASCADE")
	db.Model(PersonLanguage{}).AddForeignKey("language_id", "Language(id)", "CASCADE", "CASCADE")
	db.Save(&address)

	fmt.Println(p)
}
