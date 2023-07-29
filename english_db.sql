/*
 Navicat Premium Data Transfer

 Source Server         : root@localhost
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : english_db

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 29/07/2023 23:28:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `admin_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理员id，雪花算法生成',
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '管理员是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_admin_admin_id`(`admin_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for essay
-- ----------------------------
DROP TABLE IF EXISTS `essay`;
CREATE TABLE `essay`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `essay_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '文章id，雪花算法生成',
  `essay_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文章标题',
  `essay_author` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文章作者',
  `publish_at` datetime(3) NULL DEFAULT NULL COMMENT '发布日期',
  `essay_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '文章内容',
  `essay_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文章是否可评论，默认为0，1为可评论，0不可评论（预留字段）',
  `essay_type` tinyint(1) NULL DEFAULT NULL COMMENT '文章类型，0为英语小说，其他待定',
  `essay_collect_num` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章收藏数，默认为0',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文章是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_essay_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_essay_essay_id`(`essay_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5020 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for essay_collect
-- ----------------------------
DROP TABLE IF EXISTS `essay_collect`;
CREATE TABLE `essay_collect`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章收藏表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `essay_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '文章id，雪花算法生成',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏记录是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_essay_collect_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_essay_collect_essay_id`(`essay_id`) USING BTREE,
  INDEX `idx_essay_collect_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for essay_word
-- ----------------------------
DROP TABLE IF EXISTS `essay_word`;
CREATE TABLE `essay_word`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章单词表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `word` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词',
  `word_num` bigint(20) NULL DEFAULT NULL COMMENT '单词序号',
  `word_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '单词id，雪花算法生成',
  `word_music` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词发音链接',
  `essay_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '文章id',
  `word_meaning` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词意思',
  `word_sentence` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '单词是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_essay_word_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_essay_word_word_id`(`word_id`) USING BTREE,
  INDEX `idx_essay_word_essay_id`(`essay_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 165093 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for listen
-- ----------------------------
DROP TABLE IF EXISTS `listen`;
CREATE TABLE `listen`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '听力表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `listen_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '听力id，雪花算法生成',
  `listen_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '听力标题',
  `listen_source` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '听力来源',
  `listen_editor` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '编辑',
  `publish_at` datetime(3) NULL DEFAULT NULL COMMENT '发布日期',
  `listen_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '听力内容',
  `listen_media_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '听力视频链接',
  `listen_mp_3_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '听力音频链接',
  `listen_type` tinyint(1) NULL DEFAULT NULL COMMENT '听力类型，0为热点资讯传送门，1为国外媒体资讯，2为英语听力入门，3为可可之声，4为品牌英语听力',
  `listen_second_type` varchar(48) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '听力第二级类型',
  `listen_collect_num` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '听力收藏数，默认为0',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '听力是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_listen_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_listen_listen_id`(`listen_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13598 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for listen_collect
-- ----------------------------
DROP TABLE IF EXISTS `listen_collect`;
CREATE TABLE `listen_collect`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '听力收藏表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `listen_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '听力id，雪花算法生成',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏记录是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_listen_collect_user_id`(`user_id`) USING BTREE,
  INDEX `idx_listen_collect_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_listen_collect_listen_id`(`listen_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for listen_word
-- ----------------------------
DROP TABLE IF EXISTS `listen_word`;
CREATE TABLE `listen_word`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '听力单词表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `word_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '单词id，雪花算法生成',
  `word` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词',
  `word_phonetic` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词音标',
  `word_meaning` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词意思',
  `word_music` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词发音链接',
  `word_num` bigint(20) NULL DEFAULT NULL COMMENT '单词序号',
  `listen_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '听力id',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '单词是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_listen_word_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_listen_word_word_id`(`word_id`) USING BTREE,
  INDEX `idx_listen_word_listen_id`(`listen_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 110448 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sentence
-- ----------------------------
DROP TABLE IF EXISTS `sentence`;
CREATE TABLE `sentence`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '每日一句表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `sentence_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '句子id，雪花算法生成',
  `publish_at` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发布日期',
  `sentence_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '句子内容',
  `sentence_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '句子译文',
  `sentence_translation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '小编的话',
  `sentence_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '句子图片链接',
  `sentence_audio_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '句子语音链接',
  `sentence_collect_num` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '句子收藏数，默认为0',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '句子是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sentence_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sentence_sentence_id`(`sentence_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sentence_collect
-- ----------------------------
DROP TABLE IF EXISTS `sentence_collect`;
CREATE TABLE `sentence_collect`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '句子收藏表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `sentence_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '句子id，雪花算法生成',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏记录是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sentence_collect_user_id`(`user_id`) USING BTREE,
  INDEX `idx_sentence_collect_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sentence_collect_sentence_id`(`sentence_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成，无符号',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户状态，0为正常，1为禁用，默认为0',
  `last_login_time` datetime(3) NULL DEFAULT NULL COMMENT '上次登录时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_user_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_date
-- ----------------------------
DROP TABLE IF EXISTS `user_date`;
CREATE TABLE `user_date`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户打卡情况表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成，无符号',
  `date` datetime(3) NULL DEFAULT NULL COMMENT '时间',
  `word_learn_number` bigint(20) NULL DEFAULT NULL COMMENT '在这一天新学多少单词',
  `word_review_number` bigint(20) NULL DEFAULT NULL COMMENT '在这一天复习多少单词',
  `remark` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '在这一天的心情感悟',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户信息是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_date_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_user_date_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_feedback
-- ----------------------------
DROP TABLE IF EXISTS `user_feedback`;
CREATE TABLE `user_feedback`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户反馈表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `feedback_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '反馈内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_feedback_user_id`(`user_id`) USING BTREE,
  INDEX `idx_user_feedback_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户信息表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `le_xue_app_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '交友id，也是唯一标识，初始为手机号',
  `gender` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户性别，1为男性，2为女性，0为未知，默认为0',
  `school` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '学校',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地区',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称，注册后为随机',
  `head_sculpture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像云存储地址',
  `integral` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '积分，默认为0，无符号',
  `word_need_recite_num` bigint(20) NOT NULL DEFAULT 0 COMMENT '每日需要背单词的数量',
  `eng_level` tinyint(1) NOT NULL DEFAULT 0 COMMENT '词书等级，默认为0。1为四级，2为六级等',
  `last_start_time` datetime(3) NULL DEFAULT NULL COMMENT '上次学习的时间',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户权限，0为正常，1为VIP，默认为0',
  `invitation_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邀请码，独一无二的邀请码',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个性签名',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户信息是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_info_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_user_info_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for word
-- ----------------------------
DROP TABLE IF EXISTS `word`;
CREATE TABLE `word`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '单词表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `word_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '单词id，雪花算法生成',
  `word` varchar(48) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词',
  `phonetic_trans_eng` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词音标（英）',
  `phonetic_trans_ame` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词音标（美）',
  `word_meaning` varchar(144) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '单词意思',
  `mnemonic_aid` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词助记',
  `chi_etymology` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词中文词源',
  `sentence_eng_1` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句英文1',
  `sentence_chi_1` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句中文1',
  `sentence_eng_2` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句英文2',
  `sentence_chi_2` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句中文2',
  `sentence_eng_3` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句英文3',
  `sentence_chi_3` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '单词例句中文3',
  `word_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '单词类型（所属词书），默认为0，1为四级词汇，2为六级词汇，3为英专四级，4为英专八级，5为考研词汇，6为GRE词汇，7为托福词汇，8为雅思词汇',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '单词是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_word_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_word_word_id`(`word_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26610 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for word_collect
-- ----------------------------
DROP TABLE IF EXISTS `word_collect`;
CREATE TABLE `word_collect`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '单词收藏表主键',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `word_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '单词id，雪花算法生成',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id，雪花算法生成',
  `delete_isok` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏记录是否删除，默认为0，1是删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_word_collect_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_word_collect_word_id`(`word_id`) USING BTREE,
  INDEX `idx_word_collect_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
