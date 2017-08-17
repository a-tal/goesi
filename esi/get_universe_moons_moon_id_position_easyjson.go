// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseMoonsMoonIdPositionList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseMoonsMoonIdPositionList, 0, 5)
			} else {
				*out = GetUniverseMoonsMoonIdPositionList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseMoonsMoonIdPosition
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseMoonsMoonIdPositionList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseMoonsMoonIdPositionList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseMoonsMoonIdPositionList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdPositionList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdPositionList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseMoonsMoonIdPosition) {
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
		case "x":
			out.X = float32(in.Float32())
		case "y":
			out.Y = float32(in.Float32())
		case "z":
			out.Z = float32(in.Float32())
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
func easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseMoonsMoonIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Float32(float32(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Float32(float32(in.Y))
	}
	if in.Z != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"z\":")
		out.Float32(float32(in.Z))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseMoonsMoonIdPosition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseMoonsMoonIdPosition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ee63b03EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdPosition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseMoonsMoonIdPosition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ee63b03DecodeGithubComAntihaxGoesiEsi1(l, v)
}