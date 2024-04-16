/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	syslog "log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/hiimnhan/post-office/ui"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "post-office",
		Short: "An API platform for building and managing APIs",
	}
	defaultConfigFile = "./config-example.toml"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func createModel(cfgFile string, debug bool) (tea.Model, *os.File) {
	var loggerFile *os.File

	if debug {
		var fileErr error
		newConfigFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if fileErr == nil {
			log.SetOutput(newConfigFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging to debug.log")
		} else {
			loggerFile, _ = tea.LogToFile("debug.log", "debug")
			syslog.Print("Failed setting up logging", fileErr)
		}
	}

	m := ui.NewModel(cfgFile)
	m.InitStyles()

	return m, loggerFile
}

func initConfig() {
	fmt.Println("Running")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Cannot get user home directory", err)
			return
		}

		path := homeDir + "/.config/post-office"
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(path)

		// if file config.yml does not exist, create it
		if _, err := os.Stat(path + "/config.toml"); os.IsNotExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				log.Fatal("Cannot create config directory", err)
				return
			}

			sourceFile, err := os.Open(defaultConfigFile)
			if err != nil {
				log.Fatal("Cannot open default config file", err)
				return
			}
			defer sourceFile.Close()

			file, err := os.Create(path + "/config.toml")
			if err != nil {
				log.Fatal("Cannot create config file", err)
				return
			}
			defer file.Close()
			_, err = io.Copy(file, sourceFile)
			if err != nil {
				log.Fatal("Cannot copy default config file", err)
				return
			}

		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		cfgFile = viper.ConfigFileUsed()
	} else {
		log.Debug("No config file found")
		fmt.Println(err)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		"use this configuration file (default is $XDG_CONFIG_HOME/post-office/config.toml)",
	)

	err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	if err != nil {
		log.Fatal("Cannot bind config flag", err)
	}

	rootCmd.Flags().Bool(
		"debug",
		false,
		"passing this flag will allow writing debug output to debug.log",
	)

	rootCmd.Run = func(_ *cobra.Command, _ []string) {
		initConfig()
		debug, err := rootCmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal("Cannot parse debug flag", err)
		}

		lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

		m, logger := createModel(cfgFile, debug)

		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program", err)
			os.Exit(1)
		}

		if logger != nil {
			defer logger.Close()
		}
	}

}
