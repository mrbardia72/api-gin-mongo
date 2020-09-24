# api-gin-mongo
this is a api with gin and mongodb

## structur main
```
.
├── app.go
├── controllers
│   ├── auth
│   │   ├── login.go
│   │   ├── PasswordReset.go
│   │   ├── RefreshToken.go
│   │   ├── ResetLink.go
│   │   ├── signup.go
│   │   ├── VerifyAccount.go
│   │   └── VerifyLink.go
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── go.mod
├── go.sum
├── helpers
│   ├── bcrypt.go
│   └── error.go
├── Makefile
├── models
│   ├── config.go
│   └── user.go
├── README.md
├── repository
│   └── user
│       ├── alluser.go
│       ├── GetUserByEmail.go
│       ├── removeuser.go
│       ├── signup.go
│       ├── updateuser.go
│       └── UpdateUserPass.go
├── routers
│   └── router.go
└── services
    ├── jwt.go
    └── SendMail.go
 ```