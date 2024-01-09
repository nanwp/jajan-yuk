package helper

import "fmt"

func RegisterEmail(name, role, url string) string {
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
			<head>
				<title>Jajan Yuk - Email Verification</title>
			</head>
			<body>
				<h2>Jajan Yuk - Email Verification</h2>
				<p>Hi, %s</p>
				<p>Terima kasih telah mendaftar di Jajan Yuk Apps sebagai %s</p>
				<p>Untuk melanjutkan proses registrasi, silahkan klik link berikut</p>
				<p><a href="%s">Klik disini</a></p>
				</br>
				<p>Terima kasih,</p>
				<p>HiLine</p>
			</body>
		</html>
	`, name, role, url)

	return body
}
