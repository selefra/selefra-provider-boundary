package main

// import (
// 	"fmt"

// 	"github.com/selefra/selefra-provider-boundary/constants"
// 	"github.com/selefra/selefra-provider-boundary/provider"
// 	"github.com/selefra/selefra-provider-sdk/doc_gen"

// 	"os"
// 	"testing"
// )

// func Test(t *testing.T) {

// 	fmt.Println(constants.Begin)
// 	docOutputDirectory := os.Getenv(constants.SELEFRADOCOUTPUTDIRECTORY)
// 	if docOutputDirectory == constants.Constants_33 {

// 		docOutputDirectory = constants.Tables

// 	}
// 	fmt.Println(docOutputDirectory)
// 	err := doc_gen.New(provider.GetProvider(), docOutputDirectory).Run()

// 	if err != nil {
// 		panic(err)

// 	}

// 	fmt.Println(constants.Done)

// }