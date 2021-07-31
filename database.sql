-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 31, 2021 at 06:18 AM
-- Server version: 10.4.10-MariaDB
-- PHP Version: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crud_golang`
--

-- --------------------------------------------------------

--
-- Table structure for table `peoples`
--

CREATE TABLE `peoples` (
  `nik` varchar(255) NOT NULL,
  `name` varchar(150) NOT NULL,
  `birthday` date NOT NULL,
  `gender` enum('man','woman') NOT NULL DEFAULT 'man',
  `address` text NOT NULL,
  `photo` varchar(255) NOT NULL,
  `rt` varchar(20) NOT NULL,
  `rw` varchar(20) NOT NULL,
  `subdistrict` varchar(150) NOT NULL,
  `district` varchar(150) NOT NULL,
  `city` varchar(150) NOT NULL,
  `province` varchar(150) NOT NULL,
  `age` int(11) NOT NULL,
  `date_added` datetime NOT NULL DEFAULT current_timestamp(),
  `date_updated` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `peoples`
--

INSERT INTO `peoples` (`nik`, `name`, `birthday`, `gender`, `address`, `photo`, `rt`, `rw`, `subdistrict`, `district`, `city`, `province`, `age`, `date_added`, `date_updated`) VALUES
('1234567', 'Arief', '2001-07-14', 'man', 'JL. Kamboja 1', '', '04', '03', 'Nusa Jaya', 'Karawaci', 'Tangerang', 'Banten', 20, '2021-07-31 09:32:35', '2021-07-31 09:32:35');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(150) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('admin','user') NOT NULL DEFAULT 'user',
  `date_added` datetime DEFAULT current_timestamp(),
  `date_modified` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `username`, `password`, `role`, `date_added`, `date_modified`) VALUES
(1, 'admin', 'admin', '$2a$10$XsCPRKToKhbpanT12CtMye5PkRjvBiKV/L9JJp37o228n3k8ELLdG', 'admin', '2021-07-30 16:55:34', '2021-07-30 16:55:46');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `peoples`
--
ALTER TABLE `peoples`
  ADD PRIMARY KEY (`nik`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
