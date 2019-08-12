package goSdk

import (
	"os"
	"text/tabwriter"
	"fmt"
)

func TabWriter() {
	w := new(tabwriter.Writer)
	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 5, 1, '.',tabwriter.Debug)
	fmt.Fprintln(w, "a	b	c	d	")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	//w.Flush()
}
