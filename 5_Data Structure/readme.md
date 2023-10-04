1. Array:
    -Array adalah tipe data yang digunakan untuk menyimpan elemen-elemen dengan tipe data yang sama.
    -Ukuran array tidak dapat berubah setelah deklarasinya.
    -Indeks array dimulai dari 0.
    Contoh implementasi:
    ```
    package main

    import "fmt"

    func main() {
        // Deklarasi array dengan ukuran 5
        var myArray [5]int

        // Inisialisasi array
        myArray[0] = 1
        myArray[1] = 2
        myArray[2] = 3
        myArray[3] = 4
        myArray[4] = 5

        // Mengakses elemen-elemen array
        fmt.Println(myArray[0]) // Output: 1
        fmt.Println(myArray[3]) // Output: 4
    }
    ```
2.Slice:
    -Slice adalah struktur data yang fleksibel dan dinamis untuk menyimpan koleksi elemen dengan tipe data yang sama.
    -Ukuran slice dapat berubah sesuai kebutuhan.
    -Slice dibuat dengan menggunakan operator make atau dengan mengambil sebagian dari array atau slice yang ada.
    Contoh implementasi:
    ```
    package main

    import "fmt"

    func main() {
        // Membuat slice kosong
        mySlice := []int{}

        // Menambahkan elemen ke slice
        mySlice = append(mySlice, 1)
        mySlice = append(mySlice, 2, 3, 4)

        // Mengakses elemen-elemen slice
        fmt.Println(mySlice) // Output: [1 2 3 4]
    }

    ```

3. Map:
    -Map adalah tipe data yang digunakan untuk menyimpan pasangan nilai kunci (key-value) yang tidak memiliki urutan tertentu.
    -Setiap kunci harus unik dalam satu map.
    -Map dapat digunakan untuk mengaitkan nilai dengan kunci tertentu.
    Contoh implementasi:
    ```
        package main

        import "fmt"

        func main() {
            // Deklarasi dan inisialisasi map
            myMap := make(map[string]int)

            // Menambahkan pasangan nilai kunci ke map
            myMap["satu"] = 1
            myMap["dua"] = 2
            myMap["tiga"] = 3

            // Mengakses nilai dengan kunci
            fmt.Println(myMap["satu"]) // Output: 1
            fmt.Println(myMap["tiga"]) // Output: 3
        }

    ```