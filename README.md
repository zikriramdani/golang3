# CRUD Article By Golang

This project was generated with [`Golang`] version go1.16.6 darwin/amd64.

- go run main.go
- go tidy
- go build

## Development server

Run `go run main.go` for a dev server. Navigate to `http://localhost:8080/`. The app will automatically reload if you change any of the source files.

## Cara menjalankan Aplikasi : 

- Simpan Project di /htdocs (kalau pake xampp)
- import database (db_golang3.sql / automigrate(Buatlah database dahulu dengan nama db_golang3.sql)
- buka folder project, copykan .env-example, menjadi .env
- isi DB_DATABASE, DB_USERNAME, DB_PASSWORD, sesuaikan dengan settingan database kamu pada config.go
- di dalam directory project buka terminal, ketikan "go run main.go"
- buka browser, ketikan url "localhost:8080"

# Url API
- localhost:8080/api/v1/article/add ===> createArticle,
- localhost:8080/api/v1/articleList ===> readArticle,
- localhost:8080/api/v1/articleList/:id ===> readArticle ByID,
- localhost:8080/api/v1/article/:id ===> updateArticle ByID,
- localhost:8080/api/v1/article/:id ===> deleteArticle ByID
- localhost:8080/api/v1/articleList?limits=2&page=1 ===> paginationArticle;

Document Postman : `https://documenter.getpostman.com/view/8820668/UVXdMyCy`
