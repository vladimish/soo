package node

import (
	"github.com/telf01/soo/pkg/configurator"
	"github.com/telf01/soo/pkg/logger"
	"github.com/telf01/soo/pkg/public_node/auth"
	"github.com/telf01/soo/pkg/public_node/network"
	"github.com/telf01/soo/pkg/public_node/network/containers"
	"github.com/telf01/soo/pkg/public_node/network/interfaces"
	"github.com/telf01/soo/pkg/public_node/persistence/auth_db"
	"net/http"
)

type ConnectionManager struct {
	n *network.Network
	a *auth.Auth
}

func NewConnectionManager(n *network.Network, a *auth.Auth) *ConnectionManager {
	db, err := auth_db.NewDB(configurator.Cfg.Persistence.Auth.ConnectionString)
	if err != nil {
		logger.L.Sugar().Fatal(err)
	}
	cm := ConnectionManager{
		n: network.NewNetwork(configurator.Cfg.NetworkConfiguration.Address, &http.Server{}),
		a: auth.NewAuth(db),
	}
	return &cm
}

func (cm *ConnectionManager) StartServer() {
	go cm.n.StartListening()
}

// BindConnectionHandlers links endpoints and their parsers.
func (cm *ConnectionManager) BindConnectionHandlers(handlers map[string]interfaces.RequestContainer) {
	for i := range handlers {
		cm.n.BindParser(i, handlers[i])
	}
}

func (cm *ConnectionManager) HandleRegister(c chan interface{}) {
	for {
		something := <-c
		rc := something.(containers.RegisterWrapper)

		node, err := cm.a.GetNodeOrNil(rc.R.NickName)
		if err != nil {
			logger.L.Sugar().Error(err)
		}
		if node != nil {
			authMessage, err := cm.a.BuildAuthMessage(node)
			if err != nil {
				logger.L.Sugar().Error(err)
			}
			err = cm.n.SendMessage(node, authMessage)
			if err != nil {
				logger.L.Sugar().Error(err)
			}
		} else {
			// TODO: Save auth and send it to user.
		}
	}
}
