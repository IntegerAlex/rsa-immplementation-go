package main
import( 
	//	"fmt"
	//"bytes"
	"math"
	"math/rand"
)


func main(){
	var p, q uint32 = primeGenerator()
    	println(p, q)
}

func primeGenerator() (uint32, uint32){
	var p uint32 
	var q uint32
	for{
		p = rand.Uint32()
		if primeCheck(p){
			break
		}
	}
	for{
		q = rand.Uint32()
		if primeCheck(q){
			break
		}
	}
	if p == q{
		primeGenerator()
	}
	return p, q
}




func primeCheck(num uint32) bool {
    if num < 2 {
        return false
    }
    if num == 2 {
        return true
    }
    if num % 2 == 0 {
        return false
    }
    // Only check for factors up to the square root of num
    sqrtNum := uint32(math.Sqrt(float64(num)))
    for i := uint32(3); i <= sqrtNum; i += 2 {
        if num % i == 0 {
            return false
        }
    }
    return true
}












