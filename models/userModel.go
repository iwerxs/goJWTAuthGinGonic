package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// convert json to go, and go to json
type User struct {
	ID							primitive.ObjectID		`bson:"_id"`
	First_name			*string								`json:"first_name" validate:"required,min=2,max=50"`
	Last_name				*string								`json:"last_name" validate:"required,min=2,max=50"`
	Company					*string								`json:"company" validate:"required"`
	Email						*string								`json:"email" validate:"email,required"`
	Password				*string								`json:"password" validate:"required,min=8,lowercase,uppercase,number,specialchar"`
	Token						*string								`json:"token"`
	User_Type				*string								`json:"user_type" validate:"required,ut=SADMIN|ut=PADMIN|ut=CADMIN|ut=USER"`
	Refresh_token		*string								`json:"refresh_token"`
	Created_at			time.Time							`json:"created_at"`
	Updated_at			time.Time							`json:"updated_at"`		
	User_id					string								`json:"user_id"`
}

//Validate a Password

// func (p *Password) Validate() error {
// 	if len(p.Password) < 8 {
// 		return errors.New("password must be min of 8 characters")
// 	}
// 	if !regexp.MustCompile(`[a-z]`).MatchString(p.Password){
// 		return errors.New("password must contain a lowercase letter")
// 	}
// 	if !regexp.MustCompile(`[A-Z]`).MatchString(p.Password){
// 		return errors.New("password must contain an uppercase letter")
// 	}
// 	if !regexp.MustCompile(`[0-9]`).MatchString(p.Password){
// 		return errors.New("password must contain a number")
// 	}
// 	if !regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(p.Password){
// 		return errors.New("password must contain a special character")
// 	}
// 	return nil
// }

//Test Check Password meets Validation Criteria
// password := &Password{Password: "MyP@ssw0rd"}
// if err := password.Validate(); err != nil {
// 	log.Println(err)
// }else{
// 	log.Println("password is valid")
// }