/*
 * Wikipedia API
 *
 * This is an api that searches a wikipage
 *
 * API version: 1.0.1
 * Contact: bchristian14@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
    "log"
    "net/http"
    "time"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
