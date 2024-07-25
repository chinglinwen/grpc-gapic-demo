// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "google.golang.org/api/iterator"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var ListDisksInput serverpb.ListDisksRequest

var ListDisksFromFile string












func init() {
	ServerServiceCmd.AddCommand(ListDisksCmd)
	
	
	ListDisksCmd.Flags().StringVar(&ListDisksInput.Parent, "parent", "", "Required. The name of the server whose disks we'd like to...")
	
	ListDisksCmd.Flags().Int32Var(&ListDisksInput.PageSize, "page_size", 10, "Default is 10. Requested page size. Server may return fewer...")
	
	ListDisksCmd.Flags().StringVar(&ListDisksInput.PageToken, "page_token", "", "A token identifying a page of results the server...")
	
	
	
	ListDisksCmd.Flags().StringVar(&ListDisksFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var ListDisksCmd = &cobra.Command{
  Use:   "list-disks",
  Short: "Lists disks in a server. The order is unspecified...",
	Long: "Lists disks in a server. The order is unspecified but deterministic. Newly  created disks will not necessarily be added to the end of this list. ...",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if ListDisksFromFile == "" {
			
			
			cmd.MarkFlagRequired("parent")
			
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if ListDisksFromFile != "" {
			in, err = os.Open(ListDisksFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &ListDisksInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "ListDisks", &ListDisksInput)
		}
		iter := ServerClient.ListDisks(ctx, &ListDisksInput)
		
		
		// populate iterator with a page
		_, err = iter.Next()
		if err != nil && err != iterator.Done {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(iter.Response)
		

		
		
		return err
  },
}


