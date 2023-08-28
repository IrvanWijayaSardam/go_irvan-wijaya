5. String Advance Function Pointer

Array : 
- Array adalah sebuah struktur data yang mengandung group of element. tipe data pada array harus sudah ditentukan seperti Numeric,String,Boolean. 
    var <variable_name> [<size_of_array>] <tipe_variable>
    --Example : var primes [5] int

- Slice : Slice is datastructre that contains a group of elements, can contain one type of variable(Like Array) but have dynamic allocation size. 
    var orunes = [5]int{2,3,5,7,11}
    var part_primes []int = primes[1:4]


- Map : Map is a datastructure stores data in the form of key and value paris where every key is unique.
    example :
        var salary = map[string]int{}
        var salary_a = map[string]int{"umam":1000,"iswanul":2000}
        
        salary_b := map[string]int{}
        fmt.Println(salary_b)

        var salary_c = make(map[sring]int)
        salary_c["doe"] = 7000
        fmt.Println(salary_c)