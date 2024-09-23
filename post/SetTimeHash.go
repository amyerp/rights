package post

import (
	"encoding/json"
	"fmt"
	. "rights/model"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func SetTimeHash(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	data := TimeHash{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	data.Created = int(time.Now().Unix())
	waittime := 300
	if args["lifetime"] != nil {
		lt := p.Sanitize(fmt.Sprintf("%v", args["lifetime"]))
		waittime64, _ := strconv.ParseInt(lt, 10, 64)
		waittime = int(waittime64)
	}

	data.Lifetime = int(time.Now().Unix()) + waittime

	err = db.Conn.Create(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	ans["lifetime"] = data.Lifetime
	response = Interfacetoresponse(t, ans)

	return response

}
