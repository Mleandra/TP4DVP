/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"calculator/internals/math"

	"github.com/spf13/cobra"
)

//var x float64
//var y float64

// divCmd représente la commande de division
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Division de x et y",
	Long:  `Division de x et y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Division(x, y)
		fmt.Printf("%f / %f = %f\n", x, y, ret)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)
	divCmd.Flags().Float64VarP(&x, "x", "a", 0, "première valeur")
	divCmd.Flags().Float64VarP(&y, "y", "b", 0, "deuxième valeur")
}
