// Package models regroup all database's related stuff
package models

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// User type define user object that will be stored in the DB
// swagger:model User
type User struct {
	// the firstname
	FirstName         string   `bson:"first_name" json:"firstName" validate:"required,max=50"`
	MiddleNames       []string `bson:"middle_names" json:"middleNames"`
	LastName          string   `bson:"last_name" json:"lastName" validate:"required,max=50"`
	Age               uint8    `bson:"age" json:"age" validate:"gte=16,lte=99"`
	Email             string   `bson:"email" json:"email" validate:"required,email"`
	Adress            adress   `bson:"adress" json:"adress"`
	Salary            salary   `bson:"salary" json:"salary"`
	Job               string   `bson:"job" json:"job" validate:"max=50"`
	JoStatus          string   `bson:"job_status" json:"jobStatus" validate:"omitempty,oneof=intern extern"`
	BeginingDate      int      `bson:"begining_date" json:"beginingDate"`
	NextInterviewDate int      `bson:"next_interview_date" json:"nextInterviewDate"`
	LastInterviewDate int      `bson:"last_interview_date" json:"lastInterViewDate"`
	ActivityStatus    string   `bson:"activity_status" json:"activityStatus" validate:"omitempty,oneof=active inactive"`
	CreatedOn         int      `bson:"created_on" json:"-"`
	UpdatedOn         int      `bson:"updated_on" json:"-"`
	DeletedOn         int      `bson:"deleted_on" json:"-"`
}

type adress struct {
	Number   uint16 `bson:"number" json:"number"`
	Street   string `bson:"street" json:"street"`
	City     string `bson:"city" json:"city"`
	Province string `bson:"province" json:"province"`
}

type salary struct {
	AmountYear int    `bson:"amount_year" json:"amountYear"`
	Bonus      string `bson:"bonus" json:"bonus"`
}

// Users type is a slice of User
type Users []*User

// Validate the User
func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

// ToJSON serialize and return a list of users
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// ToJSON serialize and return user data
func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// FromJSON deserialize and return users data
func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
