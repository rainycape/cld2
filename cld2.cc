#include <cstddef>
#include <string.h>
#include <stdio.h>
#include <string>

#include "compact_lang_det.h"
#include "cld2.h"

const char* DetectLang(char *data, int length) {

    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    bool allow_extended_lang = true;
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data, 
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    return CLD2::LanguageCode(summary_lang);
}

const LanguageDetection GetLanguageDetection(char *data, int length) {
    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data,
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    LanguageDetection lang_det = {
        CLD2::LanguageCode(language3[0]),
        CLD2::LanguageName(language3[0]),
        percent3[0],
        is_reliable
    };
    return lang_det;
}
