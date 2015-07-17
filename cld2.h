#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

const char* DetectLang(char *data, int length);

typedef struct {
  const char* code;
  const char* name;
  int   percent;
  bool is_reliable;
} LanguageDetection;

const LanguageDetection GetLanguageDetection(char *data, int length);

#ifdef __cplusplus
}
#endif
