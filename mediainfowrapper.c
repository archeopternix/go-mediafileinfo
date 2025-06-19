#include "mediainfowrapper.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

MediaInfo *Get_media_info(const char *filename) {
    AVFormatContext *fmt_ctx = NULL;
    if (avformat_open_input(&fmt_ctx, filename, NULL, NULL) < 0)
        return NULL;
    if (avformat_find_stream_info(fmt_ctx, NULL) < 0) {
        avformat_close_input(&fmt_ctx);
        return NULL;
    }

    MediaInfo *info = (MediaInfo *)calloc(1, sizeof(MediaInfo));
    strncpy(info->filename, filename, sizeof(info->filename) - 1);
    info->duration = fmt_ctx->duration;
    info->nb_streams = fmt_ctx->nb_streams;
    strncpy(info->format_name, fmt_ctx->iformat->name, sizeof(info->format_name) - 1);
    strncpy(info->format_long_name, fmt_ctx->iformat->long_name, sizeof(info->format_long_name) - 1);

    info->streams = (StreamInfo *)calloc(info->nb_streams, sizeof(StreamInfo));
    for (unsigned int i = 0; i < info->nb_streams; i++) {
        AVStream *stream = fmt_ctx->streams[i];
        AVCodecParameters *codecpar = stream->codecpar;
        StreamInfo *si = &info->streams[i];
        si->stream_index = i;
        si->codec_type = codecpar->codec_type;
        strncpy(si->codec_name, avcodec_get_name(codecpar->codec_id), sizeof(si->codec_name) - 1);
        if (codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
            si->width = codecpar->width;
            si->height = codecpar->height;
            si->field_order = codecpar->field_order;
        } else {
            si->width = 0;
            si->height = 0;
            si->field_order = -1;
        }
        if (codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
            si->sample_rate = codecpar->sample_rate;
            si->channels = codecpar->channels;
        } else {
            si->sample_rate = 0;
            si->channels = 0;
        }
        si->codecpar_copy = avcodec_parameters_alloc();
        avcodec_parameters_copy(si->codecpar_copy, codecpar);
    }

    avformat_close_input(&fmt_ctx);
    return info;
}

void Free_media_info(MediaInfo *info) {
    if (!info) return;
    if (info->streams) {
        for (unsigned int i = 0; i < info->nb_streams; i++) {
            if (info->streams[i].codecpar_copy)
                avcodec_parameters_free(&info->streams[i].codecpar_copy);
        }
        free(info->streams);
    }
    free(info);
}


const char *Field_order_to_str(int field_order) {
    switch (field_order) {
        case 0: return "Unknown";
        case 1: return "Progressive";
        case 2: return "TT (Top coded first, top displayed first)";
        case 3: return "BB (Bottom coded first, bottom displayed first)";
        case 4: return "TB (Top coded first, bottom displayed first)";
        case 5: return "BT (Bottom coded first, top displayed first)";
        default: return "N/A";
    }
}