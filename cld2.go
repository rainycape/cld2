// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://code.google.com/p/cld2/.
package cld2

// #include <stdlib.h>
// #include "cld2.h"
import "C"
import "unsafe"

// Detect returns the language code for detected language
// in the given text.
func Detect(text string) string {
	cs := C.CString(text)
	res := C.DetectLang(cs, -1)
	C.free(unsafe.Pointer(cs))
	var lang string
	if res != nil {
		lang = C.GoString(res)
	}
	return lang
}

type LanguageDetection struct {
	Code       string // e.g., "en", "ko", "xx-Brai"
	Name       string // e.g., "ENGLISH", "Korean", "X_Braille"
	Percent    int
	IsReliable bool
}

// GetLanguageDetection returns the LanguageDetection from the
// given body of text; it is useful for cases where more than
// just the language code is needed.
func GetLanguageDetection(text string) LanguageDetection {
	cs := C.CString(text)
	res := C.GetLanguageDetection(cs, -1)
	C.free(unsafe.Pointer(cs))

	return LanguageDetection{
		Code:       C.GoString(res.code),
		Name:       C.GoString(res.name),
		Percent:    int(res.percent),
		IsReliable: (res.is_reliable == true),
	}
}
