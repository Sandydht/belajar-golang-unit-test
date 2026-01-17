# Pengenalan Software Testing
- Software testing adalah salah satu disiplin ilmu dalam software engineering.
- Tujuan utama dari software testing adalah memastikan kualitas kode dan aplikasi kita baik.
- Ilmu untuk software testing sendiri sangatlah luas.

## Test Pyramid
![Test Pyramid](/src/assets/img/png/test-pryramid.png)

## Contoh High Level Architecture Aplikasi
![Contoh High Level Architecture Aplikasi](/src/assets/img/png/high-level-architecture-aplikasi.png)

## UI Test / End to End Test
![End to End Test](/src/assets/img/png/end-to-end-test.png)

## Service Test / Integration Test
![Integration Test](/src/assets/img/png/integration-test.png)

## Contoh Internal Architecture Aplikasi
![Contoh Internal Architecture Aplikasi](/src/assets/img/png/integration-test.png)

## Unit Test
![Unit Test](/src/assets/img/png/unit-test.png)
- Akan selalu fokus menguji bagian kode program terkecil, biasanya menguji sebuah method.
- Unit test biasanya dibuat kecil dan cepat, oleh karena itu kode unit test terkadang lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unit test.
- Unit test bisa digunakan sebagai cara untuk meningkatkan kualitas kode program kita.

# Pengenalan Testing Package
## testing Package
- Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test.
- Berbeda dengan di Go-Lang, di Go-Lang untuk unit test sudah disediakan sebuah package khusus bernama testing.
- Selain itu untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintah-nya.
- Hal ini membuat implementasi unit testing di Go-Lang sangat mudah dibanding dengan bahasa pemrograman lain.
- [https://golang.org/pkg/testing/](https://golang.org/pkg/testing/).

## testing.T
- Go-Lang menyediakan sebuah ```struct``` yang bernama ```testing.T```
- ```struct``` ini digunakan untuk unit test di Go-Lang.

## testing.M
- ```testing.M``` adalah ```struct``` yang disediakan di Go-Lang untuk mengatur lifecycle testing.
- Di materi ini kita akan bahas di chapter main.

## testing.B
- ```testing.B``` adalah ```struct``` yang disediakan Go-Lang untuk melakukan benchmarking.
- Di Go-Lang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan.

# Membuat Unit Test
## Kode Program Hello World Function
```go
func HelloWorld(name string) string {
  return "Hello " + name
}
```

# Aturan File Test
- Go-Lang memiliki aturan cara membuat file khusus untuk unit test.
- Untuk membuat file unit test, kita harus menggunakan akhiran **_test**.
- Jadi misal kita membuat file hello_world.go, artinya untuk membuat unit test-nya, kita harus membuat file hello_world_test.go.

## Aturan Function Unit Test
- Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test-nya.
- Nama function untuk unit test harus diawali dengan nama **Test**.
- Misal jika kita ingin mengetes function **HelloWorld**, maka kita harus membuat function unit test dengan nama **TestHelloWorld**.
- Tidak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawali dengan kata **Test**.
- Selanjutnya harus memiliki parameter **(t \*testing.T)** dan tidak mengembalikan return value.

## Kode Program Hello World Unit Test
```go
import "testing"

func TestHelloWorld(t *testing.T) {
  result := HelloWorld("Sandy")
  if result != "Hello Sandy" {
    panic("Result is not Hello Sandy")
  }
}
```

# Menjalankan Unit Test
- Untuk menjalankan unit test kita bisa menggunakan perintah: ```go test```.
- Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah: ```go test -v```.
- Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah: ```go test -v -run TestNamaFunction```.
- Jika kita ingin menjalankan semua unit test dari top folder module-nya, kita bisa gunakan perintah: ```go test ./...```

# Menggagalkan Test
- Menggagalkan unit test menggunakan ```panic``` bukanlah hal yang bagus.
- Go-Lang sendiri sudah menyediakan cara untuk mengagalkan unit test menggunakan testing.T.
- Terdapat function ```Fail()```, ```FailNow()```, ```Error()```, dan ```Fatal()``` jika kita ingin menggagalkan unit test.