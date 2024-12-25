-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: lapaksiswa
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ci_sessions` (
  `id` varchar(128) NOT NULL,
  `ip_address` varchar(45) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `data` blob NOT NULL,
  KEY `ci_sessions_timestamp` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `foto_komentar` (
  `id` int NOT NULL AUTO_INCREMENT,
  `komentar_id` varchar(50) NOT NULL,
  `foto` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `foto_produk` (
  `id` int NOT NULL AUTO_INCREMENT,
  `foto` varchar(255) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kategori` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kategori` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `slug` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
