# CRUD PROJECT
Project ini Berisi App dengan beberapa endpoint yang meliputi :
1. API insert Product
2. API Get List Product
3. API View Product

apps ini juga mengimplementasikan docker untuk containerization dan redis untuk caching data sementara. kemudian handling error dalam project ini dibuat terpusat dari satu function

# ARSITEKTUR PROJECT 
project ini menggunakan arsitektur dengan struktur seperti berikut : attribute, config, dto, endpoint, model, router, service, util, main.go
projeck ini menngunakan layered architecture, yang mana setiap file dengan funngsi yang berbeda di bungkus dalam package yang berbeda, sehingga memudahkan pencarian function yang dibutuhkan.
dibuat seperti ini dikarenakan dapat membuat handling error secara terpusat kemudian, menghindari terjadinya cycle import.

- attribute : berisi function untuk initiate connection DB dan redis yang akan digunakan saat project berjalan 
- config : berisi config yang akan digunakan dalam apps 
- dto : berisi struct untuk payload saat melakukan request API
- model : berisi struct yang mewakiki model product ataupun error model
- router : berisi routing untuk endpoint 
- service : berisi function untuk melakukan pengambilan data ke Database 
- util : berisi function yang diperlukan di service2 lain 



# RUNNING PROJECT
1. Running Manual
   - Buat Database baru dengan nama 'erajaya'
   - Sesuaikan congfigurasi DB dan redis di file config.json
   - Kemudian tambahkan environtment di windows ataupun IDE yang digunakan dengan env sebagai berikut : projectconfig=./config
   - Jalankan aplikasi
   - Selamat mencoba
   - 
2. Running Dengan Docker
   - jalankan command berikut di terminal
     docker-compose up
   - sesuaikan environtment pada docker-compose.yml

# API COLLECTION
1. # Insert
   curl --location 'http://localhost:8080/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "indomie kuah",
    "price": 4000,
    "quantity": 50,
    "description" : "ini adalah product indomie"
}'

2. # Getlist
   curl --location 'http://localhost:8080/product?page=1&limit=10&orderby=created_at'

   Getlist memiliki beberapa parameter untuk melakukan sorting dan pagination :
   - page --> berisi int mewakili page berapa yang ingin diambil
   - limit --> berisi int mewakili limit jumlah data yang ditampilkan
   - orderby --> berisi beberapa field mewakili field pada database antara lain :
                 - name, price, quantity, created_at
     
   cara pemakaian param :
   1. sorting produk terbaru
      - http://localhost:8080/product?page=1&limit=10&orderby=created_at desc'
   2. sorting harga termurah
      - http://localhost:8080/product?page=1&limit=10&orderby=price'
   3. sorting harga termahal
      - http://localhost:8080/product?page=1&limit=10&orderby=price desc'
   4. sorting product name
      - http://localhost:8080/product?page=1&limit=10&orderby=name'
      - http://localhost:8080/product?page=1&limit=10&orderby=name desc'
      

4. # VIEW DETAIL
   curl --location 'http://localhost:8080/product/1'
   
   


