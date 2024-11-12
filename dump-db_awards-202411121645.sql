-- MySQL dump 10.13  Distrib 5.7.33, for Win64 (x86_64)
--
-- Host: localhost    Database: db_awards
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.37-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `kode_vote`
--

DROP TABLE IF EXISTS `kode_vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kode_vote` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `kode` char(20) NOT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `kode_vote_unique` (`kode`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kode_vote`
--

LOCK TABLES `kode_vote` WRITE;
/*!40000 ALTER TABLE `kode_vote` DISABLE KEYS */;
INSERT INTO `kode_vote` VALUES (1,'aaa',75,NULL,NULL),(3,'OspVaV5kD1Pd5mcW',0,NULL,NULL),(4,'tkPHbOsEyZu6jZ0X',0,NULL,NULL),(5,'iOdIjUapcgD2Obft',0,NULL,NULL),(6,'F7M0ri0dZMwdGYT5',75,NULL,NULL),(7,'rTtDWPmj0VgAsMGZ',75,NULL,NULL);
/*!40000 ALTER TABLE `kode_vote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nominasi`
--

DROP TABLE IF EXISTS `nominasi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `nominasi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(100) DEFAULT NULL,
  `total_vote` int(11) DEFAULT '0',
  `view` tinyint(1) DEFAULT '1',
  `logo` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nominasi`
--

LOCK TABLES `nominasi` WRITE;
/*!40000 ALTER TABLE `nominasi` DISABLE KEYS */;
INSERT INTO `nominasi` VALUES (1,'Mandaffection',0,1,'https://static.miraheze.org/jkt48wikiwiki/3/3d/Mandaffection.jpg'),(2,'Christyzer',276,1,'https://static.miraheze.org/jkt48wikiwiki/0/06/Christyzer.jpg'),(3,'Liamelior',37,1,'https://static.miraheze.org/jkt48wikiwiki/5/52/Liamelior.jpg'),(4,'Callistaverse',0,1,'https://static.miraheze.org/jkt48wikiwiki/2/2e/Callistaverse.jpg'),(5,'Onielity',0,1,'https://static.miraheze.org/jkt48wikiwiki/c/c2/Onielity.jpg'),(6,'Oracle (Olla The Miracle)',0,1,'https://static.miraheze.org/jkt48wikiwiki/d/d8/OllaTheMiracle.jpg'),(7,'Fenidelity',0,1,'https://static.miraheze.org/jkt48wikiwiki/c/c8/Fenidelity.jpg'),(8,'Symfiony',0,1,'https://static.miraheze.org/jkt48wikiwiki/1/11/Symfiony.jpg'),(9,'FloRisen',0,1,'https://static.miraheze.org/jkt48wikiwiki/9/99/Florisen.jpg'),(10,'Freyanation',0,1,'https://static.miraheze.org/jkt48wikiwiki/6/62/Freyanation.jpg'),(11,'Ellatheria',0,1,'https://static.miraheze.org/jkt48wikiwiki/5/56/Abigailuxe.jpg'),(12,'Gitroops',0,1,'https://static.miraheze.org/jkt48wikiwiki/3/3e/Gitroops.jpg'),(13,'Gracieluv',0,1,'https://static.miraheze.org/jkt48wikiwiki/f/fe/Gracieluv.jpg'),(14,'Degrees',0,1,'https://static.miraheze.org/jkt48wikiwiki/9/9e/Degrees.jpg'),(15,'Helismiley',0,1,'https://static.miraheze.org/jkt48wikiwiki/b/bd/Helismiley.png'),(16,'Interindah',0,1,'https://static.miraheze.org/jkt48wikiwiki/5/5e/Interindah.jpg'),(17,'Indiraise',0,1,'https://static.miraheze.org/jkt48wikiwiki/5/5e/Indiraise.jpg'),(18,'Jessination',0,1,'https://static.miraheze.org/jkt48wikiwiki/4/4a/Jessination.jpg'),(19,'Lynear',0,1,'https://static.miraheze.org/jkt48wikiwiki/8/85/Lynear.jpg'),(20,'KATH, Inc',0,1,'https://static.miraheze.org/jkt48wikiwiki/6/6a/KathInc.jpg'),(21,'Lunarian',0,1,'https://static.miraheze.org/jkt48wikiwiki/7/7d/Lunarian.jpg'),(22,'MarshaOshi',0,1,'https://static.miraheze.org/jkt48wikiwiki/9/9f/MarshaOshi.jpg'),(23,'Muffin',0,1,'https://static.miraheze.org/jkt48wikiwiki/5/5d/Muffin.jpg'),(24,'Raishanrise',0,1,'https://static.miraheze.org/jkt48wikiwiki/e/e7/Raishanrise.jpg'),(25,'Gracias',0,1,'https://static.miraheze.org/jkt48wikiwiki/d/d7/Gracias.jpg'),(26,'Alamanda',0,1,'https://static.miraheze.org/jkt48wikiwiki/3/30/Alamanda_logo.jpg'),(27,'Aninimous',0,1,'https://static.miraheze.org/jkt48wikiwiki/6/68/Aninimous.jpg'),(28,'Cathleenexus',0,1,'https://static.miraheze.org/jkt48wikiwiki/6/63/Cathleennexus.png'),(29,'Cellineyours',0,1,'https://static.miraheze.org/jkt48wikiwiki/a/a9/Cellineyours.jpg'),(30,'Chelsealand',0,1,'https://static.miraheze.org/jkt48wikiwiki/3/37/Chelsealand.jpg'),(31,'Cynthiaction',0,1,'https://static.miraheze.org/jkt48wikiwiki/7/7f/Cynthiaction.jpg'),(32,'Denalize',0,1,'https://static.miraheze.org/jkt48wikiwiki/d/df/Denalize.jpg'),(33,'Daisynateur',0,1,'https://static.miraheze.org/jkt48wikiwiki/f/ff/Desynetsui.jpg');
/*!40000 ALTER TABLE `nominasi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'db_awards'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-12 16:45:21
