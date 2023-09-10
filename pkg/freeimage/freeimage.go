package freeimage

import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

var fiLib *c.Lib

// init library by shared library
func InitLib(path string, mod c.LibMode) {
	if fiLib != nil {
		return
	}
	var err error
	fiLib, err = c.NewLib(path, mod)
	if err != nil {
		panic(err)
	}
}

type inArgs []interface{}
type BitMap struct{}      //FI_STRUCT (FIBITMAP) { void *data; };
type MultiBitMap struct{} //FI_STRUCT (FIMULTIBITMAP) { void *data; };
type Handle struct{}      //typedef void* fi_handle;
type Memory struct{}      //FI_STRUCT (FIMEMORY) { void *data; };
type Tag struct{}         //FI_STRUCT (FITAG) { void *data; };
type MetaData struct{}    //FI_STRUCT (FIMETADATA) { void *data; };

// color order is coupled to endianness:
//
// little-endian -> BGR
// big-endian    -> RGB
//
// typedef struct tagRGBQUAD {
// #if FREEIMAGE_COLORORDER == FREEIMAGE_COLORORDER_BGR
//
//	BYTE rgbBlue;
//	BYTE rgbGreen;
//	BYTE rgbRed;
//
// #else
//
//	BYTE rgbRed;
//	BYTE rgbGreen;
//	BYTE rgbBlue;
//
// #endif // FREEIMAGE_COLORORDER
//
//	    BYTE rgbReserved;
//	} RGBQUAD;
type RGBQUAD [4]byte

// typedef struct tagBITMAPINFOHEADER{
//     DWORD biSize;
//     LONG  biWidth;
//     LONG  biHeight;
//     WORD  biPlanes;
//     WORD  biBitCount;
//     DWORD biCompression;
//     DWORD biSizeImage;
//     LONG  biXPelsPerMeter;
//     LONG  biYPelsPerMeter;
//     DWORD biClrUsed;
//     DWORD biClrImportant;
//   } BITMAPINFOHEADER

type BitMapInfoHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}
type BitMapInfo struct {
	Header BitMapInfoHeader
	Colors RGBQUAD
}

//	FI_STRUCT (FIICCPROFILE) {
//	    WORD    flags;    //! info flag
//	    DWORD    size;    //! profile's size measured in bytes
//	    void   *data;    //! points to a block of contiguous memory containing the profile
//	};
type ICCProfile struct {
	Flags uint16         //! info flag
	Size  uint32         //! profile's size measured in bytes
	data  unsafe.Pointer //! points to a block of contiguous memory containing the profile
}

type FREE_IMAGE_TYPE int32
type FREE_IMAGE_FORMAT int32
type FREE_IMAGE_COLOR_TYPE int32
type FREE_IMAGE_QUANTIZE int32
type FREE_IMAGE_DITHER int32
type FREE_IMAGE_TMO int32
type FREE_IMAGE_MDTYPE int32
type FREE_IMAGE_MDMODEL int32
type FREE_IMAGE_JPEG_OPERATION int32
type FREE_IMAGE_FILTER int32
type FREE_IMAGE_COLOR_CHANNEL int32

type SeekOrigin int32

const (
	SEEK_CUR SeekOrigin = 1
	SEEK_END SeekOrigin = 1
	SEEK_SET SeekOrigin = 0
)

const (
	FICC_RGB   FREE_IMAGE_COLOR_CHANNEL = 0 //! Use red green and blue channels
	FICC_RED   FREE_IMAGE_COLOR_CHANNEL = 1 //! Use red channel
	FICC_GREEN FREE_IMAGE_COLOR_CHANNEL = 2 //! Use green channel
	FICC_BLUE  FREE_IMAGE_COLOR_CHANNEL = 3 //! Use blue channel
	FICC_ALPHA FREE_IMAGE_COLOR_CHANNEL = 4 //! Use alpha channel
	FICC_BLACK FREE_IMAGE_COLOR_CHANNEL = 5 //! Use black channel
	FICC_REAL  FREE_IMAGE_COLOR_CHANNEL = 6 //! Complex images: use real part
	FICC_IMAG  FREE_IMAGE_COLOR_CHANNEL = 7 //! Complex images: use imaginary part
	FICC_MAG   FREE_IMAGE_COLOR_CHANNEL = 8 //! Complex images: use magnitude
	FICC_PHASE FREE_IMAGE_COLOR_CHANNEL = 9 //! Complex images: use phase

	FILTER_BOX        FREE_IMAGE_FILTER = 0 //! Box pulse Fourier window 1st order (constant) b-spline
	FILTER_BICUBIC    FREE_IMAGE_FILTER = 1 //! Mitchell & Netravali's two-param cubic filter
	FILTER_BILINEAR   FREE_IMAGE_FILTER = 2 //! Bilinear filter
	FILTER_BSPLINE    FREE_IMAGE_FILTER = 3 //! 4th order (cubic) b-spline
	FILTER_CATMULLROM FREE_IMAGE_FILTER = 4 //! Catmull-Rom spline Overhauser spline
	FILTER_LANCZOS3   FREE_IMAGE_FILTER = 5 //! Lanczos3 filter

	FIJPEG_OP_NONE       FREE_IMAGE_JPEG_OPERATION = 0 //! no transformation
	FIJPEG_OP_FLIP_H     FREE_IMAGE_JPEG_OPERATION = 1 //! horizontal flip
	FIJPEG_OP_FLIP_V     FREE_IMAGE_JPEG_OPERATION = 2 //! vertical flip
	FIJPEG_OP_TRANSPOSE  FREE_IMAGE_JPEG_OPERATION = 3 //! transpose across UL-to-LR axis
	FIJPEG_OP_TRANSVERSE FREE_IMAGE_JPEG_OPERATION = 4 //! transpose across UR-to-LL axis
	FIJPEG_OP_ROTATE_90  FREE_IMAGE_JPEG_OPERATION = 5 //! 90-degree clockwise rotation
	FIJPEG_OP_ROTATE_180 FREE_IMAGE_JPEG_OPERATION = 6 //! 180-degree rotation
	FIJPEG_OP_ROTATE_270 FREE_IMAGE_JPEG_OPERATION = 7 //! 270-degree clockwise (or 90 ccw)

	FIMD_NODATA         FREE_IMAGE_MDMODEL = -1
	FIMD_COMMENTS       FREE_IMAGE_MDMODEL = 0  //! single comment or keywords
	FIMD_EXIF_MAIN      FREE_IMAGE_MDMODEL = 1  //! Exif-TIFF metadata
	FIMD_EXIF_EXIF      FREE_IMAGE_MDMODEL = 2  //! Exif-specific metadata
	FIMD_EXIF_GPS       FREE_IMAGE_MDMODEL = 3  //! Exif GPS metadata
	FIMD_EXIF_MAKERNOTE FREE_IMAGE_MDMODEL = 4  //! Exif maker note metadata
	FIMD_EXIF_INTEROP   FREE_IMAGE_MDMODEL = 5  //! Exif interoperability metadata
	FIMD_IPTC           FREE_IMAGE_MDMODEL = 6  //! IPTC/NAA metadata
	FIMD_XMP            FREE_IMAGE_MDMODEL = 7  //! Abobe XMP metadata
	FIMD_GEOTIFF        FREE_IMAGE_MDMODEL = 8  //! GeoTIFF metadata
	FIMD_ANIMATION      FREE_IMAGE_MDMODEL = 9  //! Animation metadata
	FIMD_CUSTOM         FREE_IMAGE_MDMODEL = 10 //! Used to attach other metadata types to a dib
	FIMD_EXIF_RAW       FREE_IMAGE_MDMODEL = 11 //! Exif metadata as a raw buffer

	FIDT_NOTYPE    FREE_IMAGE_MDTYPE = 0  //! placeholder
	FIDT_BYTE      FREE_IMAGE_MDTYPE = 1  //! 8-bit unsigned integer
	FIDT_ASCII     FREE_IMAGE_MDTYPE = 2  //! 8-bit bytes w/ last byte null
	FIDT_SHORT     FREE_IMAGE_MDTYPE = 3  //! 16-bit unsigned integer
	FIDT_LONG      FREE_IMAGE_MDTYPE = 4  //! 32-bit unsigned integer
	FIDT_RATIONAL  FREE_IMAGE_MDTYPE = 5  //! 64-bit unsigned fraction
	FIDT_SBYTE     FREE_IMAGE_MDTYPE = 6  //! 8-bit signed integer
	FIDT_UNDEFINED FREE_IMAGE_MDTYPE = 7  //! 8-bit untyped data
	FIDT_SSHORT    FREE_IMAGE_MDTYPE = 8  //! 16-bit signed integer
	FIDT_SLONG     FREE_IMAGE_MDTYPE = 9  //! 32-bit signed integer
	FIDT_SRATIONAL FREE_IMAGE_MDTYPE = 10 //! 64-bit signed fraction
	FIDT_FLOAT     FREE_IMAGE_MDTYPE = 11 //! 32-bit IEEE floating point
	FIDT_DOUBLE    FREE_IMAGE_MDTYPE = 12 //! 64-bit IEEE floating point
	FIDT_IFD       FREE_IMAGE_MDTYPE = 13 //! 32-bit unsigned integer (offset)
	FIDT_PALETTE   FREE_IMAGE_MDTYPE = 14 //! 32-bit RGBQUAD
	FIDT_LONG8     FREE_IMAGE_MDTYPE = 16 //! 64-bit unsigned integer
	FIDT_SLONG8    FREE_IMAGE_MDTYPE = 17 //! 64-bit signed integer
	FIDT_IFD8      FREE_IMAGE_MDTYPE = 18 //! 64-bit unsigned integer (offset)

	FITMO_DRAGO03    FREE_IMAGE_TMO = 0 //! Adaptive logarithmic mapping (F. Drago, 2003)
	FITMO_REINHARD05 FREE_IMAGE_TMO = 1 //! Dynamic range reduction inspired by photoreceptor physiology (E. Reinhard, 2005)
	FITMO_FATTAL02   FREE_IMAGE_TMO = 2 //! Gradient domain high dynamic range compression (R. Fattal, 2002)

	FID_FS           FREE_IMAGE_DITHER = 0 //! Floyd & Steinberg error diffusion
	FID_BAYER4x4     FREE_IMAGE_DITHER = 1 //! Bayer ordered dispersed dot dithering (order 2 dithering matrix)
	FID_BAYER8x8     FREE_IMAGE_DITHER = 2 //! Bayer ordered dispersed dot dithering (order 3 dithering matrix)
	FID_CLUSTER6x6   FREE_IMAGE_DITHER = 3 //! Ordered clustered dot dithering (order 3 - 6x6 matrix)
	FID_CLUSTER8x8   FREE_IMAGE_DITHER = 4 //! Ordered clustered dot dithering (order 4 - 8x8 matrix)
	FID_CLUSTER16x16 FREE_IMAGE_DITHER = 5 //! Ordered clustered dot dithering (order 8 - 16x16 matrix)
	FID_BAYER16x16   FREE_IMAGE_DITHER = 6 //! Bayer ordered dispersed dot dithering (order 4 dithering matrix)

	FIQ_WUQUANT  FREE_IMAGE_QUANTIZE = 0 //! Xiaolin Wu color quantization algorithm
	FIQ_NNQUANT  FREE_IMAGE_QUANTIZE = 1 //! NeuQuant neural-net quantization algorithm by Anthony Dekker
	FIQ_LFPQUANT FREE_IMAGE_QUANTIZE = 2 //! Lossless Fast Pseudo-Quantization Algorithm by Carsten Klein

	FIC_MINISWHITE FREE_IMAGE_COLOR_TYPE = 0 //! min value is white
	FIC_MINISBLACK FREE_IMAGE_COLOR_TYPE = 1 //! min value is black
	FIC_RGB        FREE_IMAGE_COLOR_TYPE = 2 //! RGB color model
	FIC_PALETTE    FREE_IMAGE_COLOR_TYPE = 3 //! color map indexed
	FIC_RGBALPHA   FREE_IMAGE_COLOR_TYPE = 4 //! RGB color model with alpha channel
	FIC_CMYK       FREE_IMAGE_COLOR_TYPE = 5 //! CMYK color model

	FIT_UNKNOWN FREE_IMAGE_TYPE = 0  //! unknown type
	FIT_BITMAP  FREE_IMAGE_TYPE = 1  //! standard image            : 1-, 4-, 8-, 16-, 24-, 32-bit
	FIT_UINT16  FREE_IMAGE_TYPE = 2  //! array of unsigned short    : unsigned 16-bit
	FIT_INT16   FREE_IMAGE_TYPE = 3  //! array of short            : signed 16-bit
	FIT_UINT32  FREE_IMAGE_TYPE = 4  //! array of unsigned long    : unsigned 32-bit
	FIT_INT32   FREE_IMAGE_TYPE = 5  //! array of long            : signed 32-bit
	FIT_FLOAT   FREE_IMAGE_TYPE = 6  //! array of float            : 32-bit IEEE floating point
	FIT_DOUBLE  FREE_IMAGE_TYPE = 7  //! array of double            : 64-bit IEEE floating point
	FIT_COMPLEX FREE_IMAGE_TYPE = 8  //! array of FICOMPLEX        : 2 x 64-bit IEEE floating point
	FIT_RGB16   FREE_IMAGE_TYPE = 9  //! 48-bit RGB image            : 3 x 16-bit
	FIT_RGBA16  FREE_IMAGE_TYPE = 10 //! 64-bit RGBA image        : 4 x 16-bit
	FIT_RGBF    FREE_IMAGE_TYPE = 11 //! 96-bit RGB float image    : 3 x 32-bit IEEE floating point
	FIT_RGBAF   FREE_IMAGE_TYPE = 12 //! 128-bit RGBA float image    : 4 x 32-bit IEEE floating point

	FIF_UNKNOWN FREE_IMAGE_FORMAT = -1
	FIF_BMP     FREE_IMAGE_FORMAT = 0
	FIF_ICO     FREE_IMAGE_FORMAT = 1
	FIF_JPEG    FREE_IMAGE_FORMAT = 2
	FIF_JNG     FREE_IMAGE_FORMAT = 3
	FIF_KOALA   FREE_IMAGE_FORMAT = 4
	FIF_LBM     FREE_IMAGE_FORMAT = 5
	FIF_IFF     FREE_IMAGE_FORMAT = FIF_LBM
	FIF_MNG     FREE_IMAGE_FORMAT = 6
	FIF_PBM     FREE_IMAGE_FORMAT = 7
	FIF_PBMRAW  FREE_IMAGE_FORMAT = 8
	FIF_PCD     FREE_IMAGE_FORMAT = 9
	FIF_PCX     FREE_IMAGE_FORMAT = 10
	FIF_PGM     FREE_IMAGE_FORMAT = 11
	FIF_PGMRAW  FREE_IMAGE_FORMAT = 12
	FIF_PNG     FREE_IMAGE_FORMAT = 13
	FIF_PPM     FREE_IMAGE_FORMAT = 14
	FIF_PPMRAW  FREE_IMAGE_FORMAT = 15
	FIF_RAS     FREE_IMAGE_FORMAT = 16
	FIF_TARGA   FREE_IMAGE_FORMAT = 17
	FIF_TIFF    FREE_IMAGE_FORMAT = 18
	FIF_WBMP    FREE_IMAGE_FORMAT = 19
	FIF_PSD     FREE_IMAGE_FORMAT = 20
	FIF_CUT     FREE_IMAGE_FORMAT = 21
	FIF_XBM     FREE_IMAGE_FORMAT = 22
	FIF_XPM     FREE_IMAGE_FORMAT = 23
	FIF_DDS     FREE_IMAGE_FORMAT = 24
	FIF_GIF     FREE_IMAGE_FORMAT = 25
	FIF_HDR     FREE_IMAGE_FORMAT = 26
	FIF_FAXG3   FREE_IMAGE_FORMAT = 27
	FIF_SGI     FREE_IMAGE_FORMAT = 28
	FIF_EXR     FREE_IMAGE_FORMAT = 29
	FIF_J2K     FREE_IMAGE_FORMAT = 30
	FIF_JP2     FREE_IMAGE_FORMAT = 31
	FIF_PFM     FREE_IMAGE_FORMAT = 32
	FIF_PICT    FREE_IMAGE_FORMAT = 33
	FIF_RAW     FREE_IMAGE_FORMAT = 34
	FIF_WEBP    FREE_IMAGE_FORMAT = 35
	FIF_JXR     FREE_IMAGE_FORMAT = 36
)

