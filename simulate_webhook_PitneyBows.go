package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//vars that can be modified
var url = "http://127.0.0.1:4567"
var numOfReq = 5       //number of request
var threadNumber = 3   //max number of concurrency level(need to be <=10)
var lowerBoundTime = 2  //lower bound of time between every two post request(need to be >0 second)
var upperBoundTime = 10 //upper bound of time between every two post request(need to be <=30 second for test purpose)
//

//
var lowerBoundTime_min = 0
var upperBoundTime_max = 20
var numOfReq_max = 50
var threadNumber_max = 10

//

var mu sync.Mutex

func main() {
	if lowerBoundTime < lowerBoundTime_min {
		log.Fatalf("invalid lowerBoundTime")
	}
	if upperBoundTime > upperBoundTime_max {
		log.Fatalf("invalid upperBoundTime")
	}
	if threadNumber > threadNumber_max {
		log.Fatalf("invalid Thread Number")
	}
	if numOfReq > numOfReq_max {
		log.Fatalf("invalid Request Number")
	}

	cond := sync.NewCond(&mu)
	finished := 0

	for i := 0; i < threadNumber; i++ {
		//one thread start
		go func() {
			for {
				mu.Lock()
				if numOfReq <= 0 {
					mu.Unlock()
					break
				}
				numOfReq--
				fmt.Println("Send one Post Request, Remain: ", numOfReq)
				mu.Unlock()

				req, err := http.NewRequest("POST", url, bytes.NewBuffer(generateJson()))
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				goSleep()
			}
			//one thread finish
			mu.Lock()
			finished++
			mu.Unlock()
			cond.Signal()
		}()
	}
	//wait until all the thread finished
	mu.Lock()
	for finished != threadNumber {
		cond.Wait()
	}
	mu.Unlock()
}

func generateTime() int {
	diff := upperBoundTime - lowerBoundTime
	return lowerBoundTime + rand.Intn(diff)
}

func goSleep() {
	time.Sleep(time.Second * time.Duration(generateTime()))
}

//instead of just send same Json every time, this will be modified
func generateJson() []byte {
	jsonStr := []byte(`{
		"eventList": [ {
			"trackingNumber": "9400109898642304965461",
			"carrier": "USPS",
			"estimatedDeliveryDate": "20190610",
			"estimatedDeliveryTime": null,
			"scanDetails": {
				"eventDate": "20190606",
				"eventTime": "224700000",
				"eventCity": "ATLANTA",
				"eventStateOrProvince": "GA",
				"postalCode": "30304",
				"country": null,
				"scanType": null,
				"scanDescription": "ACCEPTED AT USPS FACILITY  10",
				"packageStatus": "InTransit"
			}
		},{
			"trackingNumber": "9400109898642304965461",
			"carrier": "USPS",
			"estimatedDeliveryDate": "20190610",
			"estimatedDeliveryTime": null,
			"scanDetails": {
				"eventDate": "20190606",
				"eventTime": "223200000",
				"eventCity": "MARIETTA",
				"eventStateOrProvince": "GA",
				"postalCode": "30062",
				"country": null,
				"scanType": null,
				"scanDescription": "ORIGIN ACCEPTANCE",
				"packageStatus": "InTransit"
			}
		},{
			"trackingNumber": "9400109898642304965461",
			"carrier": "USPS",
			"estimatedDeliveryDate": "20190610",
			"estimatedDeliveryTime": null,
			"scanDetails": {
				"eventDate": "20190606",
				"eventTime": "100300000",
				"eventCity": "MARIETTA",
				"eventStateOrProvince": "GA",
				"postalCode": "30062",
				"country": null,
				"scanType": null,
				"scanDescription": "SHIPPING LBL CREATED  USPS AWAITS ITEM",
				"packageStatus": "Manifest"
			}
		},{
			"trackingNumber": "9400109898642304965478",
			"carrier": "USPS",
			"estimatedDeliveryDate": "20190610",
			"estimatedDeliveryTime": null,
			"scanDetails": {
				"eventDate": "20190607",
				"eventTime": "090800000",
				"eventCity": "ATLANTA",
				"eventStateOrProvince": "GA",
				"postalCode": "30354",
				"country": null,
				"scanType": null,
				"scanDescription": "DEPART USPS FACILITY",
				"packageStatus": "InTransit"
			}
		} ],
		"totalEvents": 4
	}`)
	return jsonStr
}
