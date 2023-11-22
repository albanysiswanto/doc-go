# Struct

- Struct adalah sebuah template data yang digunakan untuk menguhubungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
- Struct biasamya representasi data dalam program aplikasi yang kita buat.
- Data di Sturct disimpan dalam field.
- Sederhananya struct adalah kumpulan data field.
```go
type name_struct struct{
	field1
	field2
	field3
}
```
Contoh:
```go
type Customer struct {
    Name, Address string
    Age int
}
```
### Membuat Data Struct
- Struct adalah template data atau prototype data.
- Struct tidak bisa langsung digunakan.
- Namun kita bisa membuat  data/object dari struct yang telah kita buat.
```go
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
    var bany Customer
    bany.Name = "Albany Siswanto"
    bany.Address = "Subang"
    bany.Age = 18
    fmt.Println(bany.Name)
    fmt.Println(bany.Address)
    fmt.Println(bany.Age)
}
```
Output:
```go
Albany Siswanto
Subang
18
```

### Struct Literals
- Sebelumnya kita telah membuat data dari _Struct_, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct.

Contoh:
```go
type Customer struct {
    Name, Address string
    Age           int
}

func main() {
budi := Customer{
    Name: "Budiman",
    Address: "Bandung",
    Age: 17,
    }
    fmt.Println(budi)
    
    // ATAU
    
    udin := Customer{"Udin", "Papua", 13}
    fmt.Println(udin)
}
```
# Struct Method
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function.
- Namun jika kita ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struct memiliki function.
- Method adalah function
```go
type Customers struct {
Name, Address string
Age           int
}

func (customer Customers) sayHello(name string) {
fmt.Println("Hello,", name, "My Name is", customer.Name)
}

func main() {
rully := Customers{Name: "Rully"}
rully.sayHello("Agus")
}
```
Output:
```go
Hello, Agus My Name is Rully
```