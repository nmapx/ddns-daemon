package ifconfig

import (
	"context"
	"github.com/go-resty/resty/v2"
	elog "github.com/labstack/gommon/log"
)

type Ifconfig struct {
	Client *resty.Client
}

type Response struct {
	Ip string `json:"ip"`
}

var Inst = &Ifconfig{}

func (i *Ifconfig) SetClient() {
	Inst.Client = resty.New().
		SetBaseURL("https://ifconfig.co").
		SetHeader("Accept", "application/json")
}

func (i *Ifconfig) Fetch(ctx context.Context) (response *Response, error error) {
	elog.Infof("Fetching IP address")
	_, err := Inst.Client.R().
		SetContext(ctx).
		SetResult(&response).
		Get("")

	if err != nil {
		error = err
	}
	return
}
