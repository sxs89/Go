package admin

import (
	"fmt"
	"gin_zhenghe/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type StudentController struct {
}

type Student struct {
	Id        int64      `xorm:"pk autoincr" json:"id"`
	Name      string     `xorm:"unique" json:"name"`
	Age       int64      `json:"age"`
	CreatedAt *time.Time `xorm:"created" json:"created_at"`
	UpdatedAt *time.Time `xorm:"updated" json:"updated_at"`
}

// 创建student表
func (con StudentController) InitStudent(c *gin.Context) {
	errResult := utils.DB.Sync(new(Student))
	if errResult != nil {
		fmt.Println("同步数据失败！")
	}
	c.String(200, "创建student成功")
}

// 查询student
func (con StudentController) StudentHome(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	studentList, _ := utils.DB.QueryString("select * from student")
	c.JSON(200, studentList)
}

// 根据条件(name=php)查询数据
func (con StudentController) GetByName(c *gin.Context) {
	var student = Student{Name: "php"}
	utils.DB.Where("o.name=?", student.Name).Get(&student)
	c.JSON(200, student)
}

// 添加数据
func (con StudentController) AddStudent(c *gin.Context) {
	var student []Student
	student = append(student, Student{Name: "Java2", Age: 10})
	student = append(student, Student{Name: "golang", Age: 11})
	n, err := utils.DB.Insert(&student)
	if n > 0 {
		c.JSON(200, n)
	} else {
		fmt.Println(err)
	}
}

// 更新数据
func (con StudentController) UpdateStudent(c *gin.Context) {
	student := Student{Id: 2, Name: "php-002, mydsql"}
	n, _ := utils.DB.Id(student.Id).Update(&student)
	if n > 0 {
		c.JSON(200, gin.H{
			"msg":    "更新数据成功",
			"number": n,
		})
	}
}

// 根据id删除数据
func (con StudentController) DelteById(c *gin.Context) {
	student := Student{}
	n, _ := utils.DB.Id(2).Delete(&student)
	if n > 0 {
		c.JSON(200, gin.H{
			"msg":    "删除数据成功",
			"number": n,
		})
	}
}
