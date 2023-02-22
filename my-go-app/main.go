package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"math"
)

var count int
var mutex sync.Mutex
var group1 int
var group2 int




// var z float64

var totalVisitors = 10000
//var conversion1 = 150
//var conversion2 = 200
var lift float64
var p float64
var zScore float64



func assignWidgetID(numVisitors int) string {

	rand.Seed(time.Now().UnixNano())
	variation := rand.Intn(2)
	if variation == 0 {
		group1++
		return "Widget A"
	} else {
		group2++
		return "Widget B"
	}

}


//calculate z test Statistical Significance 
func test(){
	// number of conversions in group 1
 n1 := 1000      // number of samples in group 1
   // number of conversions in group 2
 n2 := 1000      // number of samples in group 2
 confidence := 0.92 // confidence level
 var statisticalSig float64

 // Calculate the pooled conversion rate and standard error of difference
 p1 := float64(group1) / float64(n1)
 p2 := float64(group2) / float64(n2)
 p := (float64(group1) + float64(group2)) / (float64(n1) + float64(n2))
 stdErr := math.Sqrt(p * (1 - p) * ((1.0 / float64(n1)) + (1.0 / float64(n2))))

 // Calculate the z-score and p-value
 zScore := (p1 - p2) / stdErr
 pValue := 2 * (1 - 0.5*(1+math.Erf(math.Abs(zScore)/math.Sqrt2)))
 statisticalSig = (1-pValue)*100

 // Output the results
 fmt.Printf("Group 1: %d/%d (%.2f%%)\n", group1, n1, p1*100)
 fmt.Printf("Group 2: %d/%d (%.2f%%)\n", group2, n2, p2*100)
 fmt.Printf("Pooled conversion rate: %.2f%%\n", p*100)
 fmt.Printf("Standard error of difference: %.4f\n", stdErr)
 fmt.Printf("Z-score: %.4f\n", zScore)
 fmt.Printf("P-value: %.4f\n", pValue)
 fmt.Printf("Statistical significance: %.4f\n", statisticalSig)

 // Check if the difference is statistically significant
 if pValue < 1-confidence {
	 fmt.Println("The difference between the two groups is statistically significant.")
 } else {
	 fmt.Println("The difference between the two groups is not statistically significant.")
 }
}

//calculate confidence intervall
var pointES float64
var pg1 float64
var pg2 float64
var standardER float64
var CI1 float64
var CI2 float64
func confidenceInt (){
	pg1 = float64(group1)/1000
	pg2 = float64(group2)/1000
	pointES=(pg1 - pg2) / math.Sqrt((pg1+pg2)/2)
    var se1 float64 = pg1*(1-pg1)/float64(1000)
var se2 float64 = pg2*(1-pg2)/float64(1000)

if se1 <= 0 || se2 <= 0 {
    fmt.Println("Error: Standard error is undefined")
    return
}

var standardER float64 = math.Sqrt(se1 + se2)

if math.IsNaN(standardER) {
    fmt.Println("Error: Standard error is NaN")
    return
}
    CI1 = pointES + (1.96*standardER)
	CI2 = pointES - (1.96*standardER)
	fmt.Printf("the pointES is: %f\n", pointES)
	fmt.Printf("the standardER is: %f\n", standardER)
	fmt.Printf("the CI+ is: %f\n", CI1)
    fmt.Printf("the CI- is: %f\n", CI2)
}


func main() {

	// Define a handler function that increments the counter
	handler := func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		fmt.Printf("count before increment: %d", count)
		count++
		fmt.Printf("count after increment: %d", count)

		fmt.Printf("You are visitor number %d", count)
		// Example usage
		fmt.Println(assignWidgetID(5)) // Output: Widget A or Widget B (randomly assigned)
        fmt.Println("group 1 :",group1 , "group 2", group2)
		
	lift = (float64(group2 - group1) / float64(group1)) * 100
    fmt.Println("Lift: %f%%\n", lift)
	
	//updateStats()
	test()
	confidenceInt()
	}



	// Set up a web server on port 8080
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

   
}
