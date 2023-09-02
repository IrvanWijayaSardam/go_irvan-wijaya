9. Concurrent Programming

Concurrent Programming : Concurrent programming adalah paradigma pemrograman di mana Anda dapat menjalankan beberapa tugas (proses atau goroutine dalam konteks Go) secara bersamaan. Tujuan dari concurrent programming adalah untuk meningkatkan efisiensi dan kemampuan program untuk mengatasi tugas-tugas yang berjalan secara paralel, terutama dalam lingkungan multi-core atau multi-processor.

Dalam Go (bahasa pemrograman Go atau Golang), concurrent programming didukung dengan sangat baik. Go memiliki fitur bawaan untuk mendukung konkurensi dengan goroutine, yang merupakan unit ringan untuk menjalankan tugas. Goroutine sangat efisien dalam penggunaan sumber daya dan memungkinkan Anda untuk membuat dan menjalankan ribuan goroutine tanpa terlalu banyak overhead.


Sequential : Sequential setiap proses dijalankan satu persatu.

Parallel : Paralel adalah kemampuan untuk menjalankan pada beberapa jalur atau menjalankan secara stimultant pada beberapa proses sekaligus

- Goroutines 

Goroutine adalah salah satu fitur utama dalam bahasa pemrograman Go (Golang) yang digunakan untuk menjalankan tugas-tugas secara konkuren. Goroutine adalah unit ringan (lightweight) yang dijadwalkan oleh runtime Go, dan banyak goroutine dapat berjalan secara bersamaan di dalam satu program Go tanpa banyak overhead. Mereka memungkinkan Anda untuk menjalankan tugas-tugas secara konkuren tanpa perlu mengelola threading secara langsung.