// Init / Error routines ----------------------------------------------------

var _func_FreeImage_Initialise_ = &c.FuncPrototype{Name: "FreeImage_Initialise", OutType: c.Void, InTypes: []c.Type{c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_Initialise(BOOL load_local_plugins_only FI_DEFAULT(FALSE));
func Initialise(load_local_plugins_only bool) {
	b := c.CBool(load_local_plugins_only)
	fiLib.Call(_func_FreeImage_Initialise_, inArgs{&b})
}

var _func_FreeImage_DeInitialise_ = &c.FuncPrototype{Name: "FreeImage_DeInitialise", OutType: c.Void, InTypes: nil}

// DLL_API void DLL_CALLCONV FreeImage_DeInitialise(void);
func DeInitialise() {
	fiLib.Call(_func_FreeImage_DeInitialise_, nil)
}

// Version routines ---------------------------------------------------------

var _func_FreeImage_GetVersion_ = &c.FuncPrototype{Name: "FreeImage_GetVersion", OutType: c.Pointer, InTypes: nil}

// DLL_API const char *DLL_CALLCONV FreeImage_GetVersion(void);
func GetVersion() string {
	return fiLib.Call(_func_FreeImage_GetVersion_, nil).StrFree()
}

var _func_FreeImage_GetCopyrightMessage_ = &c.FuncPrototype{Name: "FreeImage_GetCopyrightMessage", OutType: c.Pointer, InTypes: nil}

// DLL_API const char *DLL_CALLCONV FreeImage_GetCopyrightMessage(void);
func GetCopyrightMessage() string {
	return fiLib.Call(_func_FreeImage_GetCopyrightMessage_, nil).StrFree()
}

// // Message output functions -------------------------------------------------

// typedef void (*FreeImage_OutputMessageFunction)(FREE_IMAGE_FORMAT fif, const char *msg);
// typedef void (DLL_CALLCONV *FreeImage_OutputMessageFunctionStdCall)(FREE_IMAGE_FORMAT fif, const char *msg);

// DLL_API void DLL_CALLCONV FreeImage_SetOutputMessageStdCall(FreeImage_OutputMessageFunctionStdCall omf);
// DLL_API void DLL_CALLCONV FreeImage_SetOutputMessage(FreeImage_OutputMessageFunction omf);
// DLL_API void DLL_CALLCONV FreeImage_OutputMessageProc(int fif, const char *fmt, ...);

// Allocate / Clone / Unload routines ---------------------------------------

var _func_FreeImage_Allocate_ = &c.FuncPrototype{Name: "FreeImage_Allocate", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.U32, c.U32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Allocate(int width, int height, int bpp, unsigned red_mask FI_DEFAULT(0), unsigned green_mask FI_DEFAULT(0), unsigned blue_mask FI_DEFAULT(0));
func Allocate(width, height int32, bpp int32, red, green, blue uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Allocate_, inArgs{&width, &height, &bpp, &red, &green, &blue}).PtrFree())
}

var _func_FreeImage_AllocateT_ = &c.FuncPrototype{Name: "FreeImage_AllocateT", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.U32, c.U32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_AllocateT(FREE_IMAGE_TYPE type, int width, int height, int bpp FI_DEFAULT(8), unsigned red_mask FI_DEFAULT(0), unsigned green_mask FI_DEFAULT(0), unsigned blue_mask FI_DEFAULT(0));
//
// bpp : default = 8;
func AllocateT(typ FREE_IMAGE_TYPE, width, height int32, bpp int32, red, green, blue uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_AllocateT_, inArgs{&typ, &width, &height, &bpp, &red, &green, &blue}).PtrFree())
}

