package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
)

func main() {
	
	var wg sync.WaitGroup

	for i := 0; i <=1; i++ {
        wg.Add(1)
		test1:= go worker(i, &wg)

    }
	
	wg.Wait()
	
 
}
func worker(id int, wg *sync.WaitGroup) int {
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
	 var total int
	 total = 0
	
		defer wg.Done()
		fmt.Printf("Worker %d starting\n", id)
		for i:=0;i<10;i++ {
		for _,test := range tests {
			output5 := Calculateginprice(test.date,test.frame,test.handlebar,test.gear,test.geargrip,test.seating,test.seatingbottle,test.wheels,test.spokes,test.rim,test.tube,test.tyre,test.chain)
				total = total + output5
			
	
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}
	return total
}
	

}
func Calculateginprice(date string,frame string,handlebar string,gear int64,geargrip int64,seating int64,seatingbottle int64,wheels string,spokes int64,rim int64,tube int64,tyre string,chain string)(int64) {

	var tyred,frameprice,chainprice,tseatingprice,wheelprice,handlebarprice,totalcyclengineprice int64
	Date,_ := strconv.Atoi(date[0:2])
	Year,_ := strconv.Atoi(date[3:7])
	if(Date>=1 && Date<=11 && Year==2016) {
		tyred = 200
	}else {
		tyred = 230
	}

	if(frame=="steel") {
		frameprice = 150
	} else if(frame=="Aluminium") {
		frameprice = 100
	}

	if(handlebar=="shockabsorb") {
		handlebarprice = 200 + (gear*100) + geargrip
	}else {
		handlebarprice = 100 + (gear*100) + geargrip
	}
	if(seating==1) {
		tseatingprice = 100 + seatingbottle
	} else {
		tseatingprice = 200 + seatingbottle
	}
	if(wheels=="steel") {
		wheelprice = 150 + spokes + rim + tyred + 200 + tube
	} else if(wheels=="Aluminium") {
		wheelprice = 200 + spokes + rim + tyred + 100 + tube
	}
	if(chain=="onespeed") {
		chainprice = 150
	} else if(chain=="twospeed") {
		chainprice = 200
	}

	totalcyclengineprice = frameprice + handlebarprice + wheelprice + chainprice + tseatingprice
	return totalcyclengineprice

	
}