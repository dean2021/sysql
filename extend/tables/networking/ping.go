package networking

import (
	"errors"
	"github.com/dean2021/sysql/table"
	probing "github.com/prometheus-community/pro-bing"
	"log"
)

func GenPing(context *table.QueryContext) (table.TableRows, error) {

	address := context.Constraints.GetAll("addr", table.EQUALS)
	if len(address) == 0 {
		return nil, errors.New("IP parameter cannot be empty")
	}

	countField := context.Constraints.GetAll("count", table.EQUALS)
	if len(countField) > 1 {
		return nil, errors.New("can only accept a single user_agent")
	}
	var count = 3
	if len(countField) > 0 {
		if c, ok := countField[0].(int64); ok {
			count = int(c)
		}
	}

	var results table.TableRows
	for _, addr := range address {
		ping, err := probing.NewPinger(addr.(string))
		if err != nil {
			log.Println(err)
			continue
		}
		ping.SetPrivileged(true)
		ping.Count = count
		ping.OnRecv = func(pkt *probing.Packet) {
			results = append(results, table.TableRow{
				"addr":     addr,
				"ip":       pkt.IPAddr.String(),
				"count":    count,
				"bytes":    pkt.Nbytes,
				"icmp_seq": pkt.Seq,
				"time":     pkt.Rtt.String(),
			})
		}
		err = ping.Run()
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return results, nil
}
