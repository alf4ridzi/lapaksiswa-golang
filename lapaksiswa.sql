-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: lapaksiswa
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ci_sessions`
--

DROP TABLE IF EXISTS `ci_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `ci_sessions` (
  `id` varchar(128) NOT NULL,
  `ip_address` varchar(45) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `data` blob NOT NULL,
  KEY `ci_sessions_timestamp` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ci_sessions`
--

LOCK TABLES `ci_sessions` WRITE;
/*!40000 ALTER TABLE `ci_sessions` DISABLE KEYS */;
INSERT INTO `ci_sessions` VALUES ('ci_session:jk7c2fg548fq7d7e7jjb97bskmdp0kep','::1','2024-12-05 07:28:02',_binary '__ci_last_regenerate|i:1733383682;_ci_previous_url|s:53:\"http://localhost:8080/index.php/produk/mouse-wireless\";'),('ci_session:ne2fn1hmm62rohg7ea7osh7rjnq7k7j1','::1','2024-12-05 06:51:43',_binary '__ci_last_regenerate|i:1733381503;_ci_previous_url|s:43:\"http://localhost:8080/index.php/search/kaos\";'),('ci_session:h8u2ih4nuqs7s0pe5beeavdmn0he2p58','::1','2024-12-05 06:57:43',_binary '__ci_last_regenerate|i:1733381863;_ci_previous_url|s:43:\"http://localhost:8080/index.php/search/test\";'),('ci_session:7lrhud8lsf9itlpgh4t6obhcakababok','::1','2024-12-05 07:03:23',_binary '__ci_last_regenerate|i:1733382203;_ci_previous_url|s:43:\"http://localhost:8080/index.php/search/test\";'),('ci_session:5om60akfj7t463i8pugesetikve5iorr','::1','2024-12-05 07:08:40',_binary '__ci_last_regenerate|i:1733382520;_ci_previous_url|s:55:\"http://localhost:8080/index.php/produk/kaos-polos-putih\";'),('ci_session:td4mq0vgdstfsupalnhoj5t87n7is3j0','::1','2024-12-05 07:15:05',_binary '__ci_last_regenerate|i:1733382905;_ci_previous_url|s:46:\"http://localhost:8080/index.php/search/pakaian\";'),('ci_session:u9pvel1vjvtgijlqghrdto110god57ad','::1','2024-12-05 07:21:56',_binary '__ci_last_regenerate|i:1733383316;_ci_previous_url|s:46:\"http://localhost:8080/index.php/search/pakaian\";'),('ci_session:evqhspv6l0ccnlnc4i5aos2jnunccou2','::1','2024-12-05 07:18:24',_binary '__ci_last_regenerate|i:1733383064;_ci_previous_url|s:46:\"http://localhost:8080/index.php/search/pakaian\";'),('ci_session:og60fgkgelraebkhplb4dle3pmrd953l','::1','2024-12-05 07:28:12',_binary '__ci_last_regenerate|i:1733383692;_ci_previous_url|s:46:\"http://localhost:8080/index.php/search/pakaian\";'),('ci_session:lud7tin7n5lp4b5avfhhqkmcbl3lu0dg','::1','2024-12-05 07:36:15',_binary '__ci_last_regenerate|i:1733384175;_ci_previous_url|s:54:\"http://localhost:8080/index.php/search?keyword=pakaian\";'),('ci_session:ka7877c0mbvenijvn5qbp3a23etoo6dv','::1','2024-12-05 07:28:44',_binary '__ci_last_regenerate|i:1733383692;_ci_previous_url|s:54:\"http://localhost:8080/index.php/search?keyword=pakaian\";'),('ci_session:f5thmbn302unght3fdau3aj2v93aj52g','::1','2024-12-05 07:52:07',_binary '__ci_last_regenerate|i:1733385127;_ci_previous_url|s:51:\"http://localhost:8080/index.php/search?keyword=test\";'),('ci_session:2rldn5cov4o3e22h0euabo57l1hbb7q4','::1','2024-12-05 07:55:21',_binary '__ci_last_regenerate|i:1733385127;_ci_previous_url|s:54:\"http://localhost:8080/index.php/search?keyword=pakaian\";'),('ci_session:0147jqlfb8r9f01fu598e1m6uo1sm8i9','::1','2024-12-05 12:26:58',_binary '__ci_last_regenerate|i:1733401618;_ci_previous_url|s:54:\"http://localhost:8080/index.php/search?keyword=pakaian\";'),('ci_session:uil1ikn1e097buqc79e7o974nrpmm5sa','::1','2024-12-05 12:44:49',_binary '__ci_last_regenerate|i:1733402689;_ci_previous_url|s:60:\"http://localhost:8080/index.php/produk/asus-rog-zephyrus-g14\";'),('ci_session:1mbsruv7d3u1s726kje4mvgfa4bdaj19','::1','2024-12-05 14:28:24',_binary '__ci_last_regenerate|i:1733408904;_ci_previous_url|s:56:\"http://localhost:8080/index.php/search?keyword=Zephyrus+\";'),('ci_session:hq2a0jj71reh4iec7f6uh9okt8b8to1r','::1','2024-12-05 14:23:42',_binary '__ci_last_regenerate|i:1733408622;_ci_previous_url|s:118:\"http://localhost:8080/index.php/search?keyword=&kategori=makanan&harga_min=&harga_max=&kondisi=&rating=&urutan=terbaru\";'),('ci_session:3srbnr67e7f8j5qphds9upbiot89doek','::1','2024-12-05 14:26:09',_binary '__ci_last_regenerate|i:1733408622;_ci_previous_url|s:54:\"http://localhost:8080/index.php/search?keyword=pakaian\";'),('ci_session:v9m1j7puq4acas15qamu6cbfpj8vmll0','::1','2024-12-05 14:28:24',_binary '__ci_last_regenerate|i:1733408904;_ci_previous_url|s:56:\"http://localhost:8080/index.php/search?keyword=Zephyrus+\";'),('ci_session:111ljrmslpom3nop6fbnn6t310t5gh7r','::1','2024-12-10 00:44:18',_binary '__ci_last_regenerate|i:1733791458;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";'),('ci_session:rvj3f540es1uesh512ugms47tlvc0fb0','::1','2024-12-10 01:30:58',_binary '__ci_last_regenerate|i:1733794258;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";'),('ci_session:t314m2neei6ivt758lo23qn6n8cug76s','::1','2024-12-10 01:36:46',_binary '__ci_last_regenerate|i:1733794606;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";'),('ci_session:1b603ofadccfjg7v26pv2s65bc70r4v9','::1','2024-12-10 01:55:35',_binary '__ci_last_regenerate|i:1733795735;_ci_previous_url|s:37:\"http://localhost:8080/index.php/login\";'),('ci_session:o1fboeef9970k0jtoot85ik6sbpv4el7','::1','2024-12-10 01:47:40',_binary '__ci_last_regenerate|i:1733795260;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";'),('ci_session:sqh020ge3hp3ac97pp0onhte99bnsm3a','::1','2024-12-10 01:55:00',_binary '__ci_last_regenerate|i:1733795700;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:7ldk688dudd22f133okomr84nc6oshl2','::1','2024-12-10 02:06:35',_binary '__ci_last_regenerate|i:1733796395;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:dk50dh7574iuvdfi70mhafussqmeievo','::1','2024-12-10 02:03:36',_binary '__ci_last_regenerate|i:1733796216;_ci_previous_url|s:37:\"http://localhost:8080/index.php/login\";'),('ci_session:dh3eoc26perc4hod684m07dgq3fqsg5c','::1','2024-12-10 02:08:59',_binary '__ci_last_regenerate|i:1733796539;_ci_previous_url|s:32:\"http://localhost:8080/index.php/\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:2n7l3u9eejenk03f3tnvn97kgs5j1fre','::1','2024-12-10 02:11:25',_binary '__ci_last_regenerate|i:1733796395;_ci_previous_url|s:49:\"http://localhost:8080/index.php/search?keyword=22\";'),('ci_session:lovk6s38i98h7t6ae71ov03gehi8ea3m','::1','2024-12-10 02:16:06',_binary '__ci_last_regenerate|i:1733796966;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:p7fo22cbl7p8orcldfeacbv0r07q6ck8','::1','2024-12-10 02:16:06',_binary '__ci_last_regenerate|i:1733796966;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:bqdpp09fshj3utes4bs7l2de8ln4537b','::1','2024-12-11 03:30:54',_binary '__ci_last_regenerate|i:1733887854;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:d5va2ucgrb5k9qotcuednjeklac8g5h5','::1','2024-12-11 03:36:42',_binary '__ci_last_regenerate|i:1733888202;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:uocclrnol8qv7gfb496oomvn44v5nhaa','::1','2024-12-11 03:43:15',_binary '__ci_last_regenerate|i:1733888595;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:sk8mne207mtlalbeeaguv254n06lu82l','::1','2024-12-11 03:47:15',_binary '__ci_last_regenerate|i:1733888595;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:kchdvhvhmhm01qncjb39ba7eh9bid43t','::1','2024-12-12 02:05:25',_binary '__ci_last_regenerate|i:1733969125;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:dl26q65rsm2qhnftk25egqde7baaaet4','::1','2024-12-12 02:14:51',_binary '__ci_last_regenerate|i:1733969691;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:26cuj1m71o9h6g7355gbk10l4tb0mks9','::1','2024-12-12 03:08:42',_binary '__ci_last_regenerate|i:1733972922;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:1mpbrcadaedgp2gmsgd48q4kc97et4s8','::1','2024-12-12 03:24:50',_binary '__ci_last_regenerate|i:1733973890;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:q5u39jmo5tfhnjqk748263m3pousgpic','::1','2024-12-12 04:33:02',_binary '__ci_last_regenerate|i:1733977982;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:mvlgc3f71so5vngo83atdnvrgm0ci1vn','::1','2024-12-12 04:38:38',_binary '__ci_last_regenerate|i:1733978318;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:lnnorakmuucehfdk8qiff2kdciud2s1d','::1','2024-12-12 04:48:01',_binary '__ci_last_regenerate|i:1733978881;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";'),('ci_session:99lmb103c9djk7rephgoqeljtfae5k8h','::1','2024-12-12 04:51:36',_binary '__ci_last_regenerate|i:1733978881;_ci_previous_url|s:36:\"http://localhost:8080/index.php/user\";isLogin|b:1;role|s:4:\"user\";username|s:9:\"alfaridzi\";nama|s:4:\"test\";');
/*!40000 ALTER TABLE `ci_sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `foto_komentar`
--

DROP TABLE IF EXISTS `foto_komentar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `foto_komentar` (
  `id` int NOT NULL AUTO_INCREMENT,
  `komentar_id` varchar(50) NOT NULL,
  `foto` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foto_komentar`
--

LOCK TABLES `foto_komentar` WRITE;
/*!40000 ALTER TABLE `foto_komentar` DISABLE KEYS */;
/*!40000 ALTER TABLE `foto_komentar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `foto_produk`
--

DROP TABLE IF EXISTS `foto_produk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `foto_produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `foto` varchar(255) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foto_produk`
--

LOCK TABLES `foto_produk` WRITE;
/*!40000 ALTER TABLE `foto_produk` DISABLE KEYS */;
/*!40000 ALTER TABLE `foto_produk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kategori`
--

DROP TABLE IF EXISTS `kategori`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `kategori` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kategori` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `slug` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kategori`
--

LOCK TABLES `kategori` WRITE;
/*!40000 ALTER TABLE `kategori` DISABLE KEYS */;
INSERT INTO `kategori` VALUES (1,'Makanan','Cari makanan fav mu disini!','makanan'),(2,'Barang Second','Cari barang second disini!','barang-second'),(3,'Pakaian','Cari Pakaian disini','pakaian ');
/*!40000 ALTER TABLE `kategori` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `komentar`
--

DROP TABLE IF EXISTS `komentar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `komentar` (
  `id` int NOT NULL AUTO_INCREMENT,
  `komentar_id` varchar(50) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `komentar` text NOT NULL,
  `rating` tinyint DEFAULT NULL,
  `tipe` enum('diskusi','ulasan') DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `edit` enum('yes','no') DEFAULT 'no',
  `anonymous` enum('yes','no') DEFAULT 'no',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `komentar`
--

LOCK TABLES `komentar` WRITE;
/*!40000 ALTER TABLE `komentar` DISABLE KEYS */;
/*!40000 ALTER TABLE `komentar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pembayaran`
--

DROP TABLE IF EXISTS `pembayaran`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `pembayaran` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama` varchar(50) NOT NULL,
  `logo` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pembayaran`
--

LOCK TABLES `pembayaran` WRITE;
/*!40000 ALTER TABLE `pembayaran` DISABLE KEYS */;
INSERT INTO `pembayaran` VALUES (1,'dana','assets/metode/dana.webp'),(2,'bca','assets/metode/bca.webp'),(3,'qris','assets/metode/qris.webp'),(4,'Shoope Pay','assets/metode/shoopepay.webp');
/*!40000 ALTER TABLE `pembayaran` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `produk`
--

DROP TABLE IF EXISTS `produk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `produk_id` varchar(50) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `slug` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `terjual` int DEFAULT '0',
  `kategori` varchar(100) DEFAULT NULL,
  `rating` float DEFAULT '0',
  `harga` bigint NOT NULL,
  `stok` int DEFAULT '0',
  `deskripsi` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `varian` varchar(255) DEFAULT NULL,
  `diskon` decimal(10,0) DEFAULT '0',
  `status` enum('tersedia','habis','tidak_dijual') CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'tersedia',
  `unit` enum('pcs','kg','liter') DEFAULT NULL,
  `foto` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'default.png',
  `kondisi` enum('baru','bekas') NOT NULL DEFAULT 'baru',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `produk`
--

LOCK TABLES `produk` WRITE;
/*!40000 ALTER TABLE `produk` DISABLE KEYS */;
INSERT INTO `produk` VALUES (1,'P001','Buku Tulis Spiral','tokobuku','buku-tulis-spiral',50,'Alat Tulis',5,15000,100,'Buku tulis spiral ukuran A5, cocok untuk pelajar.','Merah,Biru,Hijau',10,'tersedia','pcs','https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c','baru','2024-11-24 07:39:08','2024-11-24 14:52:08'),(2,'P002','Kaos Polos Putih','lapakkaos','kaos-polos-putih',120,'pakaian',5,45000,200,'Kaos polos putih berbahan katun yang nyaman.','M,L,XL',15,'tersedia','pcs','https://images.unsplash.com/photo-1512436991641-6745cdb1723f','baru','2024-11-24 07:39:08','2024-12-01 18:28:46'),(3,'P003','Mouse Wireless','tokobuku','mouse-wireless',80,'Elektronik',4,75000,50,'Mouse wireless dengan desain ergonomis.','Hitam,Abu-abu',5,'tersedia','pcs','https://images.unsplash.com/photo-1517433456452-f9633a875f6f','baru','2024-11-24 07:39:08','2024-12-24 13:22:58'),(4,'P004','Tas Ransel Kulit','','tas-ransel-kulit',30,'Makanan Enak',5,250000,20,'Tas ransel berbahan kulit asli, cocok untuk aktivitas sehari-hari.','Coklat,Hitam',20,'tersedia','pcs','https://images.unsplash.com/photo-1579517289805-cf068d9666a7','baru','2024-11-24 07:39:08','2024-12-01 18:55:18'),(5,'P005','Lampu LED Meja','','lampu-led-meja',60,'Perlengkapan Rumah',4,95000,70,'Lampu meja LED dengan tingkat pencahayaan yang bisa diatur.','Putih',0,'tersedia','pcs','https://images.unsplash.com/photo-1574169208507-843761448847','baru','2024-11-24 07:39:08',NULL),(6,'P006','Botol Minum Stainless','','botol-minum-stainless',150,'Perlengkapan Outdoor',5,120000,300,'Botol minum stainless steel tahan panas dan dingin.','500ml,750ml',10,'tersedia','liter','https://images.unsplash.com/photo-1519671482749-fd09be7ccebf','baru','2024-11-24 07:39:08',NULL),(7,'P007','Notebook Planner','','notebook-planner',90,'Alat Tulis',5,85000,50,'Notebook planner untuk mencatat rencana harian.','A5',5,'tersedia','pcs','https://images.unsplash.com/photo-1571850553904-7a4b282f1ea0','baru','2024-11-24 07:39:08',NULL),(8,'P008','Sneakers Casual','','sneakers-casual',40,'Sepatu',4,320000,25,'Sneakers kasual untuk aktivitas sehari-hari.','39,40,41,42',25,'tersedia','pcs','https://images.unsplash.com/photo-1517598024396-9f2158a1639a','baru','2024-11-24 07:39:08',NULL),(9,'P009','Smartwatch Fit','','smartwatch-fit',70,'Elektronik',5,450000,35,'Smartwatch dengan fitur olahraga dan kesehatan.','Hitam,Putih',20,'tersedia','pcs','https://images.unsplash.com/photo-1504540982659-79520a58a86d','baru','2024-11-24 07:39:08',NULL),(10,'P010','Meja Lipat Portable','','meja-lipat-portable',25,'Furniture',4,300000,10,'Meja lipat portable, praktis untuk bekerja atau belajar.','Kayu,Coklat',0,'tersedia','pcs','https://images.unsplash.com/photo-1505691723518-36a499c86e73','baru','2024-11-24 07:39:08',NULL),(69,'1','Samsung Galaxy S23','user1','samsung-galaxy-s23',50,'Smartphones',4.8,12000000,15,'Flagship smartphone with advanced features','Color: Phantom Black',5,'tersedia','pcs','https://example.com/images/samsung-galaxy-s23.jpg','baru','2024-12-05 12:26:51',NULL),(70,'2','Apple iPhone 14 Pro','user2','iphone-14-pro',30,'Smartphones',4.9,18000000,10,'Latest iPhone with ProMotion display and triple-camera system','Color: Space Black',10,'tersedia','pcs','https://example.com/images/iphone-14-pro.jpg','baru','2024-12-05 12:26:51',NULL),(71,'3','Sony WH-1000XM5','user3','sony-wh-1000xm5',70,'Headphones',4.7,4500000,25,'Noise-canceling headphones with premium sound quality','Color: Silver',15,'tersedia','pcs','https://example.com/images/sony-wh-1000xm5.jpg','baru','2024-12-05 12:26:51',NULL),(72,'4','LG OLED TV 55C2','user4','lg-oled-tv-55c2',20,'Televisions',4.8,15000000,5,'55-inch OLED TV with 4K resolution and smart features','Size: 55 inches',5,'tersedia','pcs','https://example.com/images/lg-oled-tv-55c2.jpg','baru','2024-12-05 12:26:51',NULL),(73,'5','Adidas Ultraboost 22','user5','adidas-ultraboost-22',100,'Footwear',4.6,2500000,50,'High-performance running shoes with responsive cushioning','Size: 42',20,'tersedia','pcs','https://example.com/images/adidas-ultraboost-22.jpg','baru','2024-12-05 12:26:51',NULL),(74,'6','Dyson V15 Detect','user6','dyson-v15-detect',40,'Vacuum Cleaners',4.9,9500000,8,'Advanced vacuum cleaner with laser dust detection','Color: Yellow',10,'tersedia','pcs','https://example.com/images/dyson-v15-detect.jpg','baru','2024-12-05 12:26:51',NULL),(75,'7','NVIDIA GeForce RTX 4090','user7','nvidia-geforce-rtx-4090',15,'Graphics Cards',4.9,30000000,2,'High-end GPU for gaming and creative professionals','VRAM: 24GB',0,'tersedia','pcs','https://example.com/images/nvidia-geforce-rtx-4090.jpg','baru','2024-12-05 12:26:51',NULL),(76,'8','ASUS ROG Zephyrus G14','user8','asus-rog-zephyrus-g14',25,'Laptops',4.8,20000000,10,'Compact gaming laptop with powerful performance','RAM: 16GB',5,'tersedia','pcs','https://example.com/images/asus-rog-zephyrus-g14.jpg','baru','2024-12-05 12:26:51',NULL),(77,'9','Rolex Submariner','user9','rolex-submariner',5,'Watches',5,150000000,1,'Luxury dive watch with iconic design','Material: Stainless Steel',0,'tersedia','pcs','https://example.com/images/rolex-submariner.jpg','baru','2024-12-05 12:26:51',NULL),(78,'10','Canon EOS R5','user10','canon-eos-r5',12,'Cameras',4.9,60000000,3,'Professional mirrorless camera with 8K video recording','Lens: 24-105mm',5,'tersedia','pcs','https://example.com/images/canon-eos-r5.jpg','baru','2024-12-05 12:26:51',NULL),(86,'P223445','alibi','tokobuku','alibi-P223445',0,'test',0,123,0,'asdasd','asdas',0,'tersedia','pcs','ZBZDPRKNZC.jpg','baru','2024-12-27 02:32:58','2024-12-27 09:32:58'),(87,'P511989','alibi','tokobuku','alibi-P511989',0,'test',0,123,0,'asdasd','asdas',0,'tersedia','pcs','public/img/product/GUHUUHPNWO.jpg','baru','2024-12-27 02:34:35','2024-12-27 09:34:35'),(88,'P218164','alibi','tokobuku','alibi-P218164',0,'test',0,123,0,'asdasd','asdas',0,'tersedia','pcs','THEDZINEUH.jpg','baru','2024-12-27 02:34:43','2024-12-27 09:34:43'),(89,'P997642','alibi','tokobuku','alibi-P997642',0,'test',0,123,0,'asdasd','asdas',0,'habis','pcs','UYPTISQYKN.jpg','baru','2024-12-27 02:35:15','2024-12-27 09:46:48');
/*!40000 ALTER TABLE `produk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `produk_statistik`
--

DROP TABLE IF EXISTS `produk_statistik`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `produk_statistik` (
  `id` int NOT NULL AUTO_INCREMENT,
  `produk_id` varchar(50) NOT NULL,
  `dilihat` int DEFAULT '0',
  `disukai` int DEFAULT '0',
  `tanggal` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `produk_statistik`
--

LOCK TABLES `produk_statistik` WRITE;
/*!40000 ALTER TABLE `produk_statistik` DISABLE KEYS */;
/*!40000 ALTER TABLE `produk_statistik` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `replies`
--

DROP TABLE IF EXISTS `replies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `replies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `komentar_id` varchar(50) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `komentar` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `nama_toko` varchar(100) DEFAULT NULL,
  `edit` enum('yes','no') DEFAULT 'no',
  `anonymous` enum('yes','no') DEFAULT 'no',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `replies`
--

LOCK TABLES `replies` WRITE;
/*!40000 ALTER TABLE `replies` DISABLE KEYS */;
/*!40000 ALTER TABLE `replies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `settings_web`
--

DROP TABLE IF EXISTS `settings_web`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `settings_web` (
  `id` int NOT NULL AUTO_INCREMENT,
  `web_title` varchar(255) NOT NULL,
  `web_icon` varchar(255) NOT NULL,
  `web_logo` varchar(255) NOT NULL,
  `web_author` varchar(255) NOT NULL,
  `web_keywords` varchar(255) NOT NULL,
  `web_description` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `settings_web`
--

LOCK TABLES `settings_web` WRITE;
/*!40000 ALTER TABLE `settings_web` DISABLE KEYS */;
INSERT INTO `settings_web` VALUES (1,'Lapak Siswa','favicon.ico','assets/logo/logo.png','Muhammad Alfaridzi','lapaksiswa, jual beli butun','Lapak Siswa adalah website jual beli butun');
/*!40000 ALTER TABLE `settings_web` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `toko`
--

DROP TABLE IF EXISTS `toko`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `toko` (
  `id` int NOT NULL AUTO_INCREMENT,
  `domain` varchar(100) NOT NULL,
  `username` varchar(50) NOT NULL,
  `nama` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `kategori` varchar(100) DEFAULT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `deskripsi` text,
  `email` varchar(100) DEFAULT NULL,
  `no_hp` varchar(15) DEFAULT NULL,
  `alamat` text,
  `rating` tinyint DEFAULT NULL,
  `status` enum('aktif','nonaktif') CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'aktif',
  `saldo` bigint NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `toko`
--

LOCK TABLES `toko` WRITE;
/*!40000 ALTER TABLE `toko` DISABLE KEYS */;
INSERT INTO `toko` VALUES (1,'tokobuku','admin','Toko Buku Butun','Alat Tulis','https://i.pinimg.com/736x/80/c0/f6/80c0f658ab849ed80a74b3539ed72d58.jpg','Toko Buku Terpecaya','tokobuku@gmail.com','08589411','Jln Butun, Kota Bekasi',3,'aktif',20000,'2024-12-17 13:13:37','2024-12-24 07:53:44');
/*!40000 ALTER TABLE `toko` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaksi`
--

DROP TABLE IF EXISTS `transaksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `transaksi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `produk_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `domain` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `alamat` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `qty` bigint NOT NULL,
  `status` enum('waiting','proses','sukses') CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'waiting',
  `harga` bigint DEFAULT NULL,
  `metode` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `note` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `no_hp` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaksi`
--

LOCK TABLES `transaksi` WRITE;
/*!40000 ALTER TABLE `transaksi` DISABLE KEYS */;
INSERT INTO `transaksi` VALUES (1,'M123','P001','tokobuku','adminmarket','Jln Rawa Bango',10,'sukses',12000,'DANA','Ga tau','085894110892','2024-12-23 17:02:46','2024-12-23 17:08:16');
/*!40000 ALTER TABLE `transaksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `no_hp` bigint DEFAULT NULL,
  `role` enum('admin','user','seller') DEFAULT 'user',
  `foto` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'default.png',
  `alamat` text,
  `kelas` varchar(50) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `token_reset_password` varchar(255) DEFAULT NULL,
  `jenis_kelamin` enum('pria','wanita') NOT NULL DEFAULT 'pria',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `last_online` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'adminmarket',NULL,NULL,'aalalfaridzi9@gmail.com',NULL,'user',NULL,NULL,NULL,'alfaridzi123',NULL,'pria','2024-11-26 07:31:59',NULL,'2024-11-26 07:37:04'),(2,'alfaridzi','Muhammad Alfaridzi',NULL,'nfex.ghost@gmail.com',85894110892,'user','test123.png',NULL,NULL,'123',NULL,'pria','2024-11-26 08:02:48',NULL,'2024-12-15 16:50:06'),(3,'admin','admin market 123','2024-12-05','admin@gmail.com',85894110982,'seller','NFFHQMDWOJ.jpg','','','202cb962ac59075b964b07152d234b70','','pria','2024-12-15 16:50:51','2024-12-06 17:42:12','2024-12-22 17:42:56');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login`
--

DROP TABLE IF EXISTS `user_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `user_login` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) NOT NULL,
  `ip_address` varchar(50) DEFAULT NULL,
  `device` varchar(100) DEFAULT NULL,
  `waktu_login` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login`
--

LOCK TABLES `user_login` WRITE;
/*!40000 ALTER TABLE `user_login` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `voucher`
--

DROP TABLE IF EXISTS `voucher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `voucher` (
  `id` int NOT NULL AUTO_INCREMENT,
  `toko_id` varchar(50) NOT NULL,
  `kode` varchar(50) DEFAULT NULL,
  `potongan` decimal(5,2) DEFAULT NULL,
  `stock` int DEFAULT '0',
  `max_potongan` decimal(5,2) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `voucher`
--

LOCK TABLES `voucher` WRITE;
/*!40000 ALTER TABLE `voucher` DISABLE KEYS */;
/*!40000 ALTER TABLE `voucher` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-12-27 19:00:29
