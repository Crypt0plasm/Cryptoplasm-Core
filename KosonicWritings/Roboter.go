package KosonicWritings

import (
	b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
	"errors"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	r "github.com/go-vgo/robotgo"
	bmp "github.com/vcaesar/bitmap"
	"math/rand"
	"os"
	"strings"
	"time"
)

type MouseCoordinates struct {
	x int
	y int
}

type CornerStroke struct {
	UL int
	UR int
	LL int
	LR int
}

type PhotoShopRectangle struct {
	Rectangle Rectangle
	Stroke    int
	Colour    RGBColor
	Corners   CornerStroke
}

// MoveNoClick Moves Mouse without clicking it
func MoveNoClick(x, y int) {
	r.MilliSleep(100)
	r.MoveSmooth(x, y, 0.1, 0.1, 5)
}

// MoveLeftClick Moves Mouse to x,y and left clicks it
func MoveLeftClick(x, y int) {
	r.MilliSleep(250)
	fmt.Println("Moving Mouse to: ", x, y)
	r.MoveSmooth(x, y, 0.1, 0.1, 5)

	r.MilliSleep(250)
	fmt.Println("Mouse LEFT Click")
	r.Click()
}

// MoveRightClick Moves Mouse to x,y and right clicks it
func MoveRightClick(x, y int) {
	r.MilliSleep(250)
	fmt.Println("Moving Mouse to: ", x, y)
	r.MoveSmooth(x, y, 0.1, 0.1, 5)

	r.MilliSleep(250)
	fmt.Println("Mouse RIGHT Click")
	r.Click("right")
}

// MinimizeGoland minimzez Goland Windows to reveal the Photoshop window
func MinimizeGoland() {
	//0001)Minimize Golang windows
	MoveLeftClick(3770, 30)
	return
}

// ClickPhotoshopWindow Clicks Photoshop window making sure it is active
func ClickPhotoshopWindow() {
	//0001)Minimize Golang windows
	fmt.Println("Activating Photoshop window")
	MoveLeftClick(2560, 30)
	return
}

// ClickContiguous Click the Contiguous tick on the Magic Wand Tool option (Default is off)
func ClickContiguous() {
	//Select Magic Wand Tool
	fmt.Println("Clicking Contiguous Magic Wand option")
	r.KeyTap("w")
	r.MilliSleep(1000)
	MoveLeftClick(1840, 64)
	r.MilliSleep(1000)
}

// SaveSelection Saves current selection with Name.
// Type OK means click OK Button
func SaveSelection(Name string, Load bool) {
	fmt.Println("Saving Selection ", Name)
	r.KeyTap("s", "alt")
	r.MilliSleep(250)
	r.KeyTap("v")
	r.MilliSleep(250)

	r.TypeStr(Name)
	r.MilliSleep(1000)

	Proceed(Load)

	r.KeyTap("d", "ctrl")
}

// LoadSelection loads the selection at the given Position
// Type = Add, Subtract, Intersect
// Load = OK
func LoadSelection(Position int, Type string, Load bool) {
	r.KeyTap("s", "alt")
	r.MilliSleep(250)
	r.KeyTap("o")
	r.MilliSleep(1500)

	KeyTapMany("tab", 2)

	//x0, y0 := ImageSearch("D:\\.AutoHotkey\\Source.png", "All")
	//x := x0 + 140
	//y := y0 + 65
	//MoveLeftClick(x, y)

	KeyTapMany("down", Position)
	//r.KeyTap("enter")

	//Check Type of Load
	switch LoadType := Type; {
	case LoadType == "New":
		fmt.Println("Loading New Selection")
	case LoadType == "Add":
		fmt.Println("Adding to Selection")
		KeyTapMany("tab", 2)
		r.MilliSleep(1000)
		r.KeyTap("down")
		r.MilliSleep(1000)
		r.KeyTap("space")
	case LoadType == "Subtract":
		fmt.Println("Subtracting from Selection")
		KeyTapMany("tab", 2)
		r.MilliSleep(1000)
		KeyTapMany("down", 2)
		r.MilliSleep(1000)
		r.KeyTap("space")
	case LoadType == "Intersect":
		fmt.Println("Intersecting with Selection")
		KeyTapMany("tab", 2)
		r.MilliSleep(1000)
		KeyTapMany("down", 3)
		r.MilliSleep(1000)
		r.KeyTap("space")
	}

	Proceed(Load)
}

