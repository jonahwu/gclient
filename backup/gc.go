package main

import (
	"encoding/json"
	"fmt"
	"github.com/jonah/gomobilelib"
	"github.com/tidwall/gjson"
	"math"
	"time"
	//"time"
)

func FromStringToList(locx string) []float64 {
	c := make(map[string]interface{})
	e := json.Unmarshal([]byte(locx), &c)
	if e != nil {
		panic(e)
	}

	k := make([]string, len(c))
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	// output result to STDOUT
	fmt.Printf("%#v\n", k)
	var pos []float64
	for _, kk := range k {
		fv := float64(c[kk].(float64))
		//fmt.Println(fv)
		fmt.Println(fv)
		pos = append(pos, fv)
	}
	return pos

}
func CalGPSDistance(long1 float64, lati1 float64, long2 float64, lati2 float64) float64 {
	d2r := 0.0174532925199433
	dlong := (long2 - long1) * d2r
	dlat := (lati2 - lati1) * d2r
	a := math.Pow(math.Sin(dlat/2.0), 2) + math.Cos(lati1*d2r)*math.Cos(lati2*d2r)*math.Pow(math.Sin(dlong/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := 6373 * c
	//return meter
	return d * 1000.

}

func CalculateDistance(cx float64, cy float64, tarx float64, tary float64) float64 {
	dist := math.Sqrt((cx-tarx)*(cx-tarx) + (cy-tary)*(cy-tary))
	return dist
}

func CalculateTargetGPS(posx []float64, posy []float64) {
	//tary := 121.681260
	//tarx := 25.068420
	tary := 121.664570
	tarx := 25.066710
	fmt.Println(tarx, tary)
	for i := 0; i < len(posx); i++ {
		dist := CalGPSDistance(posx[i], posy[i], tarx, tary)
		fmt.Println("calculated dist:", dist, "meter", posx[i], posy[i], tarx, tary)
		time.Sleep(2 * time.Second)
	}

}

func main() {
	/*
		timestamp := time.Now().Unix()
		strTimeStamp := fmt.Sprintf("%d", timestamp)
		strid := "123jjoo"
		strlong := "23.3333"
		strlati := "123.3333"
		//gomobilelib.SendGPS(strTimeStamp, strlati, strlong, strid)
	*/
	//target 121.681260,25.068420
	getdata()
	locx := gjson.Get(cameralocx, "dps").String()
	locy := gjson.Get(cameralocy, "dps").String()
	posx := FromStringToList(locx)
	posy := FromStringToList(locy)
	fmt.Println(posx)
	fmt.Println(posy)
	fmt.Println(len(posx), len(posy))
	CalculateTargetGPS(posx, posy)
}

func mainold() {
	fmt.Println("vim-go")
	gomobilelib.Test()
	//	gomobilelib.Send()
	ss := gomobilelib.SendtoGCPTest()
	fmt.Println(ss)
	ff := gomobilelib.Testpass()
	fmt.Println(ff)
}
