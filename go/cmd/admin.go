/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"os"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin commands",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var makeAdminCmd = &cobra.Command{
	Use:   "make",
	Short: "Gives a user admin privileges",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		role := user.RoleAdmin
		if cmd.Flag("remove").Value.String() == "true" {
			role = user.RoleUser
		}
		println(cmd.Flag("remove").Value.String())
		entUser, err := database.NewDefaultDbManagers().UserManager.SetRole(args[0], role)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("User %s is now an `%v`\n", entUser.Name, entUser.Role)
	},
}

var listAdminsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all admins",
	Run: func(cmd *cobra.Command, args []string) {
		admins, err := database.NewDefaultDbManagers().UserManager.GetAdmins()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		text := "Admins:\n"
		for _, admin := range admins {
			text += fmt.Sprintf("%s\n", admin.Name)
		}
		fmt.Print(text)
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(makeAdminCmd)
	adminCmd.AddCommand(listAdminsCmd)

	makeAdminCmd.PersistentFlags().Bool("remove", false, "Remove admin privileges")
}
