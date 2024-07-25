// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var MergeServersInput serverpb.MergeserversRequest

var MergeServersFromFile string










func init() {
	ServerServiceCmd.AddCommand(MergeServersCmd)
	
	
	MergeServersCmd.Flags().StringVar(&MergeServersInput.Name, "name", "", "Required. The name of the server we're adding disks to.")
	
	MergeServersCmd.Flags().StringVar(&MergeServersInput.OtherServer, "other_server", "", "Required. The name of the server we're removing disks from...")
	
	
	
	MergeServersCmd.Flags().StringVar(&MergeServersFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var MergeServersCmd = &cobra.Command{
  Use:   "merge-servers",
  Short: "Merges two servers by adding all disks from the...",
	Long: "Merges two servers by adding all disks from the server named  `other_server_name` to server `name`, and deletes  `other_server_name`. Returns the...",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if MergeServersFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
			cmd.MarkFlagRequired("other_server")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if MergeServersFromFile != "" {
			in, err = os.Open(MergeServersFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &MergeServersInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "MergeServers", &MergeServersInput)
		}
		resp, err := ServerClient.MergeServers(ctx, &MergeServersInput)
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


