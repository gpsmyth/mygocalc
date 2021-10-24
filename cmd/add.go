/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"
	"encoding/json"
	"log"
	"strings"
	"reflect"

	"github.com/spf13/cobra"
)

// Placeholder for future thoughts
type TestConfig struct {
	Servers []string
}

var testConfig TestConfig
var rangeFlag []string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Range Values", rangeFlag, "type value", reflect.TypeOf(rangeFlag))
		
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("float")
		istatus, _ := cmd.Flags().GetBool("int")
		// No tenary operator in Golng!
		rstatus := false
		if len(rangeFlag) > 0 { rstatus = true } else { rstatus = false }
		fmt.Println("fstatus is ", fstatus, "istatus is", istatus, "rstatus is", rstatus)
		
		// Previousyl I tried
		// `if rangeFlag != nil` - however, this []string is never nil due to nothing
		// being assigned at call `addCmd.Flags().StringSliceVar(P)`
		// nil will only occur if `rangeFlag` doesn't exist
		if len(rangeFlag) > 0 {
			// e.g go run mygocalc add  --range 1:4
			addRange(rangeFlag)
		}

		if fstatus { // if float status is true, call addFloat
			// e.g. go run mygocalc add 1.2 1.8 -f
			addFloats(args)
		}

		if istatus {
			// go run mygocalc add 2 8 -i
			addNums(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
	addCmd.Flags().BoolP("int", "i", false, "Add integers")

	// Whilst not strictly correct in the usage here, i wanted to use
	// StringSliceVar(P) functionality
	// addCmd.Flags().StringSliceVar( &rangeFlag, "ss", []string{}, "Add a string of ints")
	// addCmd.Flags().StringSliceVarP(&rangeFlag, "range", "r", []string{"1:20"}, "Range of numbers. Optional")
	addCmd.Flags().StringSliceVar(&rangeFlag, "range", []string{}, "Range values")
}

// add.go
// Purpose: Add 2 digits, return sum
func addNums (args []string){
	var sum int = 0
	
	fmt.Printf("addNums function\n")

	for _, ival := range args {
		fmt.Printf("%s\n", ival)
		// Using strconv to perform atoi processing
		itemp, err := strconv.Atoi(ival)

		if (err != nil) {
			fmt.Println("Err is", err)
		}
		sum += itemp
	}
	fmt.Println("sum is", sum)
}

// add.go
// Purpose: Add 2 digits, return sum
func addFloats (args []string){
	var sum float64 = 0
	
	fmt.Printf("addFloats function\n")

	for _, ival := range args {
		fmt.Printf("%s\n", ival)
		// Using strconv to perform atof processing
		itemp, err := strconv.ParseFloat(ival, 64)

		if (err != nil) {
			fmt.Println("Err is", err)
		}
		sum += itemp
	}
	fmt.Println("sum is", sum)
}

// add.go
// Purpose: Add range of digits, return sum
func addRange (args []string) {
	var sum int = 0

	justString := strings.Join(args," ")

	// Replace : with space - strings are mutable at package level
	justString = strings.Replace(justString, ":", " ", -1)

	fmt.Println("addRange function", justString)
	
	var l,h int
	if _, err := fmt.Sscanf(justString, "%2d %2d", &l, &h); err == nil {
    fmt.Println("range from", l, "to", h)

		for i := l; i <= h; i++ {  
			sum += i
		}
	}
	fmt.Println("sum is", sum)

	// Here, I've got []string which is sliced and I want to rejoin to a string for []int processing
	// str := "[2,23]" // works
	// str := "2,23"   // not
	// str := fmt.Sprint(args, ",")
	str := strings.Join(args, ",")
	fmt.Println("str:", str, reflect.TypeOf(str))

	// astr := "[" + str + "]"
	// Replacing : with , and concatenating string
	astr := "[" + strings.Replace(str, ":", ",", -1) + "]"
	
	var ints []int
	err := json.Unmarshal([]byte(astr), &ints)
	if err != nil {
			log.Fatal(err)
	}
	fmt.Printf("%v", ints)
	
	result := 0
	for _, numb := range ints {  
		result += numb  
	 }  
	 fmt.Println("result -", result)
}