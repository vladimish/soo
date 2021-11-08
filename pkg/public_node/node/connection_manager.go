package node

import (
	"github.com/vladimish/soo/pkg/configurator"
	"github.com/vladimish/soo/pkg/logger"
	"github.com/vladimish/soo/pkg/public_node/auth"
	"github.com/vladimish/soo/pkg/public_node/auth/models/responses"
	"github.com/vladimish/soo/pkg/public_node/network"
	"github.com/vladimish/soo/pkg/public_node/network/containers"
	"github.com/vladimish/soo/pkg/public_node/network/interfaces"
	node_models "github.com/vladimish/soo/pkg/public_node/node/models"
	"github.com/vladimish/soo/pkg/public_node/persistence/auth_db"
	"github.com/vladimish/soo/pkg/validation"
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

		// Validate request data
		err := validation.V.Struct(rc.R)
		if err != nil {
			logger.L.Sugar().Error(err)
			r := responses.Error{
				Code:    responses.BAD_REQUEST,
				Message: err.Error(),
			}

			rc.W.WriteHeader(http.StatusBadRequest)
			rc.W.Write([]byte(r.ToJSON()))
			rc.WG.Done()
			continue
		}

		node, err := cm.a.GetNodeOrNil(rc.R.Nickname)
		if err != nil {
			logger.L.Sugar().Error(err)
		}
		if node == nil {
			n := &node_models.Node{
				NickName: rc.R.Nickname,
				Status:   node_models.REGISTRATION,
			}
			cm.a.SaveNode(n)
			node = n
		}

		ad, err := cm.a.CreateAuth(node, rc.R.Key)
		authMessage, err := cm.a.BuildAuthMessage(*ad)
		if err != nil {
			logger.L.Sugar().Error(err)
		}
		err = cm.n.SendMessage(rc.W, authMessage)
		if err != nil {
			logger.L.Sugar().Error(err)
		}
		rc.WG.Done()
	}
}

func (cm *ConnectionManager) HandleVerifyRegister(c chan interface{}) {
	for {
		something := <-c
		vc := something.(containers.VerifyRegisterWrapper)

		authResult, err := cm.a.CheckAuth(vc.R.Signature, vc.R.CheckoutMessage)
		if err != nil {
			logger.L.Error(err.Error())
		}

		logger.L.Sugar().Info(authResult)

		vc.WG.Done()
	}
}
