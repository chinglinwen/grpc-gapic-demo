// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var MoveDiskInput serverpb.MoveDiskRequest

var MoveDiskFromFile string










func init() {
	ServerServiceCmd.AddCommand(MoveDiskCmd)
	
	
	MoveDiskCmd.Flags().StringVar(&MoveDiskInput.Name, "name", "", "Required. The name of the book to move.")
	
	MoveDiskCmd.Flags().StringVar(&MoveDiskInput.OtherServerName, "other_server_name", "", "Required. The name of the destination server.")
	
	
	
	MoveDiskCmd.Flags().StringVar(&MoveDiskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var MoveDiskCmd = &cobra.Command{
  Use:   "move-disk",
  Short: "Moves a book to another server, and returns the...",
	Long: "Moves a book to another server, and returns the new book. The book  id of the new book may not be the same as the original book.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if MoveDiskFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
			cmd.MarkFlagRequired("other_server_name")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if MoveDiskFromFile != "" {
			in, err = os.Open(MoveDiskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &MoveDiskInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "MoveDisk", &MoveDiskInput)
		}
		resp, err := ServerClient.MoveDisk(ctx, &MoveDiskInput)
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


