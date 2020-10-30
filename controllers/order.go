package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"

	//"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"

	"CrudAPI/models"
)

/////////Add Order////////////
func OrderAdd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var order models.Orders
	c.BindJSON(&order)
  
	t := time.Now()
	formattedDateNow := fmt.Sprintf("%d%02d%02d%d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())  
	order_id := "OD-" + formattedDateNow
	addOrder := models.Orders{Create_dtm: t, Order_id: order_id, Phone: order.Phone, Name: order.Name, Address: order.Address, Menu: order.Menu, Total_item: order.Total_item, Pay: order.Pay,}
	if err := models.MPosGORM.Create(&addOrder).Error; err != nil {
	  fmt.Printf("error add Order: %3v \n", err)
	  c.AbortWithStatus(http.StatusInternalServerError)
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{
	  "phone": order.Phone,
	  "order_id": order_id,
	})
}

/////////detele Order Transaction///////////
func OrderDelete(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var order models.Orders
	c.BindJSON(&order)

	if err := models.MPosGORM.Where("phone = ? AND order_id = ?", order.Phone, order.Order_id).Delete(&models.Orders{}).Error; err != nil {
		fmt.Printf("error delete order: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
	  "phone": order.Phone,
	  "message": "delete ok",
	})
}


/////////Show Order///////////
func OrderShowByDate(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var order models.Orders
	c.BindJSON(&order)
  
	var orders []models.Orders
  
	if err := models.MPosGORM.Raw("SELECT * from orders where phone = ? AND date(create_dtm) = to_date(?, 'DD-Mon-YYYY')", order.Phone, order.Create_dtm).Scan(&orders).Error; err != nil {
		fmt.Printf("error show order by date: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (orders != nil) {
	  c.JSON(http.StatusOK, orders)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}


func OrderShowByPhone(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	phone := c.Query("phone")

	var orders []models.Orders
  
	if err := models.MPosGORM.Raw("SELECT * from orders where phone = ?", phone).Scan(&orders).Error; err != nil {
		fmt.Printf("error show order by phone: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (orders != nil) {
	  c.JSON(http.StatusOK, orders)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	


}

func OrderShowByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	order_id := c.Param("orderid")

	var orders []models.Orders
  
	if err := models.MPosGORM.Raw("SELECT * from orders where order_id = ?", order_id).Scan(&orders).Error; err != nil {
		fmt.Printf("error show order by id: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (orders != nil) {
	  c.JSON(http.StatusOK, orders)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	

}
  


/////////Edit Order///////////
func OrderEdit(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	//t := time.Now()
	//paymentDate := t.Format("02-Jan-2006 15:04")
  
	var order models.Orders
	c.BindJSON(&order)
  
	if err := models.MPosGORM.Model(&order).Where("phone = ? AND order_id = ?", order.Phone, order.Order_id).Updates(models.Orders{Menu: order.Menu, Total_item: order.Total_item, Pay: order.Pay}).Error; err != nil {
		fmt.Printf("error update sales payoff : %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//if err := models.MPosGORM.Model(&order).Where("user_id = ? AND outlet_id = ? AND sku_id = ?", sales.User_id, sales.Outlet_id, v.Sku_id).Update("number_sold", gorm.Expr("number_sold + ?", v.Number_orders)).Error; err != nil {
	//	fmt.Printf("error increasing product sold : %3v \n", err)
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}		
  
  
	c.JSON(http.StatusOK, gin.H{
	  "phone": order.Phone,
	  "message": "sales pay success",
	})
}