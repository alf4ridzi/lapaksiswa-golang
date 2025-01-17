-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 09, 2025 at 03:10 AM
-- Server version: 8.0.31
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lapaksiswa`
--
CREATE DATABASE IF NOT EXISTS `lapaksiswa` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `lapaksiswa`;

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int NOT NULL,
  `produk_id` text COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `cart`
--

INSERT INTO `cart` (`id`, `produk_id`, `username`, `created_at`) VALUES
(5, 'P129428', 'hanma', '2025-01-08 17:25:38'),
(6, 'P129428', 'hanma', '2025-01-09 07:25:09'),
(7, 'P085217', 'hanma', '2025-01-09 08:19:05'),
(8, 'P001', 'hanma', '2025-01-09 09:08:58');

-- --------------------------------------------------------

--
-- Table structure for table `checkout`
--

DROP TABLE IF EXISTS `checkout`;
CREATE TABLE `checkout` (
  `id` int NOT NULL,
  `produk_id` text NOT NULL,
  `checkout` text NOT NULL,
  `total` bigint NOT NULL,
  `username` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `checkout`
--

INSERT INTO `checkout` (`id`, `produk_id`, `checkout`, `total`, `username`) VALUES
(1, 'P085217', 'SA99892DF4US10M3', 30000, 'hanma'),
(2, 'P085217', 'QFP83YDUM1E0CZRE', 30000, 'hanma'),
(3, 'P085217', 'YCREUW4X4I6D92ES', 30000, 'hanma');

-- --------------------------------------------------------

--
-- Table structure for table `foto_komentar`
--

DROP TABLE IF EXISTS `foto_komentar`;
CREATE TABLE `foto_komentar` (
  `id` int NOT NULL,
  `komentar_id` varchar(50) NOT NULL,
  `foto` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `foto_produk`
--

DROP TABLE IF EXISTS `foto_produk`;
CREATE TABLE `foto_produk` (
  `id` int NOT NULL,
  `foto` varchar(255) NOT NULL,
  `produk_id` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `invoice`
--

DROP TABLE IF EXISTS `invoice`;
CREATE TABLE `invoice` (
  `id` int NOT NULL,
  `order_id` text COLLATE utf8mb4_general_ci NOT NULL,
  `metode` text COLLATE utf8mb4_general_ci NOT NULL,
  `total` bigint NOT NULL,
  `status` enum('waiting','sukses','gagal','proses') COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `kategori`
--

DROP TABLE IF EXISTS `kategori`;
CREATE TABLE `kategori` (
  `id` int NOT NULL,
  `kategori` varchar(100) DEFAULT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `slug` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `kategori`
--

INSERT INTO `kategori` (`id`, `kategori`, `deskripsi`, `slug`) VALUES
(1, 'Makanan', 'Cari makanan fav mu disini!', 'makanan'),
(2, 'Barang Second', 'Cari barang second disini!', 'barang-second'),
(3, 'Pakaian', 'Cari Pakaian disini', 'pakaian ');

-- --------------------------------------------------------

--
-- Table structure for table `komentar`
--

DROP TABLE IF EXISTS `komentar`;
CREATE TABLE `komentar` (
  `id` int NOT NULL,
  `komentar_id` varchar(50) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `komentar` text NOT NULL,
  `rating` tinyint DEFAULT NULL,
  `tipe` enum('diskusi','ulasan') DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `edit` enum('yes','no') DEFAULT 'no',
  `anonymous` enum('yes','no') DEFAULT 'no'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `pembayaran`
--

DROP TABLE IF EXISTS `pembayaran`;
CREATE TABLE `pembayaran` (
  `id` int NOT NULL,
  `nama` varchar(50) NOT NULL,
  `metode` varchar(255) NOT NULL,
  `logo` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `pembayaran`
--

INSERT INTO `pembayaran` (`id`, `nama`, `metode`, `logo`) VALUES
(1, 'Dana', 'dana', 'dana.webp'),
(2, 'BCA', 'bca', 'bca.webp'),
(3, 'Qris', 'qris', 'qris.webp'),
(4, 'Shoope Pay', 'shoopepay', 'shoopepay.webp');

-- --------------------------------------------------------

--
-- Table structure for table `produk`
--

DROP TABLE IF EXISTS `produk`;
CREATE TABLE `produk` (
  `id` int NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `domain` varchar(255) NOT NULL,
  `slug` varchar(100) NOT NULL,
  `terjual` int DEFAULT '0',
  `kategori` varchar(100) DEFAULT NULL,
  `rating` float DEFAULT '0',
  `harga` bigint NOT NULL,
  `stok` int DEFAULT '0',
  `deskripsi` varchar(255) DEFAULT NULL,
  `varian` varchar(255) DEFAULT NULL,
  `diskon` decimal(10,0) DEFAULT '0',
  `status` enum('tersedia','habis','nonaktif') DEFAULT 'tersedia',
  `unit` enum('pcs','kg','liter') DEFAULT NULL,
  `foto` varchar(255) NOT NULL DEFAULT 'default.png',
  `kondisi` enum('baru','bekas') NOT NULL DEFAULT 'baru',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `produk`
--

INSERT INTO `produk` (`id`, `produk_id`, `nama`, `domain`, `slug`, `terjual`, `kategori`, `rating`, `harga`, `stok`, `deskripsi`, `varian`, `diskon`, `status`, `unit`, `foto`, `kondisi`, `created_at`, `updated_at`) VALUES
(1, 'P001', 'Buku Tulis Spiral', 'tokobuku', 'buku-tulis-spiral', 50, 'Alat Tulis', 5, 15000, 100, 'Buku tulis spiral ukuran A5, cocok untuk pelajar.', 'Merah,Biru,Hijau', 10, 'tersedia', 'pcs', 'buku-tulis.jpeg', 'baru', '2024-11-24 07:39:08', '2025-01-08 14:44:11'),
(2, 'P002', 'Kaos Polos Putih', 'lapakkaos', 'kaos-polos-putih', 120, 'pakaian', 5, 45000, 200, 'Kaos polos putih berbahan katun yang nyaman.', 'M,L,XL', 15, 'tersedia', 'pcs', 'kaos-putih.jpeg', 'baru', '2024-11-24 07:39:08', '2025-01-08 14:49:59'),
(3, 'P003', 'Mouse Wireless', 'tokobuku', 'mouse-wireless', 80, 'Elektronik', 4, 75000, 50, 'Mouse wireless dengan desain ergonomis.', 'Hitam,Abu-abu', 5, 'tersedia', 'pcs', 'mouse.jpeg', 'baru', '2024-11-24 07:39:08', '2025-01-08 14:43:16'),
(86, 'P223445', 'alibi', 'tokobuku', 'alibi-P223445', 0, 'test', 0, 123, 0, 'asdasd', 'asdas', 0, 'tersedia', 'pcs', 'ZBZDPRKNZC.jpg', 'baru', '2024-12-27 02:32:58', '2024-12-27 09:32:58'),
(88, 'P218164', 'alibi', 'tokobuku', 'alibi-P218164', 0, 'test', 0, 123, 0, 'asdasd', 'asdas', 0, 'tersedia', 'pcs', 'THEDZINEUH.jpg', 'baru', '2024-12-27 02:34:43', '2024-12-27 09:34:43'),
(89, 'P997642', 'alibi', 'tokobuku', 'alibi-P997642', 0, 'test', 0, 123, 0, 'asdasd', 'asdas', 0, 'habis', 'pcs', 'UYPTISQYKN.jpg', 'baru', '2024-12-27 02:35:15', '2024-12-27 09:46:48'),
(90, 'P022627', 'hello', 'tokokoko123', 'hello-P022627', 0, 'test', 0, 1200, 0, 'Test aja', 'L', 0, 'tersedia', 'pcs', 'WUADWPEICV.jpeg', 'baru', '2024-12-28 17:33:14', '2024-12-29 00:33:14'),
(93, 'P115032', 'hello', 'tokoapa', 'hello-P115032', 0, 'test', 0, 123, 0, 'Test produk', 'XL', 0, 'tersedia', 'pcs', 'AOHDEJJYBE.jpg', 'baru', '2024-12-30 05:01:16', '2024-12-30 12:01:16'),
(94, 'P085217', 'Bakso Solo', 'tokoapa', 'bakso-solo-P085217', 0, 'test', 0, 10000, 10, 'Bakso Solo Wuenak cik langsung order', 'XL', 0, 'tersedia', 'pcs', 'CSAIAANYVG.jpeg', 'baru', '2024-12-30 09:54:36', '2024-12-30 16:54:36'),
(95, 'P795098', 'Loh', 'tokoapa', 'loh-P795098', 0, 'pakaian ', 0, 1200, 1, 'Loh test', 'XL', 0, 'tersedia', 'pcs', 'YJMPNZXNJZ.png', 'baru', '2024-12-30 10:32:59', '2024-12-30 17:32:59'),
(98, 'P129428', 'Motor CBR Bagus', 'tokokoko', 'motor-cbr-bagus-P129428', 0, 'barang-second', 0, 1200, 10, 'Motor CBR', 'XL', 0, 'tersedia', 'pcs', 'KCEZFEWPWN.jpg', 'bekas', '2025-01-06 02:22:59', '2025-01-08 14:51:22');

-- --------------------------------------------------------

--
-- Table structure for table `produk_statistik`
--

DROP TABLE IF EXISTS `produk_statistik`;
CREATE TABLE `produk_statistik` (
  `id` int NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `dilihat` int DEFAULT '0',
  `disukai` int DEFAULT '0',
  `tanggal` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `replies`
--

DROP TABLE IF EXISTS `replies`;
CREATE TABLE `replies` (
  `id` int NOT NULL,
  `komentar_id` varchar(50) NOT NULL,
  `produk_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `komentar` text NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `nama_toko` varchar(100) DEFAULT NULL,
  `edit` enum('yes','no') DEFAULT 'no',
  `anonymous` enum('yes','no') DEFAULT 'no',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `settings_web`
