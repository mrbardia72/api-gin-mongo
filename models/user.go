package models

import (
    "gopkg.in/mgo.v2/bson"
)

// User defines user object structure
type User struct {
    ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
    Name       string        `json:"name" bson:"name"`
    Email      string        `json:"email" bson:"email"`
    Password   string        `json:"password" bson:"password"`
    IsVerified bool          `json:"is_verified" bson:"is_verified"`
}
