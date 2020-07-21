package main

import (
	"fmt"

	config "github.com/bhambri94/voluum-apis/configs"
	"github.com/bhambri94/voluum-apis/sheets"
	"github.com/bhambri94/voluum-apis/voluum"
)

func main() {
	config.SetConfig()
	
	values, _, SheetName := voluum.GetStandardVoluumReport()
	sheets.ClearSheet(SheetName)
	sheets.BatchWrite(SheetName, values)
}
