package delete

import (
	"fmt"
	. "rights/model"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func DeleteTimeHash(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["hash"] == nil {
		return ErrorReturn(t, 406, "000027", "Missing hash")
	}
	hash := p.Sanitize(fmt.Sprintf("%v", args["hash"]))

	param := ""
	if args["email"] != nil {
		param = p.Sanitize(fmt.Sprintf("%v", args["email"]))
	}
	if args["uid"] != nil {
		param = p.Sanitize(fmt.Sprintf("%v", args["uid"]))
	}

	if param == "" {
		return ErrorReturn(t, 406, "000027", "Missing Email or User ID")
	}

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

	db.Conn.Debug().Where("hash = ? AND (uid = ? OR email = ?)", hash, param, param).Delete(&data)

	ans["answer"] = "OK"

	response = Interfacetoresponse(t, ans)
	return response

}
