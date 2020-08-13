package main
import ( "fmt"
	 "strings" 
) 

func WordCount(s string) map[string]int {
	replace := strings.NewReplacer(",","")
	s = replace.Replace(s)
	words := strings.Fields(s) 
	m := make(map[string]int) 
    for _, word := range words { 
		m[word] += 1 
	} 
	return m 
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"
	r := WordCount(s)
	fmt.Println(r)

}
