package main

import (
    "fmt"
    "os"
)

type shopping interface {
    Adduser([]customer) (bool, []customer)
    Placeorder([]customer, []Products) []Products
    Getorder([]customer)
    Getproduct([]Products)
    Modifyorder( []customer) []customer
    Cancelorder( []customer) []customer
}

//customer is
type customer struct {
    name   string
    mobile string
    place  string
    orders []Products
}

//Products is...
type Products struct {
    pname string
    qty   int
}

type deliveryrep struct {
    dname  string
    dnum   string
    dplace string
}

//gv
func Adddeliveryrep(d deliveryrep,drep []deliveryrep) (bool, []deliveryrep) {
    dr := deliveryrep{dname: d.dname, dnum: d.dnum, dplace: d.dplace}
    for i := 0; i < len(drep); i++ {
        if dr.dnum == drep[i].dnum {
            fmt.Println("Number exists")
            return false, drep
        }
    }
    drep = append(drep, dr)
    fmt.Println(drep)
    fmt.Println("delivery rep Succesfully registered")
    return true, drep
}

//Adduser is...
func (c customer) Adduser(users []customer) (bool, []customer) {
    u := customer{name: c.name, mobile: c.mobile, place: c.place}
	fmt.Println(u)
    for i := 0; i < len(users); i++ {
        if u.mobile == users[i].mobile {
            fmt.Println("Number exists")
            return false, users
        }
    }
    users = append(users, u)
    fmt.Println("User Succesfully registered")
    return true, users
}

//Getproduct is
func (c customer) Getproduct(productlist []Products) {
    fmt.Println("the store offers the following items in the store:")
    fmt.Println("Product", "->", "Quantity")
    for _, j := range productlist {
        fmt.Println(j.pname, "->", j.qty)
    }
}

//Placeorder is
func (c customer) Placeorder(users []customer, plist []Products) []Products {
    pt := ""
    pq := 0
    fmt.Println("enter the type you want to buy")
    fmt.Scanf("%s", &pt)
    fmt.Scanln()
    fmt.Println("enter quantity")
    fmt.Scanf("%d", &pq)
    fmt.Scanln()
    ord := Products{pt, pq}
    for i := 0; i < len(users); i++ {
        if c.mobile == users[i].mobile {
            users[i].orders = append(users[i].orders, ord)
            for j := 0; j < len(plist); j++ {
                if plist[j].pname == pt {
                    plist[j].qty -= pq
                }
            }
        }
    }
    fmt.Println(users)
    return plist
}

//Getorder is
func (c customer) Getorder(users []customer) {
    for i := 0; i < len(users); i++ {
        if c.mobile == users[i].mobile {
            fmt.Println("Customer ", users[i].name, " ordered the follwing items:")
            fmt.Println(users[i].orders)
        }
    }
}

//Modifyorder is
func (c customer) Modifyorder(users []customer) []customer {
    var j int
    for i := 0; i < len(users); i++ {
        if c.mobile == users[i].mobile {
            j = i
            break
        }
    }
    fmt.Println("Modify product type or change?")
    fmt.Println("enter ptype or pqty or both ")
    var t, opt string
    var q int
    fmt.Scanf("%s", &opt)
    fmt.Scanln()
    if opt == "ptype" {
        fmt.Println("enter the new product type")
        fmt.Scanf("%s", &t)
        fmt.Scanln()
        users[j].orders[0].pname = t

    } else if opt == "pqty" {
        fmt.Println("enter the new product Quantity")
        fmt.Scanf("%d", &q)
        fmt.Scanln()
        users[j].orders[0].qty = q
    } else {
        fmt.Println("enter the new product type")
        fmt.Scanf("%s", &t)
        fmt.Scanln()
        users[j].orders[0].pname = t
        fmt.Println("enter the new product Quantity")
        fmt.Scanf("%d", &q)
        fmt.Scanln()
        users[j].orders[0].qty = q
    }
    return users
}

//Cancelorder is
func (c customer) Cancelorder(users []customer) []customer {
    var j int
    for i := 0; i < len(users); i++ {
        if c.mobile == users[i].mobile {
            j = i
            break
        }
    }
    fmt.Println("Cancelling the order")
    users[j].orders[0] = users[j].orders[len(users[j].orders)-1] // Copy last element to index i.
    users[j].orders[len(users[j].orders)-1] = Products{}         // Erase last element (write zero value).
    users[j].orders = users[j].orders[:len(users[j].orders)-1]   // Truncate slice.
    return users
}
func main() {
    var i shopping
    var c customer
    i = c
    users := make([]customer, 0)
    u1 := customer{name: "vamsi", mobile: "9123456780", place: "a"}
    u2 := customer{name: "krishna", mobile: "1234567890", place: "b"}
    u3 := customer{name: "jaya", mobile: "2345678901", place: "c"}
    users = append(users, u1, u2, u3)
    productlist := []Products{}
    p1 := Products{"tv", 600}
    p2 := Products{"mobile", 800}
    p3 := Products{"laptop", 750}
    p4 := Products{"xbox", 340}
    productlist = append(productlist, p1, p2, p3, p4)
	var d deliveryrep
    dreps := []deliveryrep{}
    dr1 := deliveryrep{dname: "lewis", dnum: "2323214527", dplace: "ws"}
    dr2 := deliveryrep{dname: "john", dnum: "2356214527", dplace: "fcv"}
    dr3 := deliveryrep{dname: "carl", dnum: "2323276527", dplace: "ws"}
    dreps = append(dreps, dr1, dr2, dr3)
    var choice int
    for {
        fmt.Println("")
        fmt.Println("---------OPTIONS----------")
        fmt.Println("1.Register")
        fmt.Println("2.Get Product List")
        fmt.Println("3.Place Order")
        fmt.Println("4.Get Orders")
        fmt.Println("5.Modify Order")
        fmt.Println("6.Cancel Order")
        fmt.Println("7.Exit")
        fmt.Println("8.Add Delivery Rep")
        fmt.Println("9.View Delivery Rep")
        fmt.Println("")
        fmt.Println("Enter your choice")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
			fmt.Println("enter name")
    		fmt.Scanf("%s", &c.name)
    		fmt.Scanln()
    		fmt.Println("enter mobile")
    		fmt.Scanf("%s", &c.mobile)
    		fmt.Scanln()
    		fmt.Println("enter place")
    		fmt.Scanf("%s", &c.place)
			fmt.Scanln()
			_, b := c.Adduser(users)
            fmt.Println(b)
            users = b
        case 2:
            i.Getproduct(productlist)
        case 3:
            productlist := c.Placeorder( users, productlist)
            fmt.Println("Remaining Products:", productlist)
        case 4:
            c.Getorder( users)
        case 5:
            users = c.Modifyorder( users)
        case 6:
            users = c.Cancelorder( users)
            fmt.Println("the order has been cancelled", users)
        case 7:
            fmt.Println("Exiting...")
            os.Exit(1)
        case 8:
            fmt.Println("enter drep name")
            fmt.Scanf("%s", &d.dname)
            fmt.Scanln()
            fmt.Println("enter drep mobile")
            fmt.Scanf("%s", &d.dnum)
            fmt.Scanln()
            fmt.Println("enter drep place")
            fmt.Scanf("%s", &d.dplace)
			fmt.Scanln()
            a, b := Adddeliveryrep(d,dreps)
            fmt.Println(a)
            dreps = b
        case 9:
            fmt.Println(dreps)

        }
    }
}
