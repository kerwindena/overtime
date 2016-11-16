package http

import (
	_ "github.com/kerwindena/overtime"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func apiGetDateRange(srv *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonenc := json.NewEncoder(w)
		query := r.URL.Query()

		// make sure the begin is set
		begin, err := strconv.Atoi(query.Get("begin"))
		if err != nil {
			fmt.Fprintln(w, "Must supply begin")
			return
		}

		// make sure the end is set
		end, err := strconv.Atoi(query.Get("end"))
		if err != nil {
			fmt.Fprintln(w, "Must supply end")
			return
		}

		// make sure the begin and the end are lower and upper bounds
		if begin > end {
			fmt.Fprintln(w, "begin must be lower than end")
			return
		}

		// initialize the slice to return
		size := end - begin + 1
		dates := make([]ApiDate, size, size)

		// fill the sice to send
		for i := begin; i <= end; i++ {
			date := srv.dateStorage.Date(i)
			dates[i-begin] = ApiDate{
				Date: date.String(),
				Day:  date.Day(),
			}
		}

		// send it as json
		jsonenc.Encode(dates)
	}
}
