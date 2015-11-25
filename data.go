package main

import (
	"container/list"
    "bufio"
    "os"
	"strconv"
	"strings"
)

// Fetches data from provided file.

func countries(fname string) map[string]string{
	f, _ := os.Open(fname)
	cCodes := make(map[string]string)
    scanner := bufio.NewScanner(f)
	
    for scanner.Scan() {
		result := strings.Split(scanner.Text(), ":")
		cc , _ := strconv.Atoi(result[0])
		s := strconv.Itoa(cc)
		cCodes[s]=result[1]
    }
	return cCodes
}

// Fetches data from provided file.

func mncode(fname string) map[string]*list.List{
	f, _ := os.Open(fname)
	data := make(map[string]*list.List)
    scanner := bufio.NewScanner(f)
	
    for scanner.Scan() {
		result := strings.Split(scanner.Text(), ":")
		cc , _ := strconv.Atoi(result[4])
		s := strconv.Itoa(cc)
		if(data[s]==nil){
			data[s] = list.New()
		}
		data[s].PushBack(result[1])
	}
	return data
}

// Checks if Operator is valid

func hasOperator(data *list.List,opearator string) bool{
	for e := data.Front(); e != nil; e = e.Next() {
		if(e.Value.(string)==opearator){
			return true
		}
	}
	return false
}

//Checks if country code is valid

func hasCCode(data map[string]string,country string) bool{
	if(data[country]==""){
		return false
	}
	return data[country] != "--"
}
