DROP DATABASE IF EXISTS online_shop;
CREATE DATABASE online_shop;
USE online_shop;

DROP TABLE IF EXISTS goods_brand;
CREATE TABLE goods_brand
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    name   VARCHAR(255) NOT NULL COMMENT '品牌名',
    image  VARCHAR(255) NOT NULL COMMENT '品牌图片',
    enable BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

INSERT INTO goods_brand(name, image)
VALUES ('苹果', 'https://iphone.image');

DROP TABLE IF EXISTS goods_category;
CREATE TABLE goods_category
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    name   VARCHAR(255) NOT NULL COMMENT '分类名',
    enable BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

INSERT INTO goods_category(name)
VALUES ('手机');


DROP TABLE IF EXISTS goods_store;
CREATE TABLE goods_store
(
    id        INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    brand_ids VARCHAR(255) NOT NULL COMMENT '品牌 ID',
    name      VARCHAR(255) NOT NULL COMMENT '店铺名',
    address   VARCHAR(255) NOT NULL COMMENT '店铺地址',
    phone     VARCHAR(255) NOT NULL COMMENT '联系方式',
    enable    BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

INSERT INTO goods_store(brand_ids, name, address, phone)
VALUES ('1', '苹果官方旗舰店', '湖南郴州', '8888-8888');

DROP TABLE IF EXISTS goods_spu;
CREATE TABLE goods_spu
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    brand_id    INT          NOT NULL COMMENT '品牌 ID',
    category_id INT          NOT NULL COMMENT '类别 ID',
    store_id    INT          NOT NULL COMMENT '店铺 ID',
    name        VARCHAR(255) NOT NULL COMMENT '商品名',
    enable      BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_attr;
CREATE TABLE goods_attr
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    spu_id INT          NOT NULL COMMENT 'SPU ID',
    name   VARCHAR(255) NOT NULL COMMENT '属性名',
    enable BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_sku;
create table goods_sku
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    spu_id INT     NOT NULL COMMENT 'SPU ID',
    stock  INT     NOT NULL COMMENT '库存',
    enable BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

-- 一个属性有多个值, 例如颜色: 红色，蓝色，绿色等
DROP TABLE IF EXISTS goods_attr_option;
CREATE TABLE goods_attr_option
(
    id      INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    attr_id INT          NOT NULL COMMENT '属性 ID',
    value   VARCHAR(255) NOT NULL COMMENT '属性值',
    enable  BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_sku_attr_option;
CREATE TABLE goods_sku_attr_option
(
    id             INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    sku_id         INT     NOT NULL COMMENT 'SKU ID',
    attr_option_id INT     NOT NULL COMMENT '属性值ID',
    enable         BOOLEAN NOT NULL DEFAULT TRUE COMMENT '是否可用'
);


DROP TABLE IF EXISTS goods_sku_specification_group;
CREATE TABLE goods_sku_specification_group
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    sku_id INT          NOT NULL COMMENT 'SKU ID',
    name   VARCHAR(255) NOT NULL COMMENT '组名',
    enable BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS goods_sku_specification;
CREATE TABLE goods_sku_specification
(
    id            INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    group_id      INT          NOT NULL COMMENT '分组 ID',
    specification VARCHAR(255) NOT NULL COMMENT '规格名',
    value         VARCHAR(255) NOT NULL COMMENT '规格值',
    enable        BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS ums_user;
CREATE TABLE ums_user
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    name   VARCHAR(255) NOT NULL COMMENT '用户名',
    email  VARCHAR(255) NOT NULL COMMENT '邮箱',
    phone  VARCHAR(255) NOT NULL COMMENT '手机号',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    avatar VARCHAR(255) NOT NULL COMMENT '头像',
    login_time DATETIME NOT NULL COMMENT '登录时间',
    create_time DATETIME NOT NULL COMMENT '创建时间',
    update_time DATETIME NOT NULL COMMENT '更新时间',
    enable BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否可用'
);

DROP TABLE IF EXISTS ums_position;
CREATE TABLE ums_position(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '职位'
);

DROP TABLE IF EXISTS ums_organization;
CREATE TABLE ums_organization(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '组织名'
);

DROP TABLE IF EXISTS ums_organization_leader;
CREATE TABLE ums_organization_leader(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  user_id INT NOT NULL COMMENT '用户ID',
  organization_id INT NOT NULL COMMENT '组织ID'
);

DROP TABLE IF EXISTS ums_user_organization;
CREATE TABLE ums_user_organization(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  user_id INT NOT NULL COMMENT '用户 ID',
  organization_id INT NOT NULL COMMENT '组织 ID'
);

DROP TABLE IF EXISTS ums_user_position;
CREATE TABLE ums_user_position(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  user_id INT NOT NULL COMMENT '用户 ID',
  position_id INT NOT NULL COMMENT '职位 ID'
);

DROP TABLE IF EXISTS ums_role;
CREATE TABLE ums_role(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '角色名'
);

DROP TABLE IF EXISTS ums_user_role;
CREATE TABLE ums_user_role(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  user_id INT NOT NULL COMMENT '用户 ID',
  role_id INT NOT NULL COMMENT '角色 ID'
);

DROP TABLE IF EXISTS ums_menu;
CREATE TABLE ums_menu(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  name VARCHAR(255) NOT NULL COMMENT '菜单名',
  url VARCHAR(255) NOT NULL COMMENT 'URL',
  parent INT NOT NULL COMMENT '父级菜单 ID'
);

DROP TABLE IF EXISTS ums_role_menu;
CREATE TABLE ums_role_menu(
  id INT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  role_id INT NOT NULL COMMENT '角色 ID',
  menu_id INT NOT NULL COMMENT '菜单 ID'
);


