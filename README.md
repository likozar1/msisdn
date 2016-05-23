# MSISDN reader

Countries and operators are defined under directory /data.
Country codes are listed in countries.txt
Operators are listed in operators.txt


## To set everything up run next command:
  go build msisdn.go data.go

## Example of running msisdn.exe file:

  msisdn.exe +38640753845

### will output:

  [386 40 753845 SI]

## If Country code is not detected or is invalid:

  msisdn.exe +99940753845

### will output empty array:

  [   ]

## TESTS

### For tests run command:

  go test -v
