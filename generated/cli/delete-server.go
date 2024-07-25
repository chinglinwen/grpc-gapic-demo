// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var DeleteServerInput serverpb.DeleteServerRequest

var DeleteServerFromFile string








func init() {
	ServerServiceCmd.AddCommand(DeleteServerCmd)
	
	
	DeleteServerCmd.Flags().StringVar(&DeleteServerInput.Name, "name", "", "Required. The name of the server to delete.")
	
	
	
	DeleteServerCmd.Flags().StringVar(&DeleteServerFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var DeleteServerCmd = &cobra.Command{
  Use:   "delete-server",
  Short: "Deletes a server. Returns NOT_FOUND if the server...",
	Long: "Deletes a server. Returns NOT_FOUND if the server does not exist.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if DeleteServerFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if DeleteServerFromFile != "" {
			in, err = os.Open(DeleteServerFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &DeleteServerInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "DeleteServer", &DeleteServerInput)
		}
		err = ServerClient.DeleteServer(ctx, &DeleteServerInput)
		if err != nil {
			return err
		}
		
		return err
  },
}


