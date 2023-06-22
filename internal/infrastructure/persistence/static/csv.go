package static

import (
	"encoding/csv"
	"os"
)

type CSVManger struct {
	FilePath string
}

func (cm CSVManger) ReadAll() ([][]string, error) {
	f, err := os.Open(cm.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (cm CSVManger) Data() ([]string, error) {
	records, err := cm.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []string
	for i := range records {
		col := records[i]

		if len(col) == 0 {
			continue
		}

		for j := range col {
			cell := col[j]
			data = append(data, cell)
		}
	}

	return data, nil
}
