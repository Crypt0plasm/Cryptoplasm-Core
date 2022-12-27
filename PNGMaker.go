package main

import (
	"Cryptoplasm-Core/Auxiliary"
	b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
	kw "Cryptoplasm-Core/KosonicWritings"
	"flag"
	"fmt"
	r "github.com/go-vgo/robotgo"
	"time"
)

func CreatePNG(Day string, TxTaxAddress b.ElrondAddress) {
	start := time.Now()
	//WriteListOneByOne("R1.txt", R1)

	//1.Make SVG
	_, SVGName, SVGExtension, Rectangles := kw.DrawUnitySVG(Day, TxTaxAddress)

	//2.Copy SVG to Folder
	B, _ := Auxiliary.MyCopy(SVGName+SVGExtension, "_NFTMaker\\"+SVGName+SVGExtension)
	fmt.Println(B, " bytes copied!")

	//3.Process in CorelDraw
	kw.ConvertCorelDraw("d:\\.GO_workspace\\src\\Cryptoplasm-Core\\_NFTMaker\\", SVGName, SVGExtension)

	//4.Open PNG In Photoshop
	//Bring to front PhotoshopWindow - assumes CorelDraw is already opened.
	kw.MoveLeftClick(470, 1420)
	r.MilliSleep(3000)
	kw.OpenPNG("d:\\.GO_workspace\\src\\Cryptoplasm-Core\\_NFTMaker\\", SVGName)

	//5.Process in Photoshop
	kw.ConvertPhotoshop(Rectangles)
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)
}

func PartialCreatePNG(Day string, TxTaxAddress b.ElrondAddress) {
	start := time.Now()

	//1.Make SVG
	_, _, _, Rectangles := kw.DrawUnitySVG(Day, TxTaxAddress)

	//2.Copy SVG to Folder
	//B, _ := Auxiliary.MyCopy(SVGName, "_NFTMaker\\"+SVGName)
	//fmt.Println(B, " bytes copied!")

	//3.Process in CorelDraw
	//kw.ConvertCorelDraw(SVGName, "d:\\.GO_workspace\\src\\Cryptoplasm-Core\\_NFTMaker\\")

	//4.Open PNG In Photoshop
	//Bring to front PhotoshopWindow - assumes CorelDraw is already opened.
	//kw.MoveLeftClick(470, 1420)
	//r.MilliSleep(3000)
	//kw.OpenPNG(SVGName, "d:\\.GO_workspace\\src\\Cryptoplasm-Core\\_NFTMaker\\")

	//5.Process in Photoshop
	kw.ConvertPhotoshop(Rectangles)
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)
}

func main() {
	var (
		MakePNGH = `--make-png	DAY Number;
Starts computing the Unity NFT Value using the day number as input and creating the NFT.
`
		TestPNGH = `--test-png	DAY Number;
Starts testing the Unity NFT creation using only the PS part.
`
	)

	const (
		MakePNG = "make-png" // String - Day Number
		TestPNG = "test-png" // String - Day Number
	)

	FlagMakePNG := flag.String(MakePNG, "0", MakePNGH)
	FlagTestPNG := flag.String(TestPNG, "0", TestPNGH)
	//
	flag.Parse()

	if *FlagMakePNG != "0" {
		fmt.Println("")
		fmt.Println("Starting png Generation")
		fmt.Println("")
		Day := *FlagMakePNG
		TxTaxAddy := b.ElrondAddress("erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s")
		CreatePNG(Day, TxTaxAddy)
	}
	if *FlagTestPNG != "0" {
		fmt.Println("")
		fmt.Println("Starting testing png Generation")
		fmt.Println("")
		Day := *FlagMakePNG
		TxTaxAddy := b.ElrondAddress("erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s")
		PartialCreatePNG(Day, TxTaxAddy)
	}

}
