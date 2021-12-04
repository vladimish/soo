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
	"sync"
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

func genericValidation(d interface{}, w http.ResponseWriter, wg *sync.WaitGroup) bool {
	err := validation.V.Struct(d)
	logger.L.Sugar().Info("Validating: ", d)
	if err != nil {
		logger.L.Sugar().Error(err)
		r := responses.Error{
			Code:    responses.BAD_REQUEST,
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(r.ToJSON()))
		wg.Done()

		return false
	}
	return true
}

func (cm *ConnectionManager) HandleRegister(c chan interface{}) {
	for {
		something := <-c
		rc := something.(containers.RegisterWrapper)

		// Validate request data
		if !genericValidation(rc.R, rc.W, rc.WG) {
			continue
		}
		//err := validation.V.Struct(rc.R)
		//if err != nil {
		//	logger.L.Sugar().Error(err)
		//	r := responses.Error{
		//		Code:    responses.BAD_REQUEST,
		//		Message: err.Error(),
		//	}
		//
		//	rc.W.WriteHeader(http.StatusBadRequest)
		//	rc.W.Write([]byte(r.ToJSON()))
		//	rc.WG.Done()
		//	continue
		//}

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

		rr := responses.RegistrationResult{
			Result: authResult,
		}

		if authResult {
			err := cm.a.ChangeNodeStatus(vc.R.CheckoutMessage, node_models.ACTIVE)
			if err != nil {
				logger.L.Sugar().Error(err)
			}
		}

		vc.W.Write([]byte(rr.ToJSON()))

		vc.WG.Done()
	}
}

func (cm *ConnectionManager) HandleFindUser(c chan interface{}) {
	for {
		something := <-c
		fu := something.(containers.FindUserWrapper)

		if !genericValidation(fu.R, fu.W, fu.WG) {
			continue
		}

		authResult, err := cm.a.CheckAuth(fu.R.VerifyRegister.Signature, fu.R.VerifyRegister.CheckoutMessage)
		if err != nil {
			logger.L.Error(err.Error())
		}

		if authResult {
			nodes, err := cm.a.GetNodesLikeOrNil(fu.R.SearchQuery, fu.R.ResultsAmount)
			if err != nil {
				logger.L.Sugar().Error(err)
			}

			urls := make([]string, len(nodes))
			for i := range nodes {
				urls[i] = nodes[i].NickName
			}
			rr := responses.SearchResult{
				URL: urls,
			}

			fu.W.Write([]byte(rr.ToJSON()))
		} else {
			fu.W.WriteHeader(http.StatusUnauthorized)
		}
		fu.WG.Done()
	}
}
