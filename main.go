package main
import (
	"os"
	"shyiran/my-gin-vue/cmd"
)
func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}