func AddSelections(FirstPosition, LastPosition int) {
	for i := FirstPosition; i <= LastPosition; i++ {
		if i == FirstPosition {
			LoadSelection(i, "New", true)
		} else {
			LoadSelection(i, "Add", true)
		}
	}
}

func Proceed(Value bool) {
	if Value == true {
		r.KeyTap("enter")
	} else {
		r.KeyTap("escape")
	}
}

// KeyTapMany taps a key multiple times
func KeyTapMany(Key string, Times int) {
	for i := 1; i <= Times; i++ {
		r.KeyTap(Key)
		if i == Times {
			break
		}
	}
}

// ImageSearch searches for an image, Type = "Canvas" means only the Photoshop maximize Canvas is used
// If Type is anything else, all Screen is used.
func ImageSearch(Path, Type string) (x, y int) {
	var (
		Screen r.CBitmap
	)
	Image := bmp.Open(Path)
	if Type == "Canvas" {
		Screen = r.CaptureScreen(1665, 118, 1238, 1238)
		//ScreenTP := r.ToImage(Screen)
		//_ = r.Save(ScreenTP, "Screenshot.png")
		x0, y0 := bmp.Find(Image, Screen)
		x = x0 + 1665
		y = y0 + 118
	} else {
		x, y = bmp.Find(Image)
	}
	return x, y
}

// SelectAndSave searches for a color, makes a selection, and saves it
// Magic Wand Tool (w) must be selected beforehand
func SelectAndSave(Name string) MouseCoordinates {
	var MC MouseCoordinates
	Path := "D:\\.AutoHotkey\\" + Name + ".png"
	x, y := ImageSearch(Path, "Canvas")
	MC.x = x + 2
	MC.y = y + 2
	MoveLeftClick(MC.x, MC.y)
	r.MilliSleep(6000)
	SaveSelection(Name, true)
	return MC
}

// MakeInnerBorder loads a selection on position "Position",
// make a border of size "Size", modifies it through intersection with the original selection
// and saves it under the "Name" name
func MakeInnerBorder(Position int, Size, Name string) {
	r.MilliSleep(1000)
	LoadSelection(Position, "New", true)
	r.MilliSleep(2000)

	fmt.Println("Typing: ALT+s")
	r.KeyTap("s", "alt")
	r.MilliSleep(1000)

	fmt.Println("Typing: m")
	r.KeyTap("m")
	r.MilliSleep(1000)

	fmt.Println("Typing: b")
	r.KeyTap("b")
	r.MilliSleep(1000)

	fmt.Println("Typing: ", Size)
	r.TypeStr(Size)
	Proceed(true)
	r.MilliSleep(3000)

	//Intersecting with selection
	LoadSelection(Position, "Intersect", true)
	r.MilliSleep(3000)

	//Save Selection
	SaveSelection(Name, true)
	r.MilliSleep(2000)
	r.KeyTap("d", "ctrl")
	r.MilliSleep(2000)

}

func MakeNewLayer(Name string, Load bool) {
	fmt.Println("Making new Layer named ", Name)

	MoveLeftClick(3730, 1340)
	r.MilliSleep(1000)

	r.KeyTap("n", "shift", "ctrl")
	r.MilliSleep(1000)

	r.TypeStr(Name)
	r.MilliSleep(1000)

	Proceed(Load)
	r.MilliSleep(1000)
}

func SelectLayer(Position int) {
	fmt.Println("Selectin Layer ", Position)
	x0 := 3840
	y0 := 225
	y := y0 + 41*(Position-1)
	MoveLeftClick(x0, y)
	r.MilliSleep(1000)
}

func HideLayer(Position int) {
	SelectLayer(Position)
	r.KeyTap(",", "ctrl")
	r.MilliSleep(1000)
}

func SelectPaintBucketForeground() {
	r.KeyTap("g")
	r.MilliSleep(500)

	MoveLeftClick(1420, 64)
	r.MilliSleep(1000)

	MoveLeftClick(1420, 80)
	r.MilliSleep(1000)
}

