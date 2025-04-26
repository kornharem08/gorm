# GORM SQL Server Example

ต่อไปนี้คือตัวอย่างการใช้งานแต่ละฟังก์ชันใน ISQL interface

**1.Find(ctx context.Context, result any, conds ...any) error**

ใช้เพื่อดึงข้อมูลหลายรายการจากฐานข้อมูล

```bash
var users []User
err := sql.Find(ctx, &users, "age > ?", 20)
if err != nil {
	fmt.Println("Error finding users:", err)
	return
}
fmt.Println("Found users:", users)
```

คำอธิบาย:

- ดึงข้อมูลผู้ใช้ทั้งหมดที่มี age > 20

- result ต้องเป็น pointer ไปยัง slice เพื่อเก็บผลลัพธ์

- conds ใช้สำหรับเงื่อนไข เช่น WHERE age > 20

**2.First(ctx context.Context, result any, conds ...any) error**

ใช้เพื่อดึงข้อมูลรายการแรกที่ตรงกับเงื่อนไข

```bash
var user User
err := sql.First(ctx, &user, "name = ?", "John")
if err != nil {
	fmt.Println("Error finding first user:", err)
	return
}
fmt.Println("First user:", user)
```

คำอธิบาย:

- ดึงผู้ใช้คนแรกที่ชื่อ "John"
- หากไม่มีข้อมูลจะคืน error gorm.ErrRecordNotFound
- result ต้องเป็น pointer ไปยัง struct

**3.reate(ctx context.Context, value any) error**

ใช้เพื่อสร้าง record ใหม่ในฐานข้อมูล

```bash
var user User
err := sql.First(ctx, &user, "name = ?", "John")
if err != nil {
	fmt.Println("Error finding first user:", err)
	return
}
fmt.Println("First user:", user)
```

คำอธิบาย:

- สร้างผู้ใช้ใหม่ในตาราง users
- value ต้องเป็น pointer ไปยัง struct หรือ slice ของ struct
- GORM จะกำหนด primary key (เช่น ID) อัตโนมัติ

**4.Update(ctx context.Context, column string, value any) error**

ใช้เพื่ออัปเดต column เดียวของ record ที่เลือก

```bash
err := sql.Where(ctx, "name = ?", "Alice").Update(ctx, "age", 26)
if err != nil {
	fmt.Println("Error updating user:", err)
	return
}
fmt.Println("Updated user's age")
```

คำอธิบาย:

- อัปเดต column age เป็น 26 สำหรับผู้ใช้ที่ชื่อ "Alice"
- ต้องใช้ Where เพื่อกำหนด record ที่ต้องการอัปเดต
- หากไม่ระบุ Where จะอัปเดตทุก record ในตาราง (ระวัง!)

**5.Updates(ctx context.Context, value any) error**

ใช้เพื่ออัปเดตหลาย column โดยใช้ struct หรือ map

```bash
updates := map[string]interface{}{
	"age":   27,
	"email": "alice.new@example.com",
}
err := sql.Where(ctx, "name = ?", "Alice").Updates(ctx, updates)
if err != nil {
	fmt.Println("Error updating user:", err)
	return
}
fmt.Println("Updated user's age and email")
```

คำอธิบาย:

- อัปเดต age และ email สำหรับผู้ใช้ที่ชื่อ "Alice"
- value สามารถเป็น map หรือ struct
- GORM จะอัปเดตเฉพาะ field ที่ระบุใน value

**6.Delete(ctx context.Context, conds ...any) error**

ใช้เพื่อลบ record ที่ตรงกับเงื่อนไข

```bash
err := sql.Where(ctx, "name = ?", "Alice").Delete(ctx)
if err != nil {
	fmt.Println("Error deleting user:", err)
	return
}
fmt.Println("Deleted user")
```

คำอธิบาย:

- ลบผู้ใช้ที่ชื่อ "Alice"
- ต้องใช้ Where เพื่อระบุ record ที่จะลบ
- หากไม่ระบุ Where จะลบทุก record ในตาราง (ระวัง!)

**7.Joins(ctx context.Context, query string, args ...any) ISQL**

ใช้เพื่อเพิ่ม JOIN clause ในการ query

```bash
var users []User
err := sql.Joins(ctx, "INNER JOIN profiles ON profiles.user_id = users.id").
	Where(ctx, "profiles.active = ?", true).
	Find(ctx, &users)
if err != nil {
	fmt.Println("Error joining and finding users:", err)
	return
}
fmt.Println("Users with active profiles:", users)
```

คำอธิบาย:

- ทำ JOIN กับตาราง profiles และดึงผู้ใช้ที่มี profiles.active = true
- คืนค่า ISQL เพื่อให้ chain กับ method อื่น เช่น Where, Find
- args ใช้สำหรับ parameterized queries ใน JOIN

