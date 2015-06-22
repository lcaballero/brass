package fortesting

import (
	"os"
	"testing"
	"fmt"
	"path/filepath"
)

const TESTING_PWD = "TESTING_PWD"

func Pwd() string {
	return os.Getenv(TESTING_PWD)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func PwdForTesting(f string) string {
	return filepath.Join(Pwd(), f)
}

func Dump(err error, parms ...interface{}) {
	for _,e := range parms {
		fmt.Println(e)
	}
	fmt.Println(err)
}

func Join(root, file string) string {
	return PwdForTesting(filepath.Join(root, file))
}


func init() {
	if Pwd() == "" {
		msg := fmt.Sprintf(
			"Need to set %s, typically done with `export %s=$(pwd)`\n",
			TESTING_PWD, TESTING_PWD)
		panic(msg)
	}
}