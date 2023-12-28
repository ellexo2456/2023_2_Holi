// // Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.
package domain

//
//import (
//	json "encoding/json"
//	easyjson "github.com/mailru/easyjson"
//	jlexer "github.com/mailru/easyjson/jlexer"
//	jwriter "github.com/mailru/easyjson/jwriter"
//)
//
//// suppress unused package warning
//var (
//	_ *json.RawMessage
//	_ *jlexer.Lexer
//	_ *jwriter.Writer
//	_ easyjson.Marshaler
//)
//
//func easyjson6ff3ac1dDecode20232HoliEa(in *jlexer.Lexer, out *Response) {
//	isTopLevel := in.IsStart()
//	if in.IsNull() {
//		if isTopLevel {
//			in.Consumed()
//		}
//		in.Skip()
//		return
//	}
//	in.Delim('{')
//	for !in.IsDelim('}') {
//		key := in.UnsafeFieldName(false)
//		in.WantColon()
//		if in.IsNull() {
//			in.Skip()
//			in.WantComma()
//			continue
//		}
//		switch key {
//		case "body":
//			if m, ok := out.Body.(easyjson.Unmarshaler); ok {
//				m.UnmarshalEasyJSON(in)
//			} else if m, ok := out.Body.(json.Unmarshaler); ok {
//				_ = m.UnmarshalJSON(in.Raw())
//			} else {
//				out.Body = in.Interface()
//			}
//		case "err":
//			out.Err = string(in.String())
//		default:
//			in.SkipRecursive()
//		}
//		in.WantComma()
//	}
//	in.Delim('}')
//	if isTopLevel {
//		in.Consumed()
//	}
//}
//func easyjson6ff3ac1dEncode20232HoliEa(out *jwriter.Writer, in Response) {
//	out.RawByte('{')
//	first := true
//	_ = first
//	if in.Body != nil {
//		const prefix string = ",\"body\":"
//		first = false
//		out.RawString(prefix[1:])
//		if m, ok := in.Body.(easyjson.Marshaler); ok {
//			m.MarshalEasyJSON(out)
//		} else if m, ok := in.Body.(json.Marshaler); ok {
//			out.Raw(m.MarshalJSON())
//		} else {
//			out.Raw(json.Marshal(in.Body))
//		}
//	}
//	if in.Err != "" {
//		const prefix string = ",\"err\":"
//		if first {
//			first = false
//			out.RawString(prefix[1:])
//		} else {
//			out.RawString(prefix)
//		}
//		out.String(string(in.Err))
//	}
//	out.RawByte('}')
//}
//
//// MarshalJSON supports json.Marshaler interface
//func (v Response) MarshalJSON() ([]byte, error) {
//	w := jwriter.Writer{}
//	easyjson6ff3ac1dEncode20232HoliEa(&w, v)
//	return w.Buffer.BuildBytes(), w.Error
//}
//
//// MarshalEasyJSON supports easyjson.Marshaler interface
//func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
//	easyjson6ff3ac1dEncode20232HoliEa(w, v)
//}
//
//// UnmarshalJSON supports json.Unmarshaler interface
//func (v *Response) UnmarshalJSON(data []byte) error {
//	r := jlexer.Lexer{Data: data}
//	easyjson6ff3ac1dDecode20232HoliEa(&r, v)
//	return r.Error()
//}
//
//// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
//func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
//	easyjson6ff3ac1dDecode20232HoliEa(l, v)
//}
