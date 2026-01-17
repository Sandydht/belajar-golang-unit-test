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