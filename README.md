   <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a>
    <br>
Proje Raporu: Blog Platformu <br>
Proje Adı: Blog Platformu <br>
Proje Amacı: <br>
Bu projenin amacı, kullanıcıların yazı yazabileceği, bu yazıları paylaşabileceği ve diğer kullanıcıların bu yazılara yorum yapabileceği bir blog platformu geliştirmektir. Platform, kullanıcıların kendi içeriklerini oluşturma ve başkalarının içeriklerini görüntüleyip etkileşimde bulunma süreçlerini kolaylaştırmayı hedeflemektedir. Proje, temel web geliştirme becerilerini öğrenmek ve uygulamak isteyen bireyler için pratik bir örnek teşkil etmektedir.
Proje Özeti 

Blog Platformu, kullanıcıların kayıt olup giriş yapabilecekleri, yeni blog yazıları oluşturabilecekleri, mevcut yazıları düzenleyip silebilecekleri ve bu yazılara yorum yapabilecekleri bir web uygulamasıdır. Uygulama, kullanıcıların içerik oluşturma, düzenleme ve diğer kullanıcılarla etkileşime geçme gibi temel işlevleri gerçekleştirmelerine olanak tanır. Platform, basit ve kullanıcı dostu bir arayüzle, yazıların ve yorumların yönetimini sağlar.
Temel Özellikler 

    Kullanıcı Kayıt ve Giriş Sistemi: Kullanıcılar platforma kayıt olabilir ve mevcut hesaplarıyla giriş yapabilirler.
    Yazı Ekleme, Düzenleme ve Silme: Kullanıcılar yeni blog yazıları oluşturabilir, bu yazıları düzenleyebilir ve gerekirse silebilirler.
    Yorum Yapma ve Listeleme: Kullanıcılar blog yazılarına yorum yapabilir ve diğer kullanıcıların yorumlarını görüntüleyebilirler.

Teknolojiler ve Araçlar


    Programlama Dili: Go (Golang)
    Go, yüksek performanslı ve verimli uygulamalar geliştirmek için kullanılan modern bir programlama dilidir. Go'nun basitliği ve güçlü kütüphane desteği, bu projenin temel yapı taşlarından biridir.
    Web Çerçevesi: Gorilla Mux
    Gorilla Mux, HTTP yönlendirmelerini ve URL eşleştirmelerini kolaylaştıran popüler bir Go paketi olup, RESTful API geliştirme süreçlerini basitleştirir.
    Şablon Motoru: HTML Templates
    HTML şablonları, dinamik içerik oluşturma ve kullanıcıya sunma sürecini yönetmek için kullanılır. Bu şablonlar, kullanıcı arayüzünün yapılandırılmasında temel rol oynar.
    Oturum Yönetimi: Gorilla Sessions
    Kullanıcı oturumlarını yönetmek için Gorilla Sessions kullanılır. Bu paket, kullanıcıların oturum açma durumlarını ve kimlik doğrulama bilgilerini güvenli bir şekilde saklamayı sağlar.

Proje Yapısı

Blog Platformu projesi, çeşitli işlevleri ayrı dosya ve modüller halinde düzenleyerek yapılandırılmıştır:

    main.go: Uygulamanın başlangıç noktasıdır. HTTP sunucusunu başlatır ve yönlendirme işlevlerini tanımlar.
    handlers.go: HTTP isteklerini işleyen ve uygun şablonları render eden fonksiyonları içerir.
    models.go: Kullanıcılar, yazılar ve yorumlar gibi veri modellerini tanımlar.
    templates: HTML şablon dosyalarının bulunduğu klasördür. Bu şablonlar, kullanıcı arayüzünü oluşturmak için kullanılır.

Kullanıcı İş Akışı

    Kayıt ve Giriş: Kullanıcılar platforma kaydolur ve giriş yapar. Giriş yapmış kullanıcılar, kişisel oturum bilgilerini saklayan güvenli çerezlerle tanımlanır.
    Yazı Oluşturma: Giriş yapmış kullanıcılar, yeni blog yazıları oluşturabilir. Oluşturulan yazılar, başlık ve içerik bilgilerini içerir.
    Yazı Düzenleme ve Silme: Kullanıcılar, kendi oluşturdukları yazıları düzenleyebilir veya silebilirler.
    Yorum Yapma: Kullanıcılar, blog yazılarına yorum yapabilir. Her bir yazıya yapılan yorumlar, yazı detay sayfasında listelenir.

Sonuç

Blog Platformu, temel bir blog sistemi geliştirmek için gerekli olan tüm bileşenleri içeren kapsamlı bir projedir. Kullanıcı kayıt ve giriş sistemi, yazı oluşturma, düzenleme ve silme, yorum yapma ve listeleme gibi işlevler, bu projenin temel özelliklerindendir. Proje, Go programlama dilini kullanarak modern web uygulamaları geliştirmek isteyen kişiler için uygun bir öğrenme ve uygulama fırsatı sunmaktadır. Projenin ilerleyen aşamalarında, veri tabanı entegrasyonu, gelişmiş kimlik doğrulama yöntemleri ve daha zengin kullanıcı arayüzleri gibi ek özellikler eklenerek platformun işlevselliği artırılabilir.

 <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a>
    <br>
