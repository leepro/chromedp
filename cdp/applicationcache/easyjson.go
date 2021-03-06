// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package applicationcache

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache(in *jlexer.Lexer, out *Resource) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
		case "size":
			out.Size = int64(in.Int64())
		case "type":
			out.Type = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache(out *jwriter.Writer, in Resource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.URL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"url\":")
		out.String(string(in.URL))
	}
	if in.Size != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"size\":")
		out.Int64(int64(in.Size))
	}
	if in.Type != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type\":")
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Resource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Resource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Resource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Resource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache1(in *jlexer.Lexer, out *GetManifestForFrameReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestURL":
			out.ManifestURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache1(out *jwriter.Writer, in GetManifestForFrameReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ManifestURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"manifestURL\":")
		out.String(string(in.ManifestURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetManifestForFrameReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetManifestForFrameReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetManifestForFrameReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetManifestForFrameReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache2(in *jlexer.Lexer, out *GetManifestForFrameParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache2(out *jwriter.Writer, in GetManifestForFrameParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"frameId\":")
	out.String(string(in.FrameID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetManifestForFrameParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetManifestForFrameParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetManifestForFrameParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetManifestForFrameParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache3(in *jlexer.Lexer, out *GetFramesWithManifestsReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameIds":
			if in.IsNull() {
				in.Skip()
				out.FrameIds = nil
			} else {
				in.Delim('[')
				if out.FrameIds == nil {
					if !in.IsDelim(']') {
						out.FrameIds = make([]*FrameWithManifest, 0, 8)
					} else {
						out.FrameIds = []*FrameWithManifest{}
					}
				} else {
					out.FrameIds = (out.FrameIds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *FrameWithManifest
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(FrameWithManifest)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.FrameIds = append(out.FrameIds, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache3(out *jwriter.Writer, in GetFramesWithManifestsReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.FrameIds) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"frameIds\":")
		if in.FrameIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.FrameIds {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFramesWithManifestsReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFramesWithManifestsReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFramesWithManifestsReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFramesWithManifestsReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache4(in *jlexer.Lexer, out *GetFramesWithManifestsParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache4(out *jwriter.Writer, in GetFramesWithManifestsParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFramesWithManifestsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFramesWithManifestsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFramesWithManifestsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFramesWithManifestsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache5(in *jlexer.Lexer, out *GetApplicationCacheForFrameReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "applicationCache":
			if in.IsNull() {
				in.Skip()
				out.ApplicationCache = nil
			} else {
				if out.ApplicationCache == nil {
					out.ApplicationCache = new(ApplicationCache)
				}
				(*out.ApplicationCache).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache5(out *jwriter.Writer, in GetApplicationCacheForFrameReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ApplicationCache != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"applicationCache\":")
		if in.ApplicationCache == nil {
			out.RawString("null")
		} else {
			(*in.ApplicationCache).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetApplicationCacheForFrameReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetApplicationCacheForFrameReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetApplicationCacheForFrameReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetApplicationCacheForFrameReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache6(in *jlexer.Lexer, out *GetApplicationCacheForFrameParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache6(out *jwriter.Writer, in GetApplicationCacheForFrameParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"frameId\":")
	out.String(string(in.FrameID))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetApplicationCacheForFrameParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetApplicationCacheForFrameParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetApplicationCacheForFrameParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetApplicationCacheForFrameParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache7(in *jlexer.Lexer, out *FrameWithManifest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "manifestURL":
			out.ManifestURL = string(in.String())
		case "status":
			out.Status = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache7(out *jwriter.Writer, in FrameWithManifest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FrameID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"frameId\":")
		out.String(string(in.FrameID))
	}
	if in.ManifestURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"manifestURL\":")
		out.String(string(in.ManifestURL))
	}
	if in.Status != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.Int64(int64(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FrameWithManifest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FrameWithManifest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FrameWithManifest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FrameWithManifest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache8(in *jlexer.Lexer, out *EventNetworkStateUpdated) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "isNowOnline":
			out.IsNowOnline = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache8(out *jwriter.Writer, in EventNetworkStateUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsNowOnline {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isNowOnline\":")
		out.Bool(bool(in.IsNowOnline))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventNetworkStateUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventNetworkStateUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventNetworkStateUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventNetworkStateUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache8(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache9(in *jlexer.Lexer, out *EventApplicationCacheStatusUpdated) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "manifestURL":
			out.ManifestURL = string(in.String())
		case "status":
			out.Status = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache9(out *jwriter.Writer, in EventApplicationCacheStatusUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FrameID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"frameId\":")
		out.String(string(in.FrameID))
	}
	if in.ManifestURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"manifestURL\":")
		out.String(string(in.ManifestURL))
	}
	if in.Status != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.Int64(int64(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventApplicationCacheStatusUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventApplicationCacheStatusUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventApplicationCacheStatusUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventApplicationCacheStatusUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache9(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache10(in *jlexer.Lexer, out *EnableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache10(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache10(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache11(in *jlexer.Lexer, out *ApplicationCache) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manifestURL":
			out.ManifestURL = string(in.String())
		case "size":
			out.Size = float64(in.Float64())
		case "creationTime":
			out.CreationTime = float64(in.Float64())
		case "updateTime":
			out.UpdateTime = float64(in.Float64())
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]*Resource, 0, 8)
					} else {
						out.Resources = []*Resource{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Resource
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Resource)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Resources = append(out.Resources, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache11(out *jwriter.Writer, in ApplicationCache) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ManifestURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"manifestURL\":")
		out.String(string(in.ManifestURL))
	}
	if in.Size != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"size\":")
		out.Float64(float64(in.Size))
	}
	if in.CreationTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"creationTime\":")
		out.Float64(float64(in.CreationTime))
	}
	if in.UpdateTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"updateTime\":")
		out.Float64(float64(in.UpdateTime))
	}
	if len(in.Resources) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"resources\":")
		if in.Resources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Resources {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ApplicationCache) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ApplicationCache) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpApplicationcache11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ApplicationCache) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ApplicationCache) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpApplicationcache11(l, v)
}
