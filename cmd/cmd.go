package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/d-tsuji/clipboard"
	"github.com/robbailey3/jwt-tool/jwtTool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	userId    string
	secretKey string
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		secretKey = viper.GetString("secretKey")
		fmt.Printf("SecretKey = %s\n", secretKey)
		token, err := jwtTool.CreateToken(userId, secretKey)
		if err != nil {
			log.Fatal(err.Error())
		}
		loginStr := fmt.Sprintf("Connect.login({ \"userId\": \"%s\", \"jwt\": \"%s\" });", userId, token)
		if err := clipboard.Set(loginStr); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Login string copied to clipboard")
		fmt.Println(loginStr)
	},
}

func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&userId, "userId", GenerateRandomUserId(), "The user id of the user you want to login")
	rootCmd.PersistentFlags().StringVar(&secretKey, "secretKey", "secret", "The Identity Verification Secret taken from Web Assistant")
	viper.BindPFlag("secretKey", rootCmd.PersistentFlags().Lookup("secretKey"))
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".jwt-tool")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func GenerateRandomUserId() string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
