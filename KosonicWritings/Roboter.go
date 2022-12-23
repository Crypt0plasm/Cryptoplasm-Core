package KosonicWritings

import (
	"fmt"
	r "github.com/go-vgo/robotgo"
	bmp "github.com/vcaesar/bitmap"
)

type MouseCoordinates struct {
	x int
	y int
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
	r.MilliSleep(5000)
	SaveSelection(Name, true)
	return MC
}

// MakeInnerBorder loads a selection on position "Position",
// make a border of size "Size", modifies it through intersection with the original selection
// and saves it under the "Name" name
func MakeInnerBorder(Position int, Size, Name string) {
	r.MilliSleep(1000)
	LoadSelection(Position, "New", true)
	r.MilliSleep(1000)

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
	r.MilliSleep(1000)
	Proceed(true)

	//Intersecting with selection
	LoadSelection(Position, "Intersect", true)

	//Save Selection
	SaveSelection(Name, true)

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
	x0 := 3660
	y0 := 230
	y := y0 + 40*(Position-1)
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

// ConvertPhotoshop is the main function that manipulates the image
func ConvertPhotoshop() {

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
			MakeInnerBorder(1, "30", PositionToName(1)+"_Border")
			MakeInnerBorder(2, "20", PositionToName(2)+"_Border")
			MakeInnerBorder(3, "20", PositionToName(3)+"_Border")
			MakeInnerBorder(4, "20", PositionToName(4)+"_Border")
			MakeInnerBorder(5, "20", PositionToName(5)+"_Border")
			MakeInnerBorder(6, "20", PositionToName(6)+"_Border")
			MakeInnerBorder(7, "20", PositionToName(7)+"_Border")
			MakeInnerBorder(8, "20", PositionToName(8)+"_Border")
			MakeInnerBorder(9, "50", PositionToName(9)+"_Border")
			MakeInnerBorder(10, "25", PositionToName(10)+"_Border")
			MakeInnerBorder(11, "25", PositionToName(11)+"_Border")
			MakeInnerBorder(12, "15", PositionToName(12)+"_Border")
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
			MoveLeftClick(2287, 437)
			r.MilliSleep(5000)
		}
	}

	FillBackground := func(Position int, MO MouseCoordinates, Load bool) {
		LoadSelection(Position, "New", true)
		SelectPaintBucketPattern(Position)
		MoveLeftClick(MO.x, MO.y)
		r.MilliSleep(1000)
	}

	//MinimizeGoland()
	//r.MilliSleep(1000)
	ClickPhotoshopWindow()
	r.MilliSleep(1000)

	//FitToScreen Image
	r.KeyTap("0", "ctrl")
	//Click Layer1
	MoveLeftClick(3700, 230)
	//Select Magic Wand Tool
	r.KeyTap("w")

	//Select Units
	MC := SelectUnits(true)

	//Make Inner Border
	MakeBorders(true)

	//Make All-Border Selection
	AddSelections(14, 26)
	SaveSelection("All_Border", true)

	//Paint Backgrounds
	for i := 1; i < 14; i++ {
		MakeNewLayer(PositionToName(i), true)
		HideLayer(2)
		SelectLayer(1)
		FillBackground(i, MC[i-1], true)
	}

	//Load All_Border Selection - starts making the border black
	FillBorder(true)

	fmt.Println("Finish Script")

}
