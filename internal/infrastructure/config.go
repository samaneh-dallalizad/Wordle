package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

// ApplicationEnvironment is the required application environment variables.
type ApplicationEnvironment struct {
	ApplicationMode   string
	APIURL            string
	MobileHeaderLogo  string
	DesktopHeaderLogo string
	WebAssetsFolder   string
	BindAddr          string
}

// ConfigSetup will prepare and setup the viper instance to the correct config file.
func ConfigSetup(configName, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
}

func GetConfig() ApplicationEnvironment {
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s", err)
	}

	ValidateVariablesAreSet([]string{
		"ApplicationMode",
		"APIURL",
		"DesktopHeaderLogo",
		"MobileHeaderLogo",
		"WebAssetsFolder",
		"BindAddr",
	})

	return ApplicationEnvironment{
		ApplicationMode:   viper.GetString("ApplicationMode"),
		APIURL:            viper.GetString("APIURL"),
		MobileHeaderLogo:  viper.GetString("MobileHeaderLogo"),
		DesktopHeaderLogo: viper.GetString("DesktopHeaderLogo"),
		WebAssetsFolder:   viper.GetString("WebAssetsFolder"),
		BindAddr:          viper.GetString("BindAddr"),
	}
}

// ValidateVariablesAreSet will assert the existence of each variable,
// and kill the application when a wanted variable does not exist in the config.
func ValidateVariablesAreSet(variables []string) {
	for i := range variables {
		if !viper.IsSet(variables[i]) {
			log.Fatalf("%s variable was not set!\nAborting application start!", variables[i])
		}
	}
}
