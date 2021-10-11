-- MySQL dump 10.19  Distrib 10.3.29-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: 127.0.0.1    Database: sass_devops
-- ------------------------------------------------------
-- Server version	10.6.4-MariaDB-1:10.6.4+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Temporary table structure for view `authority_menu`
--

DROP TABLE IF EXISTS `authority_menu`;
/*!50001 DROP VIEW IF EXISTS `authority_menu`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `authority_menu` (
  `id` tinyint NOT NULL,
  `created_at` tinyint NOT NULL,
  `updated_at` tinyint NOT NULL,
  `deleted_at` tinyint NOT NULL,
  `menu_level` tinyint NOT NULL,
  `parent_id` tinyint NOT NULL,
  `path` tinyint NOT NULL,
  `name` tinyint NOT NULL,
  `hidden` tinyint NOT NULL,
  `component` tinyint NOT NULL,
  `title` tinyint NOT NULL,
  `icon` tinyint NOT NULL,
  `sort` tinyint NOT NULL,
  `authority_id` tinyint NOT NULL,
  `menu_id` tinyint NOT NULL,
  `keep_alive` tinyint NOT NULL,
  `default_menu` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('p','888','/base/login','POST','','',''),('p','888','/user/register','POST','','',''),('p','888','/api/createApi','POST','','',''),('p','888','/api/getApiList','POST','','',''),('p','888','/api/getApiById','POST','','',''),('p','888','/api/deleteApi','POST','','',''),('p','888','/api/updateApi','POST','','',''),('p','888','/api/getAllApis','POST','','',''),('p','888','/authority/createAuthority','POST','','',''),('p','888','/authority/deleteAuthority','POST','','',''),('p','888','/authority/getAuthorityList','POST','','',''),('p','888','/authority/setDataAuthority','POST','','',''),('p','888','/authority/updateAuthority','PUT','','',''),('p','888','/authority/copyAuthority','POST','','',''),('p','888','/menu/getMenu','POST','','',''),('p','888','/menu/getMenuList','POST','','',''),('p','888','/menu/addBaseMenu','POST','','',''),('p','888','/menu/getBaseMenuTree','POST','','',''),('p','888','/menu/addMenuAuthority','POST','','',''),('p','888','/menu/getMenuAuthority','POST','','',''),('p','888','/menu/deleteBaseMenu','POST','','',''),('p','888','/menu/updateBaseMenu','POST','','',''),('p','888','/menu/getBaseMenuById','POST','','',''),('p','888','/user/changePassword','POST','','',''),('p','888','/user/getUserList','POST','','',''),('p','888','/user/setUserAuthority','POST','','',''),('p','888','/user/deleteUser','DELETE','','',''),('p','888','/fileUploadAndDownload/upload','POST','','',''),('p','888','/fileUploadAndDownload/getFileList','POST','','',''),('p','888','/fileUploadAndDownload/deleteFile','POST','','',''),('p','888','/casbin/updateCasbin','POST','','',''),('p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','888','/casbin/casbinTest/:pathParam','GET','','',''),('p','888','/jwt/jsonInBlacklist','POST','','',''),('p','888','/system/getSystemConfig','POST','','',''),('p','888','/system/setSystemConfig','POST','','',''),('p','888','/system/getServerInfo','POST','','',''),('p','888','/autoCode/createTemp','POST','','',''),('p','888','/autoCode/getTables','GET','','',''),('p','888','/autoCode/getDB','GET','','',''),('p','888','/autoCode/getColumn','GET','','',''),('p','888','/user/setUserInfo','PUT','','',''),('p','888','/email/emailTest','POST','','',''),('p','888','/k8sCluster/createK8sCluster','POST','','',''),('p','888','/k8sCluster/findK8sCluster','GET','','',''),('p','888','/k8sCluster/getK8sClusterList','GET','','',''),('p','888','/k8sNodes/getK8sNodesList','GET','','',''),('p','888','/k8sNamespaces/findK8sNamespaces','GET','','',''),('p','888','/k8sNamespaces/getK8sNamespacesList','GET','','',''),('p','888','/k8sDeployments/getK8sDeploymentList','GET','','',''),('p','888','/k8sDeployments/restartK8sDeployment','PUT','','',''),('p','888','/k8sPods/getK8sPodsList','GET','','',''),('p','888','/k8sPods/getPodLog','GET','','',''),('p','888','/container/execContainer','GET','','',''),('p','8881','/base/login','POST','','',''),('p','8881','/user/register','POST','','',''),('p','8881','/api/createApi','POST','','',''),('p','8881','/api/getApiList','POST','','',''),('p','8881','/api/getApiById','POST','','',''),('p','8881','/api/deleteApi','POST','','',''),('p','8881','/api/updateApi','POST','','',''),('p','8881','/api/getAllApis','POST','','',''),('p','8881','/authority/createAuthority','POST','','',''),('p','8881','/authority/deleteAuthority','POST','','',''),('p','8881','/authority/getAuthorityList','POST','','',''),('p','8881','/authority/setDataAuthority','POST','','',''),('p','8881','/menu/getMenu','POST','','',''),('p','8881','/menu/getMenuList','POST','','',''),('p','8881','/menu/addBaseMenu','POST','','',''),('p','8881','/menu/getBaseMenuTree','POST','','',''),('p','8881','/menu/addMenuAuthority','POST','','',''),('p','8881','/menu/getMenuAuthority','POST','','',''),('p','8881','/menu/deleteBaseMenu','POST','','',''),('p','8881','/menu/updateBaseMenu','POST','','',''),('p','8881','/menu/getBaseMenuById','POST','','',''),('p','8881','/user/changePassword','POST','','',''),('p','8881','/user/getUserList','POST','','',''),('p','8881','/user/setUserAuthority','POST','','',''),('p','8881','/fileUploadAndDownload/upload','POST','','',''),('p','8881','/fileUploadAndDownload/getFileList','POST','','',''),('p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),('p','8881','/casbin/updateCasbin','POST','','',''),('p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','8881','/jwt/jsonInBlacklist','POST','','',''),('p','8881','/system/getSystemConfig','POST','','',''),('p','8881','/system/setSystemConfig','POST','','',''),('p','9528','/base/login','POST','','',''),('p','9528','/user/register','POST','','',''),('p','9528','/api/createApi','POST','','',''),('p','9528','/api/getApiList','POST','','',''),('p','9528','/api/getApiById','POST','','',''),('p','9528','/api/deleteApi','POST','','',''),('p','9528','/api/updateApi','POST','','',''),('p','9528','/api/getAllApis','POST','','',''),('p','9528','/authority/createAuthority','POST','','',''),('p','9528','/authority/deleteAuthority','POST','','',''),('p','9528','/authority/getAuthorityList','POST','','',''),('p','9528','/authority/setDataAuthority','POST','','',''),('p','9528','/menu/getMenu','POST','','',''),('p','9528','/menu/getMenuList','POST','','',''),('p','9528','/menu/addBaseMenu','POST','','',''),('p','9528','/menu/getBaseMenuTree','POST','','',''),('p','9528','/menu/addMenuAuthority','POST','','',''),('p','9528','/menu/getMenuAuthority','POST','','',''),('p','9528','/menu/deleteBaseMenu','POST','','',''),('p','9528','/menu/updateBaseMenu','POST','','',''),('p','9528','/menu/getBaseMenuById','POST','','',''),('p','9528','/user/changePassword','POST','','',''),('p','9528','/user/getUserList','POST','','',''),('p','9528','/user/setUserAuthority','POST','','',''),('p','9528','/fileUploadAndDownload/upload','POST','','',''),('p','9528','/fileUploadAndDownload/getFileList','POST','','',''),('p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),('p','9528','/casbin/updateCasbin','POST','','',''),('p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','9528','/jwt/jsonInBlacklist','POST','','',''),('p','9528','/system/getSystemConfig','POST','','',''),('p','9528','/system/setSystemConfig','POST','','',''),('p','9528','/autoCode/createTemp','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'10.png','http://qmplusimg.henrongyi.top/gvalogo.png','png','158787308910.png'),(2,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'logo.png','http://qmplusimg.henrongyi.top/1576554439myAvatar.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `jwt` text DEFAULT NULL COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/base/login','用户登录','base','POST'),(2,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/register','用户注册','user','POST'),(3,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/createApi','创建api','api','POST'),(4,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/getApiList','获取api列表','api','POST'),(5,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/getApiById','获取api详细信息','api','POST'),(6,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/deleteApi','删除Api','api','POST'),(7,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/updateApi','更新Api','api','POST'),(8,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/api/getAllApis','获取所有api','api','POST'),(9,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/createAuthority','创建角色','authority','POST'),(10,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/deleteAuthority','删除角色','authority','POST'),(11,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/getAuthorityList','获取角色列表','authority','POST'),(12,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/getMenu','获取菜单树','menu','POST'),(13,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/getMenuList','分页获取基础menu列表','menu','POST'),(14,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/addBaseMenu','新增菜单','menu','POST'),(15,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/getBaseMenuTree','获取用户动态路由','menu','POST'),(16,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','menu','POST'),(17,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/getMenuAuthority','获取指定角色menu','menu','POST'),(18,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/deleteBaseMenu','删除菜单','menu','POST'),(19,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/updateBaseMenu','更新菜单','menu','POST'),(20,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/menu/getBaseMenuById','根据id获取菜单','menu','POST'),(21,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/changePassword','修改密码','user','POST'),(22,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/getUserList','获取用户列表','user','POST'),(23,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/setUserAuthority','修改用户角色','user','POST'),(24,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/fileUploadAndDownload/upload','文件上传示例','fileUploadAndDownload','POST'),(25,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','fileUploadAndDownload','POST'),(26,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(27,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(28,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/fileUploadAndDownload/deleteFile','删除文件','fileUploadAndDownload','POST'),(29,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单','jwt','POST'),(30,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/setDataAuthority','设置角色资源权限','authority','POST'),(31,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/system/getSystemConfig','获取配置文件内容','system','POST'),(32,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/system/setSystemConfig','设置配置文件内容','system','POST'),(33,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/casbin/casbinTest/:pathParam','RESTFUL模式测试','casbin','GET'),(34,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/autoCode/createTemp','自动化代码','autoCode','POST'),(35,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/updateAuthority','更新角色信息','authority','PUT'),(36,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/authority/copyAuthority','拷贝角色','authority','POST'),(37,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/deleteUser','删除用户','user','DELETE'),(38,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/autoCode/getTables','获取数据库表','autoCode','GET'),(39,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/autoCode/getDB','获取所有数据库','autoCode','GET'),(40,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/autoCode/getColumn','获取所选table的所有字段','autoCode','GET'),(41,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/user/setUserInfo','设置用户信息','user','PUT'),(42,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/system/getServerInfo','获取服务器信息','system','POST'),(43,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/email/emailTest','发送测试邮件','email','POST'),(44,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sDeployments/restartK8sDeployment','重启Deployments','k8sDeployments','PUT'),(47,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sNamespaces/findK8sNamespaces','根据ID获取k8sNamespaces','k8sNamespaces','GET'),(48,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sNamespaces/getK8sNamespacesList','获取所有k8sNamespaces','k8sNamespaces','GET'),(49,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sDeployments/getK8sDeploymentList','获取所有k8sDeployments','k8sDeployments','GET'),(50,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sPods/getK8sPodsList','获取所有k8sPods','k8sPods','GET'),(51,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sCluster/createK8sCluster','创建k8sCluster','k8sCluster','POST'),(52,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sCluster/getK8sClusterList','获取k8sCluster','k8sCluster','GET'),(53,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sCluster/findK8sCluster','用id查询k8sCluster','k8sCluster','GET'),(54,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sPods/getPodLog','获取Pod日志','k8sPods','GET'),(55,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/k8sNodes/getK8sNodesList','获取节点','k8sNodes','GET'),(56,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'/container/execContainer','进入容器','container','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父角色ID',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'888','普通用户','0'),('2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'8881','普通用户子角色','888'),('2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'9528','测试角色','0');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_authority_menus` (
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`sys_authority_authority_id`,`sys_base_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES ('888',1),('888',2),('888',3),('888',4),('888',5),('888',6),('888',7),('888',8),('888',9),('888',10),('888',11),('888',12),('888',13),('888',14),('888',15),('888',16),('888',17),('888',18),('888',19),('888',20),('888',21),('888',22),('888',23),('888',24),('888',25),('888',26),('888',27),('888',28),('888',29),('8881',1),('8881',2),('8881',8),('9528',1),('9528',2),('9528',3),('9528',4),('9528',5),('9528',6),('9528',7),('9528',8),('9528',9),('9528',10),('9528',11),('9528',12),('9528',14),('9528',15),('9528',16),('9528',17);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_parameters`
--

DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_parameters`
--

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','dashboard','dashboard',0,'view/dashboard/index.vue',1,0,0,'首页','s-home'),(2,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','kubernetes','kubernetes',0,'view/kubernetes/index.vue',2,0,0,'Kubernetes管理','cloudy'),(3,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','admin','superAdmin',0,'view/superAdmin/index.vue',3,0,0,'超级管理员','user-solid'),(4,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'3','authority','authority',0,'view/superAdmin/authority/authority.vue',1,0,0,'角色管理','s-custom'),(5,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'3','menu','menu',0,'view/superAdmin/menu/menu.vue',2,1,0,'菜单管理','s-order'),(6,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'3','api','api',0,'view/superAdmin/api/api.vue',3,1,0,'api管理','s-platform'),(7,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'3','user','user',0,'view/superAdmin/user/user.vue',4,0,0,'用户管理','coordinate'),(8,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'3','system','system',0,'view/superAdmin/system/system.vue',5,0,0,'系统配置','s-operation'),(9,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','person','person',1,'view/person/person.vue',4,0,0,'个人信息','message-solid'),(10,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','systemTools','systemTools',0,'view/systemTools/index.vue',5,0,0,'系统工具','s-cooperation'),(11,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'10','autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,1,0,'代码生成器','cpu'),(12,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'10','formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',2,1,0,'表单生成器','magic-stick'),(13,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'10','upload','upload',0,'view/systemTools/upload/upload.vue',3,0,0,'上传下载','upload'),(14,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'0','state','state',0,'view/system/state.vue',6,0,0,'服务器状态','s-opportunity'),(15,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'2','k8sCluster','k8sCluster',0,'view/kubernetes/clusters/index.vue',0,0,0,'集群管理','menu'),(16,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'2','k8sNamespaces','k8sNamespaces',0,'view/kubernetes/namespaces/k8sNamespaces.vue',2,0,0,'命名空间管理','menu'),(17,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'2','k8sDeployments','k8sDeployments',0,'view/kubernetes/deployments/k8sDeployments.vue',3,0,0,'应用管理','s-grid'),(18,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'2','k8sPods','k8sPods',0,'view/kubernetes/pods/index.vue',4,0,0,'容器管理','s-grid'),(19,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,0,'2','k8sNodes','k8sNodes',0,'view/kubernetes/nodes/k8sNodes.vue',1,0,0,'节点管理','s-grid');
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES ('888','888'),('888','8881'),('888','9528'),('9528','8881'),('9528','9528');
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
  `authority_id` varchar(90) DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'9a0bb8a4-2bd2-4e5c-8a90-452b0bf7db3b','admin','e10adc3949ba59abbe56e057f20f883e','超级管理员','http://qmplusimg.henrongyi.top/gva_header.jpg','888'),(2,'2021-09-30 15:52:43','2021-09-30 15:52:43',NULL,'2b6fc6d5-b123-4ebb-8b3d-9433dc01d10b','a303176530','3ec063004a6f31642261936a379fde3d','QMPlusUser','http://qmplusimg.henrongyi.top/1572075907logo.png','9528');
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `authority_menu`
--

/*!50001 DROP TABLE IF EXISTS `authority_menu`*/;
/*!50001 DROP VIEW IF EXISTS `authority_menu`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on(`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-11  9:35:32
