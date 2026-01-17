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