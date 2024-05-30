package networking

import (
	"errors"
	"github.com/dean2021/sysql/table"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GenCurl(context *table.QueryContext) (table.TableRows, error) {

	requests := context.Constraints.GetAll("url", table.EQUALS)
	if len(requests) == 0 {
		return nil, errors.New("URL parameter cannot be empty")
	}

	userAgents := context.Constraints.GetAll("user_agent", table.EQUALS)
	if len(userAgents) > 1 {
		return nil, errors.New("can only accept a single user_agent")
	}

	if len(context.Constraints.GetAll("url", table.LIKE)) != 0 {
		return nil, errors.New("using LIKE clause for url is not supported")
	}

	var userAgent string

	if len(userAgents) > 0 {
		userAgent = userAgents[0].(string)
	} else {
		userAgent = "sysql"
	}

	var results table.TableRows
	for _, request := range requests {
		timeStart := time.Now().Unix()
		response, err := http.Get(request.(string))
		if err != nil {
			log.Println(err)
			continue
		}
		timeEnd := time.Now().Unix()
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("ioutil.ReadAll failed ,err:%v\n", err)
			continue
		}

		results = append(results, table.TableRow{
			"url":             request.(string),
			"method":          "GET",
			"user_agent":      userAgent,
			"response_code":   response.StatusCode,
			"round_trip_time": timeEnd - timeStart,
			"bytes":           len(body),
			"result":          string(body),
		})
	}

	return results, nil
}
