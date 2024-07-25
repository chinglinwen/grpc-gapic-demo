// Code generated. DO NOT EDIT.




package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"golang.org/x/oauth2"
	
	gapic "core.wcloud.io/generated/gapicgen/server"
	
)

var ServerConfig *viper.Viper
var ServerClient *gapic.Client
var ServerSubCommands []string = []string{
	"create-server",
	"get-server",
	"listservers",
	"delete-server",
	"merge-servers",
	"create-disk",
	"get-disk",
	"list-disks",
	"delete-disk",
	"update-disk",
	"move-disk",
	
}

func init() {
	rootCmd.AddCommand(ServerServiceCmd)

	ServerConfig = viper.New()
	ServerConfig.SetEnvPrefix("APIDEMO_SERVER")
	ServerConfig.AutomaticEnv()

	ServerServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use APIDEMO_SERVER_INSECURE. Must be used with \"address\" option")
	ServerConfig.BindPFlag("insecure", ServerServiceCmd.PersistentFlags().Lookup("insecure"))
	ServerConfig.BindEnv("insecure")

	ServerServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use APIDEMO_SERVER_ADDRESS.")
	ServerConfig.BindPFlag("address", ServerServiceCmd.PersistentFlags().Lookup("address"))
	ServerConfig.BindEnv("address")

	ServerServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use APIDEMO_SERVER_TOKEN.")
	ServerConfig.BindPFlag("token", ServerServiceCmd.PersistentFlags().Lookup("token"))
	ServerConfig.BindEnv("token")

	ServerServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use APIDEMO_SERVER_API_KEY.")
	ServerConfig.BindPFlag("api_key", ServerServiceCmd.PersistentFlags().Lookup("api_key"))
	ServerConfig.BindEnv("api_key")
}

var ServerServiceCmd = &cobra.Command{
	Use:   "server",
	Short: "This API represents a simple digital server. It...",
	Long: "This API represents a simple digital server. It lets you manage Server  resources and Disk resources in the server. It defines the following ...",
	ValidArgs: ServerSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := ServerConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if ServerConfig.GetBool("insecure"){
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := ServerConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := ServerConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		ServerClient, err = gapic.NewClient(ctx, opts...)
		return
	},
}
