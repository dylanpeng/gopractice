package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api short description",
	Long:  `api long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("api verbose: %v\n", Verbose)
		fmt.Printf("api config: %s\n", ConfigPath)
		fmt.Printf("api string: %s\n", SomeString)
	},
}

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin short description",
	Long:  `admin long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("admin verbose: %v\n", Verbose)
		fmt.Printf("admin config: %s\n", ConfigPath)
		fmt.Printf("admin string: %s\n", SomeString)
	},
}

var apiChildCmd = &cobra.Command{
	Use:   "apiChild",
	Short: "apiChild short description",
	Long:  `apiChild long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("apiChild verbose: %v\n", Verbose)
		fmt.Printf("apiChild config: %s\n", ConfigPath)
		fmt.Printf("apiChild string: %s\n", SomeString)
	},
}

var (
	Verbose    bool
	ConfigPath string
	SomeString string
)

func init() {
	rootCmd.AddCommand(apiCmd, adminCmd)
	apiCmd.AddCommand(apiChildCmd)
	// persistent是全局选项，对应的方法为PersistentFlags
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "全局版本")
	rootCmd.PersistentFlags().StringVarP(&SomeString, "string", "s", "null", "字符串")
	// local为本地选项，对应方法为Flags，只对指定的Command生效
	apiCmd.Flags().StringVarP(&ConfigPath, "config", "c", "", "读取文件路径")
}

/*
go run ./main.go -v -s aaa
root verbose: true
root config:
root string: aaa

go run ./main.go api -v -s aaa -c config
api verbose: true
api config: config
api string: aaa

go run ./main.go api apiChild -v -s aaa
apiChild verbose: true
apiChild config:
apiChild string: aaa
*/
