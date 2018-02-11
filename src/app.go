package main

/**
Go version:1.9.3

@Feb 2018
@by Ashwin Rana
*/

import (
	"flag"
	"fmt"
)

func main() {
	//most_sold
	var sold bool
	flag.BoolVar(&sold, "m", false, "")
	flag.BoolVar(&sold, "most_Sold", false, "")

	//total_Spend
	var spend string
	flag.StringVar(&spend, "t", "", "")
	flag.StringVar(&spend, "total_Spend", "", "")

	//most_Loyal
	var loyal bool
	flag.BoolVar(&sold, "l", false, "")
	flag.BoolVar(&sold, "most_Loyal", false, "")

	//help
	var gethelp bool
	flag.BoolVar(&gethelp, "help", false, "")

	//FlagSet
	setFlag(flag.CommandLine)
	flag.Parse()

	if sold {
		most_Sold()
		return
	}

	if loyal {
		most_Loyal()
		return
	}

	if gethelp {
		help()
		return
	}

	if spend != "" {
		email := ""
		total_Spend(email)
		return
	}
}

func help() {
	//return available options
	fmt.Println(`
    Usage:
    /.src command [arguments]

    The commands are:
    -m, --most_Sold	Prints the name of the most sold item.
    -t, --total_Spend	Prints the total spend of the user with this email address [EMAIL]
    -l, --most_Loyal	Prints the email address of the most loyal user (most purchases).
    -h, --help		Prints this help info
  `)
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		help()
		return
	}
}
