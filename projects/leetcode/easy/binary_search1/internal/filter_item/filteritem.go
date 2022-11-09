package filteritem
/*
Runner Config
Maximum Memory
Maximum Runtime
:
:
128 MB
1 Seconds
Filter Harga dan Kuantitas Barang
Deskripsi
Pada Algommerce, terdapat N barang, setiap barang Bi memiliki deskripsi (Vi, Qi) masing-masing menyatakan  harga barang dan kuantitas barang tersebut.

Suatu hari, terdapat M users, dengan setiap user ke-j melakukkan query berisi deskripsi (Lj, Rj, Xj, Yj) yang artinya user ke-j ini ingin melakukan filter untuk mencari banyaknya barang Bi sedemikian sehingga:

(Lj ≤ Vi ≤ Rj) dan (Xj ≤ Qi ≤ Yj)

BKarena Algommerce masih tergolong perusahaan baru, maka kuantitas setiap barang yang ada bernilai maksimal 10, dengan kata lain, Qi ≤ 10 untuk setiap barang Bi.

Anda sebagai programmer handal diminta untuk membantu menjawab setiap query tersebut.

Format Masukan (Format Input)
Baris pertama berisi bilangan bulat positif N menyatakan banyaknya barang.

Baris kedua berisi list dari deskripsi setiap barang, setiap element ke-i dari list tersebut menyatakan (Vi, Qi).

Baris ketiga berisi bilangan bulat positif M menyatakan banyaknya user.

Baris keempat berisi list dari query setiap user, setiap element ke-j dari list tersebut menyatakan (Lj, Rj, Xj, Yj).

Format Keluaran (Format Output)
Keluarkan M baris, masing-masing menyatakan hasil query dari setiap M users.

Batasan Input
1 ≤ N, M ≤ 100.000
1 ≤ Vi ≤ 2.000.000.000
1 ≤ Qi ≤ 10
1 ≤ Lj ≤ Rj ≤ 2.000.000.000
1 ≤ Xj ≤ Yj ≤ 10
Contoh
Contoh Input 	Contoh Output 
5
[[5 ,2] ,[4 ,2] ,[7 ,7] ,[8 ,3] ,[4 ,5]]
3
[[5 ,5 ,2 ,2] ,[4 ,7 ,5 ,7] ,[4 ,8 ,2 ,7]]
1
2
5
Penjelasan
Berikut merupakan penjelasan untuk setiap query:

• Untuk (Lj, Rj, Xj, Yj) = (5, 5, 2, 2) kita dapatkan 1 barang yang memenuhi yakni B1.
• Untuk (Lj, Rj, Xj, Yj) = (4, 7, 5, 7) kita dapatkan 2 barang yang memenuhi yakni B3 dan B5.
• Untuk (Lj, Rj, Xj, Yj) = (4, 8, 7, 2) kita dapatkan semua barang memenuhi, sehingga jawaban tersebut 5
*/

// a := [][]uint8{{5, 2,},{4, 2},{7, 7},{8, 3},{4, 5}}
func FilterItem(int, [][]int, int, [][]int) (int,int,int) {
	var i, j, k int
	return i, j, k
}
