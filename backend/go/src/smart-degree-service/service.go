package main

import (
	log "github.com/sirupsen/logrus"
	"encoding/hex"
)

type SmartDegreeService interface {
	Register(studentId int64, degreeHash []byte) (error)
	GetDegreeHash(studentId int64) (degreeHash []byte, err error)
	Verify(studentId int64, degreeHash []byte) (valid bool, err error)
}

func NewEthSmartDegreeService() SmartDegreeService {
	return &ethSmartDegreeService{
		logger: NewLogger("ethSmartDegreeService"),
	}
}

type ethSmartDegreeService struct {
	logger *log.Entry
}

func (service ethSmartDegreeService) Register(studentId int64, degreeHash []byte) (err error) {
	service.logger.
		WithField("studentId", studentId).
		WithField("degreeHash", hex.EncodeToString(degreeHash)).
		Debug("entering Register")
	return
}

func (service ethSmartDegreeService) GetDegreeHash(studentId int64) (degreeHash []byte, err error) {
	service.logger.
		WithField("studentId", studentId).
		Debug("entering GetDegreeHash")
	return
}

func (service ethSmartDegreeService) Verify(studentId int64, degreeHash []byte) (valid bool, err error) {
	service.logger.
		WithField("studentId", studentId).
		WithField("degreeHash", hex.EncodeToString(degreeHash)).
		Debug("entering Verify")

	return
}
