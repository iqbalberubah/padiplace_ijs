package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"password",
		"localhost",
		"3306",
		"padiplace_ijs")
)

const AllowedCORSDomain = "http://localhost"
