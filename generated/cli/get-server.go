// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var GetServerInput serverpb.GetServerRequest

var GetServerFromFile string








func init() {
	ServerServiceCmd.AddCommand(GetServerCmd)
	
	
	GetServerCmd.Flags().StringVar(&GetServerInput.Name, "name", "", "Required. The name of the server to retrieve.")
	
	
	
	GetServerCmd.Flags().StringVar(&GetServerFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var GetServerCmd = &cobra.Command{
  Use:   "get-server",
  Short: "Gets a server. Returns NOT_FOUND if the server...",
	Long: "Gets a server. Returns NOT_FOUND if the server does not exist.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if GetServerFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if GetServerFromFile != "" {
			in, err = os.Open(GetServerFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &GetServerInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "GetServer", &GetServerInput)
		}
		resp, err := ServerClient.GetServer(ctx, &GetServerInput)
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


