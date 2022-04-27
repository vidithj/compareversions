# compareversions
Simple Program to compare versions. The program takes 2 versions as commandline arguments. The program compares both the version and returns result as below : 

if version 1 > version 2 return 1
if version 1 < version 2 return -1
otherwise return 0

# Testing 
The logic to compare is present in api.go. The test cases around the function has been defined in the api_test.go

ok  	compareversions/api	0.860s	coverage: 95.2% of statements

# Execution

1. Running with main.go
    go run main.go version1 version2

Example : go run main.go 0.2.1.1.1 0.1.2

2. Running with Binary file(Darwin OS)
    goto folder cmd and run 
    ./compareversions version1 version2

Example: ./compareversions 0.2.1.1.1 0.1.2

2. Running with Binary file(Windows OS)
 goto folder cmd and run 
    ./compareversions_win version1 version2

Example: ./compareversions_win 0.2.1.1.1 0.1.2