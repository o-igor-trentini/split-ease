package repository

//func TestUserDomainRepository_Create(t *testing.T) {
//	t.Run("Insere com sucesso", func(t *testing.T) {
//		var rs repositorySuite
//		rs.Setup()
//
//		user := model.NewUserDomain(
//			"Igor",
//			"Trentini",
//			"igor@gmail.com",
//			"igor",
//			"Senha@123",
//		)
//		user.EncryptPassword()
//
//		now := time.Now()
//
//		rs.mock.ExpectBegin()
//		rs.mock.
//			ExpectExec(regexp.QuoteMeta(`INSERT INTO "us_users" ("us_created_at","us_updated_at","us_deleted_at","us_first_name","us_last_name","us_email","us_username","us_password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "us_id"`)).
//			WithArgs(
//				now,
//				now,
//				nil,
//				user.GetFirstName(),
//				user.GetLastName(),
//				user.GetEmail(),
//				user.GetUsername(),
//				user.GetPassword(),
//			).
//			WillReturnResult(sqlmock.NewResult(1, 1))
//		rs.mock.ExpectCommit()
//
//		r := NewUserDomain(rs.DB)
//
//		userEntity, err := r.Create(user)
//		assert.NoError(s.T(), err) // avalia se não houve nenhum erro na execução
//		//assert.Equal(rs.T(), rs.person, p) // verificar se as struts são iguais
//
//		fmt.Print(userEntity)
//	})
//}
