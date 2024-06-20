Proje Raporu: Blog Platformu
Proje Adı: Blog Platformu
Proje Amacı

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
