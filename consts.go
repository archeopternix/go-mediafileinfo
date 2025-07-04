// Package mediafileinfo
package mediafileinfo

// AVMediaType represents the type of media stream (video, audio, etc). To get a textual representation call String() on the const
//
//go:generate stringer -type=AVMediaType -trimprefix=AVMEDIA_TYPE_
type AVMediaType int

const (
	AVMEDIA_TYPE_UNKNOWN AVMediaType = -1
	AVMEDIA_TYPE_VIDEO   AVMediaType = iota
	AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT
	AVMEDIA_TYPE_NB
)

// AVFieldOrder indicates a progressive video or describes the field order for interlaced video. To get a textual representation call String() on the const
//   - PROGRESSIVE
//   - TT: Top coded_first, top displayed first.
//   - BB: Bottom coded first, bottom displayed first.
//   - TB: Top coded first, bottom displayed first.
//   - BT: Bottom coded first, top displayed first.
//
//go:generate stringer -type=AVFieldOrder -trimprefix=AV_FIELD_
type AVFieldOrder int

const (
	AV_FIELD_UNKNOWN     AVFieldOrder = iota // UNKNOWN
	AV_FIELD_PROGRESSIVE                     // PROGRESSIVE
	AV_FIELD_TT                              // Top coded_first, top displayed first.
	AV_FIELD_BB                              // Bottom coded first, bottom displayed first.
	AV_FIELD_TB                              // Top coded first, bottom displayed first.
	AV_FIELD_BT                              // Bottom coded first, top displayed first.
)

// CodecID identifies a codec by its integer ID. To get a textual representation call String() on the const
//
//go:generate stringer -type=CodecID -trimprefix=CODEC_ID_
type CodecID int

