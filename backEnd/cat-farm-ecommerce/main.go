package main

import (
	_ "database/sql"
	_ "log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// สร้าง struct สำหรับข้อมูลแมว
type Cat struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Breed  string `json:"breed"`
    // ... เพิ่มฟิลด์อื่นๆ เช่น อายุ, เพศ
}

func main() {
    // ... โค้ดการเชื่อมต่อฐานข้อมูล

    r := gin.Default()

    // ตั้งค่า CORS เพื่อให้ Frontend เรียกใช้งานได้
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    })

    // สร้าง API Endpoint สำหรับดึงข้อมูลแมวทั้งหมด
    r.GET("/api/cats", func(c *gin.Context) {
        // โค้ดสำหรับดึงข้อมูลจากฐานข้อมูล
        // ตัวอย่างข้อมูลจำลอง
        cats := []Cat{
            {ID: 1, Name: "Milo", Breed: "Siamese"},
            {ID: 2, Name: "Luna", Breed: "Persian"},
        }
        c.JSON(http.StatusOK, cats)
    })

    // รันเซิร์ฟเวอร์บนพอร์ต 8080
    r.Run(":8080")
}