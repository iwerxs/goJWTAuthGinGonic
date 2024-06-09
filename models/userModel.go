package models

import "time"

// convert json to go, and go to json
type User struct {
	ID							primitive.ObjectID		`bson:"_id"`
	First_name			*string								`json:"first_name" validate:"required, min=2, max=50"`
	Last_name				*string								`json:"last_name" validate:"required, min=2, max=50"`
	Company					*string								`json:"company" validate:"required"`
	Email						*string								`json:"email" validate:"email, required"`
	Password				*string								`json:"password" validate:"required, min=8"`
	Token						*string								`json:"token"`
	User_Type				*string								`json:"user_type" validate:"required, ut=SADMIN | ut=PADMIN | ut=CADMIN | ut=USER"`
	Refresh_token		*string								`json:"refresh_token"`
	Created_at			time.Time							`json:"created_at"`
	Updated_at			time.Time							`json:"updated_at"`		
	User_id					string								`json:"user_id"`
}