// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var CreateDiskInput serverpb.CreateDiskRequest

var CreateDiskFromFile string
















func init() {
	ServerServiceCmd.AddCommand(CreateDiskCmd)
	
	CreateDiskInput.Book = new(serverpb.Disk)
	
	
	CreateDiskCmd.Flags().StringVar(&CreateDiskInput.Parent, "parent", "", "Required. The name of the server in which the book is...")
	
	CreateDiskCmd.Flags().StringVar(&CreateDiskInput.Book.Name, "book.name", "", "The resource name of the book.  Disk names have...")
	
	CreateDiskCmd.Flags().StringVar(&CreateDiskInput.Book.Author, "book.author", "", "The name of the book author.")
	
	CreateDiskCmd.Flags().StringVar(&CreateDiskInput.Book.Title, "book.title", "", "The title of the book.")
	
	CreateDiskCmd.Flags().BoolVar(&CreateDiskInput.Book.Read, "book.read", false, "Value indicating whether the book has been read.")
	
	
	
	CreateDiskCmd.Flags().StringVar(&CreateDiskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var CreateDiskCmd = &cobra.Command{
  Use:   "create-disk",
  Short: "Creates a book, and returns the new Disk.",
	Long: "Creates a book, and returns the new Disk.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if CreateDiskFromFile == "" {
			
			
			cmd.MarkFlagRequired("parent")
			
			
			
			
			
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if CreateDiskFromFile != "" {
			in, err = os.Open(CreateDiskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &CreateDiskInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "CreateDisk", &CreateDiskInput)
		}
		resp, err := ServerClient.CreateDisk(ctx, &CreateDiskInput)
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


