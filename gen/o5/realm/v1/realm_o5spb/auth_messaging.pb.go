// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package realm_o5spb

// Service: AuthService
// Method: Whoami

func (msg *WhoamiRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.realm.v1.service.AuthService/Whoami",
		"grpc-message": "o5.realm.v1.service.WhoamiRequest",
	}
	return headers
}
