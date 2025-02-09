package utilClass

// 课程表
type Course struct {
	Id     uint    `gorm:"primaryKey;column:Id"`
	Name   string  `gorm:"size:20;not null;column:Name"`
	Credit float32 `gorm:"not null;column:Credit"`
}

// 专业表
type Major struct {
	MajorNum  uint   `gorm:"primaryKey;column:MajorNum"`
	MajorName string `gorm:"size:20;not null;column:MajorName"`
}

// 学生表
type Student struct {
	Id       uint    `gorm:"primaryKey;column:Id"`
	Name     string  `gorm:"size:10;not null;column:Name"`
	Password string  `gorm:"size:20;column:Password"`
	MajorNum uint    `gorm:"index;column:MajorNum"`
	Identity int8    `gorm:"default:1;column:Identity"`
	Credit   float32 `gorm:"default:0;column:Credit"`
	// 外键关系--也可以不显示的定义外键，gorm自动推断
	Major Major `gorm:"foreignKey:MajorNum;references:MajorNum"`
}

// 教师表
type Teacher struct {
	Id       uint   `gorm:"primaryKey;column:Id"`
	Name     string `gorm:"size:10; not null;column:Name"`
	Password string `gorm:"size:20;not null;column:Password"`
	Identity int8   `gorm:"default:2;column:Identity"`
}

// 教师课程表
type Tcourse struct {
	Id         uint   `gorm:"primaryKey;column:Id"`
	CourseNum  uint   `gorm:"index;not null;column:CourseNum"`
	TeacherNum uint   `gorm:"index;not null;column:TeacherNum"`
	Time       string `gorm:"size:255;default:'';column:Time"`
	Num        int    `gorm:"default:0;column:Num"`
	Total      int    `gorm:"NULL;column:Total"`
}

func (Tc Tcourse) TableName() string {
	return "t_course"
}

// 学生选课表
type StuCourse struct {
	Id         uint   `gorm:"primaryKey;column:Id"`
	StudentNum uint   `gorm:"not null;index;column:StudentNum"`
	TCourseNum uint   `gorm:"not null;index;column:TCourseNum"`
	Grade      string `gorm:"size:10;default:'暂未上传';column:Grade"`
	Time       string `gorm:"size:20;default:'';column:Time"`
	//外键
	Student Student `gorm:"foreignKey:StudentNum;references:Id"`
	Tcourse Tcourse `gorm:"foreignKey:TCourseNum;references:Id"`
}

func (stuC StuCourse) TableName() string {
	return "stu_course"
}
