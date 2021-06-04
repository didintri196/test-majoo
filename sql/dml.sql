--DML USER
INSERT INTO `tb_user` (`id_user`, `email`, `nama`, `password`, `address`, `no_telp`, `gender`, `id_mercant`, `role`, `last_login`) VALUES ('3', 'kenzo@gmail.com', 'Kenzo Trimasur', MD5('12345678'), 'Kediri, Tawa Timur', '085895567978', 'L', '1', 'admin', NULL);
SELECT * FROM `tb_user`;
UPDATE `tb_user` SET `email` = 'kenzo@gmail.coms', `nama` = 'Kenzo Trimasurs', `password` = MD5('2'), `address` = 'Kediri, Tawa Timurs', `no_telp` = '085895567974', `gender` = 'P', `id_mercant` = NULL, `role` = 'kasir' WHERE `tb_user`.`id_user` = 3;
DELETE FROM `tb_user` WHERE `tb_user`.`id_user` = 3;

--DML PRODUCT
INSERT INTO `tb_barang` (`id_barang`, `sku`, `nama`, `harga_beli`, `harga_jual`, `stock`, `id_supplier`, `id_mercant`, `foto`, `is_aktif`) VALUES (3, '7484894KCG', 'Roti Kacang Super', 10000, 90000, 10, 1, 1, '/public/images/product/b3dadd34-0dbe-494b-9f58-e9a062ffaaa5.png', 'true');
SELECT * FROM `tb_barang`;
UPDATE `tb_barang` SET `sku` = '7484894KCu', `nama` = 'Roti Kacang Superu', `harga_beli` = '1000', `harga_jual` = '9000', `stock` = '100', `id_supplier` = '0', `id_mercant` = '0', `foto` = '/public/images/product/7bf43ef04-ca78-45d9-b75b-ec059604db7b.png', `is_aktif` = 'false' WHERE `tb_barang`.`id_barang` = 5;
DELETE FROM `tb_barang` WHERE id_barang = 5;