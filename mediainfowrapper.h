// Copyright 2025 archeopternix. All rights reserved. MIT license.

#ifndef MEDIAINFOWRAPPER_H
#define MEDIAINFOWRAPPER_H

#include <libavformat/avformat.h>

#ifdef __cplusplus
extern "C" {
#endif

AVFormatContext* Get_avformat_context(const char* filename);
void Free_avformat_context(AVFormatContext* ctx);

// Optionally, helper functions to retrieve fields
int Get_stream_count(AVFormatContext* ctx);
int64_t Get_duration(AVFormatContext* ctx);
const char* Get_format_name(AVFormatContext* ctx);
AVStream* Get_stream_by_index(AVFormatContext *fmt_ctx, int index);

#ifdef __cplusplus
}
#endif

#endif // MEDIAINFOWRAPPER_H