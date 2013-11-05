package mbus

import (
	"bosh/settings"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudfoundry/yagnats"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

type natsHandler struct {
	client   yagnats.NATSClient
	settings settings.Settings
}

func newNatsHandler(client yagnats.NATSClient, s settings.Settings) (handler natsHandler) {
	handler.client = client
	handler.settings = s
	return
}

func (h natsHandler) Run(handlerFunc HandlerFunc) (err error) {
	err = h.Start(handlerFunc)
	if err != nil {
		return
	}
	defer h.Stop()

	h.runUntilInterrupted()
	return
}

func (h natsHandler) Start(handlerFunc HandlerFunc) (err error) {
	connProvider, err := h.getConnectionInfo()
	if err != nil {
		return
	}

	err = h.client.Connect(connProvider)
	if err != nil {
		return
	}

	subject := fmt.Sprintf("agent.%s", h.settings.AgentId)

	h.client.Subscribe(subject, func(natsMsg *yagnats.Message) {
		req := Request{}
		err := json.Unmarshal([]byte(natsMsg.Payload), &req)
		if err != nil {
			return
		}

		resp := handlerFunc(req)

		respBytes, err := json.Marshal(resp)
		if err != nil {
			return
		}

		h.client.Publish(req.ReplyTo, string(respBytes))
	})

	return
}

func (h natsHandler) Stop() {
	h.client.Disconnect()
}

func (h natsHandler) runUntilInterrupted() {
	keepRunning := true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-c:
			keepRunning = false
		}
	}
}

func (h natsHandler) getConnectionInfo() (connInfo *yagnats.ConnectionInfo, err error) {
	natsUrl, err := url.Parse(h.settings.Mbus)
	if err != nil {
		return
	}

	user := natsUrl.User
	if user == nil {
		err = errors.New("No username or password set for connection")
		return
	}

	password, passwordIsSet := user.Password()
	if !passwordIsSet {
		err = errors.New("No password set for connection")
		return
	}

	connInfo = new(yagnats.ConnectionInfo)
	connInfo.Password = password
	connInfo.Username = user.Username()
	connInfo.Addr = natsUrl.Host

	return
}
