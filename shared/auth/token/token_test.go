package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA++Lmn/Gh7M3yNZBAELdG
A/QvMop8aQXd3GJhTVUqgMNVCQH+s9tDEZqpCT4YUmWiFISJRS3OmEZkd++yx4Hc
OtVpgkw/d+YEJZwfXUxSaMccZRgiwE4TpehjVIPbRBvDWXia/W15lmhjFz5eg2K/
/8hfvl81IoTMYOI1kwnGJ+lCG5lHrBWQ0lGE71Py6ueq9YyaGi/2At4/14YVQWQ5
1msA28XRWTKeynZHhg1pVHfR1V3p7lzeBo4L6keXSY/0MJqZEorbHHwQFsBjNy+1
i8GoaejRi5X0ItTZv0yqpIY6ZZAL0OTHuhIxnggJBNcbljceUgZg78kaUuIUHvI2
2QIDAQAB
-----END PUBLIC KEY-----`

func TestJWTTokenVerifier_Verify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannto parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiU1pUVVJDL3NlcnZlci9hdXRoIiwic3ViIjoiNjMzZTk2MDkzNDc3ZTI3NGJjZTE3M2E2In0.ep3SP32kENh6WtVZOZAW9792xzLBm4zozAGx4p-kYYkZK5omtMPc-BxDcIp_mEdGMh8R1T1x-kzzBpbT7aGsHMBA4p79s1eiFjlGKjw32r1CIcPGqt0qP1AoTe3eUthgDA_IkxH8S7PGDsDmt8OmWSEVNbVDqa5l_7sMqYAd-qfawdi7_sx-tuvYxW8Qnx8HdIFZZAsSDk7VgcavwT4gfIAnfdLA4eN3NDBlT5qSPdWD-zS4nYceMx1GFow5iXaNhwZT7TVvRI-SSU4QUMOizO3mA2f79mnx5D4CtKJS7mq5CB01KcO9UuvrMXIh1FRCe96JONWbIQne8JpowjmPHg",
			now:     time.Unix(1516239122, 0),
			want:    "633e96093477e274bce173a6",
			wantErr: false,
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiU1pUVVJDL3NlcnZlci9hdXRoIiwic3ViIjoiNjMzZTk2MDkzNDc3ZTI3NGJjZTE3M2E2In0.ep3SP32kENh6WtVZOZAW9792xzLBm4zozAGx4p-kYYkZK5omtMPc-BxDcIp_mEdGMh8R1T1x-kzzBpbT7aGsHMBA4p79s1eiFjlGKjw32r1CIcPGqt0qP1AoTe3eUthgDA_IkxH8S7PGDsDmt8OmWSEVNbVDqa5l_7sMqYAd-qfawdi7_sx-tuvYxW8Qnx8HdIFZZAsSDk7VgcavwT4gfIAnfdLA4eN3NDBlT5qSPdWD-zS4nYceMx1GFow5iXaNhwZT7TVvRI-SSU4QUMOizO3mA2f79mnx5D4CtKJS7mq5CB01KcO9UuvrMXIh1FRCe96JONWbIQne8JpowjmPHg",
			now:     time.Unix(1517239122, 0),
			want:    "",
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			now:     time.Unix(1516239122, 0),
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiU1pUVVJDL3NlcnZlci9hdXRoIiwic3ViIjoiNjMzZTk2MDkzNDc3ZTI3NGJjZTE3M2E2In0.RTT3iPgnBpHvFYB2rajxsI9nJcGWzBcioWXkv5p-03HckeaZi5QuSFcqBlTpO5vAJTCn_d0htBmw9drHEuOUUfn5OFua-f-2gZk5yOhpFTI_ri8axEhEjs9hHORvw3RBhyxKMqG2ukMtxiKyONoIw_CNDKavahihnm7GZth9z2wCO00_3pD2R_c3MPXqqwSlIvRFXkSq_1micJpZbmoF_cEb4rlt1ySN2JcQrwbJ-cDCJZcabblSc2VreM73ztjgXLdEhjY9GGbxZtrA_UXuofsPbhnuHrntklgDWbpXJtsTZALkDpT2o0Ubs7qPmMWw8u1nmoqR4I47dOjVlBjlzg",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error;got not error")
			}
			if accountID != c.want {
				t.Errorf("wrong account id. want: %q, got: %q", c.want, accountID)
			}
		})
	}

	//tkn := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiU1pUVVJDL3NlcnZlci9hdXRoIiwic3ViIjoiNjMzZTk2MDkzNDc3ZTI3NGJjZTE3M2E2In0.ep3SP32kENh6WtVZOZAW9792xzLBm4zozAGx4p-kYYkZK5omtMPc-BxDcIp_mEdGMh8R1T1x-kzzBpbT7aGsHMBA4p79s1eiFjlGKjw32r1CIcPGqt0qP1AoTe3eUthgDA_IkxH8S7PGDsDmt8OmWSEVNbVDqa5l_7sMqYAd-qfawdi7_sx-tuvYxW8Qnx8HdIFZZAsSDk7VgcavwT4gfIAnfdLA4eN3NDBlT5qSPdWD-zS4nYceMx1GFow5iXaNhwZT7TVvRI-SSU4QUMOizO3mA2f79mnx5D4CtKJS7mq5CB01KcO9UuvrMXIh1FRCe96JONWbIQne8JpowjmPHg"
	////设定验证时的时间
	//jwt.TimeFunc = func() time.Time {
	//	return time.Unix(1516239122, 0)
	//}
	//accountID, err := v.Verify(tkn)
	//if err != nil {
	//	t.Errorf("verification failed: %v", err)
	//}
	//want := "633e96093477e274bce173a6"
	//if accountID != want {
	//	t.Errorf("wrong account id. want: %q, got: %q", want, accountID)
	//}
}
