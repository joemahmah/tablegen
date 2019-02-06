package main

import (
	"fmt"
	"flag"
)

func main() {
	var rows = flag.Int("r", 5, "The number of rows in the table.")
	var cols = flag.Int("c", 3, "The number of columns in the table.")
	var header = flag.Bool("header", false, "Include a header row in the table.")
	var tname = flag.Bool("tname", false, "Include a table name row.")

	flag.Parse()

	fmt.Println("<table>")

	if *tname {
		fmt.Println("\t<tr>\n\t\t<th colspan=\"",*cols,"\"> TITLE </th>\n\t</tr>")
	}

	if *header {
		fmt.Println("\t<tr>")
			for x := 0; x < *cols; x++ {
				fmt.Println("\t\t<th> HEAD </th>")
			}
		fmt.Println("\t</tr>")
	}

	for y := 0; y < *rows; y++ {
		fmt.Println("\t<tr>")
			for x := 0; x < *cols; x++ {
				fmt.Println("\t\t<td> x </td>")
			}
		fmt.Println("\t</tr>")
	}

	fmt.Println("</table>")

}
