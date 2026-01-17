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

## t.Fail() dan t.FailNow()
- Terdapat dua function untuk menggagalkan unit test, yaitu ```Fail()``` dan ```FailNow()```. Lantas apa bedanya ?.
- ```Fail()``` akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal.
- ```FailNow()``` akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test.

## t.Error(args...) dan t.Fatal(args...)
- Selain ```Fail()``` dan ```FailNow()```, ada juga ```Error()``` dan ```Fatal()```.
- ```Error()``` function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function ```Fail()```, sehingga mengakibatkan unit test dianggap gagal.
- Namun karena hanya memanggil ```Fail()```, artinya eksekusi unit test akan tetap berjalan sampai selesai.
- ```Fatal()``` mirip dengan ```Error()```, hanya saja setelah melakukan log error, dia akan memanggil ```FailNow()```, sehingga mengakibatkan eksekusi unit test berhenti.

## Kode Program Error
```go
func TestHelloWorldError(t * testing.T) {
  result := HelloWorld("Sandy")
  if result != "Hello Sandy" {
    t.Error("should Hello Sandy")
  }

  fmt.Println("Dieksekusi")
}
```

## Kode Program Fatal
```go
func TestHelloWorldFatal(t * testing.T) {
  result := HelloWorld("Sandy")
  if result != "Hello Sandy" {
    t.Fatal("should Hello Sandy")
  }

  fmt.Println("Tidak Dieksekusi")
}
```

# Assertion
- Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan.
- Apalagi jika result data yang harus di cek itu banyak.
- Oleh karena itu, sangat disarankan untuk menggunakan assertion untuk melakukan pengecekan.
- Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini.

## Testify
- Salah satu library assertion yang paling populer di Go-Lang adalah Testify.
- Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test.
- [https://github.com/stretchr/testify/](https://github.com/stretchr/testify/).
- Kita bisa menambahkannya di go module kita: ```go get github.com/stretchr/testify```.

## Kode Program Assertion
```go
import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestHelloWorldAssertion(t *testing.T) {
  result := HelloWorld("Sandy")
  assert.Equal(t, "Hello Sandy", result)
  fmt.Println("Dieksekusi")
}
```

## assert vs require
- Testify menyediakan dua package untuk assertion, yaitu assert dan require.
- Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil fungsi ```Fail()```, artinya eksekusi unit test akan tetap dilanjutkan jika terjadi gagal.
- Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil ```FailNow()```, artinya eksekusi unit test tidak akan dilanjutkan.

## Kode Program Require
```go
import (
  "fmt"
  "github.com/stretchr/testify/require"
  "testing"
)

func TestHelloWorldRequire(t *testing.T) {
  result := HelloWorld("Sandy")
  require.Equal(t, "Hello Sandy", result)
  fmt.Println("Dieksekusi")
}
```

# Skip Test
- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test.
- Di Go-Lang kita juga bisa membatalkan eksekusi unit test.
- Untuk membatalkan unit test kita bisa menggunakan function ```Skip()```.

## Kode Program Skip Test
```go
import (
  "fmt"
  "github.com/stretchr/testify/require"
  "testing"
)

func TestHelloWorldSkip(t *testing.T) {
  if runtime.GOOS == "darwin" {
    t.Skip("Unit test tidak bisa jalan di Mac")
  }

  result := HelloWorld("Sandy")
  require.Equal(t, "Hello Sandy", result)
}
```

# Before dan After Test
- Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi.
- Jika kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function-nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya.
- Untungnya di Go-Lang terdapat fitur yang bernama ```testing.M```.
- Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test.

## testing.M
- Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter ```testing.M```.
- Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package.
- Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau.
- Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test.

## Kode Program testing.M
```go
func TestMain(m *testing.M) {
  fmt.Println("Sebelum unit test")

  m.Run() // eksekusi semua unit test

  fmt.Println("Setelah unit test")
}
```

# Sub Test
- Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test.
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test bahasa pemrograman lain.
- Untuk membuat sub test, kita bisa menggunakan function ```Run()```.

## Kode Program Membuat Sub Test
```go
func TestSubTest(t *testing.T) {
  t.Run("Sandy", func(t *testing.T) {
    result := HelloWorld("Sandy")
    require.Equal(t, "Hello Sandy", result)
  })

  t.Run("Dwi", func(t *testing.T) {
    result := HelloWorld("Dwi")
    require.Equal(t, "Hello Dwi", result)
  })
}
```

## Menjalankan Hanya Sub Test
- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah: ```go test -run TestNamaFunction```.
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah: ```go test -run TestNamaFunction/NamaSubTest```.
- Atau untuk semua sub test di semua function, kita bisa gunakan perintah: ```go test -run /NamaSubTest```.