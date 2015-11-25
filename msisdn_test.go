package main

import "testing"

func TestSimpleMsisdn(t *testing.T){
	result := [4]string{"386", "40", "294020", "SI"}
	expected := result
	actual := parse("+38640294020")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	result = [4]string{"223", "01", "4521354", "ML"}
	expected = result
	actual = parse("+223014521354")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}

func TestInvalidMsisdn(t *testing.T){
	result := [4]string{"", "", "", ""}
	expected := result
	actual := parse("5das54564ad")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	result = [4]string{"", "", "", ""}
	expected = result
	actual = parse("%&(/#&$hjaksf45")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}

func TestValidMSISDNUnknownCountryCode(t *testing.T){
	result := [4]string{"", "", "", ""}
	expected := result
	actual := parse("+99984264481")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	result = [4]string{"", "", "", ""}
	expected = result
	actual = parse("+8095642781412")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}

func TestValidMSISDNUnknownOperator(t *testing.T){
	result := [4]string{"", "", "", ""}
	expected := result
	actual := parse("+38644294020")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	result = [4]string{"", "", "", ""}
	expected = result
	actual = parse("+223514521354")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}

func TestDifferentCountries(t *testing.T){
	result := [4]string{"386", "40", "294020", "SI"}
	expected := result
	actual := parse("+38640294020")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"223", "01", "4521354", "ML"}
	expected = result
	actual = parse("+223014521354")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"1671", "140", "8421324", "GU"}
	expected = result
	actual = parse("+16711408421324")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"244", "04", "74216414", "AO"}
	expected = result
	actual = parse("+2440474216414")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"54", "310", "4452132", "AR"}
	expected = result
	actual = parse("+543104452132")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"49", "77", "9875613", "DE"}
	expected = result
	actual = parse("+49779875613")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"33", "01", "06540125", "FR"}
	expected = result
	actual = parse("+330106540125")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"7", "16", "144556411", "RU"}
	expected = result
	actual = parse("+716144556411")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"27", "07", "84311214", "ZA"}
	expected = result
	actual = parse("+270784311214")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"1869", "50", "0454564", "KN"}
	expected = result
	actual = parse("+1869500454564")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"76", "01", "5004564", "KZ"}
	expected = result
	actual = parse("+76015004564")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	
	result = [4]string{"", "", "", ""}
	expected = result
	actual = parse("+10015004564")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}

func TestInvalidMsisdnLength(t *testing.T){
	result := [4]string{"", "", "", ""}
	expected := result
	actual := parse("+3686464402940520")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
	result = [4]string{"", "", "", ""}
	expected = result
	actual = parse("")
	if actual != expected {
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
		t.Error("Test failed")
	}else{
		t.Log("Expected:\t County Code: "+expected[0]+" MNcode: "+expected[1]+" Country ID: "+expected[3]+" Subscriber number: "+expected[2])
		t.Log("Actual:\t County Code: "+actual[0]+" MNcode: "+actual[1]+" Country ID: "+actual[3]+" Subscriber number: "+actual[2])
	}
}