--

DROP TABLE IF EXISTS `settings_web`;
CREATE TABLE `settings_web` (
  `id` int NOT NULL,
  `web_title` varchar(255) NOT NULL,
  `web_icon` varchar(255) NOT NULL,
  `web_logo` varchar(255) NOT NULL,
  `web_author` varchar(255) NOT NULL,
  `web_keywords` varchar(255) NOT NULL,
  `web_description` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `settings_web`
--

INSERT INTO `settings_web` (`id`, `web_title`, `web_icon`, `web_logo`, `web_author`, `web_keywords`, `web_description`) VALUES
(1, 'Lapak Siswa', 'favicon.ico', 'logo.png', 'Muhammad Alfaridzi', 'lapaksiswa, jual beli butun', 'Lapak Siswa adalah website jual beli butun');

-- --------------------------------------------------------

--
-- Table structure for table `toko`
--

DROP TABLE IF EXISTS `toko`;
CREATE TABLE `toko` (
  `id` int NOT NULL,
  `domain` varchar(100) NOT NULL,
  `username` varchar(50) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `kategori` varchar(100) DEFAULT NULL,
  `logo` varchar(255) DEFAULT 'default.png',
  `deskripsi` text,
  `email` varchar(100) DEFAULT NULL,
  `no_hp` varchar(15) DEFAULT NULL,
  `alamat` text,
  `rating` tinyint DEFAULT '0',
  `status` enum('aktif','nonaktif') DEFAULT 'aktif',
  `saldo` bigint NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `toko`
--

INSERT INTO `toko` (`id`, `domain`, `username`, `nama`, `kategori`, `logo`, `deskripsi`, `email`, `no_hp`, `alamat`, `rating`, `status`, `saldo`, `created_at`, `updated_at`) VALUES
(1, 'tokobuku', 'admin', 'Toko Buku Butun', 'Alat Tulis', 'https://i.pinimg.com/736x/80/c0/f6/80c0f658ab849ed80a74b3539ed72d58.jpg', 'Toko Buku Terpecaya', 'tokobuku@gmail.com', '08589411', 'Jln Butun, Kota Bekasi', 3, 'aktif', 20000, '2024-12-17 13:13:37', '2024-12-24 07:53:44'),
(8, 'tokokoko', 'hanma', 'Toko Motor Jalil 2', 'makanan', 'S9JE0KBOJE.jpeg', 'Yeah whatsapp gang im nigga', 'jalil2@gmail.com', '81234567890', 'Jln koko surabaya', 0, 'aktif', 0, '2024-12-29 14:14:42', '2025-01-08 14:50:53'),
(9, 'tokoapa', 'alfa', 'Toko Apa tau', NULL, 'default.png', 'Jual beli apa aja', 'toko@gmail.com', '08123456789', 'Butun ', 0, 'aktif', 0, '2024-12-30 12:00:48', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transaksi`
--

DROP TABLE IF EXISTS `transaksi`;
CREATE TABLE `transaksi` (
  `id` int NOT NULL,
  `order_id` varchar(50) NOT NULL,
  `produk_id` varchar(100) DEFAULT NULL,
  `domain` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `alamat` varchar(250) DEFAULT NULL,
  `qty` bigint NOT NULL,
  `status` enum('waiting','proses','sukses','gagal') DEFAULT 'waiting',
  `harga` bigint DEFAULT NULL,
  `metode` varchar(50) DEFAULT NULL,
  `note` text,
  `no_hp` varchar(15) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `transaksi`
--

INSERT INTO `transaksi` (`id`, `order_id`, `produk_id`, `domain`, `username`, `alamat`, `qty`, `status`, `harga`, `metode`, `note`, `no_hp`, `created_at`, `updated_at`) VALUES
(1, 'M123', 'P001', 'tokobuku', 'adminmarket', 'Jln Rawa Bango', 10, 'sukses', 12000, 'dana', 'Ga tau', '085894110892', '2024-12-23 17:02:46', '2025-01-08 20:54:38');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL,
  `username` varchar(50) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `no_hp` bigint DEFAULT NULL,
  `role` enum('admin','user','seller') DEFAULT 'user',
  `foto` varchar(255) DEFAULT 'default.png',
  `alamat` text,
  `kelas` varchar(50) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `token_reset_password` varchar(255) DEFAULT NULL,
  `jenis_kelamin` enum('pria','wanita') NOT NULL DEFAULT 'pria',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `last_online` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `nama`, `tanggal_lahir`, `email`, `no_hp`, `role`, `foto`, `alamat`, `kelas`, `password`, `token_reset_password`, `jenis_kelamin`, `created_at`, `last_online`, `updated_at`) VALUES
(1, 'adminmarket', NULL, NULL, 'aalalfaridzi9@gmail.com', NULL, 'user', NULL, NULL, NULL, 'alfaridzi123', NULL, 'pria', '2024-11-26 07:31:59', NULL, '2024-11-26 07:37:04'),
(2, 'alfaridzi', 'Muhammad Alfaridzi', NULL, 'nfex.ghost@gmail.com', 85894110892, 'user', 'test123.png', NULL, NULL, '123', NULL, 'pria', '2024-11-26 08:02:48', NULL, '2024-12-15 16:50:06'),
(3, 'admin', 'admin market 123', '2024-12-05', 'admin@gmail.com', 85894110982, 'seller', 'NFFHQMDWOJ.jpg', '', '', '202cb962ac59075b964b07152d234b70', '', 'pria', '2024-12-15 16:50:51', '2024-12-06 17:42:12', '2024-12-22 17:42:56'),
(4, 'hanma', 'My Problem Test', '2004-01-22', 'hanma@gmail.com', 858324532434, 'seller', 'UFXFTQPCAZ.jpeg', NULL, NULL, '4ea07e448964eea5b6dd649029764848', NULL, 'pria', '2024-12-28 13:10:55', '2024-12-28 13:10:55', '2025-01-06 10:07:15'),
(5, 'alfa', 'Muhammad Alfaridzi', '2024-12-18', 'alfa123@gmail.com', 812345678, 'seller', 'RQNHGETQOU.jpeg', NULL, NULL, '5163aece7da6f7d32627682a27062ca5', NULL, 'pria', '2024-12-30 11:59:34', '2024-12-30 11:59:34', '2024-12-30 12:00:48');

-- --------------------------------------------------------

--
-- Table structure for table `user_login`
--

DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login` (
  `id` int NOT NULL,
  `user_id` varchar(50) NOT NULL,
  `ip_address` varchar(50) DEFAULT NULL,
  `device` varchar(100) DEFAULT NULL,
  `waktu_login` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Table structure for table `voucher`
--

DROP TABLE IF EXISTS `voucher`;
CREATE TABLE `voucher` (
  `id` int NOT NULL,
  `toko_id` varchar(50) NOT NULL,
  `kode` varchar(50) DEFAULT NULL,
  `potongan` decimal(5,2) DEFAULT NULL,
  `stock` int DEFAULT '0',
  `max_potongan` decimal(5,2) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `checkout`
--
ALTER TABLE `checkout`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `foto_komentar`
--
ALTER TABLE `foto_komentar`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `foto_produk`
--
ALTER TABLE `foto_produk`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `invoice`
--
ALTER TABLE `invoice`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `komentar`
--
ALTER TABLE `komentar`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pembayaran`
--
ALTER TABLE `pembayaran`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `produk`
--
ALTER TABLE `produk`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `produk_statistik`
--
ALTER TABLE `produk_statistik`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `replies`
--
ALTER TABLE `replies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `settings_web`
--
ALTER TABLE `settings_web`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `toko`
--
ALTER TABLE `toko`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_login`
--
ALTER TABLE `user_login`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `voucher`
--
ALTER TABLE `voucher`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cart`
--
ALTER TABLE `cart`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `checkout`
--
ALTER TABLE `checkout`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `foto_komentar`
--
ALTER TABLE `foto_komentar`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `foto_produk`
--
ALTER TABLE `foto_produk`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `invoice`
--
ALTER TABLE `invoice`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `kategori`
--
ALTER TABLE `kategori`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `komentar`
--
ALTER TABLE `komentar`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pembayaran`
--
ALTER TABLE `pembayaran`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `produk`
--
ALTER TABLE `produk`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=99;

--
-- AUTO_INCREMENT for table `produk_statistik`
--
ALTER TABLE `produk_statistik`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `replies`
--
ALTER TABLE `replies`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `settings_web`
--
ALTER TABLE `settings_web`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `toko`
--
ALTER TABLE `toko`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user_login`
--
ALTER TABLE `user_login`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `voucher`
--
ALTER TABLE `voucher`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
