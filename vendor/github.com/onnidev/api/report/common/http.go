package common

import (
	"net/http"

	"github.com/xuri/excelize"
)

func SendThroughHTTP(xlsx *excelize.File, w http.ResponseWriter, filename string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename+".xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	xlsx.Write(w)
}
