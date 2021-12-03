#!/usr/bin/awk -f

# Split up skatteverkets test nin into two files, female and male.
# First argument is the original file path.

BEGIN {
    filename=ARGV[1]
    while (( getline line < filename) > 0){
        year=substr(line, 1,4)
        if (substr(line, 11, 1) % 2 == 0) 
        {
            key=("female:"year)
            hash[key]++
            femaleCounter++
            if (hash[key] == 8)
            {
                hash[key]=0
                printf "\x22%s\x22,\n", line > ("nin_female_"year".temp")
            }
            else 
             {
                printf "\x22%s\x22,", line > ("nin_female_"year".temp")
            }
        }
        else
        {
            key=("male:"year)
            hash[key]++
            maleCounter++
            if (hash[key] == 8)
            {
                hash[key]=0
                printf "\x22%s\x22,\n",line > ("nin_male_"year".temp") 
            }
            else
            {
                printf "\x22%s\x22,", line > ("nin_male_"year".temp")
            }
        }
    }
    print "males:", maleCounter, "females:", femaleCounter
}