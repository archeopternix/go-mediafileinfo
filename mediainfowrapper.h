#ifndef FFMPEG_MEDIA_INFO_H
#define FFMPEG_MEDIA_INFO_H

#include <libavformat/avformat.h>

typedef struct {
    int stream_index;
    enum AVMediaType codec_type;
    char codec_name[64];
    int width, height;        // For video
    int sample_rate;          // For audio
    int channels;             // For audio
    int field_order;          // For video: AVFieldOrder, -1 if not applicable
    AVCodecParameters *codecpar_copy;
} StreamInfo;

typedef struct {
    char filename[512];
    int64_t duration;
    unsigned int nb_streams;
    char format_name[128];
    char format_long_name[128];
    StreamInfo *streams;
} MediaInfo;

MediaInfo *get_media_info(const char *filename);
void free_media_info(MediaInfo *info);

const char *field_order_to_str(int field_order);

#endif // FFMPEG_MEDIA_INFO_H