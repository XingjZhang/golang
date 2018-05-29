/*
1  当匿名字段为值类型的时候
   S{T}     S和*S都拥有接受者为T方法，只有*S有接受者为*T的方法

2 匿名字段为指针类型 
S{*T}   S和*S都拥有接受者为T何*T的方法

拥有相关的方法，就可以赋值给相应的接口。



*/


package main

import (
    "log"
)

type User struct {
    Name  string
    Email string
}

type Admin struct {
    User
    Level string
}

func (u *User) Notify() error {
    log.Printf("User: Sending User Email To %s<%s>\n",
        u.Name,
        u.Email)

    return nil
}

type Notifier interface {
    Notify() error
}

func SendNotification(notify Notifier) error {
    return notify.Notify()
}

func main() {
   
     user := User{
        Name:  "wang",
        Email: "wang@le.com",
     }
     admin := &Admin{user,"string"}
   
    admin.Notify()
    admin.User.Notify()
    SendNotification(admin)
}

