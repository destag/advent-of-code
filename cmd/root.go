package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	yearFlag int
	dayFlag  int
)

var rootCmd = &cobra.Command{
	Use:   "go run main.go",
	Short: "runs solution for the given day",
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := solutions[yearFlag]; !ok {
			log.Fatalf("no solutions for year %d", yearFlag)
		}

		if _, ok := solutions[yearFlag][dayFlag]; !ok {
			log.Fatalf("no solution for year %d, day %d", yearFlag, dayFlag)
		}

		path := fmt.Sprintf("data/%d/%02d.txt", yearFlag, dayFlag)
		f := solutions[yearFlag][dayFlag]
		a, b := f(path)
		fmt.Println(a)
		fmt.Println(b)
	},
}

func init() {
	t := time.Now()
	rootCmd.Flags().IntVar(&yearFlag, "year", t.Year(), "year to run")
	rootCmd.Flags().IntVar(&dayFlag, "day", t.Day(), "day to run")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
