package main

import (
	"fmt"
	
)

var (
	harfNotlari      [9]string  = [9]string{"AA", "BA", "BB", "CB", "CC", "DC", "DD", "FD", "FF"}
	harfSinirlari    [8]float32 = [8]float32{90, 80, 70, 60, 50, 40, 30,20}
	sinavAgirliklari [3]float32 = [3]float32{30, 30, 40}
)

var (
	katilanlar   [100]string  = [100]string{"Öngen Gürsal Mansiz", "Bidayet Yıldırım", "Borataş Besin Sezer", "Heva Şehreban Akgündüz Ergül", "Bayan Açılay Duha Dumanlı", "Zahfer Safinaz Arslan", "Sanur Demir", "Kâzime Eraslan", "Başok Alparslan Karadeniz", "Kiyasi Aslan", "Bayan Tanyu Hançer Bilir", "Abuzar Bilge", "Kayrabay Edgübay Gül İhsanoğlu", "Tiğin Bilge", "Gök Aygönenç Şama Kısakürek", "Sami Adlan Çorlu", "Dr. Bediriye Ünsever Şensoy", "Dr. Gülgen Alize Dumanlı Sezer", "Akgöl Aytek Akar", "Çopur Ercihan İnönü", "Sayın Feza Durdu Karadeniz", "Uludağ Ersel Zengin", "Miray Pembesin Şener Mansız", "Dr. Ledün Tanak Bilge", "Fazul Akçay", "Halidun Korutürk", "Dr. Teybet Gülnaziye Seven", "Özokçu Manço", "Soyselçuk Akça", "Zaliha Behiza Yılmaz", "Seçgül Mansız", "Dr. Nurgil Ceyhun Ülker", "Savni Kısakürek", "Gümüştekin Ergül", "Geray Seracettin Akça Tevetoğlu", "Boynak Gönen Zorlu Mansız", "Dr. Sözer Akdeniz", "Bay Almus Bedri Demir", "Dr. Alaeddin Arcan Çorlu", "İnsaf Akgündüz", "Bayan Vefia Yücel Kısakürek Duran", "Nurseda Gülen", "Tekiner Tarhan", "Demiriz Şafak", "Şabettin Bilir", "Özlem Aslan", "Özdil Mansız", "Bay Altınkaya Özerdinç İnönü", "Bay Kırgız Şener", "Akgüneş Solma Akar", "Seyfullah Şevket Aslan Arslan", "Serda Hacile Yüksel Aksu", "Özger Karadeniz", "Ayçan Özge Akgündüz Aksu", "Öryürek Alemdar", "Hüda Nesfe Durdu", "Adal Aksu", "Dr. Kırgız Erdemer Akdeniz", "Mutluhan Altoğan Aksu Dumanlı", "Bay Ferat Rüşen Yıldırım", "Dr. Topuz Koçkan Ergül", "Çoğay Müfit Yılmaz", "Akkerman Şener", "Bayan Piran Rakide Aslan Karadeniz", "Gürten Duran", "Bay Tecimer Demirel", "Atasagun Şama", "Dr. Kutay Şener", "Dr. Özbilge Çetin Yıldırım", "Bayan Ticen Pekkan Soylu Soylu", "Dalan Sudi Yüksel Aksu", "Serezli Özinal İnönü", "Ayşeana Gülfeza Yıldırım", "Dr. Laze Hayrünnisa Çetin", "Mehmetzahir Erdoğan", "Dr. Nebiha Ertaş Ertaş", "Dr. Yazganalp Tunçkılıç Tarhan", "Bay İmren Duruk Dumanlı", "Bayan Muhiye Bidayet Şensoy Dumanlı", "Dr. Satıa Umuşan Sezer", "Bayan Binay Nurmelek Yaman Akçay", "Dr. Simten Dursadiye Bilir", "Tiğin Erdoğan", "Deryanur Çorlu Şener", "Naide Belgizar İnönü Karadeniz", "Rohat Sezgin", "Bayan Gülnaziye Sakarya Kısakürek", "Tuğrulhan Akça", "Dr. Goncafer Pekkan Demirel", "Dr. Aygutalp Barsen Seven", "Gökçe Kübran Zengin Sakarya", "Kiyasi Tahir Çetin", "Salkın Barak Akar", "Gülşa Ergül", "Dr. Erkan Türabi Türk", "Başok Yetişal Durmuş", "Esenbay Tekecan Yaman", "Bayan Yaşar Güllühan Şama", "Müslum Bilgen Alemdar", "Dr. Türabi Ebuakil Gülen"}
	vize1Notlari [100]float32 = [100]float32{2, 2, 7, 6, 6, 0, 9, 7, 16, 40, 78, 0, 55, 12, 53, 100, 43, 15, 31, 25, 88, 41, 35, 56, 18, 69, 5, 26, 89, 18, 81, 95, 38, 8, 28, 77, 16, 41, 15, 14, 44, 47, 5, 75, 96, 22, 85, 38, 37, 24, 5, 10, 72, 71, 34, 91, 28, 19, 85, 36, 33, 43, 100, 84, 78, 45, 19, 89, 23, 68, 2, 12, 21, 42, 95, 57, 81, 65, 25, 54, 79, 32, 18, 65, 52, 68, 26, 48, 72, 74, 51, 85, 83, 80, 33, 41, 14, 31, 58}
	vize2Notlari [100]float32 = [100]float32{16, 73, 69, 87, 92, 96, 94, 17, 46, 65, 84, 72, 2, 56, 27, 89, 56, 4, 73, 30, 100, 75, 43, 38, 26, 97, 63, 100, 2, 45, 85, 59, 39, 2, 50, 100, 10, 55, 80, 17, 2, 34, 42, 61, 42, 4, 78, 64, 16, 55, 61, 75, 84, 97, 50, 62, 20, 41, 15, 81, 38, 25, 30, 32, 69, 35, 6, 80, 6, 38, 73, 67, 80, 46, 9, 73, 75, 84, 41, 45, 30, 46, 23, 12, 13, 23, 43, 46, 37, 67, 27, 45, 1, 26, 98, 1, 73, 17, 64, 37}
	finalNotlari [100]float32 = [100]float32{79, 52, 13, 21, 10, 37, 75, 1, 9, 5, 65, 96, 70, 51, 47, 95, 82, 55, 70, 72, -1, 1, 24, 71, 20, 93, 78, 72, 40, 17, 13, 46, 72, 85, 63, 58, 11, 83, 56, 67, 68, 81, 36, 51, 68, 89, 22, 54, 12, 45, 0, 96, 69, 42, 97, 88, 23, 81, -1, 80, 38, 28, 69, 43, 47, 4, 12, 100, -1, 46, 56, 78, 23, 12, 30, 93, 100, 69, 46, 35, 83, 28, 80, 10, 16, 55, 87, 42, 87, 89, 37, 16, 20, 37, 89, 29, 71, 1, 90}
)

