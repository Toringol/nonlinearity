package http

// Test no working because of cookie logic (need to mock redis i think)
// func TestSignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userCRUD := user.NewMockUsecase(ctrl)

// 	var userJSON = `{"username": "hello", "password": "123"}`

// 	userInput := &model.User{
// 		Username: "hello",
// 		Password: "123",
// 		Avatar:   "avatars/defaultAvatar",
// 		UserPersonalData: model.UserPersonalData{
// 			DateOfBirth: time.Time{},
// 		},
// 	}

// 	userCRUD.EXPECT().SelectUserByUsername("hello").Return(nil, nil)
// 	userCRUD.EXPECT().CreateUser(userInput).Return(int64(0), nil)

// 	handler := &userHandlers{
// 		usecase: userCRUD,
// 	}

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/signup/", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, handler.handleSignUp(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 	}
// }
