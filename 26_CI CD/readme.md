CI/CD :

Continous Integration
Continous delivery
Continous deployment
how to CI & CD Improve your Productivition


Github Action dapat digunakan untuk CI/CD dengan dikolaborasikan dengan docker image , berikut contoh docker-image.yml :

name: Docker Image CI

on:
  push:
    branches: "main"

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      -
        name: Checkout
        uses: actions/checkout@v3
      -
        name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}
      -
        name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
      -
        name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKERHUB_USERNAME }}/imgmidleware:latest #nama image
      -
        name: connect ssh
        uses: appleboy/ssh-action@v0.1.9
        with:
          host: ${{ secrets.HOST }} //IP EXTERNAL
          username: ${{ secrets.USERNAME }} //USERNAME SSH
          key: ${{ secrets.KEY }} // GCP -> ISI SSH key (bukan pub)
          port: ${{ secrets.PORT }} //PORT 22
          script: |
            docker stop ${{ secrets.CNAME }} // ContainerName
            docker rm ${{ secrets.CNAME }} // ContainerName
            docker rmi ${{ secrets.DOCKERHUB_USERNAME }}/imgkmb
            docker run --name ${{ secrets.CNAME }} -p 8000:8000 -d -e DBHOST=${{secrets.DBHOST}} -e DBPORT=3306 -e DBUSER=${{secrets.DBUSER}} -e DBNAME=${{secrets.DBNAME}} -e DBPASS=${{secrets.DBPASS}} -e SECRET=${{secrets.JWTSECRET}} -e REFSECRET=${{secrets.REFSECRET}} -e SERVER=8000 ${{ secrets.DOCKERHUB_USERNAME }}/imgkmb