// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var DeleteDiskInput serverpb.DeleteDiskRequest

var DeleteDiskFromFile string








func init() {
	ServerServiceCmd.AddCommand(DeleteDiskCmd)
	
	
	DeleteDiskCmd.Flags().StringVar(&DeleteDiskInput.Name, "name", "", "Required. The name of the book to delete.")
	
	
	
	DeleteDiskCmd.Flags().StringVar(&DeleteDiskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var DeleteDiskCmd = &cobra.Command{
  Use:   "delete-disk",
  Short: "Deletes a book. Returns NOT_FOUND if the book...",
	Long: "Deletes a book. Returns NOT_FOUND if the book does not exist.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if DeleteDiskFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if DeleteDiskFromFile != "" {
			in, err = os.Open(DeleteDiskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &DeleteDiskInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "DeleteDisk", &DeleteDiskInput)
		}
		err = ServerClient.DeleteDisk(ctx, &DeleteDiskInput)
		if err != nil {
			return err
		}
		
		return err
  },
}


