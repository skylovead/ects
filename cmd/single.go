package cmd

import (
	"github.com/betterde/ects/models"
	"github.com/gorhill/cronexpr"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
)

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single node service",
	Long:  "Run a single node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		//discover.NewClient()
		//config.Conf = config.Init()
		//discover.GetConf("/ects/config")
		//pipeline.WatchPipelines("7df52971-4894-4f01-9171-7452c4ddceca")
		log.Println(cronexpr.MustParse("* * * * * *").Next(time.Now()).Format(models.DefaultTimeFormat))
		log.Println(strings.TrimPrefix("/var/local/laravel", "/var/local/"))
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
