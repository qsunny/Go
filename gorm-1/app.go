package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

/*
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
 */

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Result struct {
	ID   int
	Name string
	Age  int
}


func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database ")
	}

	createUser(db)


}

func createProduct(db *gorm.DB) {
	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)

}

func createUser(db *gorm.DB) {
	// 迁移 schema
	db.AutoMigrate(&User{})

	var now = time.Now()
	var email = "aaron.qiu"
	user := User{Name: "Jinzhu", Email: &email, Age: 18, Birthday: &now}

	result := db.Create(&user) // 通过数据的指针来创建

	fmt.Println(result)


	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: &now},
		{Name: "Jackson", Age: 19, Birthday: &now},
	}

	result = db.Create(users) // pass a slice to insert multiple row
	fmt.Println(result)

	users = []*User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)
	// batch size 100
	//db.CreateInBatches(users, 100)

	for _, user := range users {
		var uid = (*user).ID
		fmt.Println(uid) // 1,2,3
	}

	// batch insert from `[]map[string]interface{}{}`
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "jinzhu_1", "Age": 18},
	//	{"Name": "jinzhu_2", "Age": 20},
	//})

	// 获取第一条记录（主键升序）
	db.First(&user)
	fmt.Println(user)

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(user)

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	fmt.Println(user)

	result = db.First(&user)
	affected := result.RowsAffected // 返回找到的记录数
	err := result.Error             // returns error or nil
	fmt.Println(affected)


	// 检查 ErrRecordNotFound 错误
	errors.Is(err, gorm.ErrRecordNotFound)

	user = User{}
	db.Where("name = ?", "jinzhu").First(&user)
	fmt.Println(user)

	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)

	db.Where([]int64{20, 21, 22}).Find(&users)

	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)

	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)


	db.Delete(&User{}, 3)


	var results Result
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 10).Scan(&results)
	fmt.Println(results)


	db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&results)
	fmt.Println(results)

	var age int
	db.Raw("SELECT SUM(age) FROM users WHERE email = ?", "aaron.qiu").Scan(&age)
	fmt.Println(age)

	var useres []User
	db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu20", 20).Scan(&useres)
	fmt.Println(useres)

	db.Exec("select * from users")

	var result3 Result
	db.Where("name = @name OR name = @name", map[string]interface{}{"name": "jinzhu20"}).First(&users)
	fmt.Println(result3)

	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	var sql1 = stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	var sqlVars = stmt.Vars         //=> []interface{}{1}
	fmt.Println(sql1)
	fmt.Println(sqlVars)

	sql2 := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
	})
	fmt.Println(sql2)

	var name string
	// 使用 GORM API 构建 SQL
	row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Println(name,age)


	// 使用原生 SQL
	row = db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
	row.Scan(&name, &age, &email)
	fmt.Println("=======使用原生 SQL==============")
	fmt.Println(name,age,email)


	// 使用 GORM API 构建 SQL
	rows, err := db.Model(&User{}).Where("name = ?", "Jinzhu").Select("name, age, email").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name, &age, &email)
		fmt.Println(name,age,email)
		// 业务逻辑...
	}

	// 原生 SQL
	rows, err = db.Raw("select name, age, email from users where name = ?", "Jackson").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name, &age, &email)
		fmt.Println(name,age,email)

		// 业务逻辑...
	}

	fmt.Println("==============scan rows===========================")

	rows, err = db.Model(&User{}).Where("name = ?", "Jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
	defer rows.Close()

	var user2 User
	for rows.Next() {
		// ScanRows 将一行扫描至 user
		db.ScanRows(rows, &user2)
		fmt.Println(user2)
		// 业务逻辑...
	}

	// 查找 user 时预加载相关 products
	//db.Preload("products").Find(&users)
}
