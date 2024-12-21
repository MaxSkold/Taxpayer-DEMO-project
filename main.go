package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Taxpayer struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	TIN     string  `json:"tin"` // ИНН
	Address string  `json:"address"`
	Income  float64 `json:"income"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=321_Resident_123 dbname=fns_db port=5432 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Taxpayer{})

	r := gin.Default()
	r.Static("/static", "./static")

	r.Use(gin.Logger())

	//Маршруты
	r.POST("/taxpayer", addTaxpayer)
	r.GET("/taxpayer", getAllTaxpayer)
	r.PUT("/taxpayer/:id", updateTaxpayer)
	r.DELETE("/taxpayer/:id", deleteTaxpayer)

	r.Run(":8080")
}

func addTaxpayer(c *gin.Context) {
	var newTaxpayer Taxpayer
	if err := c.ShouldBindJSON(&newTaxpayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newTaxpayer.Income < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Доход не может быть отрицательным"})
		return
	}

	var existingTaxpayer Taxpayer
	if err := db.Where("tin = ?", newTaxpayer.TIN).First(&existingTaxpayer).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Налогоплательщик с таким ИНН уже существует"})
		return
	}

	if err := db.Create(&newTaxpayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить наголоплательщика"})
		return
	}
	c.JSON(http.StatusCreated, newTaxpayer)
}

func getAllTaxpayer(c *gin.Context) {
	var taxpayers []Taxpayer
	if err := db.Find(&taxpayers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}
	c.JSON(http.StatusOK, taxpayers)
}

func updateTaxpayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	var taxpayer Taxpayer
	if err := db.First(&taxpayer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Налогоплательщик не найден"})
		return
	}
	var updatedTaxpayer Taxpayer
	if err := c.ShouldBindJSON(&updatedTaxpayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taxpayer.Name = updatedTaxpayer.Name
	taxpayer.TIN = updatedTaxpayer.TIN
	taxpayer.Address = updatedTaxpayer.Address
	taxpayer.Income = updatedTaxpayer.Income

	if err := db.Save(&taxpayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить данные"})
		return
	}
	c.JSON(http.StatusOK, taxpayer)
}

func deleteTaxpayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	if err := db.Delete(&Taxpayer{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить налогоплательщика"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Налогоплательщик удалён"})
}
