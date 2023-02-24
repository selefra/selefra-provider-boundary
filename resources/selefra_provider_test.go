package resources

import (
	"os"
	"strings"
	"testing"

	"github.com/selefra/selefra-provider-sdk/test_helper"
)

func Test_Provider(t *testing.T) {
	provider := GetSelefraProvider()
	split := strings.Split(os.Getenv("SELEFRA_TEST_TABLES"), ",")
	test_helper.RunProviderPullTables(provider, "log: info", "./", split...)
}
