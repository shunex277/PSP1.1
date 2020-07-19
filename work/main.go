// Program Assignment: program4
// Name: Inoue Shunei
// Date: 2020/07/19
// Description: this is a program to caluculate relative size ranges for 
// 				very small, small, medium, large, and very large ranges using 
// 				standard deviation.
// func declarations:
// 		main                                                                                                    
// 		CalcMean
// 		CalcSd 

package main

import "fmt"
import "math"

func main() {
	buff := []float64{}
	flag := true
	for flag {
		var input float64
		fmt.Println("数字を入力してください（終了の場合は-1を入力）：")
		fmt.Scan(&input)
		if input < 0 {
			fmt.Println("終了")
			break
		} else {
			buff = append(buff, input)
		}
	}

	// 入力を自然体数値に変換
	for i, v := range buff {
		buff[i] = math.Log(v)
	}

	ave := CalcMean(buff)
	sd := CalcSd(buff)
	fmt.Printf("midpoint of the VS ranges: %.4f\n", math.Exp(ave - 2*sd))
	fmt.Printf("midpoint of the S ranges: %.4f\n", math.Exp(ave - sd))
	fmt.Printf("midpoint of the M ranges: %.4f\n", math.Exp(ave))
	fmt.Printf("midpoint of the L ranges: %.4f\n", math.Exp(ave + sd))
	fmt.Printf("midpoint of the VL ranges: %.4f\n", math.Exp(ave + 2*sd))
	
}

// Reuse instructins
// 		func CalcMean(buff []float64) float64
// 		purpose: to calculate mean value
// 		return: mean value of list
func CalcMean(buff []float64) float64 {
	var sum float64
	for _, x := range buff {
		sum += x
	}
	n := float64(len(buff))
	mean := sum / n
	return mean
}

// Reuse instructins
// 		func CalcSd(buff []float64) float64
// 		purpose: to calculate standard deviation value
// 		return: standard deviation value of list
func CalcSd(buff []float64) float64 {
	n := float64(len(buff))
	mean := CalcMean(buff)
	s := 0.0
	for _, x := range buff{
		s += (x - mean) * (x - mean)
	}
	var sd float64
	sd = math.Sqrt(s / (n-1))
	return sd
}
