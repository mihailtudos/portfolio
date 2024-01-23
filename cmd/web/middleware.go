package main

import (
	"fmt"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CSP header e used to restrict where the resources for your web page
		// (e.g. JavaScript, images, fonts etc) can be loaded from
		// TODO: get back to setting this header
		// w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com; script-src 'self'")

		// Referrer-Policy is used to control what information is included in a Referer header
		// when a user navigates away from your web page. origin-when-cross-origin instructs that full url
		// will be only used for same origin requests but for other requests the the extra into will be
		// stripped out (Url path, query string, etc.)
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// X-Content-Type-Options: nosniff -  instructs browsers to not MIME-type sniff the
		// content-type of the response, which in turn helps to prevent content-sniffing attacks
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// X-Frame-Options: deny is used to help prevent clickjacking attacks in older browsers
		// that don't support CSP headers.
		w.Header().Set("X-Frame-Options", "deny")

		// X-XSS-Protection: 0 used to disable the blocking of cross-site scripting attacks.
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		uri := r.URL.RequestURI()
		ip := r.RemoteAddr
		proto := r.Proto

		app.logger.Info(fmt.Sprintf("received request: ip = %s proto = %s method = %s uri = %s", ip, proto, method, uri))
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
