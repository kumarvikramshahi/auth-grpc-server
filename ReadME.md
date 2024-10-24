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

### Service have two endpoints for client:

#### `LogIn/LogInUser` = login user
####  `SignUp/SignUpUser` = signup user
