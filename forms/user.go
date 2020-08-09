package forms

// SignupUserCommand defines user form struct
type SignupUserCommand struct {
    // binding:"required" ensures that the field is provided
    Name string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// LoginUserCommand defines user login form struct
type LoginUserCommand struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}