// avwrapper.h
#ifndef AVPARAMS_H
#define AVPARAMS_H

#include <libavcodec/avcodec.h>

AVCodecParameters* get_video_codec_parameters(const char *filename);
void free_codec_parameters(AVCodecParameters* params);

#endif