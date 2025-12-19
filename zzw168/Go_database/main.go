package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model //内嵌gorm.Model
	Name       string
	// Age          sql.NullInt64 //零值类型
	Age          int64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小 255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 混略本字段
}

type User struct {
	gorm.Model        //内嵌gorm.Model
	Name       string `gorm:"default:'小王子'"`
	Age        int64
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	// //创建数据行
	// u1 := User{Age: 18}
	// fmt.Println(db.NewRecord(&u1))
	// db.Create(&u1)
	// fmt.Println(db.NewRecord(&u1))
	// u2 := User{Name: "小里", Age: 20}
	// db.Create(&u2)

	// 查询
	var u User
	db.Debug().First(&u)
	fmt.Println(u)

	// var us []User
	// // db.Debug().Where(&User{Age: 18}).Find(&u)
	// db.Debug().Find(&us)
	// fmt.Println(us)

	//更新
	// u.Name = "七米"
	// u.Age = 38
	// db.Debug().Save(&u)
	// m1 := map[string]interface{}{
	// 	"name": "双色球",
	// 	"age":  28,
	// }
	// db.Debug().Model(&u).Update(m1)               // m1列出的所有字段都会更新
	// db.Debug().Model(&u).Select("age").Update(m1) // 更新指定字段
	// db.Debug().Model(&u).Omit("age").Update(m1)   //排除m1中的active，更新其它字段

	// db.Debug().Model(&u).UpdateColumn("age", 30)
	// rowsNum := db.Debug().Model(User{}).Update(User{Name: "hello", Age: 18}).RowsAffected
	// fmt.Println(rowsNum)

	// 让users表中所有的用户年龄+2
	// var us User
	// db.Debug().Model(&us).Update("age", gorm.Expr("age+?", 2))

	//删除
	// db.Debug().Delete(&u)
	// db.Debug().Where(&User{Age: 22}).Delete(&User{})

	// 物理删除
	db.Debug().Unscoped().Where("name=?", "hello").Delete(User{})
}
