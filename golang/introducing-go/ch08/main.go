package main
import (
	"fmt"
	"strings"
	"bytes"
	m "learngo/introducing-go/ch08/math"
)

func main() {
	fmt.Println(strings.Contains("AAAA", "a"))
	fmt.Println(strings.Contains("AAAA", "A"))
	fmt.Println(strings.Count("AAAA", "A"))
	fmt.Println(strings.HasPrefix("ASDLFKJASDLKF", "ASD"))
	fmt.Println(strings.HasSuffix("ASDLFKJASDLKF", "F"))
	fmt.Println(strings.Index("ASDLFKJASDLKF", "FKJ"))
	fmt.Println(strings.Join([]string{"ASDL", "FKJASDLKF", "ASXXZ"}, "--"))
	fmt.Println(strings.Repeat("ASDL", 4))
	fmt.Println(strings.Replace("ASDLFKJASDLKF", "S", "111", 1))
	fmt.Println(strings.Split("ASDL--FKJASDLKF--ASXXZ", "--"))
	fmt.Println(strings.ToLower("ASDL--FKJASDLKF--ASXXZ"))
	fmt.Println(strings.ToUpper("adfasdflkasjdf"))
	// bytes <-> string
	arr := []byte("strings!")
	str := string([]byte{'A', 'B', 'C'})
	fmt.Println(arr, str)
	var buf bytes.Buffer
	buf.Write([]byte("TEST!"))
	fmt.Println(buf, buf.Bytes())

	// Test package
	fmt.Println(m.Average([]float64{1, 2, 43, 4, 5, 6}))
	//math.cannotUse()
	m.CanUse()
}
