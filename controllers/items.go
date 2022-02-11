package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Item struct {
	ID          uuid.UUID `json:"itemId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Amount      int       `json:"amount"`
}

var Inventory []Item

func CreateItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request body.",
		})
		return
	}
	newItem.ID, _ = uuid.NewUUID()

	Inventory = append(Inventory, newItem)
	c.IndentedJSON(http.StatusOK, newItem)
}

func ReadItems(c *gin.Context) {
	c.JSON(http.StatusOK, Inventory)
}

func UpdateItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	}

	for i := 0; i < len(Inventory); i++ {
		if newItem.ID == Inventory[i].ID {
			if newItem.Name != "" {
				Inventory[i].Name = newItem.Name
			}
			if newItem.Description != "" {
				Inventory[i].Description = newItem.Description
			}
			if string(newItem.Amount) != "" {
				Inventory[i].Amount = newItem.Amount
			}

			c.JSON(http.StatusOK, Inventory[i])
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid request.",
	})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The given id is not a valid uuid.",
		})
		return
	}

	for i := 0; i < len(Inventory); i++ {
		if uuid == Inventory[i].ID {
			Inventory = append(Inventory[:i], Inventory[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Item successfully deleted.",
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid request.",
	})
}
