package ovh

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	elog "github.com/labstack/gommon/log"
)

//Ovh struct
type Ovh struct {
	Client *resty.Client
}

//Inst Ovh instance
var Inst = &Ovh{}

//SetClient for Ovh requests
func (o *Ovh) SetClient(user string, pass string) {
	Inst.Client = resty.New().
		SetBaseURL("https://ovh.com").
		SetBasicAuth(user, pass)
}

//Notify Ovh with ip address
func (o *Ovh) Notify(ctx context.Context, ip string, host string) error {
	elog.Infof("Notifying host=%s ip=%s", host, ip)
	response, err := Inst.Client.R().
		SetContext(ctx).
		SetQueryParam("system", "dyndns").
		SetQueryParam("hostname", host).
		SetQueryParam("myip", ip).
		Get("/nic/update")

	if err != nil {
		return err
	}
	if response.StatusCode() == 401 {
		return errors.New(fmt.Sprintf("Invalid auth data for host=%s", host))
	}
	return nil
}
