package main

import "github.com/labstack/echo/v4"

func main()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200,map[string]string{"message":"hello"})
	})
	e.Logger.Fatal(e.Start(":8080"))
// 	return err
// 		}

// 		newUser := new(User)
// 		newUser.ID = uuid.New()
// 		newUser.Name = request.Name
// 		newUser.Email = request.Email
// 		newUser.Password = request.Password

// 		return c.JSON(200, map[string]interface{}{
// 			"message": "successfully create a user",
// 			"data":    newUser,
// 		})
// 	})

// 	e.Logger.Fatal(e.Start(":8080"))
// }
}