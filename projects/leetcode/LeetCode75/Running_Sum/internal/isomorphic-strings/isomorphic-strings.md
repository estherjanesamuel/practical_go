# 205. Isomorphic Strings

Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

 

Example 1:
Input: s = "egg", t = "add"
Output: true

Example 2:
Input: s = "foo", t = "bar"
Output: false

Example 3:
Input: s = "paper", t = "title"
Output: true
 

Constraints:

1 <= s.length <= 5 * 104
t.length == s.length
s and t consist of any valid ascii character.


Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
ChatGPT
Diberikan dua string s dan t, periksa apakah keduanya isomorfik atau tidak. Dua string dikatakan isomorfik jika karakter dalam string s dapat diganti untuk mendapatkan string t. Setiap kemunculan karakter harus diganti dengan karakter lain sambil mempertahankan urutan karakter. Tidak ada dua karakter yang mungkin dipetakan ke karakter yang sama, tetapi karakter dapat dipetakan ke dirinya sendiri.

Untuk menyelesaikan masalah ini, kita dapat menggunakan sebuah kamus (dictionary) untuk memetakan setiap karakter dalam s ke karakter dalam t. Kita juga perlu mengecek apakah setiap karakter dalam t sudah dipetakan atau belum, dan apakah ada dua karakter dalam s yang dipetakan ke karakter yang sama dalam t.

Berikut adalah contoh implementasi dalam bahasa Python:

python
Copy code
def isIsomorphic(s, t):
    if len(s) != len(t):
        return False
    
    mapping = {}
    used_chars = set()
    for i in range(len(s)):
        if s[i] in mapping:
            if mapping[s[i]] != t[i]:
                return False
        else:
            if t[i] in used_chars:
                return False
            mapping[s[i]] = t[i]
            used_chars.add(t[i])
    
    return True
Pada fungsi isIsomorphic, kita pertama-tama memeriksa apakah panjang s sama dengan panjang t. Jika tidak, maka keduanya tidak bisa isomorfik.

Kemudian kita menggunakan sebuah kamus mapping untuk memetakan setiap karakter dalam s ke karakter dalam t, dan sebuah set used_chars untuk menyimpan karakter yang sudah dipetakan.

Untuk setiap karakter dalam s, kita periksa apakah karakter tersebut sudah dipetakan ke karakter dalam t. Jika sudah, maka kita periksa apakah karakter dalam t yang sudah dipetakan sama dengan karakter di indeks i dari t. Jika tidak, maka keduanya tidak bisa isomorfik.

Jika karakter dalam s belum dipetakan ke karakter dalam t, maka kita periksa apakah karakter dalam t sudah dipetakan atau belum. Jika sudah, maka keduanya tidak bisa isomorfik. Jika belum, maka kita memetakan karakter dalam s ke karakter di indeks i dari t, dan menandai karakter tersebut sebagai sudah dipetakan.

Jika kita berhasil melakukan pemetaan untuk setiap karakter dalam s dan tidak ada dua karakter dalam s yang dipetakan ke karakter yang sama dalam t, maka s dan t isomorfik. Jika tidak, maka keduanya tidak bisa isomorfik.





User
code in go
ChatGPT
Berikut adalah contoh implementasi algoritma isIsomorphic dalam bahasa Go:

```go
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapping := make(map[byte]byte)
    usedChars := make(map[byte]bool)

    for i := 0; i < len(s); i++ {
        if _, ok := mapping[s[i]]; ok {
            if mapping[s[i]] != t[i] {
                return false
            }
        } else {
            if usedChars[t[i]] {
                return false
            }
            mapping[s[i]] = t[i]
            usedChars[t[i]] = true
        }
    }

    return true
}
```
Fungsi isIsomorphic mirip dengan implementasi dalam bahasa Python. Pertama-tama, kita memeriksa apakah panjang s sama dengan panjang t. Jika tidak, maka keduanya tidak bisa isomorfik.

Kemudian kita menggunakan sebuah map mapping untuk memetakan setiap karakter dalam s ke karakter dalam t, dan sebuah map usedChars untuk menyimpan karakter yang sudah dipetakan.

Untuk setiap karakter dalam s, kita periksa apakah karakter tersebut sudah dipetakan ke karakter dalam t. Jika sudah, maka kita periksa apakah karakter dalam t yang sudah dipetakan sama dengan karakter di indeks i dari t. Jika tidak, maka keduanya tidak bisa isomorfik.

Jika karakter dalam s belum dipetakan ke karakter dalam t, maka kita periksa apakah karakter dalam t sudah dipetakan atau belum. Jika sudah, maka keduanya tidak bisa isomorfik. Jika belum, maka kita memetakan karakter dalam s ke karakter di indeks i dari t, dan menandai karakter tersebut sebagai sudah dipetakan.

Jika kita berhasil melakukan pemetaan untuk setiap karakter dalam s dan tidak ada dua karakter dalam s yang dipetakan ke karakter yang sama dalam t, maka s dan t isomorfik. Jika tidak, maka keduanya tidak bisa isomorfik.


## Isomorphic Debug
egg add

s	t
-	-
e	a
g	d
g	d

map
used

map[{e, a},{g,d},{g,d}]
used[{a,true}, {d,true}]

foo bar

s	t
-	-
f	b
o	a
o	r

map
used

map[{f, b},{o,a},{g,r}]
used[{b,true}, {a,true}]

paper_title

s	t
-	-
p	t
a	i
p	t
e	l
r	e

map
used

map[{p, t},{a,i},{p,t},{e,l},{r,e}]
used[{t,true}, {a,true}]


cek jika len dari s dan t sama
jadikan dictionary <s,t>
simpan masing2 char ke dict 
cek jika key char(s) sudah ada(exist) maka cek value (t) apakah sama, 
jika tidak sama maka return false (tidak isomorphic)
jika sama lanjut cek berikutnya.sampai habis,

jika key tidak ada maka cek apakah value sudah ada? 
jika sudah ada maka return false (tidak isomorphic)
jika tidak ada lanjut sampai selesai dan kembalikan true (isomorphic)