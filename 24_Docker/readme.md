Docker adalah platform open-source yang digunakan untuk mengembangkan, mengirim, dan menjalankan aplikasi dalam lingkungan kontainer. Kontainer Docker mengisolasi aplikasi dan dependensinya, sehingga memungkinkan aplikasi berjalan dengan konsisten di berbagai lingkungan. Berikut adalah beberapa contoh perintah Docker yang umum digunakan:

1. **docker run**: Menjalankan kontainer dari sebuah image.
   Contoh:
   ```
   docker run -d -p 80:80 nginx
   ```

2. **docker build**: Membuat image Docker dari Dockerfile.
   Contoh:
   ```
   docker build -t nama_image:tag .
   ```

3. **docker pull**: Mengunduh image Docker dari Docker Hub atau repositori lainnya.
   Contoh:
   ```
   docker pull ubuntu:20.04
   ```

4. **docker ps**: Menampilkan daftar kontainer yang sedang berjalan.
   Contoh:
   ```
   docker ps
   ```

5. **docker stop**: Menghentikan kontainer yang sedang berjalan.
   Contoh:
   ```
   docker stop nama_kontainer
   ```

6. **docker exec**: Menjalankan perintah di dalam kontainer yang sedang berjalan.
   Contoh:
   ```
   docker exec -it nama_kontainer /bin/bash
   ```

7. **docker-compose**: Mengelola aplikasi yang terdiri dari beberapa kontainer.
   Contoh:
   ```
   docker-compose up -d
   ```

8. **docker network**: Mengelola jaringan kontainer Docker.
   Contoh:
   ```
   docker network create nama_jaringan
   ```

9. **docker volume**: Mengelola volume Docker untuk menyimpan data persisten.
   Contoh:
   ```
   docker volume create nama_volume
   ```

Dengan menggunakan perintah-perintah ini, Anda dapat mengelola kontainer Docker, membuat image kustom, mengunduh image dari repositori, dan banyak lagi, untuk mempermudah pengembangan dan distribusi aplikasi Anda dalam kontainer.