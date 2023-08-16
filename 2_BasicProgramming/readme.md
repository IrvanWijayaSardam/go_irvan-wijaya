2. Basic Programming
Mentor : Muhammad Ismail
Source : Platform Alterra Academy

- Introduction Golang 
    -golang merupakaan bahasa pemrgorman yang dikembangkan oleh google yang dapat digunakan untuk menyediakan layanan sistem lain seperti API.
    - developed by :
        1. Robert Grisemer
        2. Rob Pike
        3. Ken Thompson
        4. Ian Lance Taylor
        5. Russ Cox
    - Kelebihan Golang
        - Mudah digunakan
        - Built-in concurrency

- Variable Tipe Data
    - Variabel pada golang : variable digunakan untuk menyimpan informasi pada komputer program, mereka memberikan cara untuk melabeli data dengan deskripsi. nama dan mereka terdapat data type.
    - Data type pada golang :
        - String
        - Numeric : 
            - Integer : uint8 - uint64, int,rune,byte
            - Float : float32, float64
            - Complex : complex64, complex138
        - Boolean : true,false
    
    Contoh deklarasi variable pada golang :
        1. var <variable_name> <type_data>
        ex : var age int

    Variable Contants :
        - Single Contants : conts pi float64 = 3.14
        - Multiple constants : conts (
            pi float64 = 3.14
            pi2
            Age = 10
        )

- Operator Exppression
    --Operand(ex)
        - x := 1 + 2
        fmt.Println(x)
    --Expression(ex)
        a := 5
        b := 5
        luas := a+b
        fmt.Println(luas)
    -- String Operation
        helloword := "hello" + " " + "world"
        fmt.Println(helloworld)
- Branching
    --If else (ex)
        - if hour < 12 {fmt.Println("Selamat Pagi")}
        else if hour < 18 {fmt.Println("Selamat Sore)}
        else {fmt.Println("Selamat Malam")}
    --Switch case(ex)
        - var today int = 2
        switch today {
            case 1 :
                fmt.Println("Today is Monday")
            case 2 :
                fmt.Println("Today is Tuesday")
            default:
        }
- Looping 
    - For (ex)
    for i := 0; i < 5; i++ {
        fmtPrintln(i)
    }
    - Continue or break
    for i := 0; i < 5; i++ {
        if i == 1 {
            continue
        }
        if i > 3 {
            break
        }
    }