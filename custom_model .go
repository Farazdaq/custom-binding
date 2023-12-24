package main

type account struct {
	AcctiVation bool    `form:"boolField" binding:"required"`
	FullName    *string `form:"pointerField"`
	ID          int     `form:"intField" binding:"required"`
	Email       string  `form:"stringField" binding:"required"`
	Password    string  `form:"stringField" binding:"required"`
	AccountType rune    `form:"runeField" binding:"required"`
}
