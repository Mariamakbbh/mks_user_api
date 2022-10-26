package userLogin

import (
	"github.com/gofiber/fiber/v2"
)

// Repo provides interface for Repo (database handling) implementation.
type Repo interface {
	AddNewUser(a *UserDetails) error
	EmailExist(a *UserDetails) error
	PasswordExist(a *UserDetails) error
	// ChangePassword(a *UserDetails) error
	// DeleteExistingUser(c *fiber.Ctx) error
	// GetUserDetails(c *fiber.Ctx) error
	// UpdateUserDetails(c *fiber.Ctx) error

}

// Service provides interface for Service (Business logic) implementation.
type Service interface {
	// AuthUserDetails(c *fiber.Ctx) error
	AddNewUser(a *UserRequestBody) error
	AuthEmail(a *UserRequestBody) error
	AuthPassword(a *UserRequestBody) error
	// ChangePassword(c *UserRequestBody) error
}

// BlueprintInterface/Handler provides interface for BlueprintInterface (request handling) implementation.
type BlueprintInterface interface {
	Routes(app *fiber.App)
	AddUser(c *fiber.Ctx) error
	AuthoriseUser(c *fiber.Ctx) error
	// ForgotPassword(c *fiber.Ctx) error
	// DeleteUser(c *fiber.Ctx) error
	// GetUser(c *fiber.Ctx) error
	// UpdateUser(c *fiber.Ctx) error
}
