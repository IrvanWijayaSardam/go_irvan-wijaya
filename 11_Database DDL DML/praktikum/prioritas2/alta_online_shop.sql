-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.24 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for alta_online_shop
CREATE DATABASE IF NOT EXISTS `alta_online_shop` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `alta_online_shop`;

-- Dumping structure for table alta_online_shop.customer
CREATE TABLE IF NOT EXISTS `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `address` varchar(100) NOT NULL DEFAULT '0',
  `birth_date` varchar(50) NOT NULL,
  `status` varchar(50) NOT NULL,
  `gender` varchar(50) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

-- Dumping structure for table alta_online_shop.detail_trx
CREATE TABLE IF NOT EXISTS `detail_trx` (
  `id_trx` int(11) DEFAULT NULL,
  `id_product` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `subtotal` int(11) DEFAULT NULL,
  KEY `FK__master_trx` (`id_trx`),
  KEY `FK__product` (`id_product`),
  CONSTRAINT `FK__master_trx` FOREIGN KEY (`id_trx`) REFERENCES `master_trx` (`id`),
  CONSTRAINT `FK__product` FOREIGN KEY (`id_product`) REFERENCES `product` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

-- Dumping structure for table alta_online_shop.master_trx
CREATE TABLE IF NOT EXISTS `master_trx` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_operator` int(11) NOT NULL DEFAULT '0',
  `id_customer` int(11) NOT NULL DEFAULT '0',
  `id_payment_type` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `FK__operator` (`id_operator`),
  KEY `FK__customer` (`id_customer`),
  KEY `FK__payment_type` (`id_payment_type`),
  CONSTRAINT `FK__customer` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id`),
  CONSTRAINT `FK__operator` FOREIGN KEY (`id_operator`) REFERENCES `operator` (`id`),
  CONSTRAINT `FK__payment_type` FOREIGN KEY (`id_payment_type`) REFERENCES `payment_type` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

-- Dumping structure for table alta_online_shop.operator
CREATE TABLE IF NOT EXISTS `operator` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(70) NOT NULL,
  `address` varchar(100) NOT NULL DEFAULT '-',
  `birth_date` varchar(30) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

-- Dumping structure for table alta_online_shop.payment_type
CREATE TABLE IF NOT EXISTS `payment_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `fee` int(100) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

-- Dumping structure for table alta_online_shop.product
CREATE TABLE IF NOT EXISTS `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `description` varchar(100) NOT NULL,
  `price` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
