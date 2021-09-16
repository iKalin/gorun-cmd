package goruncmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a cli for a todo app",
}

func Execute() {
	fmt.Println("Execute function was called.")
}
