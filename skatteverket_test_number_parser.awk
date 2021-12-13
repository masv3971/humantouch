#!/usr/bin/awk -f

# Split up skatteverkets test nins into seperate files, <nin_gender_year.temp>.
# First argument is the original file path.

function help(){
    printf "usage: <inputfile>\n"
    exit 0
}

BEGIN {
    filename=ARGV[1]
    if (filename == "" || filename == "help" )
        help()

    while (( getline nin < filename) > 0){
        year=substr(nin, 1,4)
        if (substr(nin, 11, 1) % 2 == 0) 
        {
            gender="female"
            femaleCounter++
        }
        else
        {
            gender="male"
            maleCounter++
        }

         key=(gender":"year)
         hash[key]++
         if (hash[key] == 8)
         {
             hash[key]=0
             printf "\x22%s\x22,\n", nin > ("nin_"gender"_"year".temp")
         }
         else 
          {
             printf "\x22%s\x22,", nin > ("nin_"gender"_"year".temp")
         }
    }
    print "males:", maleCounter, "females:", femaleCounter, "total:" femaleCounter+maleCounter
}