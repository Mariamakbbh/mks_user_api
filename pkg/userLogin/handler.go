package userLogin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Blueprint struct for db service and auth provider.
type Blueprint struct {
	service Service
}

func (b Blueprint) Routes(app *fiber.App) {

	routes := app.Group("/user_API")

	routes.Post("/AddUser", b.AddUser)
	routes.Post("/userAuth", b.AuthoriseUser)
	// routes.Get("/forgotAuth", b.ForgotPassword)

}

// Add new user request (maybe add timestamp to request body).
func (b Blueprint) AddUser(c *fiber.Ctx) error {
	AddNewUserbody := UserRequestBody{}

	// parse body, attach to AddUserRequestBody struct
	if err := c.BodyParser(&AddNewUserbody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//Authorise user Email
	if err := b.service.AuthEmail(&AddNewUserbody); err != nil {

		if err := b.service.AddNewUser(&AddNewUserbody); err != nil {
			fmt.Println("Error in service.AddNewUser: ", err)
			return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "Database Error", Message: "Unable to add user"})
		}
	} else {
		return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "Bad", Message: "User already exists"})
	}

	return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "OK", Message: "User Added"})
}

// Get existing user request (should return a single user).
func (b Blueprint) AuthoriseUser(c *fiber.Ctx) error {
	AddNewUserbody := UserRequestBody{}

	// parse body, attach to AddUserRequestBody struct
	if err := c.BodyParser(&AddNewUserbody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//Authorise user Email
	if err := b.service.AuthEmail(&AddNewUserbody); err != nil {
		fmt.Println("Email Not Found: ", err)
		return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "Incorrect", Message: "User Email Wrong"})
	}

	//Authorise user Password
	if err := b.service.AuthPassword(&AddNewUserbody); err != nil {
		fmt.Println("Password Not Found: ", err)
		return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "Incorrect", Message: "User Password Wrong"})
	}

	return c.Status(fiber.StatusCreated).JSON(&RequestResponse{Status: "OK", Message: "User Found / Authorised"})
}

// // Forgot Password
// func (b Blueprint) ForgotPassword(c *fiber.Ctx) error {
// 	AddNewUserbody := UserRequestBody{}

// 	// parse body, attach to AddUserRequestBody struct
// 	if err := c.BodyParser(&AddNewUserbody); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	if err := b.service.AuthUser(&AddNewUserbody); err != nil {
// 		fmt.Println("Error in service.AddNewUser: ", err)
// 		return err
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(&AddNewUserbody)

// }

func NewBlueprint(service Service) *Blueprint {
	return &Blueprint{
		service: service,
	}
}
