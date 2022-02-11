package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

type Shipment struct {
	ID       uuid.UUID `json:"shipmentId,omitempty"`
	Address  string    `json:"address"`
	Contents []struct {
		ItemID uuid.UUID `json:"itemId"`
		Amount int       `json:"amount"`
	} `json:"contents"`
}

var Shipments []Shipment
var Mutex sync.Mutex

func CreateShipment(c *gin.Context) {
	var newShipment Shipment
	if err := c.BindJSON(&newShipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request body.",
		})
		return
	}
	newShipment.ID, _ = uuid.NewUUID()

	//Verify that inventory amounts can meet shipment requests and modify inventory; remove shipment contents that have unavailable amounts.
	//Lock to prevent race conditions when reading and writing to inventory.
	Mutex.Lock()
	for i := 0; i < len(newShipment.Contents); i++ {
		for j := 0; j < len(Inventory); j++ {
			if Inventory[j].ID == newShipment.Contents[i].ItemID {
				if Inventory[j].Amount >= newShipment.Contents[i].Amount {
					Inventory[j].Amount -= newShipment.Contents[i].Amount
				} else {
					newShipment.Contents = append(newShipment.Contents[:i], newShipment.Contents[i+1:]...)
				}
			}
		}
	}
	Mutex.Unlock()

	if len(newShipment.Contents) > 0 {
		Shipments = append(Shipments, newShipment)
		c.JSON(http.StatusOK, newShipment)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Stock not available to meet the request.",
		})
		return
	}
}

func ReadShipments(c *gin.Context) {
	c.JSON(http.StatusOK, Shipments)
}