const (
	// Video codecs
	CODEC_ID_NONE CodecID = iota
	CODEC_ID_MPEG1VIDEO
	CODEC_ID_MPEG2VIDEO
	CODEC_ID_H261
	CODEC_ID_H263
	CODEC_ID_RV10
	CODEC_ID_RV20
	CODEC_ID_MJPEG
	CODEC_ID_MJPEGB
	CODEC_ID_LJPEG
	CODEC_ID_SP5X
	CODEC_ID_JPEGLS
	CODEC_ID_MPEG4
	CODEC_ID_RAWVIDEO
	CODEC_ID_MSMPEG4V1
	CODEC_ID_MSMPEG4V2
	CODEC_ID_MSMPEG4V3
	CODEC_ID_WMV1
	CODEC_ID_WMV2
	CODEC_ID_H263P
	CODEC_ID_H263I
	CODEC_ID_FLV1
	CODEC_ID_SVQ1
	CODEC_ID_SVQ3
	CODEC_ID_DVVIDEO
	CODEC_ID_HUFFYUV
	CODEC_ID_CYUV
	CODEC_ID_H264
	CODEC_ID_INDEO3
	CODEC_ID_VP3
	CODEC_ID_THEORA
	CODEC_ID_ASV1
	CODEC_ID_ASV2
	CODEC_ID_FFV1
	CODEC_ID_4XM
	CODEC_ID_VCR1
	CODEC_ID_CLJR
	CODEC_ID_MDEC
	CODEC_ID_ROQ
	CODEC_ID_INTERPLAY_VIDEO
	CODEC_ID_XAN_WC3
	CODEC_ID_XAN_WC4
	CODEC_ID_RPZA
	CODEC_ID_CINEPAK
	CODEC_ID_WS_VQA
	CODEC_ID_MSRLE
	CODEC_ID_MSVIDEO1
	CODEC_ID_IDCIN
	CODEC_ID_8BPS
	CODEC_ID_SMC
	CODEC_ID_FLIC
	CODEC_ID_TRUEMOTION1
	CODEC_ID_VMDVIDEO
	CODEC_ID_MSZH
	CODEC_ID_ZLIB
	CODEC_ID_QTRLE
	CODEC_ID_TSCC
	CODEC_ID_ULTI
	CODEC_ID_QDRAW
	CODEC_ID_VIXL
	CODEC_ID_QPEG
	CODEC_ID_PNG
	CODEC_ID_PPM
	CODEC_ID_PBM
	CODEC_ID_PGM
	CODEC_ID_PGMYUV
	CODEC_ID_PAM
	CODEC_ID_FFVHUFF
	CODEC_ID_RV30
	CODEC_ID_RV40
	CODEC_ID_VC1
	CODEC_ID_WMV3
	CODEC_ID_LOCO
	CODEC_ID_WNV1
	CODEC_ID_AASC
	CODEC_ID_INDEO2
	CODEC_ID_FRAPS
	CODEC_ID_TRUEMOTION2
	CODEC_ID_BMP
	CODEC_ID_CSCD
	CODEC_ID_MMVIDEO
	CODEC_ID_ZMBV
	CODEC_ID_AVS
	CODEC_ID_SMACKVIDEO
	CODEC_ID_NUV
	CODEC_ID_KMVC
	CODEC_ID_FLASHSV
	CODEC_ID_CAVS
	CODEC_ID_JPEG2000
	CODEC_ID_VMNC
	CODEC_ID_VP5
	CODEC_ID_VP6
	CODEC_ID_VP6F
	CODEC_ID_TARGA
	CODEC_ID_DSICINVIDEO
	CODEC_ID_TIERTEXSEQVIDEO
	CODEC_ID_TIFF
	CODEC_ID_GIF
	CODEC_ID_DXA
	CODEC_ID_DNXHD
	CODEC_ID_THP
	CODEC_ID_SGI
	CODEC_ID_C93
	CODEC_ID_BETHSOFTVID
	CODEC_ID_PTX
	CODEC_ID_TXD
	CODEC_ID_VP6A
	CODEC_ID_AMV
	CODEC_ID_VB
	CODEC_ID_PCX
	CODEC_ID_SUNRAST
	CODEC_ID_INDEO4
	CODEC_ID_INDEO5
	CODEC_ID_MIMIC
	CODEC_ID_RL2
	CODEC_ID_ESCAPE124
	CODEC_ID_DIRAC
	CODEC_ID_BFI
	CODEC_ID_CMV
	CODEC_ID_MOTIONPIXELS
	CODEC_ID_TGV
	CODEC_ID_TGQ
	CODEC_ID_TQI
	CODEC_ID_AURA
	CODEC_ID_AURA2
	CODEC_ID_V210X
	CODEC_ID_TMV
	CODEC_ID_V210
	CODEC_ID_DPX
	CODEC_ID_MAD
	CODEC_ID_FRWU
	CODEC_ID_FLASHSV2
	CODEC_ID_CDGRAPHICS
	CODEC_ID_R210
	CODEC_ID_ANM
	CODEC_ID_BINKVIDEO
	CODEC_ID_IFF_ILBM
	CODEC_ID_KGV1
	CODEC_ID_YOP
	CODEC_ID_VP8
	CODEC_ID_PICTOR
	CODEC_ID_ANSI
	CODEC_ID_A64_MULTI
	CODEC_ID_A64_MULTI5
	CODEC_ID_R10K
	CODEC_ID_MXPEG
	CODEC_ID_LAGARITH
	CODEC_ID_PRORES
	CODEC_ID_JV
	CODEC_ID_DFA
	CODEC_ID_WMV3IMAGE
	CODEC_ID_VC1IMAGE
	CODEC_ID_UTVIDEO
	CODEC_ID_BMV_VIDEO
	CODEC_ID_VBLE
	CODEC_ID_DXTORY
	CODEC_ID_XWD
	CODEC_ID_CDXL
	CODEC_ID_XBM
	CODEC_ID_ZEROCODEC
	CODEC_ID_MSS1
	CODEC_ID_MSA1
	CODEC_ID_TSCC2
	CODEC_ID_MTS2
	CODEC_ID_CLLC
	CODEC_ID_MSS2
	CODEC_ID_VP9
	CODEC_ID_AIC
	CODEC_ID_ESCAPE130
	CODEC_ID_G2M
	CODEC_ID_WEBP
	CODEC_ID_HNM4_VIDEO
	CODEC_ID_HEVC
	CODEC_ID_FIC
	CODEC_ID_ALIAS_PIX
	CODEC_ID_BRENDER_PIX
	CODEC_ID_PAF_VIDEO
	CODEC_ID_EXR
	CODEC_ID_VP7
	CODEC_ID_SANM
	CODEC_ID_SGIRLE
	CODEC_ID_MVC1
	CODEC_ID_MVC2
	CODEC_ID_HQX
	CODEC_ID_TDSC
	CODEC_ID_HQ_HQA
	CODEC_ID_HAP
	CODEC_ID_DDS
	CODEC_ID_DXV
	CODEC_ID_SCREENPRESSO
	CODEC_ID_RSCC
	CODEC_ID_AVS2
	CODEC_ID_PGX
	CODEC_ID_AVS3
	CODEC_ID_MSP2
	CODEC_ID_VVC
	CODEC_ID_Y41P
	CODEC_ID_AVRP
	CODEC_ID_012V
	CODEC_ID_AVUI
	CODEC_ID_TARGA_Y216
	CODEC_ID_YUV4
	CODEC_ID_AVRN
	CODEC_ID_CPIA
	CODEC_ID_XFACE
	CODEC_ID_SNOW
	CODEC_ID_SMVJPEG
	CODEC_ID_APNG
	CODEC_ID_DAALA
	CODEC_ID_CFHD
	CODEC_ID_TRUEMOTION2RT
	CODEC_ID_M101
	CODEC_ID_MAGICYUV
	CODEC_ID_SHEERVIDEO
	CODEC_ID_YLC
	CODEC_ID_PSD
	CODEC_ID_PIXLET
	CODEC_ID_SPEEDHQ
	CODEC_ID_FMVC
	CODEC_ID_SCPR
	CODEC_ID_CLEARVIDEO
	CODEC_ID_XPM
	CODEC_ID_AV1
	CODEC_ID_BITPACKED
	CODEC_ID_MSCC
	CODEC_ID_SRGC
	CODEC_ID_SVG
	CODEC_ID_GDV
	CODEC_ID_FITS
	CODEC_ID_IMM4
	CODEC_ID_PROSUMER
	CODEC_ID_MWSC
	CODEC_ID_WCMV
	CODEC_ID_RASC
	CODEC_ID_HYMT
	CODEC_ID_ARBC
	CODEC_ID_AGM
	CODEC_ID_LSCR
	CODEC_ID_VP4
	CODEC_ID_IMM5
	CODEC_ID_MVDV
	CODEC_ID_MVHA
	CODEC_ID_CDTOONS
	CODEC_ID_MV30
	CODEC_ID_NOTCHLC
	CODEC_ID_PFM
	CODEC_ID_MOBICLIP
	CODEC_ID_PHOTOCD
	CODEC_ID_IPU
	CODEC_ID_ARGO
	CODEC_ID_CRI
	CODEC_ID_SIMBIOSIS_IMX
	CODEC_ID_SGA_VIDEO
	CODEC_ID_GEM
	CODEC_ID_VBN
	CODEC_ID_JPEGXL
	CODEC_ID_QOI
	CODEC_ID_PHM
	CODEC_ID_RADIANCE_HDR
	CODEC_ID_WBMP
	CODEC_ID_MEDIA100
	CODEC_ID_VQC
	CODEC_ID_PDV
	CODEC_ID_EVC
	CODEC_ID_RTV1
	CODEC_ID_VMIX
	CODEC_ID_LEAD
	CODEC_ID_DNXUC
	CODEC_ID_RV60
	CODEC_ID_JPEGXL_ANIM
	CODEC_ID_APV

	// Audio codecs (start at 0x10000)
	CODEC_ID_FIRST_AUDIO CodecID = 0x10000
	CODEC_ID_PCM_S16LE
	CODEC_ID_PCM_S16BE
	CODEC_ID_PCM_U16LE
	CODEC_ID_PCM_U16BE
	CODEC_ID_PCM_S8
	CODEC_ID_PCM_U8
	CODEC_ID_PCM_MULAW
	CODEC_ID_PCM_ALAW
	CODEC_ID_PCM_S32LE
	CODEC_ID_PCM_S32BE
	CODEC_ID_PCM_U32LE
	CODEC_ID_PCM_U32BE
	CODEC_ID_PCM_S24LE
	CODEC_ID_PCM_S24BE
	CODEC_ID_PCM_U24LE
	CODEC_ID_PCM_U24BE
	CODEC_ID_PCM_S24DAUD
	CODEC_ID_PCM_ZORK
	CODEC_ID_PCM_S16LE_PLANAR
	CODEC_ID_PCM_DVD
	CODEC_ID_PCM_F32BE
	CODEC_ID_PCM_F32LE
	CODEC_ID_PCM_F64BE
	CODEC_ID_PCM_F64LE
	CODEC_ID_PCM_BLURAY
	CODEC_ID_PCM_LXF
	CODEC_ID_S302M
	CODEC_ID_PCM_S8_PLANAR
	CODEC_ID_PCM_S24LE_PLANAR
	CODEC_ID_PCM_S32LE_PLANAR
	CODEC_ID_PCM_S16BE_PLANAR
	CODEC_ID_PCM_S64LE
	CODEC_ID_PCM_S64BE
	CODEC_ID_PCM_F16LE
	CODEC_ID_PCM_F24LE
	CODEC_ID_PCM_VIDC
	CODEC_ID_PCM_SGA

	// ADPCM codecs start at 0x11000
	CODEC_ID_ADPCM_IMA_QT CodecID = 0x11000
	CODEC_ID_ADPCM_IMA_WAV
	CODEC_ID_ADPCM_IMA_DK3
	CODEC_ID_ADPCM_IMA_DK4
	CODEC_ID_ADPCM_IMA_WS
	CODEC_ID_ADPCM_IMA_SMJPEG
	CODEC_ID_ADPCM_MS
	CODEC_ID_ADPCM_4XM
	CODEC_ID_ADPCM_XA
	CODEC_ID_ADPCM_ADX
	CODEC_ID_ADPCM_EA
	CODEC_ID_ADPCM_G726
	CODEC_ID_ADPCM_CT
	CODEC_ID_ADPCM_SWF
	CODEC_ID_ADPCM_YAMAHA
	CODEC_ID_ADPCM_SBPRO_4
	CODEC_ID_ADPCM_SBPRO_3
	CODEC_ID_ADPCM_SBPRO_2
	CODEC_ID_ADPCM_THP
	CODEC_ID_ADPCM_IMA_AMV
	CODEC_ID_ADPCM_EA_R1
	CODEC_ID_ADPCM_EA_R3
	CODEC_ID_ADPCM_EA_R2
	CODEC_ID_ADPCM_IMA_EA_SEAD
	CODEC_ID_ADPCM_IMA_EA_EACS
	CODEC_ID_ADPCM_EA_XAS
	CODEC_ID_ADPCM_EA_MAXIS_XA
	CODEC_ID_ADPCM_IMA_ISS
	CODEC_ID_ADPCM_G722
	CODEC_ID_ADPCM_IMA_APC
	CODEC_ID_ADPCM_VIMA
	CODEC_ID_ADPCM_AFC
	CODEC_ID_ADPCM_IMA_OKI
	CODEC_ID_ADPCM_DTK
	CODEC_ID_ADPCM_IMA_RAD
	CODEC_ID_ADPCM_G726LE
	CODEC_ID_ADPCM_THP_LE
	CODEC_ID_ADPCM_PSX
	CODEC_ID_ADPCM_AICA
	CODEC_ID_ADPCM_IMA_DAT4
	CODEC_ID_ADPCM_MTAF
	CODEC_ID_ADPCM_AGM
	CODEC_ID_ADPCM_ARGO
	CODEC_ID_ADPCM_IMA_SSI
	CODEC_ID_ADPCM_ZORK
	CODEC_ID_ADPCM_IMA_APM
	CODEC_ID_ADPCM_IMA_ALP
	CODEC_ID_ADPCM_IMA_MTF
	CODEC_ID_ADPCM_IMA_CUNNING
	CODEC_ID_ADPCM_IMA_MOFLEX
	CODEC_ID_ADPCM_IMA_ACORN
	CODEC_ID_ADPCM_XMD
	CODEC_ID_ADPCM_IMA_XBOX

	// AMR codecs start at 0x12000
	CODEC_ID_AMR_NB CodecID = 0x12000
	CODEC_ID_AMR_WB

	// RealAudio codecs start at 0x13000
	CODEC_ID_RA_144 CodecID = 0x13000
	CODEC_ID_RA_288

	// DPCM codecs start at 0x14000
	CODEC_ID_ROQ_DPCM CodecID = 0x14000
	CODEC_ID_INTERPLAY_DPCM
	CODEC_ID_XAN_DPCM
	CODEC_ID_SOL_DPCM
	CODEC_ID_SDX2_DPCM
	CODEC_ID_GREMLIN_DPCM
	CODEC_ID_DERF_DPCM
	CODEC_ID_WADY_DPCM
	CODEC_ID_CBD2_DPCM

	// audio codecs start at 0x15000
	CODEC_ID_MP2 CodecID = 0x15000
	CODEC_ID_MP3
	CODEC_ID_AAC
	CODEC_ID_AC3
	CODEC_ID_DTS
	CODEC_ID_VORBIS
	CODEC_ID_DVAUDIO
	CODEC_ID_WMAV1
	CODEC_ID_WMAV2
	CODEC_ID_MACE3
	CODEC_ID_MACE6
	CODEC_ID_VMDAUDIO
	CODEC_ID_FLAC
	CODEC_ID_MP3ADU
	CODEC_ID_MP3ON4
	CODEC_ID_SHORTEN
	CODEC_ID_ALAC
	CODEC_ID_WESTWOOD_SND1
	CODEC_ID_GSM
	CODEC_ID_QDM2
	CODEC_ID_COOK
	CODEC_ID_TRUESPEECH
	CODEC_ID_TTA
	CODEC_ID_SMACKAUDIO
	CODEC_ID_QCELP
	CODEC_ID_WAVPACK
	CODEC_ID_DSICINAUDIO
	CODEC_ID_IMC
	CODEC_ID_MUSEPACK7
	CODEC_ID_MLP
	CODEC_ID_GSM_MS
	CODEC_ID_ATRAC3
	CODEC_ID_APE
	CODEC_ID_NELLYMOSER
	CODEC_ID_MUSEPACK8
	CODEC_ID_SPEEX
	CODEC_ID_WMAVOICE
	CODEC_ID_WMAPRO
	CODEC_ID_WMALOSSLESS
	CODEC_ID_ATRAC3P
	CODEC_ID_EAC3
	CODEC_ID_SIPR
	CODEC_ID_MP1
	CODEC_ID_TWINVQ
	CODEC_ID_TRUEHD
	CODEC_ID_MP4ALS
	CODEC_ID_ATRAC1
	CODEC_ID_BINKAUDIO_RDFT
	CODEC_ID_BINKAUDIO_DCT
	CODEC_ID_AAC_LATM
	CODEC_ID_QDMC
	CODEC_ID_CELT
	CODEC_ID_G723_1
	CODEC_ID_G729
	CODEC_ID_8SVX_EXP
	CODEC_ID_8SVX_FIB
	CODEC_ID_BMV_AUDIO
	CODEC_ID_RALF
	CODEC_ID_IAC
	CODEC_ID_ILBC
	CODEC_ID_OPUS
	CODEC_ID_COMFORT_NOISE
	CODEC_ID_TAK
	CODEC_ID_METASOUND
	CODEC_ID_PAF_AUDIO
	CODEC_ID_ON2AVC
	CODEC_ID_DSS_SP
	CODEC_ID_CODEC2
	CODEC_ID_FFWAVESYNTH
	CODEC_ID_SONIC
	CODEC_ID_SONIC_LS
	CODEC_ID_EVRC
	CODEC_ID_SMV
	CODEC_ID_DSD_LSBF
	CODEC_ID_DSD_MSBF
	CODEC_ID_DSD_LSBF_PLANAR
	CODEC_ID_DSD_MSBF_PLANAR
	CODEC_ID_4GV
	CODEC_ID_INTERPLAY_ACM
	CODEC_ID_XMA1
	CODEC_ID_XMA2
	CODEC_ID_DST
	CODEC_ID_ATRAC3AL
	CODEC_ID_ATRAC3PAL
	CODEC_ID_DOLBY_E
	CODEC_ID_APTX
	CODEC_ID_APTX_HD
	CODEC_ID_SBC
	CODEC_ID_ATRAC9
	CODEC_ID_HCOM
	CODEC_ID_ACELP_KELVIN
	CODEC_ID_MPEGH_3D_AUDIO
	CODEC_ID_SIREN
	CODEC_ID_HCA
	CODEC_ID_FASTAUDIO
	CODEC_ID_MSNSIREN
	CODEC_ID_DFPWM
	CODEC_ID_BONK
	CODEC_ID_MISC4
	CODEC_ID_APAC
	CODEC_ID_FTR
	CODEC_ID_WAVARC
	CODEC_ID_RKA
	CODEC_ID_AC4
	CODEC_ID_OSQ
	CODEC_ID_QOA
	CODEC_ID_LC3

	// Subtitle codecs (start at 0x17000)
	CODEC_ID_FIRST_SUBTITLE CodecID = 0x17000
	CODEC_ID_DVD_SUBTITLE
	CODEC_ID_DVB_SUBTITLE
	CODEC_ID_TEXT
	CODEC_ID_XSUB
	CODEC_ID_SSA
	CODEC_ID_MOV_TEXT
	CODEC_ID_HDMV_PGS_SUBTITLE
	CODEC_ID_DVB_TELETEXT
	CODEC_ID_SRT
	CODEC_ID_MICRODVD
	CODEC_ID_EIA_608
	CODEC_ID_JACOSUB
	CODEC_ID_SAMI
	CODEC_ID_REALTEXT
	CODEC_ID_STL
	CODEC_ID_SUBVIEWER1
	CODEC_ID_SUBVIEWER
	CODEC_ID_SUBRIP
	CODEC_ID_WEBVTT
	CODEC_ID_MPL2
	CODEC_ID_VPLAYER
	CODEC_ID_PJS
	CODEC_ID_ASS
	CODEC_ID_HDMV_TEXT_SUBTITLE
	CODEC_ID_TTML
	CODEC_ID_ARIB_CAPTION
	CODEC_ID_IVTV_VBI

	// Other/fake codecs (start at 0x18000)
	CODEC_ID_FIRST_UNKNOWN CodecID = 0x18000
	CODEC_ID_TTF
	CODEC_ID_SCTE_35
	CODEC_ID_EPG
	CODEC_ID_BINTEXT
	CODEC_ID_XBIN
	CODEC_ID_IDF
	CODEC_ID_OTF
	CODEC_ID_SMPTE_KLV
	CODEC_ID_DVD_NAV
	CODEC_ID_TIMED_ID3
	CODEC_ID_BIN_DATA
	CODEC_ID_SMPTE_2038
	CODEC_ID_LCEVC

	// Special values
	CODEC_ID_PROBE   CodecID = 0x19000
	CODEC_ID_MPEG2TS CodecID = 0x20000
	CODEC_ID_MPEG4SYSTEMS
	CODEC_ID_FFMETADATA CodecID = 0x21000
	CODEC_ID_WRAPPED_AVFRAME
	CODEC_ID_VNULL
	CODEC_ID_ANULL
)
