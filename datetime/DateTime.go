//golang 的Time很难用, 由于IDE的提示很差, 很多时候需要阅读源代码才能知道调用的方法怎么用. 我不想老是记住"2006-01-02 15:04:05"这个
//字符串, 于是模拟了一下c#的DateTime, 提供了一些简便的方法. 例如初始化日期只需要year,month,day三个参数就可以了. golang的自带初始化
//可是要输入9个参数, 而且month居然不是int, 真实受不了
//DateTime, 简单使用

package datetime
import (
    "time"
    "fmt"
)

type DateTime struct {
    Time time.Time
}



//输入年月日得到日期
func NewDate(year,month,day int)DateTime{
    t:=time.Date(year,time.Month(month),day,0,0,0,0,time.UTC)
   // fmt.Println(t)
    dt:=DateTime{Time:t}

    return dt
}
//获取日期部分的字符串
func (x DateTime)GetDateString()string{
    return x.Time.Format("2006-01-02")
}
//获取时间部分的字符串
func (x DateTime)GetTimeString()string{
    return x.Time.Format("15:04:05")
}
//确定日期部分之后, 传入时间字符串来设置时间部分
func (x *DateTime)SetTime(timestring string){
  var str=x.GetDateString()+" "+timestring
    t,err:=time.Parse("2006-01-02 15:04:05",str)
    if err!=nil{
        fmt.Println(err)
    }
  //  fmt.Println("t:",t)
    x.Time=t
}
