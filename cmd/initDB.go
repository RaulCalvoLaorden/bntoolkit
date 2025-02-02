package cmd

import (
	"fmt"
	"log"

	"github.com/r4ulcl/bntoolkit/utils"
	"github.com/spf13/cobra"
)

// initDBCmd represents the initDB command
var initDBCmd = &cobra.Command{
	Use:   "initDB",
	Short: "Create the database and its tables",
	Long: `Create the database and it's tables. This command is required the first time the database is connected.
Recommended DB:
	mkdir ~/postgres #or any folder to store data
	sudo docker stop hashpostgres ; sudo docker rm hashpostgres #delete if it exists
	sudo docker run -d -p 5432:5432 --mount type=bind,source=$HOME/postgres/,target=/var/lib/postgresql/data --name hashpostgres -e POSTGRES_PASSWORD=postgres99 postgres
For example:
	bntoolkit initDB`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initDB called")

		database := "hash"
		err := utils.CreateDb(cfgFile, database, debug, verbose)
		if err != nil {
			log.Fatal(err)
		}

		db, err := utils.ConnectDb(cfgFile, debug, verbose)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = utils.InitDB(db, debug, verbose)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.InsertProject(db, debug, verbose, projectName) //if return error exists
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initDBCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initDBCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
