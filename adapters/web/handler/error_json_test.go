package handler

import ()

func TestHandlerJsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Hello Json"}`))
}
