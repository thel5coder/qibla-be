package excelize

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// Generate ...
func Generate(data *[]map[string]string, location string) (err error) {
	f := excelize.NewFile()
	for _, d := range *data {
		for k, v := range d {
			f.SetCellValue("Sheet1", k, v)
		}
	}

	if err := f.SaveAs(location); err != nil {
		return err
	}

	return err
}
