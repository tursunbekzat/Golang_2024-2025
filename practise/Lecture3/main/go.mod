module main

go 1.23.1

replace factorial => ../factorial
replace fibonachi => ../fibonachi

require (
    factorial v0.0.0
    fibonachi v0.0.0
)