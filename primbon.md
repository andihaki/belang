# package
1. main = executable package
2. !main = resusable package

# array
1. array = tuples (static / immutable)
2. slice = array (dynamic / mutable)

# slice
sliceName[0:2] => ambil data index sampai 2, tidak termasuk 2 (0, 1)
sliceName[:2] == sliceName[0:2] => true

#type conversion
[]byte("Hi boss!") => convert "Hi boss!" string to byte