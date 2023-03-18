package constants

import "os"

var HOST = os.Getenv("HOST")

const (
	APIServerPort  = "8080"
	GrpcServerPort = "9000"
)
