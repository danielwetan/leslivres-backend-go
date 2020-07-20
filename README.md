This is a backend app for`LesLivres` library. Built with Go, Gin and MySQL which has [features](https://github.com/danielwetan/leslivres-backend-go#features) such as login / register using JWT, pasword hashing, CORS, etc. 

## :memo: Table Of Content
* [Prerequisites](https://github.com/danielwetan/leslivres-backend-go#prerequisites)
* [Installation](https://github.com/danielwetan/leslivres-backend-go#installation)
* [Features](https://github.com/danielwetan/leslivres-backend-go#features)
* [Built wtih](https://github.com/danielwetan/leslivres-backend-go#features)
* [Author](https://github.com/danielwetan/leslivres-backend-go#author)
* [License](https://github.com/danielwetan/leslivres-backend-go#license)

## Prerequisites
- Go installed in the local machine
- MySQL intalled on the local machine (ex. XAMPP)
## Installation
1. Clone this repository:
    `git clone https://github.com/danielwetan/leslivres-backend-go`
2. Move to installation folder:
    `cd leslivres-backend-go`
3. Start XAMPP
4. Database configuration:
    * Open http://localhost/phpmyadmin in the browser
    * Create a new table with the name `leslivres`
    * Import database to current table, select `leslivres.sql` file from project folder
5. Start the server:
    `go run main.go`
6. Run app in the browser on the port http://localhost:3000

## Features
- [x] CRUD
- [x] Search, Sort, Pagination
- [x] Cors
- [x] Login/Register with JWT
- [x] Password hashing

## Built with
- [Go](https://golang.org/) - Programming Language
- [Gin](https://github.com/gin-gonic/gin/) - HTTP framework
- [MySQL](https://www.mysql.com/) Database
<!-- - [JWT](https://jwt.io/) - Login/Register authentication
- [Bcrypt](https://github.com/kelektiv/node.bcrypt.js) - Password Hashing
 -->
 
## Author
- [Daniel Saputra](https://www.linkedin.com/in/danielwetan/)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/danielwetan/leslivres-backend-go/blob/master/LICENSE) file for details
