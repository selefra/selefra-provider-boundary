package main

import (
	"github.com/selefra/selefra-provider-boundary/resources"
	"github.com/selefra/selefra-provider-sdk/doc_gen"
	"os"
	"testing"
)

func Test(t *testing.T) {

	docOutputDirectory := "./docs"
	_ = os.Mkdir(docOutputDirectory, os.ModePerm)
	err := doc_gen.New(resources.GetSelefraProvider(), docOutputDirectory).Run()

	if err != nil {
		panic(err)
	}

}
