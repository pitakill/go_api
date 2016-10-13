package main

type Person struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
	Registered int    `json:"registered"`
}

type User struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Type       string `json:"type"`
	Twitter    string `json:"twitter"`
}

type Post struct {
	Id        int    `json:"id"`
	Author_Id int    `json:"author_id"`
	Title     int    `json:"title"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	Created   int    `json:"created"`
}
