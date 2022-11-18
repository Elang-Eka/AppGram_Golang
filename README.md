
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
---
1. http://localhost:8080/users/register 	(POST) 	(formKey: username,password(value 6 character),age,email)
![Image Register](https://github.com/Elang-Eka/Image-MyGram/blob/master/Auth/Register.jpg)
2. http://localhost:8080/users/login 	(POST)	(formKey: email,password)
![Image Login](https://github.com/Elang-Eka/Image-MyGram/blob/master/Auth/Login.jpg)


User Router
---
1. http://localhost:8080/users 	(GET)
![Image Get User](https://github.com/Elang-Eka/Image-MyGram/blob/master/User/Get_User.jpg)
2. http://localhost:8080/users 	(PUT) 	(formKey: username,email,age)
![Image Update User](https://github.com/Elang-Eka/Image-MyGram/blob/master/User/Update_User.jpg)
3. http://localhost:8080/users 	(DELETE)
![Image Delete User](https://github.com/Elang-Eka/Image-MyGram/blob/master/User/Delete_User.jpg)

Social Media Router
---
1. http://localhost:8080/socialmedias 			(GET)
![Image Get Social Media](https://github.com/Elang-Eka/Image-MyGram/blob/master/Social_Media/Get_Social_Media.jpg)
2. http://localhost:8080/socialmedias 			(POST) 	(formKey: name,social_media_url)
![Image Create Social Media](https://github.com/Elang-Eka/Image-MyGram/blob/master/Social_Media/Create_Social_Media.jpg)
3. http://localhost:8080/socialmedias/:socialMediaId 	(PUT) 	(formKey: name,social_media_url)
![Image Update Social Media](https://github.com/Elang-Eka/Image-MyGram/blob/master/Social_Media/Update_Social_Media.jpg)
4. http://localhost:8080/socialmedias/:socialMediaId 	(DELETE)
![Image Delete Social Media](https://github.com/Elang-Eka/Image-MyGram/blob/master/Social_Media/Delete_Social_Media.jpg)

Photo Router
---
1. http://localhost:8080/photos 		(POST) 	(formKey: photo_url,title,caption)
![Image Create Photo](https://github.com/Elang-Eka/Image-MyGram/blob/master/Photo/Create_photo.jpg)
2. http://localhost:8080/photos 		(GET)
![Image Get Photo](https://github.com/Elang-Eka/Image-MyGram/blob/master/Photo/Get_photo.jpg)
3. http://localhost:8080/photos/:photoId 	(PUT) 	(formKey: photo_url,title,caption)
![Image Update Photo](https://github.com/Elang-Eka/Image-MyGram/blob/master/Photo/Update_photo.jpg)
4. http://localhost:8080/photos/:photoId 	(DELETE)
![Image Delete Photo](https://github.com/Elang-Eka/Image-MyGram/blob/master/Photo/Delete_photo.jpg)

Comment Router
---
1. http://localhost:8080/comments 			(POST) 	(formKey:message,photo_id)
![Image Create Comment](https://github.com/Elang-Eka/Image-MyGram/blob/master/Comment/Create_Comment.jpg)
2. http://localhost:8080/comments/:photoId 	(GET)
![Image Get Comment](https://github.com/Elang-Eka/Image-MyGram/blob/master/Comment/Get_Comment.jpg)
3. http://localhost:8080/comments/:commentId 	(PUT) 	(formKey:message)
![Image Update Comment](https://github.com/Elang-Eka/Image-MyGram/blob/master/Comment/Update_Comment.jpg)
4. http://localhost:8080/comments/:commentId 	(DELETE)
![Image Delete Comment](https://github.com/Elang-Eka/Image-MyGram/blob/master/Comment/Delete_Comment.jpg)
