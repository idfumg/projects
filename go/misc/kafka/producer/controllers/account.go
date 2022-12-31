package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{
		accountService: accountService,
	}
}

func (ac accountController) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	id, err := ac.accountService.OpenAccount(command)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "open account success",
		"id":      id,
	})
}

func (ac accountController) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = ac.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "deposit fund success",
	})
}

func (ac accountController) WithdrawFund(c *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = ac.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "withraw fund success",
	})
}

func (ac accountController) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = ac.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "close account success",
	})
}
