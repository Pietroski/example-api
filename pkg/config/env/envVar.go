package env
//
//import (
//	"github.com/joho/godotenv"
//	"github.com/kelseyhightower/envconfig"
//	"log"
//	"os"
//)
//
//const (
//	// isEnvVarsSet is used for checking whether or not
//	// the project's env variables are already in the
//	// current project's system environment.
//	isEnvVarsSet = "IS_ENV_VARS_SET"
//)
//
//var (
//	envPath string
//)
//
//func init() {
//	rootDir, err := SearchRootDir()
//	if err != nil {
//		log.Fatalf("error on init: %s", err.Error())
//	}
//
//	envPath = rootDir + "/.env"
//}
//
//// DotEnvInit Initializes the project's .env file at the project's root directory if the variables
//// are not yet in the system environment. This check is done by checking of isEnvVarsSet variable.
///*
//	As the project is being run on different environment and
//	the main.go file is not located at the project's root directory,
//	a solution was needed where to search for the very possible
//	project's root directory if no .env variable is set in the system environment.
//	Github and docker use different $GOPATH locations. This would not be a problem if
//	both of them would place this projects binaries and sources files
//	at /$GOPATH/src/KyrosId/acuris-api directory.
//	The fact is that both place the executables in different places each.
//	Accounting our local environment, there is at least three distinct environments
//	for this project.
//	To not hard-code the path from main.go, the solution is to search for its root path.
//*/
//func DotEnvInit() {
//	_, ok := os.LookupEnv(isEnvVarsSet)
//	if !ok {
//		if err := godotenv.Load(os.ExpandEnv(envPath)); err != nil {
//			log.Fatalln("Error loading .env file")
//		}
//	}
//}
//
//// DotEnvFiller fills the given struct with its corresponding env variables.
//func DotEnvFiller(envStruct interface{}) {
//	err := envconfig.Process("channel_balancer", envStruct)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//}
