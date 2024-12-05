package gateway

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func Run(gw Gateway) error {
	addr := fmt.Sprintf(":%d", gw.Config.Port)
	gw.Telemeter.Logger.Info("Listening...", zap.String("address", addr))
	if err := http.ListenAndServe(addr, gw.handler); err != nil {
		return err
	}
	return nil
}
