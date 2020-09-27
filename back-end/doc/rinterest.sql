/*
 Navicat Premium Data Transfer

 Source Server         : docker_mysql8
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : rinterest

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 27/09/2020 20:38:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blogs
-- ----------------------------
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `tags` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `summary` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `invisibility` enum('public','private','protected') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'public',
  `user_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blogs_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blogs
-- ----------------------------
INSERT INTO `blogs` VALUES (1, '2020-09-23 19:17:44', '2020-09-23 19:17:44', NULL, '', 'golang从入门到入土', 'go的一些特性非常适合做后端开发，现在小编就带你从入门到入土一步一步学习go语言吧。', '### 今天让我们来学习golang的安装和卸载\n#### 1. 安装\n#### 2. 打印HelloWorld\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, world\")\n}\n```\n#### 3. 卸载\n\n- - - \n#### 好了，大家学会了吗？', 'public', 2);
INSERT INTO `blogs` VALUES (2, '2020-09-23 20:04:23', '2020-09-23 20:04:23', NULL, '', 'jwt', 'jwt是一种介于前后端的加密传输协议，安全性高，封装简单，十分好用。', '``` go\npackage middlewares\n\nimport (\n	\"net/http\"\n	\"rinterest/models\"\n	\"strings\"\n	\"time\"\n\n	\"github.com/dgrijalva/jwt-go\"\n	\"github.com/gin-gonic/gin\"\n)\n\n// NewToken 生成Token\nfunc NewToken(user *models.User) (tokens string, err error) {\n	claim := jwt.MapClaims{\n		\"id\":        user.ID,\n		\"role\":      user.Role,\n		\"email\":     user.Email,\n		\"phone\":     user.Phone,\n		\"avatar\":    user.Avatar,\n		\"username\":  user.Username,\n		\"signature\": user.Signature,\n		\"iat\":       time.Now().Unix(),\n	}\n	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)\n	tokens, err = token.SignedString([]byte(\"you-will-never-guess\"))\n	return\n}\n\n// RefreshToken ...\nfunc RefreshToken(token string) (tokens string, err error) {\n	return\n}\n\n// JWT Auth middleware\nfunc JWT() gin.HandlerFunc {\n	return func(c *gin.Context) {\n		authorization := c.Request.Header.Get(\"Authorization\")\n		if len(authorization) == 0 {\n			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{\n				\"message\": \"Bad authorization.\",\n			})\n			return\n		}\n		bearerToken := strings.Fields(authorization)[1]\n		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {\n			return []byte(\"you-will-never-guess\"), nil\n		})\n		if err != nil {\n			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{\n				\"message\": \"parse token failed.\",\n			})\n			return\n		}\n\n		if !token.Valid {\n			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{\n				\"message\": \"invalid token.\",\n			})\n		}\n		c.Next()\n	}\n}\n\n```', 'public', 2);
INSERT INTO `blogs` VALUES (3, '2020-09-23 20:28:22', '2020-09-23 20:28:22', NULL, '', '教你一天精通Python', '想要一天就精通python吗，想要一天就拿到python的offer吗，小编现在就带你实现这个愿望，秘诀就是洗洗睡吧。', '# FunCampus\n\n##### FunCampus是一个简单的大学生交友和互助网站，项目初衷是想改善真人表白墙的一些局限性，例如时效性差，不能实时显示，数量限制，范围较小等等，有完善的用户管理，角色管理，信息流展示，评论，以及后台管理。\n\n\n\n\n\n- - -\n##### 安装\n\n1. 下载项目\n\n    `git clone https://github.com/eagony/FunCampus.git`\n    \n    `cd FunCampus`\n    \n2. 部署后端\n\n    `cd  backend`\n\n    创建虚拟环境\n\n    `python -m venv venv`\n\n    激活虚拟环境\n\n    `venv\\Scripts\\activate`\n\n    安装依赖\n\n    `pip install -r requirements.txt`\n\n    这里因为flask_uploads导入werkzeug有点小问题,需要将venv\\lib\\site-packages\\flask_uploads.py第26行的\n\n    `from werkzeug import secure_filename, FileStorage`\n\n    改成\n    ```python\n    from werkzeug.utils import secure_filename\n    from werkzeug.datastructures import FileStorage\n    ```\n\n    同步数据库\n\n    `flask db upgrade`\n\n    部署\n\n    `flask deploy`\n\n    启动\n\n    `flask run`\n\n    测试，打开浏览器输入http://localhost:5000/api/ping如果返回\"Pong\"则后端启动成功\n\n3. 部署前端\n\n    先安装Node.js，前往 [官方网站](https://nodejs.org/zh-cn/) 下载并安装 `LTS` 版本\n\n    安装好后，由于 `npm` 命令使用的国外镜像，在国内下载依赖包时很慢，这里换成 [淘宝 NPM 镜像](https://npm.taobao.org/)\n\n    `npm install -g cnpm --registry=https://registry.npm.taobao.org`\n\n\n    `cd front-end`\n\n    安装依赖\n\n    `cnpm install`\n\n    运行前端应用\n\n    `npm run dev`\n\n    浏览器访问: `http://localhost:8080`\n\n#### ps:注册前修改backend/config.py里的管理员邮箱，注册的时候将自动获得管理员权限。\n\n### 再来些go代码\n### \n```go\npackage apis\n\nimport (\n	\"net/http\"\n	\"rinterest/extensions\"\n	\"rinterest/middlewares\"\n	\"rinterest/models\"\n	\"strconv\"\n\n	\"github.com/gin-gonic/gin\"\n)\n\n// DiaryAPI ...\ntype DiaryAPI struct{}\n\n// Register ...\nfunc (d *DiaryAPI) Register(rg *gin.RouterGroup) {\n	rg.POST(\"/diaries\", d.newone)\n	rg.GET(\"/diaries/:id\", d.getone)\n\n	rg.GET(\"/users/:id/diaries\", middlewares.JWT(), d.getallbyuserid)\n}\n\nfunc (d *DiaryAPI) newone(c *gin.Context) {\n	diary := models.Diary{}\n	if err := c.ShouldBindJSON(&diary); err != nil {\n		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{\n			\"Bind error\": err.Error(),\n		})\n		return\n	}\n\n	if err := extensions.MySQL().Create(&diary).Error; err != nil {\n		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{\n			\"Create error\": err.Error(),\n		})\n		return\n	}\n\n	c.IndentedJSON(http.StatusOK, gin.H{\n		\"status\":  200,\n		\"message\": \" create success\",\n		\"data\":    diary,\n	})\n}\n\nfunc (d *DiaryAPI) getone(c *gin.Context) {\n	id, err := strconv.Atoi(c.Param(\"id\"))\n	if err != nil {\n		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{\n			\"Convert id error\": err.Error(),\n		})\n		return\n	}\n\n	diary := models.Diary{}\n	if err = extensions.MySQL().Where(\"id = ?\", id).Find(&diary).Error; err != nil {\n		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{\n			\"Query error\": err.Error(),\n		})\n		return\n	}\n	c.IndentedJSON(http.StatusOK, gin.H{\n		\"status\":  \"OK\",\n		\"message\": \"success\",\n		\"data\":    diary,\n	})\n}\n\nfunc (d *DiaryAPI) getallbyuserid(c *gin.Context) {\n	id, err := strconv.Atoi(c.Param(\"id\"))\n	if err != nil {\n		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{\n			\"Convert id error\": err.Error(),\n		})\n		return\n	}\n\n	diaries := []models.Diary{}\n	if err = extensions.MySQL().Where(\"user_id = ?\", id).Order(\"created_at desc\").Find(&diaries).Error; err != nil {\n		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{\n			\"Query error\": err.Error(),\n		})\n		return\n	}\n	c.IndentedJSON(http.StatusOK, gin.H{\n		\"status\":  \"OK\",\n		\"message\": \"success\",\n		\"data\":    diaries,\n	})\n}\n\n```\n### 完毕', 'public', 2);
INSERT INTO `blogs` VALUES (4, '2020-09-24 23:57:41', '2020-09-24 23:57:41', NULL, '', 'Flask', 'Flask是一个很好的python后端框架，比dango简单', '### Flask\n### Vue\n```python\nimport numpy as np\ndef main:\n  print(\"Hello\")\n```', 'public', 2);
INSERT INTO `blogs` VALUES (5, '2020-09-25 10:12:16', '2020-09-25 10:12:16', NULL, '', '这是一个标题', '', '# 一级标题\n- - -\n## 二级标题\n### 发  ', 'public', 1);
INSERT INTO `blogs` VALUES (6, '2020-09-25 10:17:53', '2020-09-25 10:17:53', NULL, '', '标题', '摘要摘要', '# 正文\n## 写点代码\n`sudo apt-get install golang`', 'public', 1);
INSERT INTO `blogs` VALUES (7, '2020-09-25 10:20:09', '2020-09-25 10:20:09', NULL, '', '试试', '试试有没有提醒\n', '懒得写', 'public', 1);

-- ----------------------------
-- Table structure for diaries
-- ----------------------------
DROP TABLE IF EXISTS `diaries`;
CREATE TABLE `diaries`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `tags` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_diaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of diaries
-- ----------------------------
INSERT INTO `diaries` VALUES (1, '2020-09-23 19:04:49', '2020-09-23 19:04:49', '2020-09-26 21:12:33', 'mdi-heart-outline,开心,red', '淦', 2);
INSERT INTO `diaries` VALUES (2, '2020-09-23 19:14:42', '2020-09-23 19:14:42', '2020-09-24 22:42:14', 'mdi-weather-pouring,大雨,blue-grey;mdi-movie-outline,电影,orange;mdi-heart-outline,开心,red', 'cao', 2);
INSERT INTO `diaries` VALUES (3, '2020-09-24 13:17:49', '2020-09-24 13:17:49', NULL, 'mdi-heart-outline,开心,red;mdi-white-balance-sunny,晴朗,teal', '刚去吃了午饭', 1);
INSERT INTO `diaries` VALUES (4, '2020-09-26 21:01:17', '2020-09-26 21:01:17', NULL, 'mdi-heart-outline,开心,red;mdi-movie-outline,电影,orange;mdi-weather-pouring,大雨,blue-grey', '今天我还是照常给你发消息，汇报日常工作，你终于回了我四个字：“嗯嗯，好的”你开始愿意敷衍我了，我太感动了受宠若惊。我愿意天天给你发消息。就算你天天骂我，我也不觉得烦。', 2);
INSERT INTO `diaries` VALUES (5, '2020-09-26 21:01:35', '2020-09-26 21:01:35', NULL, 'mdi-heart-outline,开心,red;mdi-cup,奶茶,brown darken-3', '你昨天晚上又没会我的消息，在我孜孜不倦的骚扰下，你终于舍得回我了，你说“滚”，这其中一定有什么含义，我想了很久，滚是三点水，这代表你对我的思念也如滚滚流水一样汹涌，我感动哭了，不知道你现在在干嘛，我很想你。', 2);
INSERT INTO `diaries` VALUES (6, '2020-09-26 21:02:35', '2020-09-26 21:02:35', NULL, 'mdi-emoticon-sad-outline,难过,grey;mdi-movie-outline,电影,orange;mdi-white-balance-sunny,晴朗,teal', '你说你想买口红，今天我去了叔叔的口罩厂做了一天的打包。拿到了两百块钱，加上我这几天省下的钱刚好能给你买一根小金条。即没有给我自己剩下一分钱，但你不用担心，因为厂里包吃包住。对了打包的时候，满脑子都是你，想着你哪天突然就接受我的橄榄枝了呢。而且今天我很棒呢，主管表扬我很能干，其实也有你的功劳啦，是你给了我无穷的力量。今天我比昨天多想你一点，比明天少想你一点。', 2);
INSERT INTO `diaries` VALUES (7, '2020-09-26 21:02:46', '2020-09-26 21:02:46', NULL, 'mdi-heart-outline,开心,red;mdi-emoticon-sad-outline,难过,grey;mdi-weather-pouring,大雨,blue-grey;mdi-white-balance-sunny,晴朗,teal;mdi-movie-outline,电影,orange;mdi-cup,奶茶,brown darken-3', '你说你想买AJ，今天我去了叔叔的口罩厂做了一天的打包。拿到了两百块钱，加上我这几天省下的钱刚好能给你买一个鞋盒。即没有给我自己剩下一分钱，但你不用担心，因为厂里包吃包住。对了打包的时候，满脑子都是你，想着你哪天突然就接受我的橄榄枝了呢。而且今天我很棒呢，主管表扬我很能干，其实也有你的功劳啦，是你给了我无穷的力量。今天我比昨天多想你一点，比明天少想你一点。', 2);

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `for` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `text` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `color` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_tags_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tags
-- ----------------------------
INSERT INTO `tags` VALUES (1, '2020-09-23 19:03:23', '2020-09-23 19:03:23', NULL, '', 'mdi-cup', '奶茶', 'brown darken-3');
INSERT INTO `tags` VALUES (2, '2020-09-23 19:04:38', '2020-09-23 19:04:38', NULL, '', 'mdi-heart-outline', '开心', 'red');
INSERT INTO `tags` VALUES (3, '2020-09-23 19:05:20', '2020-09-23 19:05:20', NULL, '', 'mdi-emoticon-sad-outline', '难过', 'grey');
INSERT INTO `tags` VALUES (4, '2020-09-23 19:06:04', '2020-09-23 19:06:04', NULL, '', 'mdi-movie-outline', '电影', 'orange');
INSERT INTO `tags` VALUES (5, '2020-09-23 19:06:34', '2020-09-23 19:06:34', NULL, '', 'mdi-white-balance-sunny', '晴朗', 'teal');
INSERT INTO `tags` VALUES (6, '2020-09-23 19:14:17', '2020-09-23 19:14:17', NULL, '', 'mdi-weather-pouring', '大雨', 'blue-grey');

-- ----------------------------
-- Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `text` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `done` tinyint(1) NULL DEFAULT NULL,
  `user_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_todos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of todos
-- ----------------------------
INSERT INTO `todos` VALUES (1, '2020-09-23 21:33:21', '2020-09-25 12:28:28', '2020-09-25 13:45:46', '睡觉', 0, 1);
INSERT INTO `todos` VALUES (2, '2020-09-23 21:33:39', '2020-09-26 13:55:20', NULL, '睡觉', 1, 2);
INSERT INTO `todos` VALUES (3, '2020-09-23 21:33:40', '2020-09-23 21:38:48', '2020-09-24 22:12:22', '睡觉', 1, 2);
INSERT INTO `todos` VALUES (4, '2020-09-23 21:36:53', '2020-09-23 21:38:47', NULL, '打游戏', 1, 2);
INSERT INTO `todos` VALUES (5, '2020-09-23 21:38:44', '2020-09-23 21:38:44', '2020-09-24 22:10:56', '发', 0, 2);
INSERT INTO `todos` VALUES (6, '2020-09-23 21:59:49', '2020-09-23 22:11:03', '2020-09-25 10:10:30', '淦分f ', 1, 1);
INSERT INTO `todos` VALUES (7, '2020-09-23 22:02:36', '2020-09-23 22:11:03', '2020-09-25 10:10:31', '掉', 1, 1);
INSERT INTO `todos` VALUES (8, '2020-09-23 22:10:37', '2020-09-23 22:13:55', '2020-09-25 10:10:32', '丰富', 1, 1);
INSERT INTO `todos` VALUES (9, '2020-09-23 22:10:39', '2020-09-23 22:13:46', '2020-09-25 10:10:33', '丰富', 1, 1);
INSERT INTO `todos` VALUES (10, '2020-09-23 22:10:43', '2020-09-23 22:13:48', '2020-09-25 10:10:34', '分为', 1, 1);
INSERT INTO `todos` VALUES (11, '2020-09-23 22:10:45', '2020-09-24 13:17:33', '2020-09-25 10:10:34', '为f', 1, 1);
INSERT INTO `todos` VALUES (12, '2020-09-23 22:10:47', '2020-09-24 13:17:35', '2020-09-25 10:10:35', '范围f', 1, 1);
INSERT INTO `todos` VALUES (13, '2020-09-23 22:16:03', '2020-09-23 22:16:03', '2020-09-25 10:10:36', 'giao', 0, 1);
INSERT INTO `todos` VALUES (14, '2020-09-24 21:59:40', '2020-09-24 21:59:40', '2020-09-24 22:12:06', '丰富', 0, 2);
INSERT INTO `todos` VALUES (15, '2020-09-24 22:12:44', '2020-09-24 22:12:44', NULL, '十点去健身', 0, 2);
INSERT INTO `todos` VALUES (16, '2020-09-25 12:27:32', '2020-09-27 14:51:01', NULL, '洗澡', 1, 1);
INSERT INTO `todos` VALUES (17, '2020-09-25 13:45:10', '2020-09-25 13:45:10', NULL, '吃饭', 0, 1);
INSERT INTO `todos` VALUES (18, '2020-09-26 13:55:13', '2020-09-26 13:55:13', '2020-09-26 22:11:28', '淦', 0, 2);
INSERT INTO `todos` VALUES (19, '2020-09-27 13:07:47', '2020-09-27 13:07:47', '2020-09-27 13:08:53', 'll', 0, 2);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2020-09-23 19:02:38', '2020-09-27 16:54:28', NULL, 'admin', '$2a$04$bd25tFk4.sCS2c9R6sPB6OAW2yvZ8InX3iIm9K1JSOY9VZnB5z7Xy', 'admin', 'fasico@qq.com', '', 'https://firebasestorage.googleapis.com/v0/b/todoteam-3263d.appspot.com/o/images%2FC6ridYN4MZfCUIw8L6bDG6wX0U32?alt=media&token=423cba6a-cfc5-4241-859e-11341cc87e62', 'simple, cool, fast');
INSERT INTO `users` VALUES (2, '2020-09-23 22:17:59', '2020-09-23 22:17:59', NULL, 'fasico', '$2a$04$mnyJd6knOa1iD9MGlIX26OiVqnOzAXWMFm00fwFc2ucKfG1z8I6aG', 'user', 'fasico@qq.com', '', '', 'simple, cool, fast');
INSERT INTO `users` VALUES (8, '2020-09-24 16:33:05', '2020-09-24 16:33:05', NULL, 'test', '$2a$04$YSdvAhHR4W5I8UpPv1IPQeIMIy5acrDoRdYJwUqa.2jMPpaSkbGUS', 'user', '', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