func SelectPaintBucketPattern(Position int) {
	r.KeyTap("g")
	r.MilliSleep(500)

	MoveLeftClick(1420, 64)
	r.MilliSleep(1000)

	MoveLeftClick(1420, 100)
	r.MilliSleep(1000)

	MoveLeftClick(1464, 64)
	r.MilliSleep(1000)

	//Select Pattern
	y := 110 + (Position-1)*34
	MoveLeftClick(1500, y)
	r.MilliSleep(1000)
}

func PositionToName(Position int) (Result string) {
	switch Variant := Position; {
	case Variant == 1:
		Result = "Period"
	case Variant == 2:
		Result = "Number0"
	case Variant == 3:
		Result = "Number1"
	case Variant == 4:
		Result = "Number2"
	case Variant == 5:
		Result = "Number3"
	case Variant == 6:
		Result = "Number4"
	case Variant == 7:
		Result = "Number5"
	case Variant == 8:
		Result = "Number6"
	case Variant == 9:
		Result = "BlackBar"
	case Variant == 10:
		Result = "Bar"
	case Variant == 11:
		Result = "Bracket"
	case Variant == 12:
		Result = "Capital"
	case Variant == 13:
		Result = "Letter"
	}
	return Result
}

func RandomColour() RGBColor {
	Seed := p.NFS("256")
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	R1 := uint8(y1.Intn(int(p.INT64(Seed))))
	R2 := uint8(y1.Intn(int(p.INT64(Seed))))
	R3 := uint8(y1.Intn(int(p.INT64(Seed))))
	Colour := RGBColor{R: R1, G: R2, B: R3}
	return Colour

}
func OpenPNG(Path, Name string) {
	CompletePath := Path + Name + ".png"
	r.KeyTap("o", "ctrl")
	r.MilliSleep(3000)
	r.TypeStr(CompletePath)
	r.KeyTap("enter")
	WaitItUp("Waiting PNG Open", 40)
}

func ConvertCorelDraw(Path, Name, Extension string) {
	CompletePath := Path + Name + Extension
	//Bring to front CorelDrawWindow - assumes CorelDraw is already opened. And no files are opened within it.
	MoveLeftClick(510, 1420)
	r.MilliSleep(3000)

	//Open File
	r.KeyTap("o", "ctrl")
	r.MilliSleep(2000)
	r.TypeStr(CompletePath)
	r.MilliSleep(2000)
	r.KeyTap("enter")
	r.MilliSleep(3000)

	//F4 and Autofit
	r.KeyTap("f4")
	r.MilliSleep(3000)
	MoveLeftClick(1555, 120)
	MoveLeftClick(1600, 180)
	r.MilliSleep(3000)

	//Export as PNG
	r.KeyTap("e", "ctrl")
	r.MilliSleep(3000)
	r.TypeStr(Path + Name + ".png")
	r.MilliSleep(1000)
	r.KeyTap("enter")
	r.MilliSleep(3000)

	//Input PNG Size
	KeyTapMany("tab", 9)
	r.MilliSleep(1000)
	KeyTapMany("down", 4)
	r.MilliSleep(1000)
	r.KeyTap("enter")
	r.MilliSleep(1000)
	r.KeyTap("tab")
	r.MilliSleep(1000)
	r.TypeStr("10800")
	r.MilliSleep(1000)
	r.KeyTap("tab")
	r.MilliSleep(2000)
	r.KeyTap("enter")
	r.MilliSleep(2000)

	//Wait Up PNG Save
	//WaitItUp("Waiting PNG Save", 40)
	WaitUntilFileExists("Waiting PNG Save", Path+Name+".png")
	return
}

func WaitUntilFileExists(Text, FilePath string) {
	for i := 0; i <= 3600; i++ {
		BaseString := Text
		StringPoint := strings.Repeat(".", i)
		StringToPrint := BaseString + StringPoint
		if _, err := os.Stat(FilePath); errors.Is(err, os.ErrNotExist) {
			fmt.Print("\r", StringToPrint)
			r.MilliSleep(1000)
		} else {
			// file exists
			r.MilliSleep(2000)
			break
		}
	}
	fmt.Println("")
}

