// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var CreateServerInput serverpb.CreateServerRequest

var CreateServerFromFile string










func init() {
	ServerServiceCmd.AddCommand(CreateServerCmd)
	
	CreateServerInput.Server = new(serverpb.Server)
	
	
	CreateServerCmd.Flags().StringVar(&CreateServerInput.Server.Name, "server.name", "", "The resource name of the server.  Server names...")
	
	CreateServerCmd.Flags().StringVar(&CreateServerInput.Server.Theme, "server.theme", "", "The theme of the server")
	
	
	
	CreateServerCmd.Flags().StringVar(&CreateServerFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var CreateServerCmd = &cobra.Command{
  Use:   "create-server",
  Short: "Creates a server, and returns the new Server.",
	Long: "Creates a server, and returns the new Server.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if CreateServerFromFile == "" {
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if CreateServerFromFile != "" {
			in, err = os.Open(CreateServerFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &CreateServerInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "CreateServer", &CreateServerInput)
		}
		resp, err := ServerClient.CreateServer(ctx, &CreateServerInput)
		if err != nil {
			return err
		}
		
		
		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)
		

		
		
		return err
  },
}


