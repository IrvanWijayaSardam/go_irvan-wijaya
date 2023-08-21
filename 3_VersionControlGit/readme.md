1. Version Control and Branch Management (Git)

Versioning yaitu mengatur versi source code.

- Git yaitu salah satu version control system populer yang digunakan para developer untuk mengembangkan software secara bersama-sama.
    -- Git dibuat oleh Linus Torvalds pada tahun (2005)

-Command pada git :
    - git status : digunakan untuk mengecek status perbedaan kode
    - git add <directory> : menambahkan direktory untuk dipush ke repository.
        -- ex : git add hello.py
        -- ex : git add .
    
    - git commit : yaitu digunakan untuk mengcommit 
        --ex : git commit -m "add config file"

    - git init : menginisialisasi git 
    - git remote add <remote_name> <remote_repo_url> : digunakan untuk menambahkan remote ke sebuah repository
    - git push - u <remote_name> <local_branch_name> : mengupdate kode ke repository.
    - git clone <link_repository> : untuk mendownload repository project github
