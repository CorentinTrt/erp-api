// Package models regroup all database's related stuff
package models

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User type define user object that will be stored in the DB
type User struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	FirstName         string             `bson:"first_name" json:"firstName"`
	MiddleNames       []string           `bson:"middle_names" json:"middleNames"`
	LastName          string             `bson:"last_name" json:"lastName"`
	Adress            adress             `bson:"adress" json:"adress"`
	Salary            salary             `bson:"salary" json:"salary"`
	Job               string             `bson:"job" json:"job"`
	JoStatus          string             `bson:"job_status" json:"jobStatus"`
	BeginingDate      int                `bson:"begining_date" json:"beginingDate"`
	NextInterviewDate int                `bson:"next_interview_date" json:"nextInterviewDate"`
	LastInterviewDate int                `bson:"last_interview_date" json:"lastInterViewDate"`
	ActivityStatus    string             `bson:"activity_status" json:"activityStatus"`
	CreatedOn         int                `bson:"created_on" json:"-"`
	UpdatedOn         int                `bson:"updated_on" json:"-"`
	DeletedOn         int                `bson:"deleted_on" json:"-"`
}

type adress struct {
	Number   int    `bson:"number" json:"number"`
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

// ToJSON serialize and return users data
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
