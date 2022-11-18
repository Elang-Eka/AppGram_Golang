
# AppMyGram

The App is My Portofolio Final Project Golang by DTS program Hacktiv8 - Scalable Web Service With Golang

## Clone & Running Locally

Create db in mysql "db-mygram"
```sh
$ git clone https://github.com/Elang-Eka/AppGram_Golang.git
$ cd AppGram_Golang
```
Open your VSCode
```sh
$ go run main.go
```
Allow access firewall

Your app should now be running on [localhost:8080](http://localhost:8080/).

## Router Local

Auth user Register and Login
1. http://localhost:8080/users/register 	(POST) 	(formKey: username,password(value 6 character),age,email)
[![Image Register](https://github.com/Elang-Eka/Image-MyGram/blob/master/Auth/Register.jpg)
2. http://localhost:8080/users/login 	(POST)	(formKey: email,password)
[![Image Login](https://github.com/Elang-Eka/Image-MyGram/blob/master/Auth/Login.jpg)

