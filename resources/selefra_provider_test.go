package resources

import (
	"os"
	"testing"

	"github.com/selefra/selefra-provider-sdk/test_helper"
)

func Test_Provider(t *testing.T) {
	provider := GetSelefraProvider()
	// split := strings.Split(os.Getenv("SELEFRA_TEST_TABLES"), ",")
	os.Setenv("SELEFRA_DATABASE_DSN", "host=127.0.0.1 user=postgres password=postgres port=5432 dbname=postgres sslmode=disable")
	test_helper.RunProviderPullTables(provider, "log: info", "./", "boundary_scope")
}
