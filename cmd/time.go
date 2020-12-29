package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wlight/tour/internal/timer"
	"log"
)

var (
	calculateTime string
	duration string
)

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s，%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}