package varios

import (
	"aplicacaoExemplo/domain"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func saveMysql(person domain.Person) {
	db, err := gorm.Open("mysql", "user:password@/person") // colocar usuario e senha num arquivo
	if err != nil {
		panic(err)
	}
	defer db.Close()

	langs := []Language{{Name: "Portugues"}, {Name: "Espanhol"}}

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

type PersonMysql struct {
	gorm.Model
	Id       string            `gorm:"primary_key"`
	Name     string            `gorm:"not null"`
	Email    string            `gorm:"not null"`
	Age      int               `gorm:"not null"`
	Birthday string            `gorm:"not null"`
	Contacts []ContactMysql    `gorm:"ForeignKey:PersonId"`
	PersonLanguages []Language `gorm:"many2many:Person_language";"ForeignKey:PersonId;"`
}

type PersonLanguage struct {
	PersonId int64
	LanguageId int64
}

type ContactMysql struct {
	Id int64 `gorm:"primary_key"`
	CountryCode int
	MobileNo string
	PersonId int
}
//one to one
type AddressMysql struct {
	gorm.Model
	Id       int64  `gorm:"primary_key"`
	Street   string `gorm:"not null"`
	City     string `gorm:"not null"`
	Person   PersonMysql
	PersonId int64 `gorm:"ForeignKey:Id"`
}

func (PersonLanguage) TableName() string {
	return "Person_language"
}

func (ContactMysql) TableName() string {
	return "Contact"
}

func (PersonMysql) TableName() string {
	return "Person"
}

type Language struct {
	Id int64 `gorm:"primary_key"`
	Name string
}

func (Language) TableName() string {
	return "Language"
}

func (AddressMysql) TableName() string {
	return "Address"
}



