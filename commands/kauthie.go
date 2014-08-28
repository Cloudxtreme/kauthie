// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kiasaki/kauthie/app"
	"github.com/kiasaki/kauthie/bg"
)

var CfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	//viper.SetDefault("key", "val")

	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/.kauthie/config.yaml)")

	rootCmd.PersistentFlags().String("mongodb_uri", "mongodb://localhost:27017/", "Uri to connect to mongoDB")
	viper.BindPFlag("mongodb_uri", rootCmd.PersistentFlags().Lookup("mongodb_uri"))

	serverCmd.Flags().Int("port", 1138, "Port to run Kauthie app server on")
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kauthie/")
	viper.AddConfigPath("$HOME/.kauthie/")

	viper.ReadInConfig()
}

// Root command
var rootCmd = &cobra.Command{
	Use:   "kauthie",
	Short: "Kauthie is an account/user management implementation suitable for your next startup idea",
	Long: `Kauthie is an account/user management implementation, it contains the necessary oauth2 endpoints
	the forgot password paths a small api and all the account managment/crud needed to operate.`,
	Run: rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	port := viper.GetString("port")
	go app.Server(port)
	go bg.Worker()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

// Server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the app server",
	Long:  `Starts Kauthie app server.`,
	Run:   serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	port := viper.GetString("port")
	app.Server(port)
}

// Worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start a Kauthie worker",
	Long:  `Starts a trusty Kauthie worker.`,
	Run:   workerRun,
}

func workerRun(cmd *cobra.Command, args []string) {
	bg.Worker()
}

func addCommands() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(workerCmd)
}

func Execute() {
	addCommands()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
