// Code generated. DO NOT EDIT.








package main

import (
	"github.com/spf13/cobra"
	
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	
	 "fmt"
	
	 "github.com/golang/protobuf/jsonpb"
	
	 "os"
	
	serverpb "core.wcloud.io/generated/grpcgen"
	
)

var UpdateDiskInput serverpb.UpdateDiskRequest

var UpdateDiskFromFile string
















func init() {
	ServerServiceCmd.AddCommand(UpdateDiskCmd)
	
	UpdateDiskInput.Book = new(serverpb.Disk)
	
	UpdateDiskInput.UpdateMask = new(fieldmaskpb.FieldMask)
	
	
	UpdateDiskCmd.Flags().StringVar(&UpdateDiskInput.Book.Name, "book.name", "", "The resource name of the book.  Disk names have...")
	
	UpdateDiskCmd.Flags().StringVar(&UpdateDiskInput.Book.Author, "book.author", "", "The name of the book author.")
	
	UpdateDiskCmd.Flags().StringVar(&UpdateDiskInput.Book.Title, "book.title", "", "The title of the book.")
	
	UpdateDiskCmd.Flags().BoolVar(&UpdateDiskInput.Book.Read, "book.read", false, "Value indicating whether the book has been read.")
	
	UpdateDiskCmd.Flags().StringSliceVar(&UpdateDiskInput.UpdateMask.Paths, "update_mask.paths", []string{}, "The set of field mask paths.")
	
	
	
	UpdateDiskCmd.Flags().StringVar(&UpdateDiskFromFile, "from_file", "", "Absolute path to JSON file containing request payload")
	
	
	
}

var UpdateDiskCmd = &cobra.Command{
  Use:   "update-disk",
  Short: "Updates a book. Returns INVALID_ARGUMENT if the...",
	Long: "Updates a book. Returns INVALID_ARGUMENT if the name of the book  is non-empty and does not equal the existing name.",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		if UpdateDiskFromFile == "" {
			
			
			
			
			
			
			
			
			
			
			
			
		}
		
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		
		in := os.Stdin
		if UpdateDiskFromFile != "" {
			in, err = os.Open(UpdateDiskFromFile)
			if err != nil {
				return err
			}
			defer in.Close()
			
			err = jsonpb.Unmarshal(in, &UpdateDiskInput)
			if err != nil {
				return err
			}
			
		} 
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		if Verbose {
			printVerboseInput("Server", "UpdateDisk", &UpdateDiskInput)
		}
		resp, err := ServerClient.UpdateDisk(ctx, &UpdateDiskInput)
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


