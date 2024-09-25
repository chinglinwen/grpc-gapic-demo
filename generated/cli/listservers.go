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

var ListServersInput serverpb.ListServersRequest

var ListServersFromFile string










func init() {
	ServerServiceCmd.AddCommand(ListServersCmd)
	
	
	ListServersCmd.Flags().Int32Var(&ListServersInput.PageSize, "page_size", 10, "Default is 10. Requested page size. Server may return fewer...")
	
	ListServersCmd.Flags().StringVar(&ListServersInput.PageToken, "page_token", "", "A token identifying a page of results the server...")
	
	
	
	ListServersCmd.Flags().StringVar(&ListServersFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var ListServersCmd = &cobra.Command{
  Use:   "listservers",
  Short: "Lists servers. The order is unspecified but...",
	Long: "Lists servers. The order is unspecified but deterministic. Newly created  servers will not necessarily be added to the end of this list.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if ListServersFromFile == "" {
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if ListServersFromFile != "" {
			in, err = os.Open(ListServersFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &ListServersInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "ListServers", &ListServersInput)
		}
		iter := ServerClient.ListServers(ctx, &ListServersInput)
		
		
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


