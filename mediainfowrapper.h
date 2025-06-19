#ifndef MEDIAINFOWRAPPER_H
#define MEDIAINFOWRAPPER_H

#include <libavformat/avformat.h>

#ifdef __cplusplus
extern "C" {
#endif

AVFormatContext* Get_avformat_context(const char* filename);
void Free_avformat_context(AVFormatContext* ctx);

// Optionally, helper functions to retrieve fields
int Get_avstreams(AVFormatContext* ctx);
int64_t Get_duration(AVFormatContext* ctx);
const char* Get_format_name(AVFormatContext* ctx);

#ifdef __cplusplus
}
#endif

#endif // MEDIAINFOWRAPPER_H