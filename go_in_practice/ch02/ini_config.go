import (
	"fmt"
	"os"
	"github.com/go-ini/ini"
)

func main() {
	config, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path, err = config.Section("Section").Key("path").String()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	enabled, err := config.Section("Section").Key("path").Bool()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}