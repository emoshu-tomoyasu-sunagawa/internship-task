-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- ホスト: emonavi-db
-- 生成日時: 2023 年 2 月 03 日 10:44
-- サーバのバージョン： 8.0.31
-- PHP のバージョン: 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- データベース: `emonavi_db`
--
CREATE DATABASE IF NOT EXISTS `emonavi_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `emonavi_db`;

-- --------------------------------------------------------

--
-- テーブルの構造 `departments`
--

CREATE TABLE `departments` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部署名',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部署';

-- --------------------------------------------------------

--
-- テーブルの構造 `employment_status`
--

CREATE TABLE `employment_status` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `employment_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '雇用形態',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='雇用形態';

-- --------------------------------------------------------

--
-- テーブルの構造 `members`
--

CREATE TABLE `members` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '社員番号',
  `profile_img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'プロフィール画像URL',
  `full_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '氏名',
  `kana_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'フリガナ',
  `motto` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '一言コメント',
  `biography` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '自己紹介文',
  `start_date` timestamp NULL DEFAULT NULL COMMENT '入社日',
  `end_date` timestamp NULL DEFAULT NULL COMMENT '退社日',
  `employment_status` int UNSIGNED DEFAULT NULL COMMENT '雇用形態',
  `status` int UNSIGNED DEFAULT NULL COMMENT 'ステータス',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='EMoshUメンバー';

-- --------------------------------------------------------

--
-- テーブルの構造 `member_department`
--

CREATE TABLE `member_department` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `member_id` int UNSIGNED NOT NULL COMMENT 'メンバーID',
  `department_id` int UNSIGNED NOT NULL COMMENT '部署ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='所属部署';

-- --------------------------------------------------------

--
-- テーブルの構造 `member_role`
--

CREATE TABLE `member_role` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `member_id` int UNSIGNED NOT NULL COMMENT 'メンバーID',
  `role_id` int UNSIGNED NOT NULL COMMENT '役職・職種ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='メンバー役職・職種';

-- --------------------------------------------------------

--
-- テーブルの構造 `roles`
--

CREATE TABLE `roles` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '役職名・職種名',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='役職・職種';

-- --------------------------------------------------------

--
-- テーブルの構造 `status`
--

CREATE TABLE `status` (
  `id` int UNSIGNED NOT NULL COMMENT '自動採番',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ステータス',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='ステータス';

--
-- ダンプしたテーブルのインデックス
--

--
-- テーブルのインデックス `departments`
--
ALTER TABLE `departments`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `employment_status`
--
ALTER TABLE `employment_status`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `members`
--
ALTER TABLE `members`
  ADD PRIMARY KEY (`id`),
  ADD KEY `no` (`no`),
  ADD KEY `members_employment_status_id_foreign` (`employment_status`),
  ADD KEY `members_status_id_foreign` (`status`);

--
-- テーブルのインデックス `member_department`
--
ALTER TABLE `member_department`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `member_department_unique` (`member_id`,`department_id`),
  ADD KEY `member_department_department_id_foreign` (`department_id`);

--
-- テーブルのインデックス `member_role`
--
ALTER TABLE `member_role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `member_role_unique` (`member_id`,`role_id`),
  ADD KEY `member_role_role_id_foreign` (`role_id`);

--
-- テーブルのインデックス `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `status`
--
ALTER TABLE `status`
  ADD PRIMARY KEY (`id`);

--
-- ダンプしたテーブルの AUTO_INCREMENT
--

--
-- テーブルの AUTO_INCREMENT `departments`
--
ALTER TABLE `departments`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `employment_status`
--
ALTER TABLE `employment_status`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `members`
--
ALTER TABLE `members`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `member_department`
--
ALTER TABLE `member_department`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `member_role`
--
ALTER TABLE `member_role`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- テーブルの AUTO_INCREMENT `status`
--
ALTER TABLE `status`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自動採番';

--
-- ダンプしたテーブルの制約
--

--
-- テーブルの制約 `members`
--
ALTER TABLE `members`
  ADD CONSTRAINT `members_employment_status_id_foreign` FOREIGN KEY (`employment_status`) REFERENCES `employment_status` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `members_status_id_foreign` FOREIGN KEY (`status`) REFERENCES `status` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- テーブルの制約 `member_department`
--
ALTER TABLE `member_department`
  ADD CONSTRAINT `member_department_department_id_foreign` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `member_department_member_id_foreign` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- テーブルの制約 `member_role`
--
ALTER TABLE `member_role`
  ADD CONSTRAINT `member_role_member_id_foreign` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `member_role_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
