package vars

import "fmt"

// var _ = godotenv.Load(".env")

// var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 	"root",
// 	"P4ssw0rd.123",
// 	"34.101.123.137",
// 	"3306",
// 	"padiplace_pos")

var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	"root",
	"root",
	"localhost",
	"3306",
	"padiplace_pos")

const AllowedCORSDomain = "http://localhost"
