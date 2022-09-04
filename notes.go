Login = "bugga" ###public variable (starts with capital L)
login = "bugga" ###private variable

##variable delcaration 

var name string = "krishna"
name := "krishna"  ##valid only inside a method, not valid outside a method..



####time
func main() {
	fmt.Println("welcome")
	presentTime := time.Now()
	fmt.Println("presentTime")
	fmt.Println(presentTime.Format("01-01-2000"))  ###time requires a format 

}

output: 
 welcome
 presentTime
 09-09-2000


### Memory
Garbage collection is automatic in go.

###Pointers

*points data 
& points address



### arrays
a := [5]int{0,1,2,3,4}
fmt.Println(a)