var _func_FreeImage_Clone_ = &c.FuncPrototype{Name: "FreeImage_Clone", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP * DLL_CALLCONV FreeImage_Clone(FIBITMAP *dib);
func (dib *BitMap) Clone() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Clone_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_Unload_ = &c.FuncPrototype{Name: "FreeImage_Unload", OutType: c.Void, InTypes: []c.Type{c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_Unload(FIBITMAP *dib);
func (dib *BitMap) Unload() {
	fiLib.Call(_func_FreeImage_Unload_, inArgs{&dib})
}

// Header loading routines

var _func_FreeImage_HasPixels_ = &c.FuncPrototype{Name: "FreeImage_HasPixels", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_HasPixels(FIBITMAP *dib);
func (dib *BitMap) HasPixels() bool {
	return fiLib.Call(_func_FreeImage_HasPixels_, inArgs{&dib}).BoolFree()
}

// Load / Save routines -----------------------------------------------------

var _func_FreeImage_Load_ = &c.FuncPrototype{Name: "FreeImage_Load", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Load(FREE_IMAGE_FORMAT fif, const char *filename, int flags FI_DEFAULT(0));
func Load(format FREE_IMAGE_FORMAT, filename string, flags int32) *BitMap {
	fn := c.CStr(filename)
	defer c.Free(fn)
	return (*BitMap)(fiLib.Call(_func_FreeImage_Load_, inArgs{&format, &fn, &flags}).PtrFree())
}

func NewBitMapFromFile(format FREE_IMAGE_FORMAT, filename string, flags int32) *BitMap {
	return Load(format, filename, flags)
}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_LoadU(FREE_IMAGE_FORMAT fif, const wchar_t *filename, int flags FI_DEFAULT(0));
// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_LoadFromHandle(FREE_IMAGE_FORMAT fif, FreeImageIO *io, fi_handle handle, int flags FI_DEFAULT(0));

var _func_FreeImage_Save_ = &c.FuncPrototype{Name: "FreeImage_Save", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_Save(FREE_IMAGE_FORMAT fif, FIBITMAP *dib, const char *filename, int flags FI_DEFAULT(0));
func (dib *BitMap) Save(format FREE_IMAGE_FORMAT, filename string, flags int32) bool {
	fn := c.CStr(filename)
	defer c.Free(fn)
	return fiLib.Call(_func_FreeImage_Save_, inArgs{&format, &dib, &fn, &flags}).Bool()
}

// DLL_API BOOL DLL_CALLCONV FreeImage_SaveU(FREE_IMAGE_FORMAT fif, FIBITMAP *dib, const wchar_t *filename, int flags FI_DEFAULT(0));
// DLL_API BOOL DLL_CALLCONV FreeImage_SaveToHandle(FREE_IMAGE_FORMAT fif, FIBITMAP *dib, FreeImageIO *io, fi_handle handle, int flags FI_DEFAULT(0));

// Memory I/O stream routines -----------------------------------------------

var _func_FreeImage_OpenMemory_ = &c.FuncPrototype{Name: "FreeImage_OpenMemory", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.U32}}

// DLL_API FIMEMORY *DLL_CALLCONV FreeImage_OpenMemory(BYTE *data FI_DEFAULT(0), DWORD size_in_bytes FI_DEFAULT(0));
func OpenMemory(data []byte) *Memory {
	d, l := (*byte)(nil), uint32(0)
	if data != nil {
		d, l = &data[0], uint32(len(data))
	}
	return (*Memory)(fiLib.Call(_func_FreeImage_OpenMemory_, inArgs{&d, &l}).PtrFree())
}

func NewMemory(data []byte) *Memory {
	return OpenMemory(data)
}

var _func_FreeImage_CloseMemory_ = &c.FuncPrototype{Name: "FreeImage_CloseMemory", OutType: c.Void, InTypes: []c.Type{c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_CloseMemory(FIMEMORY *stream);
func (stream *Memory) CloseMemory() {
	fiLib.Call(_func_FreeImage_CloseMemory_, inArgs{&stream})
}

var _func_FreeImage_LoadFromMemory_ = &c.FuncPrototype{Name: "FreeImage_LoadFromMemory", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_LoadFromMemory(FREE_IMAGE_FORMAT fif, FIMEMORY *stream, int flags FI_DEFAULT(0));
func LoadFromMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_LoadFromMemory_, inArgs{&fif, &stream, &flag}).PtrFree())
}

func NewBitMapFromMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) *BitMap {
	return LoadFromMemory(fif, stream, flag)
}

var _func_FreeImage_SaveToMemory_ = &c.FuncPrototype{Name: "FreeImage_SaveToMemory", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SaveToMemory(FREE_IMAGE_FORMAT fif, FIBITMAP *dib, FIMEMORY *stream, int flags FI_DEFAULT(0));
func SaveToMemory(fif FREE_IMAGE_FORMAT, dib *BitMap, stream *Memory, flag int32) bool {
	return fiLib.Call(_func_FreeImage_SaveToMemory_, inArgs{&fif, &dib, &stream, &flag}).BoolFree()
}

// DLL_API BOOL DLL_CALLCONV FreeImage_SaveToMemory(FREE_IMAGE_FORMAT fif, FIBITMAP *dib, FIMEMORY *stream, int flags FI_DEFAULT(0));
func (dib *BitMap) SaveToMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) bool {
	return fiLib.Call(_func_FreeImage_SaveToMemory_, inArgs{&fif, &dib, &stream, &flag}).BoolFree()
}

var _func_FreeImage_TellMemory_ = &c.FuncPrototype{Name: "FreeImage_TellMemory", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API long DLL_CALLCONV FreeImage_TellMemory(FIMEMORY *stream);
func (stream *Memory) Tell() int32 {
	return fiLib.Call(_func_FreeImage_TellMemory_, inArgs{&stream}).I32Free()
}

var _func_FreeImage_SeekMemory_ = &c.FuncPrototype{Name: "FreeImage_SeekMemory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SeekMemory(FIMEMORY *stream, long offset, int origin);
func (stream *Memory) SeekMemory(offset int32, origin SeekOrigin) bool {
	return fiLib.Call(_func_FreeImage_SeekMemory_, inArgs{&stream, &offset, &origin}).BoolFree()
}

var _func_FreeImage_AcquireMemory_ = &c.FuncPrototype{Name: "FreeImage_AcquireMemory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AcquireMemory(FIMEMORY *stream, BYTE **data, DWORD *size_in_bytes);
func (stream *Memory) AcquireMemory() (bool, unsafe.Pointer) {
	data, _size := usf.Malloc(8), uint32(0)
	usf.Memset(data, 0, 8)
	size := &_size
	defer usf.Free(data)
	ret := fiLib.Call(_func_FreeImage_AcquireMemory_, inArgs{&stream, &data, &size}).BoolFree()
	if !ret {
		return ret, nil
	}
	return ret, usf.Slice(usf.Pop(data), uint64(*size))
}

var _func_FreeImage_ReadMemory_ = &c.FuncPrototype{Name: "FreeImage_ReadMemory", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_ReadMemory(void *buffer, unsigned size, unsigned count, FIMEMORY *stream);
//
// Reads data from a memory stream.
//
// The FreeImage_ReadMemory function reads up to count items of size bytes from the input
// memory stream and stores them in buffer. The memory pointer associated with stream is
// increased by the number of bytes actually read.
//
// The function returns the number of full items actually read, which may be less than count if an
// error occurs or if the end of the stream is encountered before reaching count.
//
// e.g.,
//
// buf := make([]byte, 16, 16) // slice
//
// buf_size := len(buf[0])     // size of slice items
//
// srcMem.WriteTo(&buf[0], buf_size, len(buf))
func (stream *Memory) WriteToSlice(dstAddr interface{}, size, count uint32) uint32 {
	p := usf.AddrOf(dstAddr)
	return fiLib.Call(_func_FreeImage_ReadMemory_, inArgs{&p, &size, &count, &stream}).U32Free()
}

var _func_FreeImage_WriteMemory_ = &c.FuncPrototype{Name: "FreeImage_WriteMemory", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_WriteMemory(const void *buffer, unsigned size, unsigned count, FIMEMORY *stream);
//
// Writes data to a memory stream.
//
// The FreeImage_WriteMemory function writes up to count items, of size length each, from
// buffer to the output memory stream. The memory pointer associated with stream is
//
// incremented by the number of bytes actually written.
// The function returns the number of full items actually written, which may be less than count if
// an error occurs.
//
// e.g.,
//
// srcBuf := make([]byte, 16, 16) // slice
//
// buf_size := len(buf[0])     // size of slice items
//
// dstMem.ReadFromSlice(&srcBuf[0], buf_size, len(buf))
func (stream *Memory) ReadFromSlice(srcAddr interface{}, size, count uint32) uint32 {
	p := usf.AddrOf(srcAddr)
	return fiLib.Call(_func_FreeImage_WriteMemory_, inArgs{&p, &size, &count, &stream}).U32Free()
}

var _func_FreeImage_LoadMultiBitmapFromMemory_ = &c.FuncPrototype{Name: "FreeImage_LoadMultiBitmapFromMemory", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.I32}}

// DLL_API FIMULTIBITMAP *DLL_CALLCONV FreeImage_LoadMultiBitmapFromMemory(FREE_IMAGE_FORMAT fif, FIMEMORY *stream, int flags FI_DEFAULT(0));
func LoadMultiBitmapFromMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) *MultiBitMap {
	return (*MultiBitMap)(fiLib.Call(_func_FreeImage_LoadMultiBitmapFromMemory_, inArgs{&fif, &stream, &flag}).PtrFree())
}

// DLL_API FIMULTIBITMAP *DLL_CALLCONV FreeImage_LoadMultiBitmapFromMemory(FREE_IMAGE_FORMAT fif, FIMEMORY *stream, int flags FI_DEFAULT(0));
func NewMultiBitMapFromMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) *MultiBitMap {
	return LoadMultiBitmapFromMemory(fif, stream, flag)
}

var _func_FreeImage_SaveMultiBitmapToMemory_ = &c.FuncPrototype{Name: "FreeImage_SaveMultiBitmapToMemory", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SaveMultiBitmapToMemory(FREE_IMAGE_FORMAT fif, FIMULTIBITMAP *bitmap, FIMEMORY *stream, int flags);
func (mb *MultiBitMap) SaveToMemory(fif FREE_IMAGE_FORMAT, stream *Memory, flag int32) bool {
	return fiLib.Call(_func_FreeImage_SaveMultiBitmapToMemory_, inArgs{&fif, &mb, &stream, &flag}).BoolFree()
}

// Plugin Interface ---------------------------------------------------------

// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_RegisterLocalPlugin(FI_InitProc proc_address, const char *format FI_DEFAULT(0), const char *description FI_DEFAULT(0), const char *extension FI_DEFAULT(0), const char *regexpr FI_DEFAULT(0));
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_RegisterExternalPlugin(const char *path, const char *format FI_DEFAULT(0), const char *description FI_DEFAULT(0), const char *extension FI_DEFAULT(0), const char *regexpr FI_DEFAULT(0));
// DLL_API int DLL_CALLCONV FreeImage_GetFIFCount(void);
// DLL_API int DLL_CALLCONV FreeImage_SetPluginEnabled(FREE_IMAGE_FORMAT fif, BOOL enable);
// DLL_API int DLL_CALLCONV FreeImage_IsPluginEnabled(FREE_IMAGE_FORMAT fif);
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFIFFromFormat(const char *format);
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFIFFromMime(const char *mime);
// DLL_API const char *DLL_CALLCONV FreeImage_GetFormatFromFIF(FREE_IMAGE_FORMAT fif);
// DLL_API const char *DLL_CALLCONV FreeImage_GetFIFExtensionList(FREE_IMAGE_FORMAT fif);
// DLL_API const char *DLL_CALLCONV FreeImage_GetFIFDescription(FREE_IMAGE_FORMAT fif);
// DLL_API const char *DLL_CALLCONV FreeImage_GetFIFRegExpr(FREE_IMAGE_FORMAT fif);
// DLL_API const char *DLL_CALLCONV FreeImage_GetFIFMimeType(FREE_IMAGE_FORMAT fif);
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFIFFromFilename(const char *filename);
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFIFFromFilenameU(const wchar_t *filename);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsReading(FREE_IMAGE_FORMAT fif);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsWriting(FREE_IMAGE_FORMAT fif);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsExportBPP(FREE_IMAGE_FORMAT fif, int bpp);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsExportType(FREE_IMAGE_FORMAT fif, FREE_IMAGE_TYPE type);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsICCProfiles(FREE_IMAGE_FORMAT fif);
// DLL_API BOOL DLL_CALLCONV FreeImage_FIFSupportsNoPixels(FREE_IMAGE_FORMAT fif);

// Multipaging interface ----------------------------------------------------

var _func_FreeImage_OpenMultiBitmap_ = &c.FuncPrototype{Name: "FreeImage_OpenMultiBitmap", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.I32, c.I32, c.I32, c.I32}}

// DLL_API FIMULTIBITMAP * DLL_CALLCONV FreeImage_OpenMultiBitmap(FREE_IMAGE_FORMAT fif, const char *filename, BOOL create_new, BOOL read_only, BOOL keep_cache_in_memory FI_DEFAULT(FALSE), int flags FI_DEFAULT(0));
func OpenMultiBitmap(fif FREE_IMAGE_FORMAT, filename string, create_new, read_only, keep_cache_in_memory bool, flag int32) *MultiBitMap {
	fn, cn, ro, kcm := c.CStr(filename), c.CBool(create_new), c.CBool(read_only), c.CBool(keep_cache_in_memory)
	defer c.Free(fn)
	return (*MultiBitMap)(fiLib.Call(_func_FreeImage_OpenMultiBitmap_, inArgs{&fif, &fn, &cn, &ro, kcm, &flag}).PtrFree())
}

func NewMultiBitmapFromFile(fif FREE_IMAGE_FORMAT, filename string, create_new, read_only, keep_cache_in_memory bool, flag int32) *MultiBitMap {
	return OpenMultiBitmap(fif, filename, create_new, read_only, keep_cache_in_memory, flag)
}

// DLL_API FIMULTIBITMAP * DLL_CALLCONV FreeImage_OpenMultiBitmapFromHandle(FREE_IMAGE_FORMAT fif, FreeImageIO *io, fi_handle handle, int flags FI_DEFAULT(0));
// DLL_API BOOL DLL_CALLCONV FreeImage_SaveMultiBitmapToHandle(FREE_IMAGE_FORMAT fif, FIMULTIBITMAP *bitmap, FreeImageIO *io, fi_handle handle, int flags FI_DEFAULT(0));

var _func_FreeImage_CloseMultiBitmap_ = &c.FuncPrototype{Name: "FreeImage_CloseMultiBitmap", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_CloseMultiBitmap(FIMULTIBITMAP *bitmap, int flags FI_DEFAULT(0));
func (bitmap *MultiBitMap) Close(flag int32) bool {
	return fiLib.Call(_func_FreeImage_CloseMultiBitmap_, inArgs{&bitmap, &flag}).BoolFree()
}

var _func_FreeImage_GetPageCount_ = &c.FuncPrototype{Name: "FreeImage_GetPageCount", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API int DLL_CALLCONV FreeImage_GetPageCount(FIMULTIBITMAP *bitmap);
func (bitmap *MultiBitMap) GetPageCount() int32 {
	return fiLib.Call(_func_FreeImage_GetPageCount_, inArgs{&bitmap}).I32Free()
}

var _func_FreeImage_AppendPage_ = &c.FuncPrototype{Name: "FreeImage_AppendPage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_AppendPage(FIMULTIBITMAP *bitmap, FIBITMAP *data);
func (bitmap *MultiBitMap) AppendPage(data *BitMap) {
	fiLib.Call(_func_FreeImage_AppendPage_, inArgs{&bitmap, &data})
}

var _func_FreeImage_InsertPage_ = &c.FuncPrototype{Name: "FreeImage_InsertPage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_InsertPage(FIMULTIBITMAP *bitmap, int page, FIBITMAP *data);
func (bitmap *MultiBitMap) InsertPage(page int32, data *BitMap) {
	fiLib.Call(_func_FreeImage_InsertPage_, inArgs{&bitmap, &page, &data})
}

var _func_FreeImage_DeletePage_ = &c.FuncPrototype{Name: "FreeImage_DeletePage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_DeletePage(FIMULTIBITMAP *bitmap, int page);
func (bitmap *MultiBitMap) DeletePage(page int32) {
	fiLib.Call(_func_FreeImage_DeletePage_, inArgs{&bitmap, &page})
}

var _func_FreeImage_LockPage_ = &c.FuncPrototype{Name: "FreeImage_LockPage", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP * DLL_CALLCONV FreeImage_LockPage(FIMULTIBITMAP *bitmap, int page);
func (bitmap *MultiBitMap) LockPage(page int32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_LockPage_, inArgs{&bitmap, &page}).PtrFree())
}

var _func_FreeImage_UnlockPage_ = &c.FuncPrototype{Name: "FreeImage_UnlockPage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_UnlockPage(FIMULTIBITMAP *bitmap, FIBITMAP *data, BOOL changed);
func (bitmap *MultiBitMap) UnlockPage(data *BitMap, changed bool) {
	b := c.CBool(changed)
	fiLib.Call(_func_FreeImage_UnlockPage_, inArgs{&bitmap, &data, &b})
}

var _func_FreeImage_MovePage_ = &c.FuncPrototype{Name: "FreeImage_MovePage", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_MovePage(FIMULTIBITMAP *bitmap, int target, int source);
func (bitmap *MultiBitMap) MovePage(target, source int32) bool {
	return fiLib.Call(_func_FreeImage_MovePage_, inArgs{&bitmap, &target, &source}).BoolFree()
}

var _func_FreeImage_GetLockedPageNumbers_ = &c.FuncPrototype{Name: "FreeImage_GetLockedPageNumbers", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetLockedPageNumbers(FIMULTIBITMAP *bitmap, int *pages, int *count);
func (bitmap *MultiBitMap) GetLockedPageNumbers() (pages []int32, count int32, ok bool) {
	_pages, _count := (*int32)(nil), &count

	ok = fiLib.Call(_func_FreeImage_GetLockedPageNumbers_, inArgs{&bitmap, &_pages, &_count}).BoolFree()
	if !ok || count == 0 {
		return nil, 0, ok
	}

	pages = make([]int32, count)
	_pages = &pages[0]
	return pages, count, fiLib.Call(_func_FreeImage_GetLockedPageNumbers_, inArgs{&bitmap, &_pages, &count}).BoolFree()
}

// File type request routines ------------------------------------------------

var _func_FreeImage_GetFileType_ = &c.FuncPrototype{Name: "FreeImage_GetFileType", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFileType(const char *filename, int size FI_DEFAULT(0));
func GetFileType(filename string, size int32) FREE_IMAGE_FORMAT {
	fn := c.CStr(filename)
	defer c.Free(fn)
	return FREE_IMAGE_FORMAT(fiLib.Call(_func_FreeImage_GetFileType_, inArgs{&fn, &size}).I32Free())
}

// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFileTypeU(const wchar_t *filename, int size FI_DEFAULT(0));
// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFileTypeFromHandle(FreeImageIO *io, fi_handle handle, int size FI_DEFAULT(0));

var _func_FreeImage_GetFileTypeFromMemory_ = &c.FuncPrototype{Name: "FreeImage_GetFileTypeFromMemory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFileTypeFromMemory(FIMEMORY *stream, int size FI_DEFAULT(0));
func GetFileTypeFromMemory(stream *Memory, size int32) FREE_IMAGE_FORMAT {
	return FREE_IMAGE_FORMAT(fiLib.Call(_func_FreeImage_GetFileTypeFromMemory_, inArgs{&stream, &size}).I32Free())
}

// DLL_API FREE_IMAGE_FORMAT DLL_CALLCONV FreeImage_GetFileTypeFromMemory(FIMEMORY *stream, int size FI_DEFAULT(0));
func (stream *Memory) GetFileType() FREE_IMAGE_FORMAT {
	size := int32(0)
	return FREE_IMAGE_FORMAT(fiLib.Call(_func_FreeImage_GetFileTypeFromMemory_, inArgs{&stream, &size}).I32Free())
}

var _func_FreeImage_Validate_ = &c.FuncPrototype{Name: "FreeImage_Validate", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_Validate(FREE_IMAGE_FORMAT fif, const char *filename);
func Validate(fif FREE_IMAGE_FORMAT, filename string) bool {
	fn := c.CStr(filename)
	defer c.Free(fn)
	return fiLib.Call(_func_FreeImage_Validate_, inArgs{&fif, &fn}).BoolFree()
}

// DLL_API BOOL DLL_CALLCONV FreeImage_ValidateU(FREE_IMAGE_FORMAT fif, const wchar_t *filename);
// DLL_API BOOL DLL_CALLCONV FreeImage_ValidateFromHandle(FREE_IMAGE_FORMAT fif, FreeImageIO *io, fi_handle handle);

var _func_FreeImage_ValidateFromMemory_ = &c.FuncPrototype{Name: "FreeImage_ValidateFromMemory", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_ValidateFromMemory(FREE_IMAGE_FORMAT fif, FIMEMORY *stream);
func ValidateFromMemory(fif FREE_IMAGE_FORMAT, stream *Memory) bool {
	return fiLib.Call(_func_FreeImage_ValidateFromMemory_, inArgs{&fif, &stream}).BoolFree()
}

// DLL_API BOOL DLL_CALLCONV FreeImage_ValidateFromMemory(FREE_IMAGE_FORMAT fif, FIMEMORY *stream);
func (stream *Memory) Validate(fif FREE_IMAGE_FORMAT) bool {
	return fiLib.Call(_func_FreeImage_ValidateFromMemory_, inArgs{&fif, &stream}).BoolFree()
}

// Image type request routine -----------------------------------------------

var _func_FreeImage_GetImageType_ = &c.FuncPrototype{Name: "FreeImage_GetImageType", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API FREE_IMAGE_TYPE DLL_CALLCONV FreeImage_GetImageType(FIBITMAP *dib);
func GetImageType(dib *BitMap) FREE_IMAGE_TYPE {
	return FREE_IMAGE_TYPE(fiLib.Call(_func_FreeImage_GetImageType_, inArgs{&dib}).I32Free())
}

// DLL_API FREE_IMAGE_TYPE DLL_CALLCONV FreeImage_GetImageType(FIBITMAP *dib);
func (dib *BitMap) GetImageType() FREE_IMAGE_TYPE {
	return FREE_IMAGE_TYPE(fiLib.Call(_func_FreeImage_GetImageType_, inArgs{&dib}).I32Free())
}

// FreeImage helper routines ------------------------------------------------

var _func_FreeImage_IsLittleEndian_ = &c.FuncPrototype{Name: "FreeImage_IsLittleEndian", OutType: c.I32, InTypes: nil}

// DLL_API BOOL DLL_CALLCONV FreeImage_IsLittleEndian(void);
func IsLittleEndian() bool {
	return fiLib.Call(_func_FreeImage_IsLittleEndian_, nil).BoolFree()
}

var _func_FreeImage_LookupX11Color_ = &c.FuncPrototype{Name: "FreeImage_LookupX11Color", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_LookupX11Color(const char *szColor, BYTE *nRed, BYTE *nGreen, BYTE *nBlue);
func LookupX11Color(szColor string) (r, g, b byte, ok bool) {
	sc, _r, _g, _b := c.CStr(szColor), &r, &g, &b
	defer c.Free(sc)

	ok = fiLib.Call(_func_FreeImage_LookupX11Color_, inArgs{&sc, &_r, &_g, &_b}).BoolFree()
	return
}

var _func_FreeImage_LookupSVGColor_ = &c.FuncPrototype{Name: "FreeImage_LookupSVGColor", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_LookupSVGColor(const char *szColor, BYTE *nRed, BYTE *nGreen, BYTE *nBlue);
func LookupSVGColor(szColor string) (r, g, b byte, ok bool) {
	sc, _r, _g, _b := c.CStr(szColor), &r, &g, &b
	defer c.Free(sc)

	ok = fiLib.Call(_func_FreeImage_LookupSVGColor_, inArgs{&sc, &_r, &_g, &_b}).BoolFree()
	return
}

// Pixel access routines ----------------------------------------------------

var _func_FreeImage_GetBits_ = &c.FuncPrototype{Name: "FreeImage_GetBits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API BYTE *DLL_CALLCONV FreeImage_GetBits(FIBITMAP *dib);
func (dib *BitMap) GetBits() unsafe.Pointer {
	return fiLib.Call(_func_FreeImage_GetBits_, inArgs{&dib}).PtrFree()
}

var _func_FreeImage_GetScanLine_ = &c.FuncPrototype{Name: "FreeImage_GetScanLine", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API BYTE *DLL_CALLCONV FreeImage_GetScanLine(FIBITMAP *dib, int scanline);
func (dib *BitMap) GetScanLine(scanline int32) unsafe.Pointer {
	return fiLib.Call(_func_FreeImage_GetScanLine_, inArgs{&dib, &scanline}).PtrFree()
}

var _func_FreeImage_GetPixelIndex_ = &c.FuncPrototype{Name: "FreeImage_GetPixelIndex", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetPixelIndex(FIBITMAP *dib, unsigned x, unsigned y, BYTE *value);//
// e.g.,
//
// buf := make([]byte, 4);
//
// dib.GetPixelIndex(1, 1, &buf[0])
func (dib *BitMap) GetPixelIndex(x, y uint32, refAddr interface{}) (ok bool) {
	p := usf.AddrOf(refAddr)
	return fiLib.Call(_func_FreeImage_GetPixelIndex_, inArgs{&dib, &x, &y, &p}).BoolFree()
}

var _func_FreeImage_GetPixelColor_ = &c.FuncPrototype{Name: "FreeImage_GetPixelColor", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetPixelColor(FIBITMAP *dib, unsigned x, unsigned y, RGBQUAD *value);
func (dib *BitMap) GetPixelColor(x, y uint32) (value RGBQUAD, ok bool) {
	v := &value
	ok = fiLib.Call(_func_FreeImage_GetPixelColor_, inArgs{&dib, &x, &y, &v}).BoolFree()
	return
}

var _func_FreeImage_SetPixelIndex_ = &c.FuncPrototype{Name: "FreeImage_SetPixelIndex", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetPixelIndex(FIBITMAP *dib, unsigned x, unsigned y, BYTE *value);
func (dib *BitMap) SetPixelIndex(x, y uint32, addr interface{}) (ok bool) {
	p := usf.AddrOf(addr)
	return fiLib.Call(_func_FreeImage_SetPixelIndex_, inArgs{&dib, &x, &y, &p}).BoolFree()
}

var _func_FreeImage_SetPixelColor_ = &c.FuncPrototype{Name: "FreeImage_SetPixelColor", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetPixelColor(FIBITMAP *dib, unsigned x, unsigned y, RGBQUAD *value);
func (dib *BitMap) SetPixelColor(x, y uint32, value *RGBQUAD) (ok bool) {
	return fiLib.Call(_func_FreeImage_SetPixelColor_, inArgs{&dib, &x, &y, &value}).BoolFree()
}

// DIB info routines --------------------------------------------------------

var _func_FreeImage_GetColorsUsed_ = &c.FuncPrototype{Name: "FreeImage_GetColorsUsed", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetColorsUsed(FIBITMAP *dib);
func (dib *BitMap) GetColorsUsed() uint32 {
	return fiLib.Call(_func_FreeImage_GetColorsUsed_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetBPP_ = &c.FuncPrototype{Name: "FreeImage_GetBPP", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetBPP(FIBITMAP *dib);
func (dib *BitMap) GetBPP() uint32 {
	return fiLib.Call(_func_FreeImage_GetBPP_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetWidth_ = &c.FuncPrototype{Name: "FreeImage_GetWidth", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetWidth(FIBITMAP *dib);
func (dib *BitMap) GetWidth() uint32 {
	return fiLib.Call(_func_FreeImage_GetWidth_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetHeight_ = &c.FuncPrototype{Name: "FreeImage_GetHeight", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetHeight(FIBITMAP *dib);
func (dib *BitMap) GetHeight() uint32 {
	return fiLib.Call(_func_FreeImage_GetHeight_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetLine_ = &c.FuncPrototype{Name: "FreeImage_GetLine", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetLine(FIBITMAP *dib);
func (dib *BitMap) GetLine() uint32 {
	return fiLib.Call(_func_FreeImage_GetLine_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetPitch_ = &c.FuncPrototype{Name: "FreeImage_GetPitch", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetPitch(FIBITMAP *dib);
func (dib *BitMap) GetPitch() uint32 {
	return fiLib.Call(_func_FreeImage_GetPitch_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetDIBSize_ = &c.FuncPrototype{Name: "FreeImage_GetDIBSize", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetDIBSize(FIBITMAP *dib);
func (dib *BitMap) GetDIBSize() uint32 {
	return fiLib.Call(_func_FreeImage_GetDIBSize_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetMemorySize_ = &c.FuncPrototype{Name: "FreeImage_GetMemorySize", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetMemorySize(FIBITMAP *dib);
func (dib *BitMap) GetMemorySize() uint32 {
	return fiLib.Call(_func_FreeImage_GetMemorySize_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetPalette_ = &c.FuncPrototype{Name: "FreeImage_GetPalette", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API RGBQUAD *DLL_CALLCONV FreeImage_GetPalette(FIBITMAP *dib);
func (dib *BitMap) GetPalette() uint32 {
	return fiLib.Call(_func_FreeImage_GetPalette_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetDotsPerMeterX_ = &c.FuncPrototype{Name: "FreeImage_GetDotsPerMeterX", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetDotsPerMeterX(FIBITMAP *dib);
func (dib *BitMap) GetDotsPerMeterX() uint32 {
	return fiLib.Call(_func_FreeImage_GetDotsPerMeterX_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetDotsPerMeterY_ = &c.FuncPrototype{Name: "FreeImage_GetDotsPerMeterY", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetDotsPerMeterY(FIBITMAP *dib);
func (dib *BitMap) GetDotsPerMeterY() uint32 {
	return fiLib.Call(_func_FreeImage_GetDotsPerMeterY_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_SetDotsPerMeterX_ = &c.FuncPrototype{Name: "FreeImage_SetDotsPerMeterX", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.U32}}

// DLL_API void DLL_CALLCONV FreeImage_SetDotsPerMeterX(FIBITMAP *dib, unsigned res);
func (dib *BitMap) SetDotsPerMeterX(res uint32) {
	fiLib.Call(_func_FreeImage_SetDotsPerMeterX_, inArgs{&dib, &res})
}

var _func_FreeImage_SetDotsPerMeterY_ = &c.FuncPrototype{Name: "FreeImage_SetDotsPerMeterY", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.U32}}

// DLL_API void DLL_CALLCONV FreeImage_SetDotsPerMeterY(FIBITMAP *dib, unsigned res);
func (dib *BitMap) SetDotsPerMeterY(res uint32) {
	fiLib.Call(_func_FreeImage_SetDotsPerMeterY_, inArgs{&dib, &res})
}

var _func_FreeImage_GetInfoHeader_ = &c.FuncPrototype{Name: "FreeImage_GetInfoHeader", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API BITMAPINFOHEADER *DLL_CALLCONV FreeImage_GetInfoHeader(FIBITMAP *dib);
func (dib *BitMap) GetInfoHeader() *BitMapInfoHeader {
	return (*BitMapInfoHeader)(fiLib.Call(_func_FreeImage_GetInfoHeader_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_GetInfo_ = &c.FuncPrototype{Name: "FreeImage_GetInfo", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API BITMAPINFO *DLL_CALLCONV FreeImage_GetInfo(FIBITMAP *dib);
func (dib *BitMap) GetInfo() *BitMapInfo {
	return (*BitMapInfo)(fiLib.Call(_func_FreeImage_GetInfo_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_GetColorType_ = &c.FuncPrototype{Name: "FreeImage_GetColorType", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API FREE_IMAGE_COLOR_TYPE DLL_CALLCONV FreeImage_GetColorType(FIBITMAP *dib);
func (dib *BitMap) GetColorType() FREE_IMAGE_COLOR_TYPE {
	return (FREE_IMAGE_COLOR_TYPE)(fiLib.Call(_func_FreeImage_GetColorType_, inArgs{&dib}).I32Free())
}

var _func_FreeImage_GetRedMask_ = &c.FuncPrototype{Name: "FreeImage_GetRedMask", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetRedMask(FIBITMAP *dib);
func (dib *BitMap) GetRedMask() uint32 {
	return fiLib.Call(_func_FreeImage_GetRedMask_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetGreenMask_ = &c.FuncPrototype{Name: "FreeImage_GetGreenMask", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetGreenMask(FIBITMAP *dib);
func (dib *BitMap) GetGreenMask() uint32 {
	return fiLib.Call(_func_FreeImage_GetGreenMask_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetBlueMask_ = &c.FuncPrototype{Name: "FreeImage_GetBlueMask", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetBlueMask(FIBITMAP *dib);
func (dib *BitMap) GetBlueMask() uint32 {
	return fiLib.Call(_func_FreeImage_GetBlueMask_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetTransparencyCount_ = &c.FuncPrototype{Name: "FreeImage_GetTransparencyCount", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetTransparencyCount(FIBITMAP *dib);
func (dib *BitMap) GetTransparencyCount() uint32 {
	return fiLib.Call(_func_FreeImage_GetTransparencyCount_, inArgs{&dib}).U32Free()
}

var _func_FreeImage_GetTransparencyTable_ = &c.FuncPrototype{Name: "FreeImage_GetTransparencyTable", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API BYTE * DLL_CALLCONV FreeImage_GetTransparencyTable(FIBITMAP *dib);
func (dib *BitMap) GetTransparencyTable() unsafe.Pointer {
	return fiLib.Call(_func_FreeImage_GetTransparencyTable_, inArgs{&dib}).PtrFree()
}

var _func_FreeImage_SetTransparent_ = &c.FuncPrototype{Name: "FreeImage_SetTransparent", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_SetTransparent(FIBITMAP *dib, BOOL enabled);
func (dib *BitMap) SetTransparent(enabled bool) {
	e := c.CBool(enabled)
	fiLib.Call(_func_FreeImage_SetTransparent_, inArgs{&dib, &e})
}

var _func_FreeImage_SetTransparencyTable_ = &c.FuncPrototype{Name: "FreeImage_SetTransparencyTable", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_SetTransparencyTable(FIBITMAP *dib, BYTE *table, int count);
func (dib *BitMap) SetTransparencyTable(table []*byte) {
	t := &table[0]
	l := int32(len(table))
	fiLib.Call(_func_FreeImage_SetTransparencyTable_, inArgs{&dib, &t, &l})
}

var _func_FreeImage_IsTransparent_ = &c.FuncPrototype{Name: "FreeImage_IsTransparent", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_IsTransparent(FIBITMAP *dib);
func (dib *BitMap) IsTransparent() bool {
	return fiLib.Call(_func_FreeImage_IsTransparent_, inArgs{&dib}).BoolFree()
}

var _func_FreeImage_SetTransparentIndex_ = &c.FuncPrototype{Name: "FreeImage_SetTransparentIndex", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_SetTransparentIndex(FIBITMAP *dib, int index);
func (dib *BitMap) SetTransparentIndex(index int32) {
	fiLib.Call(_func_FreeImage_SetTransparentIndex_, inArgs{&dib, &index})
}

var _func_FreeImage_GetTransparentIndex_ = &c.FuncPrototype{Name: "FreeImage_GetTransparentIndex", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API int DLL_CALLCONV FreeImage_GetTransparentIndex(FIBITMAP *dib);
func (dib *BitMap) GetTransparentIndex() int32 {
	return fiLib.Call(_func_FreeImage_GetTransparentIndex_, inArgs{&dib}).I32Free()
}

var _func_FreeImage_HasBackgroundColor_ = &c.FuncPrototype{Name: "FreeImage_HasBackgroundColor", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_HasBackgroundColor(FIBITMAP *dib);
func (dib *BitMap) HasBackgroundColor() bool {
	return fiLib.Call(_func_FreeImage_HasBackgroundColor_, inArgs{&dib}).BoolFree()
}

var _func_FreeImage_GetBackgroundColor_ = &c.FuncPrototype{Name: "FreeImage_GetBackgroundColor", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetBackgroundColor(FIBITMAP *dib, RGBQUAD *bkcolor);
func (dib *BitMap) GetBackgroundColor() (bkcolor RGBQUAD, ok bool) {
	col := &bkcolor
	ok = fiLib.Call(_func_FreeImage_GetBackgroundColor_, inArgs{&dib, &col}).BoolFree()
	return
}

var _func_FreeImage_SetBackgroundColor_ = &c.FuncPrototype{Name: "FreeImage_SetBackgroundColor", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetBackgroundColor(FIBITMAP *dib, RGBQUAD *bkcolor);
func (dib *BitMap) SetBackgroundColor(bkcolor RGBQUAD) bool {
	col := &bkcolor
	return fiLib.Call(_func_FreeImage_SetBackgroundColor_, inArgs{&dib, &col}).BoolFree()
}

var _func_FreeImage_GetThumbnail_ = &c.FuncPrototype{Name: "FreeImage_GetThumbnail", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_GetThumbnail(FIBITMAP *dib);
func (dib *BitMap) GetThumbnail() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_GetThumbnail_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_SetThumbnail_ = &c.FuncPrototype{Name: "FreeImage_SetThumbnail", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetThumbnail(FIBITMAP *dib, FIBITMAP *thumbnail);
func (dib *BitMap) SetThumbnail(thumbnail *BitMap) bool {
	return fiLib.Call(_func_FreeImage_SetThumbnail_, inArgs{&dib, &thumbnail}).BoolFree()
}

// ICC profile routines -----------------------------------------------------

var _func_FreeImage_GetICCProfile_ = &c.FuncPrototype{Name: "FreeImage_GetICCProfile", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIICCPROFILE *DLL_CALLCONV FreeImage_GetICCProfile(FIBITMAP *dib);
func (dib *BitMap) GetICCProfile() *ICCProfile {
	return (*ICCProfile)(fiLib.Call(_func_FreeImage_GetICCProfile_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_CreateICCProfile_ = &c.FuncPrototype{Name: "FreeImage_CreateICCProfile", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API FIICCPROFILE *DLL_CALLCONV FreeImage_CreateICCProfile(FIBITMAP *dib, void *data, long size);
func (dib *BitMap) CreateICCProfile(data unsafe.Pointer, size int32) *ICCProfile {
	return (*ICCProfile)(fiLib.Call(_func_FreeImage_CreateICCProfile_, inArgs{&dib, &data, &size}).PtrFree())
}

var _func_FreeImage_DestroyICCProfile_ = &c.FuncPrototype{Name: "FreeImage_DestroyICCProfile", OutType: c.Void, InTypes: []c.Type{c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_DestroyICCProfile(FIBITMAP *dib);
func (dib *BitMap) DestroyICCProfile() {
	fiLib.Call(_func_FreeImage_DestroyICCProfile_, inArgs{&dib})
}

// Line conversion routines -------------------------------------------------

// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To4(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To4(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To4_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To4_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine24To4(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine32To4(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To8(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To8(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To8_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To8_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine24To8(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine32To8(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To16_555(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To16_555(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To16_555(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16_565_To16_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine24To16_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine32To16_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To16_565(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To16_565(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To16_565(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16_555_To16_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine24To16_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine32To16_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To24(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To24(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To24(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To24_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To24_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine32To24(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To32(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine1To32MapTransparency(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette, BYTE *table, int transparent_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To32(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine4To32MapTransparency(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette, BYTE *table, int transparent_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To32(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine8To32MapTransparency(BYTE *target, BYTE *source, int width_in_pixels, RGBQUAD *palette, BYTE *table, int transparent_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To32_555(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine16To32_565(BYTE *target, BYTE *source, int width_in_pixels);
// DLL_API void DLL_CALLCONV FreeImage_ConvertLine24To32(BYTE *target, BYTE *source, int width_in_pixels);
// Smart conversion routines ------------------------------------------------

var _func_FreeImage_ConvertTo4Bits_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo4Bits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo4Bits(FIBITMAP *dib);
func (dib *BitMap) ConvertTo4Bits() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo4Bits_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertTo8Bits_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo8Bits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo8Bits(FIBITMAP *dib);
func (dib *BitMap) ConvertTo8Bits() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo8Bits_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToGreyscale_ = &c.FuncPrototype{Name: "FreeImage_ConvertToGreyscale", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToGreyscale(FIBITMAP *dib);
func (dib *BitMap) ConvertToGreyscale() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToGreyscale_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertTo16Bits555_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo16Bits555", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo16Bits555(FIBITMAP *dib);
func (dib *BitMap) ConvertTo16Bits555() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo16Bits555_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertTo16Bits565_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo16Bits565", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo16Bits565(FIBITMAP *dib);
func (dib *BitMap) ConvertTo16Bits565() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo16Bits565_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertTo24Bits_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo24Bits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo24Bits(FIBITMAP *dib);
func (dib *BitMap) ConvertTo24Bits() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo24Bits_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertTo32Bits_ = &c.FuncPrototype{Name: "FreeImage_ConvertTo32Bits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertTo32Bits(FIBITMAP *dib);
func (dib *BitMap) ConvertTo32Bits() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertTo32Bits_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ColorQuantize_ = &c.FuncPrototype{Name: "FreeImage_ColorQuantize", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ColorQuantize(FIBITMAP *dib, FREE_IMAGE_QUANTIZE quantize);
func (dib *BitMap) ColorQuantize(quantize FREE_IMAGE_QUANTIZE) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ColorQuantize_, inArgs{&dib, &quantize}).PtrFree())
}

var _func_FreeImage_ColorQuantizeEx_ = &c.FuncPrototype{Name: "FreeImage_ColorQuantizeEx", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ColorQuantizeEx(FIBITMAP *dib, FREE_IMAGE_QUANTIZE quantize FI_DEFAULT(FIQ_WUQUANT), int PaletteSize FI_DEFAULT(256), int ReserveSize FI_DEFAULT(0), RGBQUAD *ReservePalette FI_DEFAULT(NULL));
func (dib *BitMap) ColorQuantizeEx(quantize FREE_IMAGE_QUANTIZE, PaletteSize, ReserveSize int32, ReservePalette *RGBQUAD) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ColorQuantizeEx_, inArgs{&dib, &quantize, &PaletteSize, &ReserveSize, &ReservePalette}).PtrFree())
}

var _func_FreeImage_Threshold_ = &c.FuncPrototype{Name: "FreeImage_Threshold", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.U8}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Threshold(FIBITMAP *dib, BYTE T);
func (dib *BitMap) Threshold(t byte) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Threshold_, inArgs{&dib, &t}).PtrFree())
}

var _func_FreeImage_Dither_ = &c.FuncPrototype{Name: "FreeImage_Dither", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Dither(FIBITMAP *dib, FREE_IMAGE_DITHER algorithm);
func (dib *BitMap) Dither(algorithm FREE_IMAGE_DITHER) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Dither_, inArgs{&dib, &algorithm}).PtrFree())
}

var _func_FreeImage_ConvertFromRawBits_ = &c.FuncPrototype{Name: "FreeImage_ConvertFromRawBits", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.U32, c.U32, c.U32, c.U32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertFromRawBits(BYTE *bits, int width, int height, int pitch, unsigned bpp, unsigned red_mask, unsigned green_mask, unsigned blue_mask, BOOL topdown FI_DEFAULT(FALSE));
func ConvertFromRawBits(bits *byte, width, height, pitch int32, bpp uint32, red_mask, green_mask, blue_mask uint32, topdown bool) *BitMap {
	td := c.CBool(topdown)
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertFromRawBits_, inArgs{&bits, &width, &height, &pitch, &bpp, &red_mask, &green_mask, &blue_mask, &td}).PtrFree())
}

var _func_FreeImage_ConvertFromRawBitsEx_ = &c.FuncPrototype{Name: "FreeImage_ConvertFromRawBitsEx", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.I32, c.I32, c.I32, c.I32, c.U32, c.U32, c.U32, c.U32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertFromRawBitsEx(BOOL copySource, BYTE *bits, FREE_IMAGE_TYPE type, int width, int height, int pitch, unsigned bpp, unsigned red_mask, unsigned green_mask, unsigned blue_mask, BOOL topdown FI_DEFAULT(FALSE));
func ConvertFromRawBitsEx(copySource bool, bits *byte, width, height, pitch int32, bpp uint32, red_mask, green_mask, blue_mask uint32, topdown bool) *BitMap {
	cs, td := c.CBool(copySource), c.CBool(topdown)
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertFromRawBitsEx_, inArgs{&cs, &bits, &width, &height, &pitch, &bpp, &red_mask, &green_mask, &blue_mask, &td}).PtrFree())
}

var _func_FreeImage_ConvertToRawBits_ = &c.FuncPrototype{Name: "FreeImage_ConvertToRawBits", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.U32, c.U32, c.U32, c.U32, c.I32}}

// DLL_API void DLL_CALLCONV FreeImage_ConvertToRawBits(BYTE *bits, FIBITMAP *dib, int pitch, unsigned bpp, unsigned red_mask, unsigned green_mask, unsigned blue_mask, BOOL topdown FI_DEFAULT(FALSE));
func ConvertToRawBits(bits *byte, dib *BitMap, pitch int32, bpp uint32, red_mask, green_mask, blue_mask uint32, topdown bool) {
	td := c.CBool(topdown)
	fiLib.Call(_func_FreeImage_ConvertToRawBits_, inArgs{&bits, &dib, &pitch, &bpp, &red_mask, &green_mask, &blue_mask, &td})
}

var _func_FreeImage_ConvertToFloat_ = &c.FuncPrototype{Name: "FreeImage_ConvertToFloat", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToFloat(FIBITMAP *dib);
func (dib *BitMap) ConvertToFloat() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToFloat_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToRGBF_ = &c.FuncPrototype{Name: "FreeImage_ConvertToRGBF", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToRGBF(FIBITMAP *dib);
func (dib *BitMap) ConvertToRGBF() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToRGBF_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToRGBAF_ = &c.FuncPrototype{Name: "FreeImage_ConvertToRGBAF", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToRGBAF(FIBITMAP *dib);
func (dib *BitMap) ConvertToRGBAF() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToRGBAF_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToUINT16_ = &c.FuncPrototype{Name: "FreeImage_ConvertToUINT16", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToUINT16(FIBITMAP *dib);
func (dib *BitMap) ConvertToUINT16() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToUINT16_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToRGB16_ = &c.FuncPrototype{Name: "FreeImage_ConvertToRGB16", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToRGB16(FIBITMAP *dib);
func (dib *BitMap) ConvertToRGB16() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToRGB16_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToRGBA16_ = &c.FuncPrototype{Name: "FreeImage_ConvertToRGBA16", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToRGBA16(FIBITMAP *dib);
func (dib *BitMap) ConvertToRGBA16() *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToRGBA16_, inArgs{&dib}).PtrFree())
}

var _func_FreeImage_ConvertToStandardType_ = &c.FuncPrototype{Name: "FreeImage_ConvertToStandardType", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToStandardType(FIBITMAP *src, BOOL scale_linear FI_DEFAULT(TRUE));
func (dib *BitMap) ConvertToStandardType(scale_linear bool) *BitMap {
	sl := c.CBool(scale_linear)
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToStandardType_, inArgs{&dib, &sl}).PtrFree())
}

var _func_FreeImage_ConvertToType_ = &c.FuncPrototype{Name: "FreeImage_ConvertToType", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ConvertToType(FIBITMAP *src, FREE_IMAGE_TYPE dst_type, BOOL scale_linear FI_DEFAULT(TRUE));
func (dib *BitMap) ConvertToType(dst_type FREE_IMAGE_TYPE, scale_linear bool) *BitMap {
	sl := c.CBool(scale_linear)
	return (*BitMap)(fiLib.Call(_func_FreeImage_ConvertToType_, inArgs{&dib, &dst_type, &sl}).PtrFree())
}

// Tone mapping operators ---------------------------------------------------

var _func_FreeImage_ToneMapping_ = &c.FuncPrototype{Name: "FreeImage_ToneMapping", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.F64, c.F64}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_ToneMapping(FIBITMAP *dib, FREE_IMAGE_TMO tmo, double first_param FI_DEFAULT(0), double second_param FI_DEFAULT(0));
func (dib *BitMap) ToneMapping(tom FREE_IMAGE_TMO, first_param, second_param float64) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_ToneMapping_, inArgs{&dib, &tom, &first_param, &second_param}).PtrFree())
}

var _func_FreeImage_TmoDrago03_ = &c.FuncPrototype{Name: "FreeImage_TmoDrago03", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.F64}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_TmoDrago03(FIBITMAP *src, double gamma FI_DEFAULT(2.2), double exposure FI_DEFAULT(0));
func (dib *BitMap) TmoDrago03(gamma, exposure float64) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_TmoDrago03_, inArgs{&dib, &gamma, &exposure}).PtrFree())
}

var _func_FreeImage_TmoReinhard05_ = &c.FuncPrototype{Name: "FreeImage_TmoReinhard05", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.F64}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_TmoReinhard05(FIBITMAP *src, double intensity FI_DEFAULT(0), double contrast FI_DEFAULT(0));
func (dib *BitMap) TmoReinhard05(intensity, contrast float64) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_TmoReinhard05_, inArgs{&dib, &intensity, &contrast}).PtrFree())
}

var _func_FreeImage_TmoReinhard05Ex_ = &c.FuncPrototype{Name: "FreeImage_TmoReinhard05Ex", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.F64, c.F64, c.F64}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_TmoReinhard05Ex(FIBITMAP *src, double intensity FI_DEFAULT(0), double contrast FI_DEFAULT(0), double adaptation FI_DEFAULT(1), double color_correction FI_DEFAULT(0));
func (dib *BitMap) TmoReinhard05Ex(intensity, contrast, adaptation, color_correction float64) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_TmoReinhard05Ex_, inArgs{&dib, &intensity, &contrast, &adaptation, &color_correction}).PtrFree())
}

var _func_FreeImage_TmoFattal02_ = &c.FuncPrototype{Name: "FreeImage_TmoFattal02", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.F64}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_TmoFattal02(FIBITMAP *src, double color_saturation FI_DEFAULT(0.5), double attenuation FI_DEFAULT(0.85));
func (dib *BitMap) TmoFattal02(color_saturation, attenuation float64) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_TmoFattal02_, inArgs{&dib, &color_saturation, &attenuation}).PtrFree())
}

// ZLib interface -----------------------------------------------------------

var _func_FreeImage_ZLibCompress_ = &c.FuncPrototype{Name: "FreeImage_ZLibCompress", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.Pointer, c.U32}}

// DLL_API DWORD DLL_CALLCONV FreeImage_ZLibCompress(BYTE *target, DWORD target_size, BYTE *source, DWORD source_size);
func ZLibCompress(target []byte, target_size uint32, source []byte, source_size uint32) uint32 {
	t, s := &target[0], &source[0]
	return fiLib.Call(_func_FreeImage_ZLibCompress_, inArgs{&t, &target_size, &s, &source_size}).U32Free()
}

var _func_FreeImage_ZLibUncompress_ = &c.FuncPrototype{Name: "FreeImage_ZLibUncompress", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.Pointer, c.U32}}

// DLL_API DWORD DLL_CALLCONV FreeImage_ZLibUncompress(BYTE *target, DWORD target_size, BYTE *source, DWORD source_size);
func ZLibUncompress(target []byte, target_size uint32, source []byte, source_size uint32) uint32 {
	t, s := &target[0], &source[0]
	return fiLib.Call(_func_FreeImage_ZLibUncompress_, inArgs{&t, &target_size, &s, &source_size}).U32Free()
}

var _func_FreeImage_ZLibGZip_ = &c.FuncPrototype{Name: "FreeImage_ZLibGZip", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.Pointer, c.U32}}

// DLL_API DWORD DLL_CALLCONV FreeImage_ZLibGZip(BYTE *target, DWORD target_size, BYTE *source, DWORD source_size);
func ZLibGZip(target []byte, target_size uint32, source []byte, source_size uint32) uint32 {
	t, s := &target[0], &source[0]
	return fiLib.Call(_func_FreeImage_ZLibGZip_, inArgs{&t, &target_size, &s, &source_size}).U32Free()
}

var _func_FreeImage_ZLibGUnzip_ = &c.FuncPrototype{Name: "FreeImage_ZLibGUnzip", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.U32, c.Pointer, c.U32}}

// DLL_API DWORD DLL_CALLCONV FreeImage_ZLibGUnzip(BYTE *target, DWORD target_size, BYTE *source, DWORD source_size);
func ZLibGUnzip(target []byte, target_size uint32, source []byte, source_size uint32) uint32 {
	t, s := &target[0], &source[0]
	return fiLib.Call(_func_FreeImage_ZLibGUnzip_, inArgs{&t, &target_size, &s, &source_size}).U32Free()
}

var _func_FreeImage_ZLibCRC32_ = &c.FuncPrototype{Name: "FreeImage_ZLibCRC32", OutType: c.U32, InTypes: []c.Type{c.U32, c.Pointer, c.U32}}

// DLL_API DWORD DLL_CALLCONV FreeImage_ZLibCRC32(DWORD crc, BYTE *source, DWORD source_size);
func ZLibCRC32(crc uint32, source []byte, source_size uint32) uint32 {
	s := &source[0]
	return fiLib.Call(_func_FreeImage_ZLibCRC32_, inArgs{&crc, &s, &source_size}).U32Free()
}

// --------------------------------------------------------------------------
// Metadata routines
// --------------------------------------------------------------------------

// tag creation / destruction

var _func_FreeImage_CreateTag_ = &c.FuncPrototype{Name: "FreeImage_CreateTag", OutType: c.Pointer, InTypes: nil}

// DLL_API FITAG *DLL_CALLCONV FreeImage_CreateTag(void);
func CreateTag() *Tag {
	return (*Tag)(fiLib.Call(_func_FreeImage_CreateTag_, nil).PtrFree())
}

var _func_FreeImage_DeleteTag_ = &c.FuncPrototype{Name: "FreeImage_DeleteTag", OutType: c.Void, InTypes: []c.Type{c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_DeleteTag(FITAG *tag);
func (tag *Tag) DeleteTag() {
	fiLib.Call(_func_FreeImage_DeleteTag_, inArgs{&tag})
}

var _func_FreeImage_CloneTag_ = &c.FuncPrototype{Name: "FreeImage_CloneTag", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API FITAG *DLL_CALLCONV FreeImage_CloneTag(FITAG *tag);
func (tag *Tag) CloneTag() *Tag {
	return (*Tag)(fiLib.Call(_func_FreeImage_CloneTag_, inArgs{&tag}).PtrFree())
}

var _func_FreeImage_GetTagKey_ = &c.FuncPrototype{Name: "FreeImage_GetTagKey", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API const char *DLL_CALLCONV FreeImage_GetTagKey(FITAG *tag);
func (tag *Tag) GetTagKey() string {
	return fiLib.Call(_func_FreeImage_GetTagKey_, inArgs{&tag}).StrFree()
}

var _func_FreeImage_GetTagDescription_ = &c.FuncPrototype{Name: "FreeImage_GetTagDescription", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API const char *DLL_CALLCONV FreeImage_GetTagDescription(FITAG *tag);
func (tag *Tag) GetTagDescription() string {
	return fiLib.Call(_func_FreeImage_GetTagDescription_, inArgs{&tag}).StrFree()
}

var _func_FreeImage_GetTagID_ = &c.FuncPrototype{Name: "FreeImage_GetTagID", OutType: c.U16, InTypes: []c.Type{c.Pointer}}

// DLL_API WORD DLL_CALLCONV FreeImage_GetTagID(FITAG *tag);
func (tag *Tag) GetTagID() uint16 {
	return fiLib.Call(_func_FreeImage_GetTagID_, inArgs{&tag}).U16Free()
}

var _func_FreeImage_GetTagType_ = &c.FuncPrototype{Name: "FreeImage_GetTagType", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API FREE_IMAGE_MDTYPE DLL_CALLCONV FreeImage_GetTagType(FITAG *tag);
func (tag *Tag) GetTagType() FREE_IMAGE_MDTYPE {
	return FREE_IMAGE_MDTYPE(fiLib.Call(_func_FreeImage_GetTagType_, inArgs{&tag}).I32Free())
}

var _func_FreeImage_GetTagCount_ = &c.FuncPrototype{Name: "FreeImage_GetTagCount", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API DWORD DLL_CALLCONV FreeImage_GetTagCount(FITAG *tag);
func (tag *Tag) GetTagCount() uint32 {
	return fiLib.Call(_func_FreeImage_GetTagCount_, inArgs{&tag}).U32Free()
}

var _func_FreeImage_GetTagLength_ = &c.FuncPrototype{Name: "FreeImage_GetTagLength", OutType: c.U32, InTypes: []c.Type{c.Pointer}}

// DLL_API DWORD DLL_CALLCONV FreeImage_GetTagLength(FITAG *tag);
func (tag *Tag) GetTagLength() uint32 {
	return fiLib.Call(_func_FreeImage_GetTagLength_, inArgs{&tag}).U32Free()
}

var _func_FreeImage_GetTagValue_ = &c.FuncPrototype{Name: "FreeImage_GetTagValue", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}

// DLL_API const void *DLL_CALLCONV FreeImage_GetTagValue(FITAG *tag);
func (tag *Tag) GetTagValue() unsafe.Pointer {
	return fiLib.Call(_func_FreeImage_GetTagValue_, inArgs{&tag}).PtrFree()
}

var _func_FreeImage_SetTagKey_ = &c.FuncPrototype{Name: "FreeImage_SetTagKey", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagKey(FITAG *tag, const char *key);
func (tag *Tag) SetTagKey(key string) bool {
	k := c.CStr(key)
	defer c.Free(k)
	return fiLib.Call(_func_FreeImage_SetTagKey_, inArgs{&tag, &k}).BoolFree()
}

var _func_FreeImage_SetTagDescription_ = &c.FuncPrototype{Name: "FreeImage_SetTagDescription", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagDescription(FITAG *tag, const char *description);
func (tag *Tag) SetTagDescription(description string) bool {
	des := c.CStr(description)
	defer c.Free(des)
	return fiLib.Call(_func_FreeImage_SetTagDescription_, inArgs{&tag, &des}).BoolFree()
}

var _func_FreeImage_SetTagID_ = &c.FuncPrototype{Name: "FreeImage_SetTagID", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U16}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagID(FITAG *tag, WORD id);
func (tag *Tag) SetTagID(id uint16) bool {
	return fiLib.Call(_func_FreeImage_SetTagID_, inArgs{&tag, &id}).BoolFree()
}

var _func_FreeImage_SetTagType_ = &c.FuncPrototype{Name: "FreeImage_SetTagType", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagType(FITAG *tag, FREE_IMAGE_MDTYPE type);
func (tag *Tag) SetTagType(typ FREE_IMAGE_MDTYPE) bool {
	return fiLib.Call(_func_FreeImage_SetTagType_, inArgs{&tag, &typ}).BoolFree()
}

var _func_FreeImage_SetTagCount_ = &c.FuncPrototype{Name: "FreeImage_SetTagCount", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagCount(FITAG *tag, DWORD count);
func (tag *Tag) SetTagCount(count uint32) bool {
	return fiLib.Call(_func_FreeImage_SetTagCount_, inArgs{&tag, &count}).BoolFree()
}

var _func_FreeImage_SetTagLength_ = &c.FuncPrototype{Name: "FreeImage_SetTagLength", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.U32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagLength(FITAG *tag, DWORD length);
func (tag *Tag) SetTagLength(length uint32) bool {
	return fiLib.Call(_func_FreeImage_SetTagLength_, inArgs{&tag, &length}).BoolFree()
}

var _func_FreeImage_SetTagValue_ = &c.FuncPrototype{Name: "FreeImage_SetTagValue", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetTagValue(FITAG *tag, const void *value);
func (tag *Tag) SetTagValue(value unsafe.Pointer) bool {
	return fiLib.Call(_func_FreeImage_SetTagValue_, inArgs{&tag, &value}).BoolFree()
}

// iterator

var _func_FreeImage_FindFirstMetadata_ = &c.FuncPrototype{Name: "FreeImage_FindFirstMetadata", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer}}

// DLL_API FIMETADATA *DLL_CALLCONV FreeImage_FindFirstMetadata(FREE_IMAGE_MDMODEL model, FIBITMAP *dib, FITAG **tag);
func (dib *BitMap) FindFirstMetadata(model FREE_IMAGE_MDMODEL) (metaData *MetaData, tag *Tag) {
	t := &tag
	metaData = (*MetaData)(fiLib.Call(_func_FreeImage_FindFirstMetadata_, inArgs{&model, &dib, &t}).PtrFree())
	return
}

var _func_FreeImage_FindNextMetadata_ = &c.FuncPrototype{Name: "FreeImage_FindNextMetadata", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_FindNextMetadata(FIMETADATA *mdhandle, FITAG **tag);
func (mdhandle *MetaData) FindNextMetadata() (tag *Tag, ok bool) {
	t := &tag
	ok = fiLib.Call(_func_FreeImage_FindNextMetadata_, inArgs{&mdhandle, &t}).BoolFree()
	return
}

var _func_FreeImage_FindCloseMetadata_ = &c.FuncPrototype{Name: "FreeImage_FindCloseMetadata", OutType: c.Void, InTypes: []c.Type{c.Pointer}}

// DLL_API void DLL_CALLCONV FreeImage_FindCloseMetadata(FIMETADATA *mdhandle);
func (mdhandle *MetaData) FindCloseMetadata() {
	fiLib.Call(_func_FreeImage_FindCloseMetadata_, inArgs{&mdhandle})
}

// metadata setter and getter

var _func_FreeImage_SetMetadata_ = &c.FuncPrototype{Name: "FreeImage_SetMetadata", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetMetadata(FREE_IMAGE_MDMODEL model, FIBITMAP *dib, const char *key, FITAG *tag);
func (dib *BitMap) SetMetadata(model FREE_IMAGE_MDMODEL, key string, tag *Tag) bool {
	k := c.CStr(key)
	defer c.Free(k)
	return fiLib.Call(_func_FreeImage_SetMetadata_, inArgs{&model, &dib, &k, &tag}).BoolFree()
}

var _func_FreeImage_GetMetadata_ = &c.FuncPrototype{Name: "FreeImage_GetMetadata", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetMetadata(FREE_IMAGE_MDMODEL model, FIBITMAP *dib, const char *key, FITAG **tag);
func (dib *BitMap) GetMetadata(model FREE_IMAGE_MDMODEL, key string) (tag *Tag, ok bool) {
	k, t := c.CStr(key), &tag
	defer c.Free(k)
	ok = fiLib.Call(_func_FreeImage_GetMetadata_, inArgs{&model, &dib, &k, &t}).BoolFree()
	return
}

var _func_FreeImage_SetMetadataKeyValue_ = &c.FuncPrototype{Name: "FreeImage_SetMetadataKeyValue", OutType: c.I32, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetMetadataKeyValue(FREE_IMAGE_MDMODEL model, FIBITMAP *dib, const char *key, const char *value);
func (dib *BitMap) SetMetadataKeyValue(model FREE_IMAGE_MDMODEL, key, value string) bool {
	k, v := c.CStr(key), c.CStr(value)
	defer c.Free(k)
	defer c.Free(v)
	return fiLib.Call(_func_FreeImage_SetMetadataKeyValue_, inArgs{&model, &dib, &k, &v}).BoolFree()
}

// helpers

var _func_FreeImage_GetMetadataCount_ = &c.FuncPrototype{Name: "FreeImage_GetMetadataCount", OutType: c.U32, InTypes: []c.Type{c.I32, c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_GetMetadataCount(FREE_IMAGE_MDMODEL model, FIBITMAP *dib);
func (dib *BitMap) GetMetadataCount(model FREE_IMAGE_MDMODEL) uint32 {
	return fiLib.Call(_func_FreeImage_GetMetadataCount_, inArgs{&model, &dib}).U32Free()
}

var _func_FreeImage_CloneMetadata_ = &c.FuncPrototype{Name: "FreeImage_CloneMetadata", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_CloneMetadata(FIBITMAP *dst, FIBITMAP *src);
func (src *BitMap) CloneMetadataTo(dst *BitMap) bool {
	return fiLib.Call(_func_FreeImage_CloneMetadata_, inArgs{&dst, &src}).BoolFree()
}

// tag to C string conversion

var _func_FreeImage_TagToString_ = &c.FuncPrototype{Name: "FreeImage_TagToString", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer}}

// DLL_API const char* DLL_CALLCONV FreeImage_TagToString(FREE_IMAGE_MDMODEL model, FITAG *tag, char *Make FI_DEFAULT(NULL));
func TagToString(model FREE_IMAGE_MDMODEL, tag *Tag, Make ...string) string {
	m := unsafe.Pointer(nil)
	if Make != nil {
		m = c.CStr(Make[0])
		defer c.Free(m)
	}
	return fiLib.Call(_func_FreeImage_TagToString_, inArgs{&model, &tag, &m}).StrFree()
}

// --------------------------------------------------------------------------
// JPEG lossless transformation routines
// --------------------------------------------------------------------------

var _func_FreeImage_JPEGTransform_ = &c.FuncPrototype{Name: "FreeImage_JPEGTransform", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGTransform(const char *src_file, const char *dst_file, FREE_IMAGE_JPEG_OPERATION operation, BOOL perfect FI_DEFAULT(TRUE));
func JPEGTransform(src_file, dst_file string, opera FREE_IMAGE_JPEG_OPERATION, perfect bool) bool {
	sf, df, pf := c.CStr(src_file), c.CStr(dst_file), c.CBool(perfect)
	defer c.Free(sf)
	defer c.Free(df)
	return fiLib.Call(_func_FreeImage_JPEGTransform_, inArgs{&sf, &df, &opera, &pf}).BoolFree()
}

var _func_FreeImage_JPEGCrop_ = &c.FuncPrototype{Name: "FreeImage_JPEGCrop", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.I32, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGCrop(const char *src_file, const char *dst_file, int left, int top, int right, int bottom);
func JPEGCrop(src_file, dst_file string, left, top, right, bottom int32) bool {
	sf, df := c.CStr(src_file), c.CStr(dst_file)
	defer c.Free(sf)
	defer c.Free(df)
	return fiLib.Call(_func_FreeImage_JPEGCrop_, inArgs{&sf, &df, &left, &top, &right, &bottom}).BoolFree()
}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGCropU(const wchar_t *src_file, const wchar_t *dst_file, int left, int top, int right, int bottom);
// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGTransformFromHandle(FreeImageIO* src_io, fi_handle src_handle, FreeImageIO* dst_io, fi_handle dst_handle, FREE_IMAGE_JPEG_OPERATION operation, int* left, int* top, int* right, int* bottom, BOOL perfect FI_DEFAULT(TRUE));

var _func_FreeImage_JPEGTransformCombined_ = &c.FuncPrototype{Name: "FreeImage_JPEGTransformCombined", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGTransformCombined(const char *src_file, const char *dst_file, FREE_IMAGE_JPEG_OPERATION operation, int* left, int* top, int* right, int* bottom, BOOL perfect FI_DEFAULT(TRUE));
func JPEGTransformCombined(src_file, dst_file string, perfect bool) (left, top, right, bottom int32, ok bool) {
	sf, df, pf := c.CStr(src_file), c.CStr(dst_file), c.CBool(perfect)
	defer c.Free(sf)
	defer c.Free(df)
	_left, _top, _right, _bottom := left, top, right, bottom
	ok = fiLib.Call(_func_FreeImage_JPEGTransformCombined_, inArgs{&sf, &df, &_left, &_top, &_right, &_bottom, &pf}).BoolFree()
	return
}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGTransformCombinedU(const wchar_t *src_file, const wchar_t *dst_file, FREE_IMAGE_JPEG_OPERATION operation, int* left, int* top, int* right, int* bottom, BOOL perfect FI_DEFAULT(TRUE));

var _func_FreeImage_JPEGTransformCombinedFromMemory_ = &c.FuncPrototype{Name: "FreeImage_JPEGTransformCombinedFromMemory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_JPEGTransformCombinedFromMemory(FIMEMORY* src_stream, FIMEMORY* dst_stream, FREE_IMAGE_JPEG_OPERATION operation, int* left, int* top, int* right, int* bottom, BOOL perfect FI_DEFAULT(TRUE));
func JPEGTransformCombinedFromMemory(src_stream, dst_stream *Memory, perfect bool) (left, top, right, bottom int32, ok bool) {
	pf := c.CBool(perfect)
	_left, _top, _right, _bottom := left, top, right, bottom
	ok = fiLib.Call(_func_FreeImage_JPEGTransformCombinedFromMemory_, inArgs{&src_stream, &dst_stream, &_left, &_top, &_right, &_bottom, &pf}).BoolFree()
	return
}

// rotation and flipping

var _func_FreeImage_Rotate_ = &c.FuncPrototype{Name: "FreeImage_Rotate", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.Pointer}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Rotate(FIBITMAP *dib, double angle, const void *bkcolor FI_DEFAULT(NULL));
func (dib *BitMap) Rotate(angle float64, bkcolor unsafe.Pointer) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Rotate_, inArgs{&dib, &angle, &bkcolor}).PtrFree())
}

var _func_FreeImage_RotateEx_ = &c.FuncPrototype{Name: "FreeImage_RotateEx", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.F64, c.F64, c.F64, c.F64, c.F64, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_RotateEx(FIBITMAP *dib, double angle, double x_shift, double y_shift, double x_origin, double y_origin, BOOL use_mask);
func (dib *BitMap) RotateEx(angle float64, x_shift, y_shift, x_origin, y_origin float64, use_mask bool) *BitMap {
	um := c.CBool(use_mask)
	return (*BitMap)(fiLib.Call(_func_FreeImage_RotateEx_, inArgs{&dib, &angle, &x_shift, &y_shift, &x_origin, &y_origin, &um}).PtrFree())
}

var _func_FreeImage_FlipHorizontal_ = &c.FuncPrototype{Name: "FreeImage_FlipHorizontal", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_FlipHorizontal(FIBITMAP *dib);
func (dib *BitMap) FlipHorizontal() bool {
	return fiLib.Call(_func_FreeImage_FlipHorizontal_, inArgs{&dib}).BoolFree()
}

var _func_FreeImage_FlipVertical_ = &c.FuncPrototype{Name: "FreeImage_FlipVertical", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_FlipVertical(FIBITMAP *dib);
func (dib *BitMap) FlipVertical() bool {
	return fiLib.Call(_func_FreeImage_FlipVertical_, inArgs{&dib}).BoolFree()
}

// upsampling / downsampling

var _func_FreeImage_Rescale_ = &c.FuncPrototype{Name: "FreeImage_Rescale", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Rescale(FIBITMAP *dib, int dst_width, int dst_height, FREE_IMAGE_FILTER filter FI_DEFAULT(FILTER_CATMULLROM));
func (dib *BitMap) Rescale(dst_width, dst_height int32, filter FREE_IMAGE_FILTER) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Rescale_, inArgs{&dib, &dst_width, &dst_height, &filter}).PtrFree())
}

var _func_FreeImage_MakeThumbnail_ = &c.FuncPrototype{Name: "FreeImage_MakeThumbnail", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_MakeThumbnail(FIBITMAP *dib, int max_pixel_size, BOOL convert FI_DEFAULT(TRUE));
func (dib *BitMap) MakeThumbnail(max_pixel_size int32, convert bool) *BitMap {
	cvt := c.CBool(convert)
	return (*BitMap)(fiLib.Call(_func_FreeImage_MakeThumbnail_, inArgs{&dib, &max_pixel_size, &cvt}).PtrFree())
}

var _func_FreeImage_RescaleRect_ = &c.FuncPrototype{Name: "FreeImage_RescaleRect", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_RescaleRect(FIBITMAP *dib, int dst_width, int dst_height, int left, int top, int right, int bottom, FREE_IMAGE_FILTER filter FI_DEFAULT(FILTER_CATMULLROM), unsigned flags FI_DEFAULT(0));
func (dib *BitMap) RescaleRect(dst_width, dst_height int32, left, top, right, bottom int32, filter FREE_IMAGE_FILTER, flags uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_RescaleRect_, inArgs{&dib, &dst_width, &dst_height, &left, &top, &right, &bottom, &filter, &flags}).PtrFree())
}

// color manipulation routines (point operations)

var _func_FreeImage_AdjustCurve_ = &c.FuncPrototype{Name: "FreeImage_AdjustCurve", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AdjustCurve(FIBITMAP *dib, BYTE *LUT, FREE_IMAGE_COLOR_CHANNEL channel);
func (dib *BitMap) AdjustCurve(Lut [256]byte, channel FREE_IMAGE_COLOR_CHANNEL) bool {
	l := &Lut[0]
	return fiLib.Call(_func_FreeImage_AdjustCurve_, inArgs{&dib, &l, &channel}).BoolFree()
}

var _func_FreeImage_AdjustGamma_ = &c.FuncPrototype{Name: "FreeImage_AdjustGamma", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.F64}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AdjustGamma(FIBITMAP *dib, double gamma);
func (dib *BitMap) AdjustGamma(gamma float64) bool {

	return fiLib.Call(_func_FreeImage_AdjustGamma_, inArgs{&dib, &gamma}).BoolFree()
}

var _func_FreeImage_AdjustBrightness_ = &c.FuncPrototype{Name: "FreeImage_AdjustBrightness", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.F64}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AdjustBrightness(FIBITMAP *dib, double percentage);
func (dib *BitMap) AdjustBrightness(percentage float64) bool {
	return fiLib.Call(_func_FreeImage_AdjustBrightness_, inArgs{&dib, &percentage}).BoolFree()
}

var _func_FreeImage_AdjustContrast_ = &c.FuncPrototype{Name: "FreeImage_AdjustContrast", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.F64}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AdjustContrast(FIBITMAP *dib, double percentage);
func (dib *BitMap) AdjustContrast(percentage float64) bool {
	return fiLib.Call(_func_FreeImage_AdjustContrast_, inArgs{&dib, &percentage}).BoolFree()
}

var _func_FreeImage_Invert_ = &c.FuncPrototype{Name: "FreeImage_Invert", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_Invert(FIBITMAP *dib);
func (dib *BitMap) Invert(percentage float64) bool {
	return fiLib.Call(_func_FreeImage_Invert_, inArgs{&dib, &percentage}).BoolFree()
}

var _func_FreeImage_GetHistogram_ = &c.FuncPrototype{Name: "FreeImage_GetHistogram", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_GetHistogram(FIBITMAP *dib, DWORD *histo, FREE_IMAGE_COLOR_CHANNEL channel FI_DEFAULT(FICC_BLACK));
func (dib *BitMap) GetHistogram(channel FREE_IMAGE_COLOR_CHANNEL) (histo [256]uint32, ok bool) {
	h := &histo[0]
	ok = fiLib.Call(_func_FreeImage_GetHistogram_, inArgs{&dib, &h, &channel}).BoolFree()
	return
}

var _func_FreeImage_GetAdjustColorsLookupTable_ = &c.FuncPrototype{Name: "FreeImage_GetAdjustColorsLookupTable", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.F64, c.F64, c.F64, c.I32}}

// DLL_API int DLL_CALLCONV FreeImage_GetAdjustColorsLookupTable(BYTE *LUT, double brightness, double contrast, double gamma, BOOL invert);
func GetAdjustColorsLookupTable(Lut [256]byte, brightness, contrast, gamma float64, invert bool) int32 {
	l, i := &Lut[0], c.CBool(invert)
	return fiLib.Call(_func_FreeImage_GetAdjustColorsLookupTable_, inArgs{&l, &brightness, &contrast, &gamma, &i}).I32Free()
}

var _func_FreeImage_AdjustColors_ = &c.FuncPrototype{Name: "FreeImage_AdjustColors", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.F64, c.F64, c.F64, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_AdjustColors(FIBITMAP *dib, double brightness, double contrast, double gamma, BOOL invert FI_DEFAULT(FALSE));
func (dib *BitMap) AdjustColors(brightness, contrast, gamma float64, invert bool) bool {
	ivt := c.CBool(invert)
	return fiLib.Call(_func_FreeImage_AdjustColors_, inArgs{&dib, &brightness, &contrast, &gamma, &ivt}).BoolFree()
}

var _func_FreeImage_ApplyColorMapping_ = &c.FuncPrototype{Name: "FreeImage_ApplyColorMapping", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.U32, c.I32, c.I32}}

// DLL_API unsigned DLL_CALLCONV FreeImage_ApplyColorMapping(FIBITMAP *dib, RGBQUAD *srccolors, RGBQUAD *dstcolors, unsigned count, BOOL ignore_alpha, BOOL swap);
func (dib *BitMap) ApplyColorMapping(srccolors, dstcolors []RGBQUAD, ignore_alpha, swap bool) uint32 {
	sc, dc, ia, sw := &srccolors[0], &dstcolors[0], c.CBool(ignore_alpha), c.CBool(swap)
	count := uint32(len(dstcolors))
	if len(srccolors) < len(dstcolors) {
		count = uint32(len(srccolors))
	}

	return fiLib.Call(_func_FreeImage_ApplyColorMapping_, inArgs{&dib, &sc, &dc, &count, &ia, &sw}).U32Free()
}

var _func_FreeImage_SwapColors_ = &c.FuncPrototype{Name: "FreeImage_SwapColors", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.I32}}

// DLL_API unsigned DLL_CALLCONV FreeImage_SwapColors(FIBITMAP *dib, RGBQUAD *color_a, RGBQUAD *color_b, BOOL ignore_alpha);
func (dib *BitMap) SwapColors(color_a, color_b *RGBQUAD, ignore_alpha bool) uint32 {
	ia := c.CBool(ignore_alpha)
	return fiLib.Call(_func_FreeImage_SwapColors_, inArgs{&dib, &color_a, &color_b, &ia}).U32Free()
}

var _func_FreeImage_ApplyPaletteIndexMapping_ = &c.FuncPrototype{Name: "FreeImage_ApplyPaletteIndexMapping", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.U32, c.I32}}

// DLL_API unsigned DLL_CALLCONV FreeImage_ApplyPaletteIndexMapping(FIBITMAP *dib, BYTE *srcindices,    BYTE *dstindices, unsigned count, BOOL swap);
func (dib *BitMap) ApplyPaletteIndexMapping(srcindices, dstindices []byte, swap bool) uint32 {
	si, di, sw := &srcindices[0], &dstindices[0], c.CBool(swap)
	count := uint32(len(dstindices))
	if len(srcindices) < len(dstindices) {
		count = uint32(len(srcindices))
	}
	return fiLib.Call(_func_FreeImage_ApplyPaletteIndexMapping_, inArgs{&dib, &si, &di, &count, &sw}).U32Free()
}

var _func_FreeImage_SwapPaletteIndices_ = &c.FuncPrototype{Name: "FreeImage_SwapPaletteIndices", OutType: c.U32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}

// DLL_API unsigned DLL_CALLCONV FreeImage_SwapPaletteIndices(FIBITMAP *dib, BYTE *index_a, BYTE *index_b);
func (dib *BitMap) SwapPaletteIndices(index_a, index_b byte) uint32 {
	a, b := &index_a, &index_b
	return fiLib.Call(_func_FreeImage_SwapPaletteIndices_, inArgs{&dib, &a, &b}).U32Free()
}

// // channel processing routines

var _func_FreeImage_GetChannel_ = &c.FuncPrototype{Name: "FreeImage_GetChannel", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_GetChannel(FIBITMAP *dib, FREE_IMAGE_COLOR_CHANNEL channel);
func (dib *BitMap) GetChannel(channel FREE_IMAGE_COLOR_CHANNEL) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_GetChannel_, inArgs{&dib, &channel}).PtrFree())
}

var _func_FreeImage_SetChannel_ = &c.FuncPrototype{Name: "FreeImage_SetChannel", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetChannel(FIBITMAP *dst, FIBITMAP *src, FREE_IMAGE_COLOR_CHANNEL channel);
func (dib *BitMap) SetChannel(src *BitMap, channel FREE_IMAGE_COLOR_CHANNEL) bool {
	return fiLib.Call(_func_FreeImage_SetChannel_, inArgs{&dib, &src, &channel}).BoolFree()
}

var _func_FreeImage_GetComplexChannel_ = &c.FuncPrototype{Name: "FreeImage_GetComplexChannel", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_GetComplexChannel(FIBITMAP *src, FREE_IMAGE_COLOR_CHANNEL channel);
func (dib *BitMap) GetComplexChannel(channel FREE_IMAGE_COLOR_CHANNEL) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_GetComplexChannel_, inArgs{&dib, &channel}).PtrFree())
}

var _func_FreeImage_SetComplexChannel_ = &c.FuncPrototype{Name: "FreeImage_SetComplexChannel", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_SetComplexChannel(FIBITMAP *dst, FIBITMAP *src, FREE_IMAGE_COLOR_CHANNEL channel);
func (dib *BitMap) SetComplexChannel(src *BitMap, channel FREE_IMAGE_COLOR_CHANNEL) bool {
	return fiLib.Call(_func_FreeImage_SetComplexChannel_, inArgs{&dib, &src, &channel}).BoolFree()
}

// // copy / paste / composite routines

var _func_FreeImage_Copy_ = &c.FuncPrototype{Name: "FreeImage_Copy", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Copy(FIBITMAP *dib, int left, int top, int right, int bottom);
func (dib *BitMap) Copy(left, top, right, bottom int32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_Copy_, inArgs{&dib, &left, &top, &right, &bottom}).PtrFree())
}

var _func_FreeImage_Paste_ = &c.FuncPrototype{Name: "FreeImage_Paste", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_Paste(FIBITMAP *dst, FIBITMAP *src, int left, int top, int alpha);
//
// alpha : alpha blend factor. The source and destination images are alpha blended if
// alpha=0..255. If alpha > 255, then the source image is combined to the destination image.
func (dib *BitMap) Paste(src *BitMap, left, top int32, alpha int32) bool {
	return fiLib.Call(_func_FreeImage_Paste_, inArgs{&dib, &src, &left, &top, &alpha}).BoolFree()
}

var _func_FreeImage_CreateView_ = &c.FuncPrototype{Name: "FreeImage_CreateView", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.U32, c.U32, c.U32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_CreateView(FIBITMAP *dib, unsigned left, unsigned top, unsigned right, unsigned bottom);
func (dib *BitMap) CreateView(left, top, right, bottom uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_CreateView_, inArgs{&dib, &left, &top, &right, &bottom}).PtrFree())
}

var _func_FreeImage_Composite_ = &c.FuncPrototype{Name: "FreeImage_Composite", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer, c.Pointer, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_Composite(FIBITMAP *fg, BOOL useFileBkg FI_DEFAULT(FALSE), RGBQUAD *appBkColor FI_DEFAULT(NULL), FIBITMAP *bg FI_DEFAULT(NULL));
func Composite(fg *BitMap, useFileBkg bool, appBkColor *RGBQUAD, bg *BitMap) *BitMap {
	ufb := c.CBool(useFileBkg)
	return (*BitMap)(fiLib.Call(_func_FreeImage_Composite_, inArgs{&fg, &ufb, &appBkColor, &bg}).PtrFree())
}

var _func_FreeImage_PreMultiplyWithAlpha_ = &c.FuncPrototype{Name: "FreeImage_PreMultiplyWithAlpha", OutType: c.I32, InTypes: []c.Type{c.Pointer}}

// DLL_API BOOL DLL_CALLCONV FreeImage_PreMultiplyWithAlpha(FIBITMAP *dib);
func (dib *BitMap) PreMultiplyWithAlpha() bool {
	return fiLib.Call(_func_FreeImage_PreMultiplyWithAlpha_, inArgs{&dib}).BoolFree()
}

// // background filling routines

var _func_FreeImage_FillBackground_ = &c.FuncPrototype{Name: "FreeImage_FillBackground", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32}}

// DLL_API BOOL DLL_CALLCONV FreeImage_FillBackground(FIBITMAP *dib, const void *color, int options FI_DEFAULT(0));
func (dib *BitMap) FillBackground(color interface{}, options int32) bool {
	clr := usf.AddrOf(color)
	return fiLib.Call(_func_FreeImage_FillBackground_, inArgs{&dib, &clr, &options}).BoolFree()
}

var _func_FreeImage_EnlargeCanvas_ = &c.FuncPrototype{Name: "FreeImage_EnlargeCanvas", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_EnlargeCanvas(FIBITMAP *src, int left, int top, int right, int bottom, const void *color, int options FI_DEFAULT(0));
func (dib *BitMap) EnlargeCanvas(left, top, right, bottom int32, color interface{}, options int32) *BitMap {
	clr := usf.AddrOf(color)
	return (*BitMap)(fiLib.Call(_func_FreeImage_EnlargeCanvas_, inArgs{&dib, &left, &top, &right, &bottom, &clr, &options}).PtrFree())
}

var _func_FreeImage_AllocateEx_ = &c.FuncPrototype{Name: "FreeImage_AllocateEx", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.Pointer, c.I32, c.Pointer, c.U32, c.U32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_AllocateEx(int width, int height, int bpp, const RGBQUAD *color, int options FI_DEFAULT(0), const RGBQUAD *palette FI_DEFAULT(NULL), unsigned red_mask FI_DEFAULT(0), unsigned green_mask FI_DEFAULT(0), unsigned blue_mask FI_DEFAULT(0));
func AllocateEx(width, height, bpp int32, color *RGBQUAD, options int32, palette *RGBQUAD, red_mask, green_mask, blue_mask uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_AllocateEx_, inArgs{&width, &height, &bpp, &color, &options, &palette, &red_mask, &green_mask, &blue_mask}).PtrFree())
}

var _func_FreeImage_AllocateExT_ = &c.FuncPrototype{Name: "FreeImage_AllocateExT", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer, c.I32, c.Pointer, c.U32, c.U32, c.U32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_AllocateExT(FREE_IMAGE_TYPE type, int width, int height, int bpp, const void *color, int options FI_DEFAULT(0), const RGBQUAD *palette FI_DEFAULT(NULL), unsigned red_mask FI_DEFAULT(0), unsigned green_mask FI_DEFAULT(0), unsigned blue_mask FI_DEFAULT(0));
func AllocateExT(typ FREE_IMAGE_TYPE, width, height, bpp int32, color *RGBQUAD, options int32, palette *RGBQUAD, red_mask, green_mask, blue_mask uint32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_AllocateExT_, inArgs{&typ, &width, &height, &bpp, &color, &options, &palette, &red_mask, &green_mask, &blue_mask}).PtrFree())
}

// // miscellaneous algorithms

var _func_FreeImage_MultigridPoissonSolver_ = &c.FuncPrototype{Name: "FreeImage_MultigridPoissonSolver", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}

// DLL_API FIBITMAP *DLL_CALLCONV FreeImage_MultigridPoissonSolver(FIBITMAP *Laplacian, int ncycle FI_DEFAULT(3));
func (dib *BitMap) MultigridPoissonSolver(ncycle int32) *BitMap {
	return (*BitMap)(fiLib.Call(_func_FreeImage_MultigridPoissonSolver_, inArgs{&dib, &ncycle}).PtrFree())
}
