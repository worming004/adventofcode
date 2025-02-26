package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed try.txt
var try string

func main() {
	reader := strings.NewReader(input)
	inputs := Parse(reader)
	//part2
	for _, i := range inputs {
		i.x = i.x + 10000000000000
		i.y = i.y + 10000000000000
	}
	total := 0
	var invalid []resolveRes
	for _, i := range inputs {
		r := resolveInput(*i)
		fmt.Println(r.Pretty())
		if !r.isValid {
			invalid = append(invalid, r)
		} else {
			total = total + r.Cost()
		}
	}
	fmt.Println("Total cost: ", total)
	fmt.Println("Number invalid", len(invalid))
}

type Input struct {
	a, b, c, d, x, y int
}

func Parse(r io.Reader) []*Input {
	scanner := bufio.NewScanner(r)
	var res []*Input
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lineButtonA := scanner.Text()
		scanner.Scan()
		lineButtonB := scanner.Text()
		scanner.Scan()
		linePrize := scanner.Text()
		scanner.Scan() // emtpy line

		re := regexp.MustCompile(`\d+`)
		amatches := re.FindAllString(lineButtonA, -1)
		bmatches := re.FindAllString(lineButtonB, -1)
		pmatches := re.FindAllString(linePrize, -1)

		res = append(res, &Input{
			a: atoi(amatches[0]),
			b: atoi(amatches[1]),
			c: atoi(bmatches[0]),
			d: atoi(bmatches[1]),
			x: atoi(pmatches[0]),
			y: atoi(pmatches[1]),
		})
	}

	return res
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return int(i)
}

type resolveRes struct {
	rx, ry int

	detA int

	isValid bool
}

func (rr resolveRes) Cost() int {
	return 3*int(rr.rx) + int(rr.ry)
}

func (rr resolveRes) Pretty() string {
	sb := strings.Builder{}
	sb.WriteString("Valid: ")
	sb.WriteString(fmt.Sprintf("%t\n", rr.isValid))

	sb.WriteString("Det(A): ")
	sb.WriteString(fmt.Sprintf("%d\n", rr.detA))

	sb.WriteString("A pressed ")
	sb.WriteString(fmt.Sprintf("%d\n", int(rr.rx)))

	sb.WriteString("A cost ")
	sb.WriteString(fmt.Sprintf("%d\n", 3*int(rr.rx)))

	sb.WriteString("B pressed ")
	sb.WriteString(fmt.Sprintf("%d\n", int(rr.ry)))

	sb.WriteString("B cost ")
	sb.WriteString(fmt.Sprintf("%d\n", int(rr.ry)))

	sb.WriteString("Total cost ")
	sb.WriteString(fmt.Sprintf("%d\n", rr.Cost()))
	return sb.String()
}

func resolveInput(i Input) resolveRes {
	return resolve(i.a, i.b, i.c, i.d, i.x, i.y)
}

func resolve(a, b, c, d, x, y int) resolveRes {
	detA := a*d - b*c
	if detA == 0 {
		return resolveRes{0, 0, 0, false}
	}
	detAf := float64(detA)
	rx := float64(d*x-c*y) / detAf
	ry := float64(a*y-b*x) / detAf

	isValid := true
	if rx < 0 {
		isValid = false
	}
	if ry < 0 {
		isValid = false
	}

	if math.Mod(rx, 1) != 0 {
		isValid = false
	}
	if math.Mod(ry, 1) != 0 {
		isValid = false
	}

	return resolveRes{int(math.Ceil(rx)), int(math.Ceil(ry)), detA, isValid}
}

// import (
// 	"fmt"
// 	"gonum.org/v1/gonum/mat"
// )
// func main() {
// 	// // 94x + 22y = 8400
// 	// // 34x + 67y = 5400
// 	// A := mat.NewDense(2, 2, []float64{94, 22, 34, 67})
// 	// b := mat.NewVecDense(2, []float64{8400, 5400})
//
// 	A := mat.NewDense(2, 2, []float64{26, 66, 67, 21})
// 	b := mat.NewVecDense(2, []float64{12748, 12176})
//
// 	var x mat.VecDense
// 	x.SolveVec(A, b)
// 	if err := x.SolveVec(A, b); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
//
// 	fmt.Println(x)
// 	nA := int(x.At(0, 0))
// 	nB := int(x.At(1, 0))
//
// 	r := nA + 3*nB
//
// 	fmt.Printf("A pressed %d and cost %d\n", nA, nA)
// 	fmt.Printf("B pressed %d and cost %d\n", nB, 3*nB)
// 	fmt.Printf("Total cost %d\n", r)
// }
