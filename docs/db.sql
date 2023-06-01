DROP DATABASE IF EXISTS online_shop;
CREATE DATABASE online_shop;
USE online_shop;

DROP TABLE IF EXISTS goods_spu;
CREATE TABLE goods_spu(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '商品名',
  brand_id INT NOT NULL COMMENT '品牌 ID',
  category_id INT NOT NULL COMMENT '类别 ID',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_attr;
CREATE TABLE goods_attr(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  attr VARCHAR(255) NOT NULL COMMENT '属性名',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_spu_attr;
CREATE TABLE goods_spu_attr(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  spu_id INT NOT NULL COMMENT 'SPU ID',
  attr_id INT NOT NULL COMMENT '属性 ID',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

-- 一个属性有多个值, 例如颜色: 红色，蓝色，绿色等
DROP TABLE IF EXISTS goods_attr_option;
CREATE TABLE goods_attr_option(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  attr_id INT NOT NULL COMMENT '属性 ID',
  value VARCHAR(255) NOT NULL COMMENT '属性值',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_brand;
CREATE TABLE goods_brand (
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '品牌名',
  image VARCHAR(255) NOT NULL COMMENT '品牌图片',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

INSERT INTO goods_brand(name, image) VALUES('苹果','https://iphone.image');

DROP TABLE IF EXISTS goods_category;
CREATE TABLE goods_category(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '分类名',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

INSERT INTO goods_category(name) VALUES('手机');

DROP TABLE IF EXISTS goods_sku;
CREATE TABLE goods_sku(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  spu_id INT NOT NULL COMMENT 'SPU ID',
  store_id INT NOT NULL COMMENT '店铺 ID',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_sku_attr_value;
CREATE TABLE goods_sku_attr_value(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  sku_id INT NOT NULL COMMENT 'SKU ID',
  attr_val_id INT NOT NULL COMMENT '属性值ID',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_store;
CREATE TABLE goods_store(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '店铺名',
  enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);
