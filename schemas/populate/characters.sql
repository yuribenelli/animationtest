insert into characters 
(name,race,class,base_damage,hp, background,src_image) 
values 
('Barristan','Elf','Warlock',12,120,'Nato e cresciuto in una fattoria ad ovest di Neverwinter...','http://localhost:4200/assets/barristan.jpg'),
('Damien','Human','Rogue',13,85,'Abituato a sopravvivere al gelo e al freddo si abituò presto...','http://localhost:4200/assets/damien.jpg');

UPDATE characters
SET src_image = 'http://localhost:4200/assets/damien.jpg'
WHERE id = 2;