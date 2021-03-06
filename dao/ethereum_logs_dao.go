package dao

import (
	"github.com/ninjadotorg/handshake-ethereum/models"
	"log"
	"github.com/jinzhu/gorm"
	"time"
	"strings"
)

type EthereumLogsDao struct {
}

func (contractLogsDao EthereumLogsDao) GetById(id int64) (models.EthereumLogs) {
	dto := models.EthereumLogs{}
	err := models.Database().Where("id = ?", id).First(&dto).Error
	if err != nil {
		log.Print(err)
	}
	return dto
}

func (contractLogsDao EthereumLogsDao) GetByFilter(contractAddress string, event string) (models.EthereumLogs) {
	contractAddress = strings.ToLower(contractAddress)
	dto := models.EthereumLogs{}
	err := models.Database().Where("contract_address = ? AND event = ?", contractAddress, event).Order("block_number desc").First(&dto).Error
	if err != nil {
		log.Print(err)
	}
	return dto
}

func (contractLogsDao EthereumLogsDao) Create(dto models.EthereumLogs, tx *gorm.DB) (models.EthereumLogs, error) {
	if tx == nil {
		tx = models.Database()
	}
	dto.FromAddress = strings.ToLower(dto.FromAddress)
	dto.ContractAddress = strings.ToLower(dto.ContractAddress)
	dto.Hash = strings.ToLower(dto.Hash)
	dto.DateCreated = time.Now()
	dto.DateModified = dto.DateCreated
	err := tx.Create(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}

func (contractLogsDao EthereumLogsDao) Update(dto models.EthereumLogs, tx *gorm.DB) (models.EthereumLogs, error) {
	if tx == nil {
		tx = models.Database()
	}
	dto.FromAddress = strings.ToLower(dto.FromAddress)
	dto.ContractAddress = strings.ToLower(dto.ContractAddress)
	dto.Hash = strings.ToLower(dto.Hash)
	dto.DateModified = time.Now()
	err := tx.Save(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}

func (contractLogsDao EthereumLogsDao) Delete(dto models.EthereumLogs, tx *gorm.DB) (models.EthereumLogs, error) {
	if tx == nil {
		tx = models.Database()
	}
	err := tx.Delete(&dto).Error
	if err != nil {
		log.Println(err)
		return dto, err
	}
	return dto, nil
}
