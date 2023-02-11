package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// The aliasNormalizeFunc() function is used for creating additional aliases for a
// flag—in this case an alias for the --password flag. According to the existing code,
// the --password flag can be accessed as either --pass or --ps.
func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	var str string
	switch n {
	case "pass":
		str = "password"
	case "ps":
		str = "password"
	}
	return pflag.NormalizedName(str)
}

func main() {
	// 	In the preceding code, we create a new flag called name that can also be accessed as
	// -n. Its default value is Mike and its description that appears in the usage of the utility
	// 	is Name parameter.
	pflag.StringP("name", "n", "Mike", "Name parameter")
	pflag.StringP("password", "p", "hardToGuess", "Password")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println(name, password)

	// Reading an Environment variable
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}

	// Setting an Environment variable
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
}
