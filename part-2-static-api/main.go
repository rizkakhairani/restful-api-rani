package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id			int		`json:"id" form:"id"`
	Name		string	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages" : "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"));

	if err != nil || id < 1 {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages" : "success get user by id",
				"id": user.Id,
				"name": user.Name,
				"email": user.Email,
			})	
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id < 1{
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	if id <= len(users) {
		var afterDelete[]User
		for i, user := range users {
			if user.Id == id {
				if i == len(users)-1 {
						users = users[:len(users)-1]
						return c.JSON(http.StatusOK, map[string]interface{}{
						"messages": "success delete user by id",
						"users":    users,
					})
				}
				afterDelete = append(users[:i], users[i+1:]...)
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success delete user by id",
					"users":    afterDelete,
				})
			}

		}
	}

	return c.String(http.StatusNotFound, "User not found")
}

// update user by id
func UpdateUserController(c echo.Context) error {
	u := User{}
	c.Bind(&u)
	id, err := strconv.Atoi(c.Param("id"));

	if err != nil || id < 1 {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	for i, user := range users {
		if user.Id == id {
			users[i].Name = u.Name
			users[i].Email = u.Email
			users[i].Password = u.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages" : "success update user by id",
				"id": users[i].Id,
				"name": users[i].Name,
				"email": users[i].Email,
			})	
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":		user,
	})
}

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}