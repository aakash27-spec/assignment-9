package main 

import "testing"

func TestCalculate(t *testing.T) {

	var tests = []struct {
		
		date string  

		frame string 

		handlebar string  
		gear int64
		geargrip int64

		seating int64  
		seatingbottle int64

		wheels string 
		spokes int64
		rim int64
		tube int64
		tyre string 
		
		chain string 
	
	}{
		{"10-2016","steel","steel",4,220,1,200,"steel",400,200,300,"tubeless","onespeed"},
		
	}

	for _,test := range tests {
		if output,output1,outpu2,output3,output4,output5 := Calculateginprice(test.date,test.frame,test.handlebar,test.gear,test.geargrip,test.seating,test.seatingbottle,test.wheels,test.spokes,test.rim,test.tube,test.tyre,test.chain); test.date[3:7] != "2016" {
			t.Error("test failed {} is {} is {}: {}: {}: {}:",output,output1,outpu2,output3,output4,output5)
		}else {
			t.Log("tested the price for cycle engine {}",output5)
		}
	}
}