package main

import (
	"fmt"
	"strings"
)

type User struct {
	Username string
	Password string
}

type InventoryItem struct {
	ID    int
	Name  string
	Price float64
}

type Inventory struct {
	items []*InventoryItem
}

func (inv *Inventory) addItem(item *InventoryItem) {
	inv.items = append(inv.items, item)
}

func (inv *Inventory) getItemByID(id int) *InventoryItem {
	for _, item := range inv.items {
		if item.ID == id {
			return item
		}
	}
	return nil
}

func main() {
	// Create a user
	user := User{
		Username: "admin",
		Password: "password",
	}

	// Authenticate the user
	var username string
	var password string

	fmt.Println("Inventory Management System")
	fmt.Println("---------------------------")

	for {
		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		if strings.TrimSpace(username) == user.Username && strings.TrimSpace(password) == user.Password {
			fmt.Println("Authentication successful!")
			break
		} else {
			fmt.Println("Authentication failed. Please try again.")
		}
	}

	// User authenticated, proceed with the inventory management

	inventory := &Inventory{}

	for {
		fmt.Println()
		fmt.Println("1. Add item")
		fmt.Println("2. Get item by ID")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var itemID int
			var itemName string
			var itemPrice float64

			fmt.Print("Enter item ID: ")
			fmt.Scanln(&itemID)
			fmt.Print("Enter item name: ")
			fmt.Scanln(&itemName)
			fmt.Print("Enter item price: ")
			fmt.Scanln(&itemPrice)

			item := &InventoryItem{
				ID:    itemID,
				Name:  itemName,
				Price: itemPrice,
			}
			inventory.addItem(item)
			fmt.Println("Item added successfully!")

		case 2:
			var itemID int
			fmt.Print("Enter item ID: ")
			fmt.Scanln(&itemID)

			item := inventory.getItemByID(itemID)
			if item != nil {
				fmt.Println("Item found:")
				fmt.Printf("ID: %d\nName: %s\nPrice: %.2f\n", item.ID, item.Name, item.Price)
			} else {
				fmt.Println("Item not found!")
			}

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
