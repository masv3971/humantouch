#!/usr/bin/awk -f

# Split up skatteverkets test nin into two files, female and male.
# First argument is the original file path.

BEGIN {
    filename=ARGV[1]
    while (( getline line < filename) > 0){
        if (substr(line, 11, 1) % 2 == 0) 
        {
            print "Adding nin to female list"
            print "\x22"line"\x22""," > "nin_female.temp"
        }
        else
        {
            print "Adding nin to male list"
            print "\x22"line"\x22""," > "nin_male.temp"
        }
    }
}