var (
	alinanNotlar       [100]float32 = [100]float32{}
	notOrtalamasi       [100]float32 = [100]float32{}
	alinanHarfNotlari  [100]string  = [100]string{}
	ogrenciDurumlari   [100]string  = [100]string{}
	gecenOgrenciSayisi int          = 0
	kalanOgrenciSayisi int          = 0
	sarliGegenOgrenciSayisi int = 0 
	gecerNot           float32      = 0
)

func main() {
	for i:=0; i<100; i++ {
		notOrtHesap(i)
		harfNotuBulma(i)

		harfler:=alinanHarfNotlari[i]

		if harfler==harfNotlari[8]{
			kalanOgrenciSayisi++
		} else if harfler==harfNotlari[7]{
			sarliGegenOgrenciSayisi++
		} else {
			gecenOgrenciSayisi++
		}

		fmt.Println("Ad-Soyad: ",katilanlar[i],"\nNot-Ortalamasi: ",notOrtalamasi[i],"\nHarf-Notu: ",alinanHarfNotlari[i],"\n =================================================================")
	}
/*fmt.Println(notOrtalamasi)
fmt.Println(alinanHarfNotlari)
fmt.Println(kalanOgrenciSayisi)
fmt.Println(sarliGegenOgrenciSayisi)
fmt.Println(gecenOgrenciSayisi)
*/
}


func notOrtHesap(i int){
	notOrtalamasi[i] =vize1Notlari[i]*sinavAgirliklari[0]/100+vize2Notlari[i]*sinavAgirliklari[1]/100+finalNotlari[i]*sinavAgirliklari[2]/100		
}


func harfNotuBulma(i int){ortalama:=notOrtalamasi[i]
	if finalNotlari[i] == -1 {
		alinanHarfNotlari[i] = harfNotlari[8]
	} else if ortalama>harfSinirlari[0]{
		alinanHarfNotlari[i] = harfNotlari[0]
	} else if ortalama>harfSinirlari[1]{
		alinanHarfNotlari[i] = harfNotlari[1]
	} else if ortalama>harfSinirlari[2]{
		alinanHarfNotlari[i] = harfNotlari[2]
	} else if ortalama>harfSinirlari[3]{
		alinanHarfNotlari[i] = harfNotlari[3]
	} else if ortalama>harfSinirlari[4]{
		alinanHarfNotlari[i] = harfNotlari[4]
	} else if ortalama>harfSinirlari[5]{
		alinanHarfNotlari[i] = harfNotlari[5]
	} else if ortalama > harfSinirlari[6]{
		alinanHarfNotlari[i] = harfNotlari[6]
	} else if ortalama > harfSinirlari[7]{
		alinanHarfNotlari[i] = harfNotlari[7]
	}
}
//çalışıyosa dokunma