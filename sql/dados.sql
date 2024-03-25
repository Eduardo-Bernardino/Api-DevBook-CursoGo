insert into usuarios (nome, nick, email, senha)
values
("Eduardo", "Edu_Mello", "eduardo@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario1
("Rafael", "Mell0", "rafael@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario2
("Thiago", "Thiaguim", "thiago@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Eduardo", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Rafael", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Thiago", "Essa é a publicação do usuário 3! Oba!", 3);