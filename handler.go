package godata

import (
	"encoding/xml"
	"github.com/qwerty-iot/godata/result"
	"github.com/qwerty-iot/tox"
	"net/http"
	"strings"
)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logDebug(r, nil, "request received.")
	var err error

	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(nil)
		return
	} else if len(pathParts) == 1 {
		scope := pathParts[0]
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Odata-Version", "4.0")
		w.WriteHeader(200)
		model := s.Model
		var b []byte
		if model == nil {
			model, err = s.ModelHandler.Metadata(s, scope)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
			b, _ = model.ToSimpleJson("")
		} else {
			b, _ = model.ToSimpleJson(scope)
		}
		_, _ = w.Write(b)
		return
	} else {
		scope := pathParts[0]
		entity := pathParts[1]
		if entity == "$metadata" {
			w.Header().Set("Content-Type", "application/xml")
			w.Header().Set("Odata-Version", "4.0")
			w.WriteHeader(200)
			model := s.Model
			if model == nil {
				model, err = s.ModelHandler.Metadata(s, scope)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte(err.Error()))
					return
				}
			}
			b, _ := model.ToXml()
			_, _ = w.Write([]byte(xml.Header))
			_, _ = w.Write(b)
			return
		}
		handler, ok := s.Handlers[entity]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(nil)
			return
		}

		lo := &ListOptions{}
		_ = r.ParseForm()
		lo.Top = tox.ToInt(r.Form.Get("$top"))
		lo.Skip = tox.ToInt(r.Form.Get("$skip"))
		lo.SkipToken = r.Form.Get("$skiptoken")
		lo.Expand = r.Form.Get("$expand")
		lo.ParseFilter(r.Form.Get("$filter"))

		res := result.New(s.Fqdn + "/" + scope + "/$metadata#" + entity)
		arr, skipToken, err := handler.List(s, scope, lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		res.Append(arr)
		if len(skipToken) != 0 {
			res.SetNextLink(s.Fqdn + "/" + scope + "/" + entity + "?$skiptoken=" + skipToken)
		}

		b, _ := res.ToJson()
		logDebug(r, nil, "response sent (%d bytes).", len(b))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Odata-Version", "4.0")
		w.WriteHeader(200)
		_, _ = w.Write(b)
		return
	}
}
