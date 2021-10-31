package cmd

import (
	"fmt"
	"gql/config"
	"gql/router"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&config.C.Port, "port", "p", "3000", "service port")
}

// server -
func server() {
	app := iris.Default()

	router.SetRouter(app)

	app.Run(iris.Addr(fmt.Sprintf(":%s", config.C.Port)), iris.WithoutInterruptHandler)
}
