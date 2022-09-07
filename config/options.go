package config

var (
	//MasterHost is for the operators to connect to
	MasterHost string
	//MasterEnabled is to check if the Master Host is enabled in the config
	MasterEnabled bool

	//MaxLoginTries is how many incorrect logins are allowed until the user is kicked
	MaxLoginTries int

	//CapchaEnabled checks if the capcha system should be enabled
	CapchaEnabled bool
	//CapchaMaxTries is a max tries limiter until the user is disconnected
	CapchaMaxTries int
	//ForceMfa is a check if the cnc should force mfa onto the user
	ForceMfa bool
	//MaxMfaTries is a check for how many times can the user fail the mfa
	MaxMfaTries int

	//SQLHost is used to connect to the DB this should be localhost
	SQLHost string
	//SQLUser should not be root
	SQLUser string
	//SQLPassword is used to authenticate with the user
	SQLPassword string
	//SQLDatabase is the database to use
	SQLDatabase string
	//Started
	Started int64

	//Version number of the CNC
	Version = "API Build v2.2.0"
)
