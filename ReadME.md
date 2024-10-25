# Auth grpc Server
A GRPC server that uses grpc to communicate and use Redis to store data of user.

### Quick Start
To quicly run server, make sure you have properly installed go on your local then executes below commands.

```
git clone https://github.com/kumarvikramshahi/auth-grpc-server.git
```

```
cd auth-grpc-server
```
```
go mod download
```
```
go run main.go dev
```
**Note**: *Server support Reflection, so you can automatically list endpoints in POSTMAN after starting server.(by default server runs at localhost:8000)*

### Architecture
Used Hexagonal Architecture (Port-Adapter) to build the web server.

* `configs` - contains configs & env vars related logics
* `env` - contains env vars files.
* `core` - DB connections
* `pkg` - here the main application logics relies
  
![image](https://github.com/user-attachments/assets/ff40714e-f7e0-45c3-a545-b38712521256)


### Service have two endpoints for client:

####  `SignUp/SignUpUser` = signup user
Example request:
```
{
    "email": "vikram1",
    "password": "sdfjshdfkjsa",
    "name": "sldk"
}
```
Example response:
```
{
    "data": {
        "message": "User created"
    }
}
```

#### `LogIn/LogInUser` = login user
Example request:
```
{
    "email": "vikram",
    "password": "sdfjshdfkjsa",
}
```
Example response:
```
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZpa3JhbTEiLCJleHBpcnkiOjE3Mjk5MjA3OTl9.ULei1kVLSekoklKe279ZjZOdFqoqFW5SULBlO0pX8KI",
        "expiry_timestamp": "1729920799"
    }
}
```

