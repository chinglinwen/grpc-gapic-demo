// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var GetDiskInput serverpb.GetDiskRequest

var GetDiskFromFile string








func init() {
	ServerServiceCmd.AddCommand(GetDiskCmd)
	
	
	GetDiskCmd.Flags().StringVar(&GetDiskInput.Name, "name", "", "Required. The name of the book to retrieve.")
	
	
	
	GetDiskCmd.Flags().StringVar(&GetDiskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var GetDiskCmd = &cobra.Command{
  Use:   "get-disk",
  Short: "Gets a book. Returns NOT_FOUND if the book does...",
	Long: "Gets a book. Returns NOT_FOUND if the book does not exist.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if GetDiskFromFile == "" {
			
			
			cmd.MarkFlagRequired("name")
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if GetDiskFromFile != "" {
			in, err = os.Open(GetDiskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &GetDiskInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "GetDisk", &GetDiskInput)
		}
		resp, err := ServerClient.GetDisk(ctx, &GetDiskInput)
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


