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

var ListserversInput serverpb.ListServersRequest

var ListserversFromFile string










func init() {
	ServerServiceCmd.AddCommand(ListserversCmd)
	
	
	ListserversCmd.Flags().Int32Var(&ListserversInput.PageSize, "page_size", 10, "Default is 10. Requested page size. Server may return fewer...")
	
	ListserversCmd.Flags().StringVar(&ListserversInput.PageToken, "page_token", "", "A token identifying a page of results the server...")
	
	
	
	ListserversCmd.Flags().StringVar(&ListserversFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var ListserversCmd = &cobra.Command{
  Use:   "listservers",
  Short: "Lists servers. The order is unspecified but...",
	Long: "Lists servers. The order is unspecified but deterministic. Newly created  servers will not necessarily be added to the end of this list.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if ListserversFromFile == "" {
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if ListserversFromFile != "" {
			in, err = os.Open(ListserversFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &ListserversInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "Listservers", &ListserversInput)
		}
		iter := ServerClient.Listservers(ctx, &ListserversInput)
		
		
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