**8.Where(ctx context.Context, query any, args ...any) ISQL**

ใช้เพื่อเพิ่ม WHERE clause ในการ query

```bash
var users []User
err := sql.Where(ctx, "age > ? AND email LIKE ?", 20, "%@example.com").
	Find(ctx, &users)
if err != nil {
	fmt.Println("Error finding users:", err)
	return
}
fmt.Println("Filtered users:", users)
```

คำอธิบาย:

- กรองผู้ใช้ที่มี age > 20 และ email ลงท้ายด้วย @example.com
- คืนค่า ISQL เพื่อให้ chain กับ method อื่น
- query สามารถเป็น string หรือ struct

**9.Preload(ctx context.Context, query string, args ...any) ISQL**

ใช้เพื่อ preload associations (เช่น relation ใน GORM)

```bash
var user User
err := sql.Preload(ctx, "Profile").First(ctx, &user, "name = ?", "John")
if err != nil {
	fmt.Println("Error preloading profile:", err)
	return
}
fmt.Println("User with profile:", user)
```

คำอธิบาย:

- ดึงผู้ใช้ที่ชื่อ "John" พร้อม preload ข้อมูลจากตาราง Profile (สมมติว่า User มี relation Profile)
- คืนค่า ISQL เพื่อให้ chain ได้
- args ใช้สำหรับเงื่อนไขใน preload

**10.Order(ctx context.Context, value string) ISQ**

ใช้เพื่อกำหนดลำดับของผลลัพธ์

```bash
var users []User
err := sql.Order(ctx, "age DESC").Find(ctx, &users)
if err != nil {
	fmt.Println("Error ordering users:", err)
	return
}
fmt.Println("Users ordered by age:", users)
```

คำอธิบาย:

- ดึงผู้ใช้ทั้งหมดโดยเรียงตาม age จากมากไปน้อย
- คืนค่า ISQL เพื่อให้ chain ได้
- value เช่น "age DESC", "name ASC"

**11.Limit(ctx context.Context, limit int) ISQL**

ใช้เพื่อจำกัดจำนวน record ที่ดึงมา

```bash
var users []User
err := sql.Limit(ctx, 5).Find(ctx, &users)
if err != nil {
	fmt.Println("Error limiting users:", err)
	return
}
fmt.Println("Top 5 users:", users)
```

คำอธิบาย:

- ดึงผู้ใช้ 5 คนแรก
- คืนค่า ISQL เพื่อให้ chain ได้
- มักใช้คู่กับ Order เพื่อควบคุมลำดับ

**12. Offset(ctx context.Context, offset int) ISQL**

ใช้เพื่อข้าม record จำนวนหนึ่ง (pagination)

```bash
var users []User
err := sql.Limit(ctx, 5).Offset(ctx, 10).Find(ctx, &users)
if err != nil {
	fmt.Println("Error paginating users:", err)
	return
}
fmt.Println("Users (page 3, 5 per page):", users)
```

คำอธิบาย:

- ข้าม 10 record แรกและดึง 5 record ถัดไป (เหมือน page 3 ใน pagination)
- คืนค่า ISQL เพื่อให้ chain ได้
- มักใช้คู่กับ Limit

**13. Raw(ctx context.Context, sql string, values ...any) ISQL**

ใช้เพื่อรัน raw SQL query

```bash
var users []User
err := sql.Raw(ctx, "SELECT * FROM users WHERE age > ?", 20).Find(ctx, &users)
if err != nil {
	fmt.Println("Error executing raw query:", err)
	return
}
fmt.Println("Users from raw query:", users)
```

คำอธิบาย:

- รัน SQL query โดยตรงเพื่อดึงผู้ใช้ที่มี age > 20
- คืนค่า ISQL เพื่อให้ chain กับ Find หรือ method อื่น
- values ใช้สำหรับ parameterized queries เพื่อป้องกัน SQL injection

**14. Exec(ctx context.Context) error**

ใช้เพื่อรัน query ที่สะสมไว้ (เช่น หลังจาก Raw หรือ Where)

```bash
err := sql.Raw(ctx, "UPDATE users SET age = age + 1 WHERE name = ?", "John").Exec(ctx)
if err != nil {
	fmt.Println("Error executing update:", err)
	return
}
fmt.Println("Updated age for John")
```

คำอธิบาย:

- รัน query ที่สร้างจาก Raw หรือ chain อื่น
- ใช้เมื่อ query ไม่ได้คืนผลลัพธ์ (เช่น UPDATE, DELETE)
- คืน error หาก query ล้มเหลว

**15. Transaction(ctx context.Context, fc func(tx ISQL) error) error**

