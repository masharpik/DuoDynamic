// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package photo

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

func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(in *jlexer.Lexer, out *VoiceResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "qid":
			out.Qid = string(in.String())
		case "result":
			easyjson6a975c40Decode(in, &out.Result)
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(out *jwriter.Writer, in VoiceResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"qid\":"
		out.RawString(prefix[1:])
		out.String(string(in.Qid))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		easyjson6a975c40Encode(out, in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoiceResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VoiceResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoiceResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VoiceResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto(l, v)
}
func easyjson6a975c40Decode(in *jlexer.Lexer, out *struct {
	Texts []struct {
		Text           string  `json:"text"`
		Confidence     float32 `json:"confidence"`
		PunctuatedText string  `json:"punctuated_text"`
	} `json:"texts"`
	PhraseId string `json:"phrase_id"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "texts":
			if in.IsNull() {
				in.Skip()
				out.Texts = nil
			} else {
				in.Delim('[')
				if out.Texts == nil {
					if !in.IsDelim(']') {
						out.Texts = make([]struct {
							Text           string  `json:"text"`
							Confidence     float32 `json:"confidence"`
							PunctuatedText string  `json:"punctuated_text"`
						}, 0, 1)
					} else {
						out.Texts = []struct {
							Text           string  `json:"text"`
							Confidence     float32 `json:"confidence"`
							PunctuatedText string  `json:"punctuated_text"`
						}{}
					}
				} else {
					out.Texts = (out.Texts)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						Text           string  `json:"text"`
						Confidence     float32 `json:"confidence"`
						PunctuatedText string  `json:"punctuated_text"`
					}
					easyjson6a975c40Decode1(in, &v1)
					out.Texts = append(out.Texts, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "phrase_id":
			out.PhraseId = string(in.String())
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
func easyjson6a975c40Encode(out *jwriter.Writer, in struct {
	Texts []struct {
		Text           string  `json:"text"`
		Confidence     float32 `json:"confidence"`
		PunctuatedText string  `json:"punctuated_text"`
	} `json:"texts"`
	PhraseId string `json:"phrase_id"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"texts\":"
		out.RawString(prefix[1:])
		if in.Texts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Texts {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson6a975c40Encode1(out, v3)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"phrase_id\":"
		out.RawString(prefix)
		out.String(string(in.PhraseId))
	}
	out.RawByte('}')
}
func easyjson6a975c40Decode1(in *jlexer.Lexer, out *struct {
	Text           string  `json:"text"`
	Confidence     float32 `json:"confidence"`
	PunctuatedText string  `json:"punctuated_text"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "text":
			out.Text = string(in.String())
		case "confidence":
			out.Confidence = float32(in.Float32())
		case "punctuated_text":
			out.PunctuatedText = string(in.String())
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
func easyjson6a975c40Encode1(out *jwriter.Writer, in struct {
	Text           string  `json:"text"`
	Confidence     float32 `json:"confidence"`
	PunctuatedText string  `json:"punctuated_text"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"confidence\":"
		out.RawString(prefix)
		out.Float32(float32(in.Confidence))
	}
	{
		const prefix string = ",\"punctuated_text\":"
		out.RawString(prefix)
		out.String(string(in.PunctuatedText))
	}
	out.RawByte('}')
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(in *jlexer.Lexer, out *ResponseUploadFile) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "error":
			out.Error = string(in.String())
		case "body":
			easyjson6a975c40Decode2(in, &out.Body)
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(out *jwriter.Writer, in ResponseUploadFile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		easyjson6a975c40Encode2(out, in.Body)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseUploadFile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseUploadFile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseUploadFile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseUploadFile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto1(l, v)
}
func easyjson6a975c40Decode2(in *jlexer.Lexer, out *struct {
	PathToFile string `json:"pathToFile"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "pathToFile":
			out.PathToFile = string(in.String())
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
func easyjson6a975c40Encode2(out *jwriter.Writer, in struct {
	PathToFile string `json:"pathToFile"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pathToFile\":"
		out.RawString(prefix[1:])
		out.String(string(in.PathToFile))
	}
	out.RawByte('}')
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(in *jlexer.Lexer, out *AnswerPhoto) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "error":
			out.Error = string(in.String())
		case "body":
			easyjson6a975c40Decode3(in, &out.Body)
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(out *jwriter.Writer, in AnswerPhoto) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		easyjson6a975c40Encode3(out, in.Body)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AnswerPhoto) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AnswerPhoto) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AnswerPhoto) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AnswerPhoto) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgPhoto2(l, v)
}
func easyjson6a975c40Decode3(in *jlexer.Lexer, out *struct {
	PhotoID uint `json:"photoID"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "photoID":
			out.PhotoID = uint(in.Uint())
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
func easyjson6a975c40Encode3(out *jwriter.Writer, in struct {
	PhotoID uint `json:"photoID"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"photoID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.PhotoID))
	}
	out.RawByte('}')
}