Go ile Basit Bir Web Uygulaması Geliştirme
Giriş
Bu makale, Go programlama dili kullanılarak geliştirilmiş basit bir web uygulamasının işleyişini ve bileşenlerini açıklamaktadır. Web uygulaması, kullanıcıların iki farklı sayfaya (ana sayfa ve hakkında sayfası) erişebilmesini sağlayan bir HTTP sunucusunu içermektedir. Uygulama, temel HTTP yönlendirmeleri ve şablonlar kullanılarak dinamik içerik sunma konularını ele almaktadır.

Proje Yapısı
Proje aşağıdaki dosya yapısına sahiptir:

css
Kodu kopyala
mywebsite/
├── main.go
├── templates/
│   ├── index.html
│   └── about.html
Ana Dosya: main.go
main.go dosyası, uygulamanın ana çalışma dosyasıdır ve HTTP sunucusunu başlatan, yönlendirmeleri tanımlayan ve şablonları render eden fonksiyonları içerir.

Paketlerin İçe Aktarılması
Uygulama, html/template, log ve net/http paketlerini içe aktarmaktadır. html/template paketi, HTML şablonlarını işlemek için kullanılırken, log paketi hata ve bilgi mesajlarını kaydetmek için kullanılmaktadır. net/http paketi ise HTTP sunucusunu oluşturmak ve yönetmek için kullanılır.

Şablonların Yüklenmesi
Uygulama başlatıldığında, şablon dosyaları template.Must(template.ParseGlob("templates/*.html")) komutu ile yüklenir. Bu, templates klasöründeki tüm .html dosyalarını yükler ve bir hata oluşursa programın başlatılmasını engeller.

HTTP Yönlendirmeleri
Uygulama, iki temel yönlendirme tanımlar:

/: Ana sayfa yönlendirmesi
/about: Hakkında sayfası yönlendirmesi
Bu yönlendirmeler, http.HandleFunc fonksiyonu kullanılarak tanımlanır ve her biri bir handler fonksiyonuna yönlendirilir.

Handler Fonksiyonları
Handler fonksiyonları, belirli bir URL'ye gelen HTTP isteklerini işler. Bu uygulamada iki handler fonksiyonu bulunmaktadır:


indexHandler: Ana sayfa isteklerini işler. renderTemplate fonksiyonunu çağırarak index.html şablonunu render eder.
aboutHandler: Hakkında sayfası isteklerini işler. renderTemplate fonksiyonunu çağırarak about.html şablonunu render eder.
Şablonları Render Eden Fonksiyon
renderTemplate fonksiyonu, belirli bir şablon dosyasını render eder ve sonucu HTTP yanıtına yazar. Eğer şablon render edilirken bir hata oluşursa, bu hata HTTP yanıtına yazılır ve istemciye gönderilir.

Sunucunun Başlatılması
main fonksiyonu, tüm yönlendirmeleri tanımladıktan sonra, log.Fatal(http.ListenAndServe(":8080", nil)) komutunu kullanarak HTTP sunucusunu başlatır. Bu komut, sunucunun localhost üzerindeki 8080 portunda çalışmasını sağlar. Eğer sunucu başlatılırken bir hata oluşursa, bu hata log.Fatal ile kaydedilir ve program sonlandırılır.

Şablon Dosyaları
index.html
index.html dosyası, ana sayfa içeriğini tanımlar. HTML yapısında bir başlık, bir hoş geldiniz mesajı ve hakkında sayfasına yönlendiren bir bağlantı içerir.

about.html
about.html dosyası, hakkında sayfası içeriğini tanımlar. HTML yapısında bir başlık, bir açıklama ve ana sayfaya yönlendiren bir bağlantı içerir.

Sonuç
Bu basit Go web uygulaması, HTTP yönlendirmeleri ve şablonları kullanarak dinamik içerik sunmanın temel adımlarını göstermektedir. Go dilinin sunduğu güçlü ve performanslı yapılar sayesinde, daha karmaşık web uygulamaları ve API'ler kolaylıkla geliştirilebilir. Bu örnek, temel HTTP sunucusu işleyişi, yönlendirme ve şablon işleme konularında iyi bir başlangıç noktası sunmaktadır. Gerçek dünya uygulamalarında, veritabanı entegrasyonu, güvenlik önlemleri ve ölçeklenebilirlik gibi konulara da dikkat edilmelidir.    

![Ekran görüntüsü 2024-07-22 105532](https://github.com/user-attachments/assets/08fd8ce2-368c-4de3-8bc5-8a8f1927029c)
![Ekran görüntüsü 2024-07-22 105459](https://github.com/user-attachments/assets/b4666e66-2f7b-4347-8cd7-7d3f7a005c41)

