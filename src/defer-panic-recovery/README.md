# Defer, Panic, dan Recovery
Di bahasa pemrograman lain mungkin konsep ini mirip dengan _try catch_.


## Defer
- Defer function adalah function yang kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi.
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi.

```go
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Running Application")
}

func main() {
	runApplication()
}
```
Output:
```go
Running Application
Selesai memanggil function
```

## Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program (benar-benar menghentikan program).
- Punic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan.
- Saat panic function dipanggil, program akan berhenti, namun _defer function_ tetap akan dieksekusi.

```go
func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runApp(false)
}
```

Output:
```go
1. End App // jika kondisi false

1. End App
2. Panic:Error // jika kondisi true
```

## Recover 
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic.
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.

1. Kode Program Recover yang Salah :
```go
func RunApp(error bool) {
	defer endApp()
	if error {
		panic("Error") // Program selesai sampai sini
	}
    
	// Sehingga kode program dibawah ini tidak akan dieksekusi
	message := recover() 
	fmt.Println("Terjadi Panic:", message)
}
```
2. Kode Program Recover yang Benar :
```go
/*
Penjelsan:
Cara yang benar menggunakan recover function
dengan cara recover function nya digunakan didalam defer function nya.
*/

func EndApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic:", message)
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	RunApp(true)
}
```