ใช้เพื่อรัน operation ภายใน transaction

```bash
err := sql.Transaction(ctx, func(tx sqlwrap.ISQL) error {
	if err := tx.Create(ctx, &User{Name: "Bob", Age: 30, Email: "bob@example.com"}); err != nil {
		return err
	}
	return tx.Where(ctx, "name = ?", "Bob").Update(ctx, Playa del Rey
	if err != nil {
		return err
	}
	fmt.Println("User created and updated in transaction")
	return nil
})
if err != nil {
	fmt.Println("Transaction failed:", err)
}
```

คำอธิบาย:

- รัน operation ภายใน transaction เพื่อให้ทุกอย่างสำเร็จหรือยกเลิกทั้งหมด
- fc รับ ISQL เป็น parameter ทำให้สามารถใช้ method อื่นใน transaction ได้
- หากมี error จะ rollback อัตโนมัติ


** โค้ดรวมทั้งหมด **

```bash
package main

import (
	"context"
	"fmt"
	"yourmodule/sqlwrap"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Age   int
	Email string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&User{})
	sql := sqlwrap.NewSQLTable(db)
	ctx := context.Background()

	// Find
	var users []User
	err = sql.Find(ctx, &users, "age > ?", 20)
	if err != nil {
		fmt.Println("Error finding users:", err)
		return
	}
	fmt.Println("Found users:", users)

	// First
	var user User
	err = sql.First(ctx, &user, "name = ?", "John")
	if err != nil {
		fmt.Println("Error finding first user:", err)
		return
	}
	fmt.Println("First user:", user)

	// Create
	newUser := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	err = sql.Create(ctx, &newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("Created user:", newUser)

	// Update
	err = sql.Where(ctx, "name = ?", "Alice").Update(ctx, "age", 26)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Println("Updated user's age")

	// Updates
	updates := map[string]interface{}{
		"age":   27,
		"email": "alice.new@example.com",
	}
	err = sql.Where(ctx, "name = ?", "Alice").Updates(ctx, updates)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Println("Updated user's age and email")

	// Delete
	err = sql.Where(ctx, "name = ?", "Alice").Delete(ctx)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	fmt.Println("Deleted user")

	// Joins
	err = sql.Joins(ctx, "INNER JOIN profiles ON profiles.user_id = users.id").
		Where(ctx, "profiles.active = ?", true).
		Find(ctx, &users)
	if err != nil {
		fmt.Println("Error joining and finding users:", err)
		return
	}
	fmt.Println("Users with active profiles:", users)

	// Where
	err = sql.Where(ctx, "age > ? AND email LIKE ?", 20, "%@example.com").
		Find(ctx, &users)
	if err != nil {
		fmt.Println("Error finding users:", err)
		return
	}
	fmt.Println("Filtered users:", users)

	// Preload
	err = sql.Preload(ctx, "Profile").First(ctx, &user, "name = ?", "John")
	if err != nil {
		fmt.Println("Error preloading profile:", err)
		return
	}
	fmt.Println("User with profile:", user)

	// Order
	err = sql.Order(ctx, "age DESC").Find(ctx, &users)
	if err != nil {
		fmt.Println("Error ordering users:", err)
		return
	}
	fmt.Println("Users ordered by age:", users)

	// Limit
	err = sql.Limit(ctx, 5).Find(ctx, &users)
	if err != nil {
		fmt.Println("Error limiting users:", err)
		return
	}
	fmt.Println("Top 5 users:", users)

	// Offset
	err = sql.Limit(ctx, 5).Offset(ctx, 10).Find(ctx, &users)
	if err != nil {
		fmt.Println("Error paginating users:", err)
		return
	}
	fmt.Println("Users (page 3, 5 per page):", users)

	// Raw
	err = sql.Raw(ctx, "SELECT * FROM users WHERE age > ?", 20).Find(ctx, &users)
	if err != nil {
		fmt.Println("Error executing raw query:", err)
		return
	}
	fmt.Println("Users from raw query:", users)

	// Exec
	err = sql.Raw(ctx, "UPDATE users SET age = age + 1 WHERE name = ?", "John").Exec(ctx)
	if err != nil {
		fmt.Println("Error executing update:", err)
		return
	}
	fmt.Println("Updated age for John")

	// Transaction
	err = sql.Transaction(ctx, func(tx sqlwrap.ISQL) error {
		if err := tx.Create(ctx, &User{Name: "Bob", Age: 30, Email: "bob@example.com"}); err != nil {
			return err
		}
		return tx.Where(ctx, "name = ?", "Bob").Update(ctx, "age", 31).Exec(ctx)
	})
	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}
	fmt.Println("User created and updated in transaction")
}
```