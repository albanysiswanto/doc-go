# Documentation Go Basic

Memuat isi Go-Lang basic yang sering atau penting untuk di pakai.

### Type Declaration

Di Go, dimungkinkan untuk mengikat suatu tipe ke pengenal untuk membuat tipe bernama baru yang dapat direferensikan dan digunakan di mana pun tipe tersebut diperlukan. Mendeklarasikan suatu tipe mengambil format umum sebagai berikut:
```type <name identifier> <underlying type name>```

Note:
Deklarasi tipe juga dapat menggunakan literal tipe komposit sebagai tipe dasarnya. Tipe komposit meliputi array, irisan, peta, dan struct. Bagian ini berfokus pada tipe non-komposit.

```go
func main() {
	type numberPhone string

	var phone numberPhone = "0823232323"
	var contohPhone = "0823232321"
	fmt.Println(phone)
	fmt.Println(numberPhone(contohPhone))
}
```
Output:
```go
1. 0823232323
2. 0823232321
```