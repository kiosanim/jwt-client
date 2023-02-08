package client

//func init() {
//	viper.SetDefault("USERNAME", "juvenal")
//	viper.SetDefault("PASSWORD", "Dp}btpMyA(j3GK)nca2H{zX7h{v2BoK5]QTXJm7")
//	viper.SetDefault("CLIENT_ID", "clientid.xpto")
//	viper.SetDefault("CLIENT_SECRET", "clientsecret.xpto")
//	viper.SetDefault("AUTHENTICATION_URI", "http://localhost/realms/xpto/protocol/openid-connect/token")
//	viper.SetDefault("LOGOUT_URI", "http://localhost/realms/xpto/protocol/openid-connect/token/logout")
//	viper.SetDefault("FAKE_REFRESH_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImU3MmRjYTZiLTdjNDMtNDQ0ZS04YmMzLTk2NGJlMzFlMWE3MCJ9.eyJleHAiOjQwNzUzNjE0NTEsImlhdCI6NDA3NTM1OTY1MSwianRpIjoiODcyZWRmYzktNjVmYi00MzI3LTk1NzAtM2ExNWFkY2JkYjAzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9yZWFsbXMvdGVuYW50IiwiYXVkIjoiaHR0cDovL2xvY2FsaG9zdC9yZWFsbXMvdGVuYW50Iiwic3ViIjoiOGJhMWM4ZDAtZTY2ZC00ODhhLWIwZmQtNWUwMzVhZGI5OWQ4IiwidHlwIjoiUmVmcmVzaCIsImF6cCI6InRlc3RlLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIxZDFjZTUxZS1mYmUxLTRiNTctOGQ3NC04NGNlZjQ0NTc0NDEiLCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiIxZDFjZTUxZS1mYmUxLTRiNTctOGQ3NC04NGNlZjQ0NTc0NDEifQ.pyIzlz-g2mxzYnDOnRWqJpOuhZff0F-0lkNv2WnMOMM")
//	viper.SetDefault("FAKE_ACCESS_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImU3MmRjYTZiLTdjNDMtNDQ0ZS04YmMzLTk2NGJlMzFlMWE3MCJ9.eyJleHAiOjQwNzUzNjE0NTEsImlhdCI6NDA3NTM1OTY1MSwianRpIjoiODcyZWRmYzktNjVmYi00MzI3LTk1NzAtM2ExNWFkY2JkYjAzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9yZWFsbXMvdGVuYW50IiwiYXVkIjoiaHR0cDovL2xvY2FsaG9zdC9yZWFsbXMvdGVuYW50Iiwic3ViIjoiOGJhMWM4ZDAtZTY2ZC00ODhhLWIwZmQtNWUwMzVhZGI5OWQ4IiwidHlwIjoiUmVmcmVzaCIsImF6cCI6InRlc3RlLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIxZDFjZTUxZS1mYmUxLTRiNTctOGQ3NC04NGNlZjQ0NTc0NDEiLCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiIxZDFjZTUxZS1mYmUxLTRiNTctOGQ3NC04NGNlZjQ0NTc0NDEifQ.pyIzlz-g2mxzYnDOnRWqJpOuhZff0F-0lkNv2WnMOMM")
//}
//
//func TestNewTokenClient(t *testing.T) {
//	got := NewTokenClient()
//	want := TokenClient{}
//	if reflect.DeepEqual(got, want) {
//		t.Errorf("NewTokenClient() = %v, want %v", got, want)
//	}
//}
//
//func TestTokenClient_GetAccessToken(t *testing.T) {
//	tokenClient := NewTokenClient()
//	tokenClient.currentTokens.AccessToken = viper.GetString("FAKE_ACCESS_TOKEN")
//	token, err := tokenClient.GetAccessToken()
//	if err != nil {
//		t.Errorf("GetAccessToken() = %v", token)
//	}
//
//	//TODO Validar o access token (Mas não perder muito tempo com ele
//
//}

//
//func TestTokenClient_GetAccessToken(t1 *testing.T) {
//	type fields struct {
//		currentTokens *model.Tokens
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := &TokenClient{
//				currentTokens: tt.fields.currentTokens,
//			}
//			got, err := t.GetAccessToken()
//			if (err != nil) != tt.wantErr {
//				t1.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t1.Errorf("GetAccessToken() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTokenClient_Logout(t1 *testing.T) {
//	type fields struct {
//		currentTokens *model.Tokens
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := &TokenClient{
//				currentTokens: tt.fields.currentTokens,
//			}
//			if err := t.Logout(); (err != nil) != tt.wantErr {
//				t1.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestTokenClient_authenticate(t1 *testing.T) {
//	type fields struct {
//		currentTokens *model.Tokens
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := &TokenClient{
//				currentTokens: tt.fields.currentTokens,
//			}
//			if err := t.authenticate(); (err != nil) != tt.wantErr {
//				t1.Errorf("authenticate() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestTokenClient_isTokenValid(t1 *testing.T) {
//	type fields struct {
//		currentTokens *model.Tokens
//	}
//	type args struct {
//		token string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := &TokenClient{
//				currentTokens: tt.fields.currentTokens,
//			}
//			if err := t.isTokenValid(tt.args.token); (err != nil) != tt.wantErr {
//				t1.Errorf("isTokenValid() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestTokenClient_refreshTokens(t1 *testing.T) {
//	type fields struct {
//		currentTokens *model.Tokens
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := &TokenClient{
//				currentTokens: tt.fields.currentTokens,
//			}
//			if err := t.refreshTokens(); (err != nil) != tt.wantErr {
//				t1.Errorf("refreshTokens() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
