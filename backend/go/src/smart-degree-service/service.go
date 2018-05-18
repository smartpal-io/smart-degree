package main

import (
	log "github.com/sirupsen/logrus"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type SmartDegreeService interface {
	Register(studentId uint32, degreeHash [32]byte) (error)
	GetDegreeHash(studentId uint32) (degreeHash [32]byte, err error)
	Verify(studentId uint32, degreeHash [32]byte) (valid bool, err error)
}

func NewEthSmartDegreeService() SmartDegreeService {
	rpcClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	smartDegreeService := &ethSmartDegreeService{
		logger:          NewLogger("ethSmartDegreeService"),
		rpcClient:       rpcClient,
		contractAddress: "0x5af00cb5fa6405c1730eb00b5d94488d451db104",
	}
	smartDegreeService.logger.Debug("interacting with contract")
	smartDegreeService.connect()
	return smartDegreeService
}

type ethSmartDegreeService struct {
	logger          *log.Entry
	rpcClient       *ethclient.Client
	contractAddress string
	smartDegree     *SmartDegree
}

func (service *ethSmartDegreeService) connect() {
	smartDegree, err := NewSmartDegree(common.HexToAddress(service.contractAddress), service.rpcClient)

	if err != nil {
		service.logger.Error(err.Error())
		panic(err)
	}
	service.smartDegree = smartDegree

}

func (service *ethSmartDegreeService) Register(studentId uint32, degreeHash [32]byte) (err error) {
	h := make([]byte, 32)
	copy(h, degreeHash[:])
	service.logger.
		WithField("studentId", studentId).
		WithField("degreeHash", hex.EncodeToString(h)).
		Debug("entering Register")
	return
}

func (service *ethSmartDegreeService) GetDegreeHash(studentId uint32) (degreeHash [32]byte, err error) {
	service.logger.
		WithField("studentId", studentId).
		Debug("entering GetDegreeHash")
	degreeHash, err = service.smartDegree.GetHash(nil, studentId)
	h := make([]byte, 32)
	copy(h, degreeHash[:])
	service.logger.Debugf("GetHash output : %s", hex.EncodeToString(h))
	return
}

func (service *ethSmartDegreeService) Verify(studentId uint32, degreeHash [32]byte) (valid bool, err error) {
	h := make([]byte, 32)
	copy(h, degreeHash[:])
	service.logger.
		WithField("studentId", studentId).
		WithField("degreeHash", hex.EncodeToString(h)).
		Debug("entering Verify")

	return
}
