package cmd

import (
	"context"
	elog "github.com/labstack/gommon/log"
	config2 "github.com/nmapx/ddns-daemon/config"
	"github.com/nmapx/ddns-daemon/ifconfig"
	"github.com/nmapx/ddns-daemon/ovh"
	"github.com/spf13/cobra"
	"time"
)

var (
	configFilepath string
	sleepError     int
	sleepSuccess   int
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Background process",
	Long: `It's fetching your IP address from https://ifconfig.co API
then updating all your OVH DynHost hosts based on the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		handleSignals()

		elog.Infof("Loading config")

		var config = config2.Config{}
		config.Load(configFilepath)

		elog.Infof("Starting")

		script(&config)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	daemonCmd.PersistentFlags().StringVar(&configFilepath, "config-filepath", "./config.yml", "Config filepath [yaml]")
	daemonCmd.PersistentFlags().IntVar(&sleepError, "sleep-error", 1, "Sleep after error [minutes]")
	daemonCmd.PersistentFlags().IntVar(&sleepSuccess, "sleep-success", 5, "Sleep after success [minutes]")
}

func script(config *config2.Config) {
	var prevIp string
	ifconfig.Inst.SetClient()

	for {
		ifConfigResponse, err := ifconfig.Inst.Fetch(context.Background())
		if err != nil {
			elog.Errorf("Problem with ifconfig request: %v", err)
			time.Sleep(time.Minute * time.Duration(sleepError))
			continue
		}

		elog.Infof("IP prev=%s current=%s", prevIp, ifConfigResponse.Ip)

		if len(prevIp) == 0 {
			elog.Infof("First run, missing data to compare, updating")
			for _, host := range config.Hosts {
				if err := hostProcess(context.Background(), host, ifConfigResponse.Ip); err != nil {
					elog.Error(err)
					time.Sleep(time.Minute * time.Duration(sleepError))
					continue
				}
			}
			prevIp = ifConfigResponse.Ip
			elog.Infof("Sleeping")
			time.Sleep(time.Minute * time.Duration(sleepSuccess))
			continue
		}

		if prevIp == ifConfigResponse.Ip {
			elog.Infof("No change, skipping")
			time.Sleep(time.Minute * time.Duration(sleepSuccess))
			continue
		}

		for _, host := range config.Hosts {
			if err := hostProcess(context.Background(), host, ifConfigResponse.Ip); err != nil {
				elog.Error(err)
				time.Sleep(time.Minute * time.Duration(sleepError))
				continue
			}
		}
		prevIp = ifConfigResponse.Ip
		elog.Infof("Sleeping")
		time.Sleep(time.Minute * time.Duration(sleepSuccess))
	}
}

func hostProcess(ctx context.Context, config config2.HostConfig, ip string) error {
	ovh.Inst.SetClient(config.User, config.Pass)
	err := ovh.Inst.Notify(ctx, ip, config.Host)
	if err != nil {
		return err
	}
	return nil
}
