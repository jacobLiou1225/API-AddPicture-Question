package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, _ := excelize.OpenFile("cat.xlsx")
	f.AddPicture("cat", "A1", "cat.png",
		&excelize.GraphicOptions{
			AutoFit:     true,
			ScaleX:      1,
			ScaleY:      1,
			OffsetX:     10,
			OffsetY:     20,
			Positioning: "absolute",
		},
	)
	f.SaveAs("output.xlsx")
}
