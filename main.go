package main

import "fmt"

type user struct {
    name    string
    request int
    money   int
}

type Coffee struct {
    name            string
    specification   string
    price           int
}

func (c *Coffee) createCoffee() {
       fmt.Println("your coffee is ready", c.name)
}

func coffeeBuy(userInfo *user, coffee *Coffee)  {
    if(userInfo.money >= coffee.price) {
        coffee.createCoffee()
    } else {
        fmt.Println("you don't have enough money")
    }
}

func main()  {
    var user1 user

    fmt.Print("Username: ")
    fmt.Scanln(&user1.name)


    coffeeList := []Coffee{
        {
            name: "Türk Kahvesi",
            specification: "Öğütülmüş Türk Kahvesi + kaynamış su",
            price: 40,
        },
        {
            name: "Americano",
            specification: "Espresso + su",
            price: 65,
        },
        {
            name: "Espresso",
            specification: "Espresso kahvesi + Espresso makinesi + kaynamış su",
            price: 55,
        },
    }

    for i, coffee := range coffeeList{
        if(i == len(coffeeList) -1){
            fmt.Printf("%d- coffe name: %s, specification: %s, price: %d \n", i,coffee.name, coffee.specification, coffee.price)
        } else {
            fmt.Printf("%d- coffe name: %s, specification: %s, price: %d \n",i, coffee.name, coffee.specification, coffee.price)
        }  
    }

    fmt.Print("Request: ")
    fmt.Scanln(&user1.request)

    fmt.Print("Money: ")
    fmt.Scanln(&user1.money)

    coffeeBuy(&user1, &coffeeList[user1.request])
}

