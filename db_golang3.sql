-- MySQL dump 10.13  Distrib 8.0.21, for macos10.15 (x86_64)
--
-- Host: localhost    Database: db_golang3
-- ------------------------------------------------------
-- Server version	5.7.9

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
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(2,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(3,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(4,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(5,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(6,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(7,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(8,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(9,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(10,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(11,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(12,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(13,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(14,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(15,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(16,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(17,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish'),(18,'2021-07-13 09:49:19','2021-07-13 09:49:19',NULL,'title','content','golang','Publish');
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-08-02 15:08:39
