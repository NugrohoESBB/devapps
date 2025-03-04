-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 04 Mar 2025 pada 02.57
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_apps`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `logs`
--

CREATE TABLE `logs` (
  `id` int(11) NOT NULL,
  `d` date DEFAULT current_timestamp(),
  `t` time DEFAULT current_timestamp(),
  `n` varchar(50) NOT NULL,
  `l` varchar(50) NOT NULL,
  `k` double NOT NULL,
  `i` double NOT NULL,
  `f` double NOT NULL,
  `a` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `logs`
--

INSERT INTO `logs` (`id`, `d`, `t`, `n`, `l`, `k`, `i`, `f`, `a`) VALUES
(1, '2025-01-08', '23:53:42', 'Beverlie', 'Tobii', 36.93, 63.88, 55.96, 90.9),
(2, '2024-11-06', '04:37:22', 'Alessandra', 'Clemmie', 11.72, 72.97, 80.61, 79.07),
(3, '2024-02-11', '18:56:08', 'Noelani', 'Abbie', 21.83, 64.32, 9.39, 40.16),
(4, '2024-11-05', '05:56:02', 'Fidelia', 'Phillipe', 69.19, 38.51, 2.92, 1.98),
(5, '2024-06-22', '08:19:36', 'Vincenty', 'Ashli', 93.15, 54.05, 68.32, 85.51),
(6, '2024-11-23', '01:13:02', 'Christoforo', 'Carmelle', 97.58, 77.94, 97.89, 72.65),
(7, '2024-09-29', '17:02:19', 'Janina', 'Melloney', 91.68, 99.7, 78.27, 50.03),
(8, '2024-07-25', '13:05:28', 'Eran', 'Vivian', 49.24, 34.71, 63.93, 72.29),
(9, '2024-04-07', '03:15:27', 'Bobbi', 'Jean', 66.45, 65.53, 84.02, 53),
(10, '2024-04-09', '21:50:53', 'Miran', 'Erminie', 26.43, 94.53, 77.13, 29.04);

-- --------------------------------------------------------

--
-- Struktur dari tabel `logsessions`
--

CREATE TABLE `logsessions` (
  `id` int(11) NOT NULL,
  `d` date DEFAULT current_timestamp(),
  `t` time DEFAULT current_timestamp(),
  `tn` varchar(255) NOT NULL,
  `s` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `logsessions`
--

INSERT INTO `logsessions` (`id`, `d`, `t`, `tn`, `s`) VALUES
(1, '2025-02-15', '13:57:01', 'session-96-1739948221', 'Login'),
(2, '2025-02-16', '14:02:25', 'session-94-1739948545', 'Login'),
(3, '2025-02-16', '14:11:13', 'session-96-1739949073', 'Login'),
(4, '2025-02-17', '14:11:34', 'session-96-1739949092', 'Login'),
(5, '2025-02-18', '14:12:32', 'session-96-1739949152', 'Login'),
(6, '2025-02-19', '14:12:34', 'session-96-1739949154', 'Login'),
(7, '2025-02-19', '14:12:34', 'session-96-1739949154', 'Login'),
(8, '2025-02-19', '14:22:09', 'session-96-1739949729', 'Login'),
(9, '2025-02-19', '14:23:52', 'session-96-1739949831', 'Login'),
(10, '2025-02-20', '09:02:11', 'session-96-1740016931', 'Login'),
(11, '2025-02-20', '09:08:00', 'session-96-1740017280', 'Login'),
(12, '2025-02-20', '09:13:24', 'session-96-1740017604', 'Login'),
(13, '2025-02-20', '09:14:41', 'session-96-1740017680', 'Login'),
(14, '2025-02-20', '09:20:33', 'session-96-1740018032', 'Login'),
(15, '2025-02-20', '09:45:22', 'session-96-1740019521', 'Login'),
(16, '2025-02-20', '10:28:49', 'session-96-1740022129', 'Login'),
(17, '2025-02-20', '10:33:12', 'session-96-1740022392', 'Login'),
(18, '2025-02-20', '11:32:11', 'session-96-1740025930', 'Login'),
(19, '2025-02-20', '13:23:53', 'session-96-1740032633', 'Login'),
(21, '2025-02-20', '13:26:13', 'session-96-1740032773', 'Login'),
(23, '2025-02-20', '13:30:18', 'session-96-1740033017', 'Login'),
(24, '2025-02-20', '13:37:36', 'session-96-1740033456', 'Login'),
(25, '2025-02-20', '13:40:10', 'session-96-1740033610', 'Login'),
(26, '2025-02-20', '13:41:44', 'session-96-1740033702', 'Login'),
(27, '2025-02-20', '13:47:36', 'session-96-1740034056', 'Login'),
(28, '2025-02-20', '14:00:19', 'session-94-1740034819', 'Login'),
(29, '2025-02-20', '14:00:48', 'session-96-1740034846', 'Login'),
(30, '2025-02-20', '14:25:48', 'session-96-1740036348', 'Login'),
(31, '2025-02-21', '10:17:22', 'session-96-1740107842', 'Login'),
(32, '2025-02-21', '10:42:11', 'session-96-1740109331', 'Login'),
(33, '2025-02-27', '10:03:59', 'session-96-1740625439', 'Login'),
(34, '2025-02-27', '10:20:02', 'session-94-1740626402', 'Login'),
(35, '2025-02-27', '11:03:43', 'session-94-1740629023', 'Login'),
(36, '2025-02-27', '12:50:34', 'session-94-1740635434', 'Login'),
(37, '2025-02-27', '13:44:31', 'session-96-1740638671', 'Login'),
(38, '2025-02-27', '13:45:17', 'session-94-1740638715', 'Login'),
(39, '2025-02-27', '13:45:38', 'session-96-1740638738', 'Login'),
(40, '2025-02-27', '15:28:07', 'session-96-1740644887', 'Login'),
(41, '2025-02-28', '08:21:39', 'session-96-1740705699', 'Login'),
(42, '2025-02-28', '08:52:42', 'session-96-1740707561', 'Login'),
(43, '2025-02-28', '09:03:17', 'session-96-1740708196', 'Login'),
(44, '2025-02-28', '09:04:36', 'session-96-1740708276', 'Login'),
(45, '2025-03-03', '08:46:20', 'session-96-1740966380', 'Login'),
(46, '2025-03-03', '09:22:34', 'session-96-1740968554', 'Login'),
(47, '2025-03-03', '09:32:03', 'session-96-1740969123', 'Login'),
(48, '2025-03-03', '09:44:07', 'session-96-1740969847', 'Login'),
(49, '2025-03-04', '08:54:49', 'session-96-1741053288', 'Login');

-- --------------------------------------------------------

--
-- Struktur dari tabel `logtasks`
--

CREATE TABLE `logtasks` (
  `id` int(11) NOT NULL,
  `d` date NOT NULL DEFAULT current_timestamp(),
  `t` time NOT NULL DEFAULT current_timestamp(),
  `r` varchar(50) NOT NULL,
  `dc` varchar(255) NOT NULL,
  `rt` varchar(50) NOT NULL,
  `s` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `logtasks`
--

INSERT INTO `logtasks` (`id`, `d`, `t`, `r`, `dc`, `rt`, `s`) VALUES
(14, '2025-02-27', '10:06:54', 'admin', 'Tolong buatkan notifications API untuk setiap role user, setiap account akan menerima notifications yang berbeda beda tergantung rolenya. Terima kasih', '/logsessions', 'pending'),
(20, '2025-02-27', '12:49:34', 'user', 'Tolong ini notification untuk user', '/logs', 'done'),
(29, '2025-03-03', '13:07:30', 'user', 'alala', '/informationLog', 'done'),
(30, '2025-03-03', '15:38:32', 'user', 'test', '/logsessions', 'done'),
(31, '2025-03-03', '15:39:33', 'user', 'akaka', '/invoice', 'done');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sessions`
--

