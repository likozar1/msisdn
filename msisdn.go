package main

import (
	"strconv"
	"strings"
	"fmt"
	"os"
)

type MSISDN struct {
    msisdn string
}

//Validation of input
func (w *MSISDN) validate() string {
		trimmed := strings.TrimSpace(w.msisdn)
		if(len(trimmed)>15){
			return ""
		}
		if (trimmed == "") {
            return ""
        }
        if _, err := strconv.Atoi(trimmed); err != nil {
			return ""
		}else{
			msd , _ := strconv.Atoi(trimmed)
			return strconv.Itoa(msd)
		}	
}

// Parse msisdn depending on values in data files.

func (w *MSISDN) check(msisdn string) [4]string{
	result := [4]string{"", "", "", ""}
	oCodes := mncode("data/operators.txt")
	cCodes := countries("data/countries.txt")
	operatorLen := 0
	for i := 0; i < len(msisdn); i++ {
	
		if( i>4 && hasCCode(cCodes,result[0]) != true || i>7 && hasOperator(oCodes[result[0]],result[1]) != true){
			break
		}
		
		if(i<5 && hasCCode(cCodes,result[0]) != true){
			result[0]= msisdn[0:i]
			result[3]=cCodes[result[0]]
		}	
		
		if(oCodes[result[0]]!=nil){
			if(operatorLen>=len(oCodes[result[0]].Front().Value.(string))){
				if(operatorLen == 2){
					result[0]= msisdn[0:i-1]
				}else if(operatorLen == 1){
					result[0]= msisdn[0:i-2]
				}else{
					result[0]= msisdn[0:i]
				}
				result[3]=cCodes[result[0]]
				operatorLen=0	
			}
		}
		
		if(hasCCode(cCodes,result[0]) == true && i < 8 && hasOperator(oCodes[result[0]],result[1]) != true){
			operatorLen+=1
			result[1]= msisdn[len(result[0]):i+1]
			if(hasOperator(oCodes[result[0]],result[1]) == true){
				result[2] = msisdn[i+1:len(msisdn)]
				break					
			}
		}
    }
	return result
}

// Ensure that the result is not null or that 
// any part of msisdn is missing.
// If msisdn is invalid it returns empty array.
// ex: ["","","",""]

func (w *MSISDN) getResult() [4]string{
		result := [4]string{"", "", "", ""}
		msisd := w.validate()
		
		if(msisd == ""){
			return result
		}else{
			result = w.check(msisd)
		}
		
		if(result[0] == "" || result[1] == "" || result[2] == "" || result[3] == ""){
			return [4]string{"", "", "", ""}
		}else{
			return result
		}
}

// Get string as input and parses it into an array.
// ex: "+38640258450" -> ["386","40","258450","SI"]
// result: ["CC","MNC","SN","CID"]

func parse(msisdn string)[4]string{
	var result MSISDN = MSISDN{msisdn}
	return result.getResult()
}

// Takes msisdn as first argument and returns array of it's info.

func main(){
	if(len(os.Args)>1){
		fmt.Println(parse(os.Args[1]))
	}else{
		fmt.Println("missing msisdn argument")
	}
}

//RPC 

type Args struct {
    M string
}

type Result [4]string

func (t *MSISDN) Multiply(args Args, result *Result) error {
    return Multiply(args, result)
}

// Stores product of args.M in result pointer

func Multiply(args Args, result *Result) error {
    *result = Result(parse(args.M))
    return nil
}


