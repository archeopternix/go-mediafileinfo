#include "mediainfowrapper.h"
#include <libavformat/avformat.h>
#include <libavutil/avutil.h>
#include <stdlib.h>
#include <string.h>

// Returns a pointer to AVFormatContext or NULL on error
AVFormatContext* Get_avformat_context(const char* filename) {
    AVFormatContext* ctx = NULL;
    if (avformat_open_input(&ctx, filename, NULL, NULL) < 0)
        return NULL;
    if (avformat_find_stream_info(ctx, NULL) < 0) {
        avformat_close_input(&ctx);
        return NULL;
    }
    return ctx;
}

// Frees the AVFormatContext
void Free_avformat_context(AVFormatContext* ctx) {
    if (ctx) {
        avformat_close_input(&ctx);
    }
}

// Helper: number of streams
int Get_stream_count(AVFormatContext* ctx) {
    return ctx ? ctx->nb_streams : -1;
}

// Helper: duration
int64_t Get_duration(AVFormatContext* ctx) {
    return ctx ? ctx->duration : -1;
}

// Helper: format name
const char* Get_format_name(AVFormatContext* ctx) {
    return (ctx && ctx->iformat && ctx->iformat->name) ? ctx->iformat->name : "";
}

// Gibt einen Zeiger auf den AVStream mit dem angegebenen Index zurück.
// Gibt NULL zurück, falls der Index ungültig ist.
AVStream* Get_stream_by_index(AVFormatContext *fmt_ctx, int index) {
    if (!fmt_ctx || index < 0 || index >= fmt_ctx->nb_streams) {
        return NULL;
    }
    return fmt_ctx->streams[index];
}