func WaitItUp(Text string, Time int) {
	for i := 0; i <= Time; i++ {
		BaseString := Text
		StringPoint := strings.Repeat(".", i)
		StringToPrint := BaseString + StringPoint
		fmt.Print("\r", StringToPrint)
		r.MilliSleep(1000)

	}
	fmt.Println("")
}

// ConvertPhotoshop is the main function that manipulates the image
func ConvertPhotoshop(Rectangles []Rectangle) {
	MainRectangle := Rectangle{p.NFS("150"), p.NFS("150"), p.NFS("10500"), p.NFS("10500")}
	MainRectangleCornerStroke := CornerStroke{50, 50, 50, 50}
	SecondaryRectangleCornerStroke := CornerStroke{20, 0, 20, 0}
	MainPSRectangle := PhotoShopRectangle{MainRectangle, 40, RandomColour(), MainRectangleCornerStroke}

	DrawPSRectangle := func(RectangleToDraw PhotoShopRectangle, Pattern int, StrokeType, Load bool) {
		if Load == true {
			x := 3335
			r.KeyTap("u")
			r.MilliSleep(1000)

			MoveLeftClick(1796, 121)
			r.MilliSleep(1000)
			r.KeyTap("enter")
			r.MilliSleep(1000)

			MoveLeftClick(x, 386)
			r.MilliSleep(1000)

			r.TypeStr(b.DTS(RectangleToDraw.Rectangle.Length) + " px")
			r.KeyTap("tab")
			r.MilliSleep(1000)
			r.TypeStr(b.DTS(RectangleToDraw.Rectangle.Height) + " px")
			r.KeyTap("tab")
			r.MilliSleep(1000)

			r.TypeStr(b.DTS(RectangleToDraw.Rectangle.X) + " px")
			r.KeyTap("tab")
			r.MilliSleep(1000)
			r.TypeStr(b.DTS(RectangleToDraw.Rectangle.Y) + " px")
			r.KeyTap("tab")
			r.MilliSleep(1000)

			r.TypeStr(b.DTS(p.NFI(int64(RectangleToDraw.Stroke))) + " px")
			r.KeyTap("tab")

			UL := b.DTS(p.NFI(int64(RectangleToDraw.Corners.UL)))
			UR := b.DTS(p.NFI(int64(RectangleToDraw.Corners.UR)))
			LL := b.DTS(p.NFI(int64(RectangleToDraw.Corners.LL)))
			LR := b.DTS(p.NFI(int64(RectangleToDraw.Corners.LR)))
			CornersString := UL + "px" + UR + "px" + LR + "px" + LL + "px"
			r.TypeStr(CornersString)
			r.MilliSleep(1000)

			//Fill
			MoveLeftClick(x, 440)
			r.MilliSleep(1000)
			if Pattern == 0 {
				//Choose Solid Color
				MoveLeftClick(x+60, 480)
				r.MilliSleep(1000)
				MoveLeftClick(x+200, 480)
				r.MilliSleep(1000)
				r.TypeStr(HexFromRGB(RectangleToDraw.Colour).Hex)
				r.MilliSleep(1000)
				r.KeyTap("enter")
			} else if Pattern != 0 {
				//Choose Pattern
				MoveLeftClick(x+140, 480)
				r.MilliSleep(1000)
				if Pattern == 1 {
					MoveLeftClick(x+50, 620)
					r.MilliSleep(1000)
				} else if Pattern == 2 {
					MoveLeftClick(x+80, 620)
					r.MilliSleep(1000)
				}
			}

			//Stroke
			MoveLeftClick(x+40, 440)
			r.MilliSleep(1000)
			if StrokeType == true {
				//Solid Color
				MoveLeftClick(x+60, 480)
				r.MilliSleep(1000)
				//Choose Color
				MoveLeftClick(x+200, 480)
				r.MilliSleep(1000)
				//Enter Color
				r.TypeStr("ffffff")
				r.MilliSleep(1000)
				//Confirm Color
				r.KeyTap("enter")
			} else {
				//No Stroke
				MoveLeftClick(x+20, 480)
				r.MilliSleep(1000)
			}

			//Modify Opacity
			if Pattern == 0 {
				//Opacity 25 if Pattern is 0
				SelectLayer(1)
				MoveLeftClick(x+535, 170)
				r.MilliSleep(3000)
				MoveLeftClick(x+535, 170)
				r.TypeStr("25")
				r.MilliSleep(1000)
				r.KeyTap("enter")
			} else {
				SelectLayer(1)
				MoveLeftClick(x+535, 170)
				r.MilliSleep(3000)
				MoveLeftClick(x+535, 170)
				r.TypeStr("50")
				r.MilliSleep(1000)
				r.KeyTap("enter")
			}
		}
	}

	SelectUnits := func(Load bool) []MouseCoordinates {
		var MCS = make([]MouseCoordinates, 13)
		if Load == true {
			MCS[0] = SelectAndSave(PositionToName(1))
			MCS[1] = SelectAndSave(PositionToName(2))
			MCS[2] = SelectAndSave(PositionToName(3))
			MCS[3] = SelectAndSave(PositionToName(4))
			MCS[4] = SelectAndSave(PositionToName(5))
			MCS[5] = SelectAndSave(PositionToName(6))
			MCS[6] = SelectAndSave(PositionToName(7))
			MCS[7] = SelectAndSave(PositionToName(8))
			MCS[8] = SelectAndSave(PositionToName(9))
			MCS[9] = SelectAndSave(PositionToName(10))
			MCS[10] = SelectAndSave(PositionToName(11))
			MCS[11] = SelectAndSave(PositionToName(12))
			MCS[12] = SelectAndSave(PositionToName(13))
		}
		return MCS
	}

	MakeBorders := func(Load bool) {
		if Load == true {
			MakeInnerBorder(1, "80", PositionToName(1)+"_Border")
			MakeInnerBorder(2, "30", PositionToName(2)+"_Border")
			MakeInnerBorder(3, "30", PositionToName(3)+"_Border")
			MakeInnerBorder(4, "30", PositionToName(4)+"_Border")
			MakeInnerBorder(5, "30", PositionToName(5)+"_Border")
			MakeInnerBorder(6, "30", PositionToName(6)+"_Border")
			MakeInnerBorder(7, "30", PositionToName(7)+"_Border")
			MakeInnerBorder(8, "30", PositionToName(8)+"_Border")
			MakeInnerBorder(9, "100", PositionToName(9)+"_Border")
			MakeInnerBorder(10, "50", PositionToName(10)+"_Border")
			MakeInnerBorder(11, "50", PositionToName(11)+"_Border")
			MakeInnerBorder(12, "20", PositionToName(12)+"_Border")
			MakeInnerBorder(13, "15", PositionToName(13)+"_Border")
		}
	}

	FillBorder := func(Load bool) {
		if Load == true {
			MakeNewLayer("Inner_Border", true)
			HideLayer(2)
			SelectLayer(1)
			LoadSelection(27, "New", true)

			SelectPaintBucketForeground()
			MoveLeftClick(2261, 434)
			r.MilliSleep(5000)
			r.KeyTap("d", "ctrl")
			HideLayer(1)
			r.MilliSleep(5000)
		}
	}

	FillBackground := func(Position int, MO MouseCoordinates, Load bool) {
		LoadSelection(Position, "New", true)
		SelectPaintBucketPattern(Position)
		MoveLeftClick(MO.x, MO.y)
		r.MilliSleep(1000)
	}

	PaintBackgrounds := func(MC []MouseCoordinates, Load bool) {
		if Load == true {
			for i := 1; i < 14; i++ {
				MakeNewLayer(PositionToName(i), true)
				HideLayer(2)
				SelectLayer(1)
				FillBackground(i, MC[i-1], true)
			}
		}
	}

	//MinimizeGoland()
	//r.MilliSleep(1000)
	ClickPhotoshopWindow()
	r.MilliSleep(1000)

	//FitToScreen Image
	r.KeyTap("0", "ctrl")
	//Click Layer1
	MoveLeftClick(3700, 230)
	//Select Magic Wand Too

	//Draw Rectangles, Round Rectangle must be preselected
	fmt.Println("HERE")
	PrimaryUpperPSRectangle := PhotoShopRectangle{Rectangles[0], 40, RandomColour(), MainRectangleCornerStroke}
	SecondaryUpperPSRectangle := PhotoShopRectangle{Rectangles[1], 40, RandomColour(), SecondaryRectangleCornerStroke}
	PrimaryMiddlePSRectangle := PhotoShopRectangle{Rectangles[2], 40, RandomColour(), MainRectangleCornerStroke}
	SecondaryMiddlePSRectangle := PhotoShopRectangle{Rectangles[3], 40, RandomColour(), SecondaryRectangleCornerStroke}
	PrimaryLowerPSRectangle := PhotoShopRectangle{Rectangles[4], 40, RandomColour(), MainRectangleCornerStroke}
	SecondaryLowerPSRectangle := PhotoShopRectangle{Rectangles[5], 40, RandomColour(), SecondaryRectangleCornerStroke}
	MintingPSRectangle := PhotoShopRectangle{Rectangles[6], 40, RandomColour(), MainRectangleCornerStroke}
	NodePSRectangle := PhotoShopRectangle{Rectangles[7], 40, RandomColour(), SecondaryRectangleCornerStroke}

	DrawAllRectangles := func(Load bool) {
		if Load == true {
			DrawPSRectangle(MainPSRectangle, 0, true, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(PrimaryUpperPSRectangle, 2, true, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(SecondaryUpperPSRectangle, 1, false, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(PrimaryMiddlePSRectangle, 2, true, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(SecondaryMiddlePSRectangle, 1, false, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(PrimaryLowerPSRectangle, 2, true, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(SecondaryLowerPSRectangle, 1, false, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(MintingPSRectangle, 2, true, true)
			r.KeyTap(",", "ctrl")
			DrawPSRectangle(NodePSRectangle, 2, true, true)
			r.KeyTap(",", "ctrl")
		}
	}
	GroupRectangles := func(Load bool) {
		if Load == true {
			//Group all Rectangles
			SelectLayer(10)
			r.KeyTap(",", "ctrl")
			r.MilliSleep(1000)
			r.KeyTap("a", "ctrl", "alt")
			r.KeyTap(",", "ctrl")

			_ = r.KeyToggle("ctrl")
			SelectLayer(10)
			_ = r.KeyToggle("ctrl", "up")
			r.MilliSleep(1500)
			r.KeyTap("g", "ctrl")
			r.KeyTap(",", "ctrl")
			SelectLayer(1)
			r.KeyTap("l", "alt")
			r.MilliSleep(500)
			r.KeyTap("a")
			r.MilliSleep(500)
			r.KeyTap("b")
			r.MilliSleep(500)
		}
	}

	GroupBackgrounds := func(Load bool) {
		if Load == true {
			//Group all Rectangles
			r.KeyTap("a", "ctrl", "alt")
			r.KeyTap(",", "ctrl")

			_ = r.KeyToggle("ctrl")
			SelectLayer(15)
			SelectLayer(16)
			_ = r.KeyToggle("ctrl", "up")
			r.KeyTap("g", "ctrl")

			r.KeyTap("l", "alt")
			r.MilliSleep(500)
			r.KeyTap("a")
			r.MilliSleep(500)
			r.KeyTap("k")
			r.MilliSleep(500)
		}
	}

	GroupLastGroups := func(Load bool) {
		if Load == true {
			r.KeyTap("a", "ctrl", "alt")
			_ = r.KeyToggle("ctrl")
			SelectLayer(1)
			_ = r.KeyToggle("ctrl", "up")
			r.MilliSleep(1000)
			r.KeyTap("g", "ctrl")
		}
	}

	BringToBack := func(Load bool) {
		if Load == true {
			SelectLayer(1)
			r.KeyTap("l", "alt")
			r.MilliSleep(500)
			r.KeyTap("a")
			r.MilliSleep(500)
			r.KeyTap("k")
			r.MilliSleep(500)
		}
	}

	DrawAllRectangles(true)
	GroupRectangles(true)

	r.KeyTap("w")
	SelectLayer(1)
	//Select Units
	MC := SelectUnits(true)

	//Make Inner Border
	MakeBorders(true)

	//Make All-Border Selection
	AddSelections(14, 26)
	SaveSelection("All_Border", true)
	r.KeyTap("d", "ctrl")

	//Paint Backgrounds
	PaintBackgrounds(MC, true)

	//Load All_Border Selection - starts making the border black
	FillBorder(true)

	//Group Backgrounds
	GroupBackgrounds(true)
	GroupLastGroups(true)
	BringToBack(true)

	fmt.Println("Finish Script")

}
