package main	
		//package 声明文件属于哪个包  main: 入口函数 没有返回值 代码会编译成一个可执行文件 
import "fmt"    导入语句 
		//函数外只能放标识符(变量 常量 函数 类型)的声明  每个语句都必须以"关键字"开头 

func main(){  	 //入口函数 
## 声明变量  
	初始值为空  非全局变量声明之后必须使用 
	var A string       
	var s1="hello"  //类型推导
	var s2 int32=1
	var s3=int8(2)     //强制转换
	s4:=12     //省略var关键字  只能在函数内使用 
	s5:=int8(2)

	//批量声明
	var(
		A string
		B int
		C bool
	)  
//变量名格式    
	推荐使用小驼峰
	var studentName string //小驼峰 
	var StudentName string //大驼峰 
	var student_name string//下划线 
//定义常量  
	定义之后不能修改 不会改变 
	const pi=3.1415926
	const (
		s1=100
		s2    //默认和上一行值一样 
		s3
	) 
    iota : 计数器 //在const出现时被置为零 每新增一行常量声明就计数一次 加一 
	const(
		b1=iota  //0
		b2       //1
		_        //2
		b3       //3
	) 
	//定义数量级
	const(
		_ =iota        //iota=0
		KB=1<<(10*iota)//iota=1 二进制1左移10*1位 2^10 10000000000 1024 B 
		MB=1<<(10*iota)//iota=2 二进制1左移10*2位 2^20   1024*1024 
		GB=1<<(10*iota)
		TB=1<<(10*iota)
	) 
## 输入与输出 
	//输出语句 
	fmt.Print(A) 直接输出要打印的内容 
	fmt.Printf("name:%s\n",A)  %s:占位符 字符串类型  \n:换行符 
    fmt.Println(a) 末尾输出一个换行符 
	//输入语句 
	fmt.Scan(&a,&b)  //键盘输入时以空格作为分隔符  
	fmt.Scanf("%d",&a) 
	fmt.Scanln(&a)   
## 基本数据类型
//占位符:
	fmt.Printf("%T\n",A)	 
	%T 变量的类型 
	%v 默认类型值（万能）%#v 标注变量的类型
	%b 二进制  %d十进制  %o八进制  %x十六进制 
	%f 浮点数  %9.2f  
	%c 字符 （数对应的unicode码值） 
	%s 字符串
	%q 单/双引号括起来的字符/字符串字面值，必要时采用安全的转义表示
	%p 内存地址（指针）  & ：取(变量)地址 
	%t 布尔值 
	%U Unicode格式:U+1234 等价于 U+%04X 
	%% 百分号
		 
//整形      int或uint在32位操作系统上就是int32或uint32 在64位操作系统上就是int64或uint64
	int //有符号整型 
		int8    //八位（ -128到127） 
		int16
		int32 //rune 
		int64
	uint//无符号整型 
		uint8   //八位（0到255）
		uint16
	uintptr     //无符号整型 存放指针 
八进制和十六进制  
	A:=10101    //十进制 
	A:=0777     //八进制 
	A:=0x12345  //十六进制  
	fmt.Printf("%d%b%o%x",A) //%d 十进制 %b 二进制 %o 八进制 %x 十六进制   
//浮点型小数 
	float32	float64    //两个完全不同的类型值  不能直接相互赋值 
	math.MaxFloat32	  math.MaxFloat64  //最大值
	f1:=1.23456 // 默认为float64 
	f2:=float32(1.23456) 
	fmt.Printf("%T\n", f2) //32位 
//复数
	var a1 complex64       var a2 complex128
	a1=1+2i
	fmt.Println(a1)
//布尔值            1            0
	bool  只有true（真）和false（假）两个值 
	默认值为false     
	不允许将整形（0 1）强制转换为布尔型 
	无法参与数值运算，也无法与其他类型进行转换 
//字符串   
	字符串:%s "双引号" 类型：string 
	字符： %c '单引号' 类型：byte（int8）代表ASCII码字符  以及 rune（int32）代表UTF-8字符   
	多行字符串： ` `  原样输出  
	A:=`田某某 
	大帅B `
    //字符串相关操作：
	B:="理想"
	C:="大帅B" 
	len(B)  //字符串长度
	D:=B+C  //字符串拼接
	D:=fmt.Sprintf(B,C) //字符串拼接，不打印，同上 
	fmt.Printf("%s%s",B,C) //直接打印出B和C，理想大帅B  
    //修改字符串：
	s2:=[]rune(s1)  //切割为[]rune或[]byte类型的切片  
	fmt.Println(string(s2))  //再转换成字符串类型输出
	//字符串切割
	s:="how do you do"
	s3:=strings.Split(s," ") //按空格切割 得到s3切片类型 
//自定义类型与类型别名 (用type)
	type a int	//自定义类型 编译后类型:main.a 
	type b=int  //类型别名 	 编译后还是int类型  如byte是int8的类型别名 而rune是int32 
//匿名变量 
    _   多用来忽略某个值  不占用命名空间 不分配内存 	
	 
//字符串转义符 
\	\r 回车符 \n 换行符 \t 制表符（Tab键） \" " 双引号 \\ 反斜杠 

## 复合数据类型
//数组   （array） 
	必须指定元素的类型和容量 
	//初始化
	var A [3]bool //默认为零值 [false false false] 		
	A=[3]bool{true,true,true}
	B:=[5]int{0:1,4:2}//根据索引初始化[1 0 0 0 2]
	C:=[...]int{1,2,3,4,5}//自动推断长度 
	注意：多维数组只有第一层可以使用[...] 
	//遍历   for range
	citys:=[3]string{"北京","上海","深圳"} 	
	for _,value:= range citys{//匿名变量也可以用index(索引)接收 
		fmt.Println(value) 
	} 
	citys:=[2][3]string{  //多维数组 
		{"北京","上海","保定"},
		{"深圳","广州","佛山"}, 
	} 
	for _,v1:= range citys{ //多维数组的遍历
		fmt.Println(v1)
		for _,v2:=range v1{
			fmt.Println(v2)
		} 
	}
//切片  (slice)
	首指针、长度、容量	
	"引用类型",切片不存值 
	对切片内值的任何改变都是改变"底层数组"内的值 
	//定义
	var s1 []int  //nil 此时还未分配内存 
		fmt.Println(s1==nil)//true  不能直接比较，只能和nil比较
	//初始化
	s2=[]string{"北京","保定"}
		fmt.Println(s1==nil)//false 初始化之后就不为nil了 （首指针就不是nil了） 
	//由数组构造切片 
	s1:=s[:] 
	//切割字符串
	s1:=[]rune(s) //将字符串s切割成切片s1（s1中是字符） 
	s1:=[]byte(s) 
	//make()    同时分配内存 
	a:=make([]int,2,10)   长度：2 容量：10 
	//长度和容量
	a:=[5]int{1,2,3,4,5}
	s3:=a[i:j] //从i到j 不包含j (含左不含右) 
	s3:=a[1:4] //{2,3,4}  长度：3 容量：4 
	用len( )求长度  cap( )求容量
	fmt.Printf("len(s3):%d cap(s3):%d",len(s3),cap(s3))
	容量：按底层数组算 从i开始有几个数容量就是几 
	//再分割
	s4:=s3[0:2] //{2,3} 长度：2 容量：4  （含左不含右） 
	//判断空切片 （长度为0就是空切片） 
	用len(s) == 0	 不要用 s==nil//长度为零就是空切片，但是空切片分配内存之后就不是nil了
	一个nil值的切片长度和容量都是0  但一个长度和容量都是0的切片不一定是nil	 
	//赋值拷贝
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1 //将s1直接赋值给s2，共用一个底层数组 
	//复制 copy( )	 
	s3:=make([]int,2,3) //提前定义好s3的长度和容量，区别于s2  
	copy(s3,s1) //不共用底层数组 只把数一一对应放到s3中 s3的容量和长度不变   
	//遍历
	方式和数组一致，支持索引遍历和for range遍历。 
	//追加元素 append（）   自动初始化切片 
	s:=[]int{1,2}   
	s=append(s,3,4) //{1,2,3,4}
	容量不够时按照一定策略进行“扩容”//先判断新申请容量如果大于两倍旧容量直接给新申请容量  否则旧切片长度为1024以下就乘以2  大于等于1024就循环乘以1.25直到容量大于新申请容量   
	  s2:=[]int{3,4}
	s=append(s,s2...) //{1,2,3,4,3,4}  ... ：拆开s2
	//删除元素   (先分割再追加）用append  
	s:=[]int{1,2,3,4,5}
	s=append(s[:2],s[3:]...)//{1,2}+4,5  即{1,2,4,5} （分割含左不含右） 
	//排序
	a:=[...]int{3,7,8,9,1}
	sort.Ints(a[:]) // a[:] 由数组构造切片 
//指针	  只能读不能改	    
	a:=&b //a取出b的地址，此时a为指针 类型是*int
	c:=*a //根据地址找值，c的类型是a指向的int类型 
	数据类型：*int  *string 等
	//地址 
	fmt.Printf("%p",&b)  //b的地址 
	fmt.Printf("%p", a)  //b的地址（指针a的值） 
	fmt.Printf("%p",&a)  //a的地址
	//new 
	a:=new(int) //返回一个指向该类型的指针   分配一个新地址 
	b:=new(bool)  
	//make
	用来给slice、map以及channel申请内存 //返回这三个类型本身 
//map  无序的映射  键值对  必须设定长度，分配内存 
	map[KeyType]ValueType //KeyType:键的类型 ValueType:对应值的类型
	map[int]string        //此时还未分配内存 
	//初始化 
	(1)a:=make(map[int]string,5) //make 构造map并分配内存
	   a[100]="田子恒"
	(2)b:=map[string]string{    //声明同时初始化 
		"username": "沙河小王子",
		"password": "123456",
	}
	//判断是否存在    (一般用ok) 
	v := Map[key]    // v: 对应值或零
	v, ok := Map[key]//ok：布尔值(true  false)
	_, ok := Map[key]//单纯判断 
	if !ok {                   
		fmt.Println("查无此人")
	} else {                
		fmt.Println(v) 
	}
	//遍历  for range
	for key,value := range scoreMap {
		fmt.Println(key,value)
	}
	for key := range scoreMap {     //只遍历key 
		fmt.Println(key)
	}
	for _,value := range scoreMap { //只遍历value
		fmt.Println(value)
	}	
	//按一定顺序遍历 
	var keys = make([]string, 0, 200) //定义一个切片 
	for key := range scoreMap {		//取出map中的所有key存入切片keys
		keys = append(keys, key)
	}
	sort.Strings(keys)	//对切片进行排序
	for _ , key := range keys {		//按照排序后的key遍历map
		fmt.Println(key, scoreMap[key])
	}
	//删除   delete  
	delete(scoreMap,"田子恒")//删除元素,若nil或无此元素,则无操作
	aaa:="北京"
	delete(scoreMap,aaa) 
	// 元素为map类型的切片
	slice := make ([]map[int]string,2,3) 
	slice[0]=make(map[int]string,9) //未分配内存的map需要先用make分配内存 
	slice[0][100]="QQQ" // slice:[map[100:QQQ] map[]]
	//值为切片类型的map
	m1:=make(map[string][]int,10)
	m1["北京"]=[]int{1,2,3} // m1: map[北京：[1,2,3] ]	
//结构体  （struck）	
	//定义  
	type qqq struct{//定义名为qqq的结构体类型 
		x int    
		y int 		//占用一块连续的内存空间
	}
	//初始化 
	(1)var p qqq   //声明变量 p
		p.x=1	   //访问p中的某一个字段 
		p.y=2	
	(2)p:=qqq{     //键值对初始化 key-value 
		x:1,    
		y:2,       // :  ， 
	} 
	(3)p:=qqq{ 	   //值列表初始化   一一对应  
		18,		
		24,
	} 
	//构造函数     用new开头命名  
	func newQqq (x,y int) (*qqq) { //指针省内存 也可以不用指针直接qqq 
		return &qqq{		
			x:x,
			y:y,
		} 		//返回qqq类型的结构体变量(或指针) 
	}	
	p:=newQqq(1,2)  //调用newQqq 
	//匿名结构体 
	var s struct{	//多用于临时场景
		a int
		b string
	}{10,"战斗机"}  //同时赋值 
	//指针结构体
	若结构体内容较多 用指针可以极大节省运行时的内存占用
	(1)var p qqq 
	    a:=&p  	 //类型：*main.qqq
	  (*a).x=1   //根据内存地址找到原变量  修改指向的结构体中的值
        a.x=1    //简写 Go语法糖，指针a自动找原变量 不用加*  
	(2)a:=new(qqq) //定义指针a  指向qqq类型 并分配地址 
	(3)a:=&qqq{  //定义并赋值
		x:2, 
		y:4,		//赋值必须加逗号 
	}
	(4)a:=&qqq{   
		1,			
		2,
	}
	//结构体嵌套
	type aaa struct{
		a int
		b qqq	//嵌套别的结构体 
		qqq     //匿名嵌套 
	} 	
	 
//方法（和接收者）
	给类型添加方法 
	//值类型(传拷贝值，副本)
	func (q qqq)wang(){ //只作用于qqq类型的函数 
		fmt,Printf("%v：汪汪汪~",q.x)
	} 	 //接收者 使用类型首字母的小写命名 
	q1:=newQqq(1,2)
	//调用方法
	q1.wang()   
	//指针类型(传地址)    	  	 //以下情况用指针 
	func (q *qqq)jiayi(){ 		 //(1)需要修改值
		q.x++       			 //(2)内存太大 
	}  					 		 //(3)保证一致性(其他方法使用了指针)
	q1.jiayi()//(&q1).jiayi()的简写形式 
	//给自定义类型添加方法  
	type myInt int //先定义自己包里的类型 
	func (m myInt)hello(){ //变相实现
		fmt.Println("我是一个int")
	} 

//接口   interface 
	接口是一个(抽象的)类型，只约定拥有哪些方法 
	实现所有规定方法的结构体就是此接口类型 
	//定义
	type aaa interface{ //定义接口类型aaa
		b()  
		c() //满足所有方法的变量属于aaa类型
	} 
	var a1 aaa//包含两部分（1）动态类型（2）动态值 
	//指针类型与值类型
	使用值接收者实现接口，结构体类型和结构体指针类型的变量都可以存 
	用指针接收者实现接口，只有结构体指针类型的变量可以存 
	type 
	
	//一个结构体可以实现多个接口
	type mover interface{ move() } 
	type eater interface{ eat(string) } 
	func (c *cat) move(){fmt.Println("走猫步")}
	func (c *cat) eat(food string){fmt.Printf("吃%s",food)}
	//接口嵌套
	type animal interface{
		mover  //接口 
		eater
	}
	//作为函数参数 
	func da(x,aaa){ //参数为aaa类型
		x.b() 
	}
	//空接口作为函数参数
	func f1(a,interface){
		fmt.Printf("type:%T,value:%v\n",a,a)
	} 	 
	//空接口类型  
	interface{}  //任何类型都是空接口类型  直接使用
	m1:=make(map[int]interface{},10)
	map[1]=100086
	map[2]="qqq"
	map[3]=true
	map[4]=[...]string{"唱","跳","rap"}
	//类型断言   判断空接口接收的是什么类型的值 
	str,ok:=a.(string)   //1 str为对应值或零值 ok为布尔值bool 
	switch t:=a.(type){  //2 判断类型 且能输出值 
	case string:
		fmt.Println("是一个字符串：",t)
	case int:
		fmt.Println("是一个int：",t)			
	} 
	
	 
##函数
	函数是一段代码的封装，把一段代码抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了，能够让代码更清晰 更简洁。	
	//定义    ( 参数 )   (返回值) 
	func sum(x int,y int)(ret int){ //参数和返回值可以命名也可以不命名  同时在此函数内部定义了x,y和ret 
		return x+y //return后若省略  返回定义的ret 
	} 
	//没有返回值
	func f1(x int,y int){
		fmt.Println(x+y)
	} 
	//没有参数
	func f2( ) int{ //返回值的类型可以不用加括号 
		return 3
	} 
	//没有参数也没有返回值
	func f3(){
		fmt.Println(f3)
	}  
	//多个返回值 
	func f4()(int,string){
		return 1，"北京" //多个返回值的时候 必须用多个变量接收 
	}   
	//参数类型的简写 
	func f5(x,y,z int，a，b string)int{
		return x+y//当参数中多个连续的类型一致的时候，可以将非最后一个参数的类型省略 
	} 
	//可变长参数
	func f6(x string,y ...int){ //可变长参数必须放在参数的最后 
		fmt.Println(y)   //y是切片[]int  可传任意长度的值 
	} 
//函数作为参数和返回值
	函数也是变量的类型 func(int)int 等 
	func f1( x func(int)int ) ( y func(int)int ) {
		printf("%T",x) //func(int)int
		return y
	} 
//匿名函数
	函数内部不能再命名有名字的函数  匿名函数没有函数名 
	匿名函数在定义时需要保存到某个变量中
	f1:=func (x int,y int){
		fmt.Println(x+y)
	} 
	f1(4,5) 		//调用 
	//立即执行函数 
	如果只调用一次  定义完后可以加（）直接执行 
	func(x,y int){
		fmt.Println(x+y)
	}(4,5) 		//直接执行
//闭包
	闭包=函数+引用环境（外部变量） 
	//f2 利用 f3 输入到 f1 中运行 
	func f1(f func()){f()}//f1需要func()类型的参数传进去 
	func f2(x,y int) {fmt.Println(x+y)}//f2想进f1 
	func f3(f2.1 func(int,int),x,y int) func(){ 
		return func(){  f2.1(x,y)	}
		}//构造函数，将f2传进去生成一个伪装的函数（f4） 
	func main(){
		f4:=f3(f2,100,200) //此时f4就是一个闭包 包含函数本身以及其外部（f3）中的变量 f2.1()、x、y 
		f1(f4) 	
	}	
//defer
	延迟处理其后跟随的语句  一般用于释放资源
	若有多个defer语句 先进后出 压栈操作
			fmt.Print(1)	
	defer 	fmt.Print(2)
			fmt.Print(3) // 1 3 2
	//defer与函数返回值
	return 分两步 1 给返回值赋值 2 RET （不是原子操作） 
	defer则在return的两步之间执行 可以覆盖返回值变量 
		//返回值未声明变量
		defer不会改变返回值
		//返回值声明了变量名 
		在defer中改变此变量值就会改变输出值 
	同时要注意作用域（此x非彼x）可能改变的是副本（如defer后又有函数且又定义了个x接收原来x的值） 不会改变输出值 
	//defer后的语句有参数
	先确定参数再将语句压栈   返回时直接用确定的参数运行语句 
	fmt.Print(a,b)         //压栈时压入确定的值  fmt.Print(1，2) 
	//函数语句前的defer  
	defer f1(a,b,f2(a,b))  //压栈时压入 defer f1（2，2，4） 
	若此函数输入的参数中还有函数   确定参数时把这个函数（参数）运行完再压入	
//内置函数
	close	主要用来关闭channel
	len		用来求长度，比如string、array、slice、map、channel
	new		用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
	make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
	append	用来追加元素到数组、slice中
	panic和recover	用来做错误处理
//panic/recover      （一般不写） 
	panic()  程序出现错误导致崩溃
	recover() 恢复程序继续执行 必须在defer语句中才能生效（基本不用）	
	defer 一定要在可能引发panic的语句之前定义 
//递归   ：在函数中重复调用本函数 
	func QQQQQ(n int)int{  //数几种走法  一步1或2个台阶 
		if n==1{return 1}
		if n==2{return 2}
		return QQQQQ(n-1)+QQQQQ(n-2) 
	} 

	
 






	 
}

	//go mod init +文件名     创建go.mod文件

编译同目录的多个源码文件时，可以在 go build 的后面提供多个文件名，会输出可执行文件
在代码代码所在目录中使用 go build，在 go build 后添加要编译的源码文件名，

	跨平台编译 
 SET CGO_ENABLED=0  禁用CGO 
 SET GOOS=darwin    目标平台 如 Linux  Windows  darwin(mac) 
 SET GOARCH=amd64   目标处理器架构 

//编译运行 
 go build         遍历文件夹内的所有.go文件 编译生成.exe文件
 go run  main.go  执行main.go文件 
 go install       先编译后拷贝 
 
 
 
 
	
