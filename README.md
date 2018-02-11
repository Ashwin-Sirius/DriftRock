# DriftRock 

This app is designed to use the DriftRock API and answer some questions about the data held in their public API. This involved having to work out how best to merge data returned from two different end points.   

## Getting Started
To run the programme I have compile it as a binary therefore requires no further installations of Go or any dependencies

### Run (I didn't manage to finish the commands however, most_sold simply returns 50 items)

```
git clone https://github.com/Ashwin-Sirius/DriftRock.git
cd driftrock/bin
run "./src -h" for help

Usage:
/.src command [arguments]

The commands are:
-m, --most_Sold    Prints the name of the most sold item.
-t, --total_Spend    Prints the total spend of the user with this email address [EMAIL]
-l, --most_Loyal    Prints the email address of the most loyal user (most purchases).
-h, --help        Prints this help info
```

## Approach & Pitfalls
I spent quite a while simply thinking about the problem at hand and playing with the endpoints provided on documention online. The documentation was very helpful to understand what the data looked like and the code example was good enough to get me started with my research. I decided to use Go for this project as I was fairly familier with, further to that I really wanted to avoid any dependency problems further down the line with Ruby considering I didn't have a lot of time.

Most of the time was spent parsing the Json data across properly and creating the CLI. I found extracting all the data from the paginated endpoints fairly hard. I tried getting the maximum results per_page to the most I can in order to avoid going through multiple pages but I was only allowed no more than 100 results per page.

I have few ideas for solving the commands but the algorithms I tried need more tweeking and more knowlegde on how maps actually work on Go.

I didn't manage to finish the actual tasks but regardless thoroughly enjoyed learning a lot of new materials and got fairly better at Go.
