CREATE TABLE Missao(
	id_missao int,
	nome_missao varchar(100), 
	rank_missao char
);

CREATE TABLE Aventureiro (
	id_avent int,
	nome_avent varchar (100),
	rank_avent char
);
CREATE TABLE Aventura (
	id_aventura int,
	id_missao int,
	nome_missao varchar(100),
	id_avent int,
	nome_avent varchar (100)
);
insert into Aventureiro ( id_avent,nome_avent,rank_avent)
values (123, 'Lucio','S');
insert into Missao (id_missao, nome_missao,rank_missao)
values (1, 'Resgate da princesa','D');
insert into Aventura ( id_aventura,id_missao, nome_missao, id_avent, nome_avent)
values (1,1, 'Resgate da princesa',123, 'Lucio');

drop table Aventureiro;
drop table Missao;
drop table Aventura;