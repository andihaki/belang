# cheat sheet
[https://github.com/a8m/go-lang-cheat-sheet]https://github.com/a8m/go-lang-cheat-sheet

# soalan
- Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?
    it creates a copy of each argument, and these copies are used inside of the function
- func (d data) changeName(){
    (*d).firstName = "jojon"
    }
    *location = type data yang akan dikembalikan dari reciver

# package
1. main = executable package
2. !main = resusable package

# array
1. array = tuples (static / immutable)
2. slice = array (dynamic / mutable)

# slice
sliceName[0:2] => ambil data index sampai 2, tidak termasuk 2 (0, 1)
sliceName[:2] == sliceName[0:2] => true

# type conversion
[]byte("Hi boss!") => convert "Hi boss!" string to byte

# tipe data: struct
struktur data. koleksi dari properties yang berhubungan satu sama lain.
contain suit and value
*sepertinya mirip dictionary / object
*fmt.Printf("%+v", namaStruct) // print zero value

"""
type person struct {
	firstName string
	lastName  string
}

func main() {
	spongebob := person{firstName: "Spongebob", lastName: "Squarepants"}
	fmt.Println(spongebob)
}
"""

# pointer
- &namaStruct = ambil lokasi penyimpanan di memory
    ubah **value** ke **address**
- *namaPointer = ambil nilai dari memory address
    ubah **address** ke **value**

# zero value
di go "katanya" tidak panya null / nill, tapi:
string = ""
int = 0
float = 0
bool = false

# data type
- value types: int, float, string, bool, struct
    pake pointer untuk ubah nilai di dalam fungsi
- reference types: slices, maps, channel, pointer, function
    ubah aja, bebas