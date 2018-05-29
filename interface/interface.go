/*
指针类型可以拥有接收者为指针或者值类型的方法
值类型只能拥有接受者为值类型的方法
*/


package main

import (
    "log"
)

type User struct {
    Name  string
    Email string
}

func (u User) Notify() error {
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
    user := &User{
        Name:  "wang",
        Email: "wang@163.com",
    }
    user.Notify()
    SendNotification(user)
}