CREATE TABLE `sessions` (
  `id` int(11) NOT NULL,
  `u_id` int(11) NOT NULL,
  `r` varchar(50) NOT NULL,
  `tn` varchar(255) NOT NULL,
  `t` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `sessions`
--

INSERT INTO `sessions` (`id`, `u_id`, `r`, `tn`, `t`) VALUES
(43, 96, 'admin', 'session-96-1741053288', '2025-03-04 01:54:48');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `d` date DEFAULT current_timestamp(),
  `t` time DEFAULT current_timestamp(),
  `n` varchar(50) NOT NULL,
  `r` varchar(50) NOT NULL,
  `e` varchar(50) NOT NULL,
  `lt` varchar(50) NOT NULL,
  `ln` varchar(50) NOT NULL,
  `p` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `d`, `t`, `n`, `r`, `e`, `lt`, `ln`, `p`) VALUES
(94, '2025-02-17', '13:36:37', 'mur', 'user', 'mur@gmail.com', '-3.238190', '116.227859', '$2a$10$4lrwCvSzgZ9ZfzYhzXx7yOCcQgnVK0OFagHXhGDn0Zrj/JN6hxn.O'),
(95, '2025-02-17', '13:37:41', 'kiki', 'user', 'kiki@gmail.com', '1.725922', '128.010010', '$2a$10$I8iVnn37Msahm6rpmIllCu/tcC1XILkpofogKGUVaW/0wZ4t2VgX2'),
(96, '2025-02-17', '13:38:17', 'ade sayangg', 'admin', 'ade@gmail.com', '-4.429552', '102.893082', '$2a$10$tqS9W4oN28CakxnybII6u.VpcbWN/aWjS38ydJAoPJUrW8Dp9DvfC'),
(97, '2025-02-17', '13:39:05', 'nugroho', 'user', 'nug@gmail.com', '0.728557', '124.271858', '$2a$10$/CWIqxF6zJWVyJmdD6TkzOCFy3U4.7IRoV4qAnAY879XAnCI.BI0m'),
(98, '2025-02-17', '13:40:02', 'zora', 'user', 'zor@gmail.com', '-8.652933', '117.361649', '$2a$10$LQYkxbD3A09W2Cmd48NENOyjdl6xvqgP7ghnJgOl1YkujcZHT0RRa');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `logs`
--
ALTER TABLE `logs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `logsessions`
--
ALTER TABLE `logsessions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `logtasks`
--
ALTER TABLE `logtasks`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `token` (`tn`),
  ADD KEY `user_id` (`u_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `n` (`n`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `logs`
--
ALTER TABLE `logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT untuk tabel `logsessions`
--
ALTER TABLE `logsessions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=50;

--
-- AUTO_INCREMENT untuk tabel `logtasks`
--
ALTER TABLE `logtasks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- AUTO_INCREMENT untuk tabel `sessions`
--
ALTER TABLE `sessions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=44;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=128;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `sessions`
--
ALTER TABLE `sessions`
  ADD CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
