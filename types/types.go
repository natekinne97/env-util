package types

import (
	"os"
	"time"
)

type User struct {
	Id         string    `bson:"_id,omitempty", json:"id,omitempty"`
	Email      string    `bson:"email,omitempty", json:"email"`
	Password   string    `bson:"password", json:"password"`
	Firstname  string    `bson:"firstname,omitempty", json:"firstname,omitempty"`
	Role       string    `bson:"role,omitempty", json:"role,omitempty"`
	IsActive   bool      `bson:"isActive,omitempty", json:"isActive,omitempty"`
	CreatedAt  time.Time `bson:"createdAt,omitempty", json:"createdAt"`
	ModifiedAt time.Time `bson:"modifiedAt,omitEmpty", json:"modifiedAt"`
}

type Token struct {
	ID     string `bson:"id,omitempty", json:"id,omitempty"`
	Email  string `bson:"email", json:"email"`
	Expiry int64  `bson:"expiry", json:"expiry"`
}



// Helpers
type EnvUtilHelper struct{}

func (_ EnvUtilHelper) Getenv (str string) string{
	return os.Getenv(str)
}

type MongoClient struct{}

func (_ MongoClient) GetTokenByEmail(email string)(*User, error){
	return nil, nil
}

func (_ MongoClient) InsertToken(email string)error{
	return nil
}
